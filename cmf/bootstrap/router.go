package bootstrap

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/gincmf/cmf/controller"
	"github.com/gincmf/cmf/data"
	"github.com/gorilla/websocket"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"strings"
)

type GroupMapStruct data.GroupMapStruct

var (
	g    errgroup.Group
	conn *websocket.Conn
)

// ws router

var wsRouterMap []data.WsRouterStruct
var routerMap []data.RouterMapStruct
var TemplateMap data.TemplateMapStruct

var engine *gin.Engine
var theme, path, themePath string

var HandleFunc []gin.HandlerFunc

func Start(port ...string) {

	appPort := getPort(port...)

	server := &http.Server{
		Addr:    ":" + appPort,
		Handler: register(),
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	// 捕获err
	if err := g.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Get all num successfully!")
	}
}

func getPort(port ...string) string {
	appPort := config.App.Port
	if appPort == "" {
		if len(port) > 0 {
			appPort = port[0]
		} else {
			appPort = "8000"
		}
	}
	return appPort
}

func register() http.Handler {
	//注册路由
	engine = gin.Default()

	engine.Use(HandleFunc...)

	var store cookie.Store
	var err error
	sessionStore := cookie.NewStore([]byte(config.Database.AuthCode))
	if config.Redis.Host != "" && config.Redis.Enabled {
		store, err = redis.NewStore(10, "tcp", config.Redis.Host+":"+config.Redis.Port, config.Redis.Port, []byte(config.Database.AuthCode))
		if err != nil {
			fmt.Println("[ERROR]", fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(91), err.Error()))
			store = sessionStore
		}
	} else {
		store = sessionStore
	}
	if store != nil {
		engine.Use(sessions.Sessions("mySession", store))
	}
	// engine.Delims("${", "}")
	LoadTemplate() //加载模板
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rangeRouter(routerMap)
	rangeSocket(wsRouterMap)

	//扫描主题路径
	if path != "" {
		files := scanThemeDir(path)
		for _, t := range files {
			//扫描项目模板下的全部模块
			engine.StaticFS(path+"/"+t.name+"/"+"assets", http.Dir(t.path+"/public/assets"))
		}
		//加载uploads静态资源
		engine.StaticFS("/uploads", http.Dir("public/uploads"))
		engine.StaticFS("/exports", http.Dir("public/exports"))
	}
	appPort := getPort()
	_ = engine.Run(":" + appPort)
	//配置路由端口
	return engine
}

func rangeRouter(routerMap []data.RouterMapStruct) {
	for _, router := range routerMap {
		switch router.Method {
		case "GET":
			engine.GET(router.RelativePath, router.Handlers...)
		case "POST":
			engine.POST(router.RelativePath, router.Handlers...)
		case "PUT":
			engine.PUT(router.RelativePath, router.Handlers...)
		case "DELETE":
			engine.DELETE(router.RelativePath, router.Handlers...)
		default:
		}
	}
}

func rangeSocket(routerMap []data.WsRouterStruct) {
	for _, router := range wsRouterMap {
		engine.GET(router.RelativePath, router.Handlers...)
	}
}

func Socket(relativePath string, handler gin.HandlerFunc, handlers ...gin.HandlerFunc) {

	socketHandler := func(c *gin.Context) {
		var upgrader = websocket.Upgrader{
			// 允许跨域
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		} // use default options

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(0, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Set("websocket", conn)
	}

	handlersMap := []gin.HandlerFunc{socketHandler}

	for _, v := range handlers {
		handlersMap = append(handlersMap, v)
	}

	handlersMap = append(handlersMap, handler)

	wsRouterMap = append(wsRouterMap, data.WsRouterStruct{
		RelativePath: relativePath,
		Handlers:     handlersMap,
	})
}

//抛出对外注册路由方法
func Get(relativePath string, handler gin.HandlerFunc, handlers ...gin.HandlerFunc) {
	handlers = append(handlers, handler)
	routerMap = append(routerMap, data.RouterMapStruct{RelativePath: relativePath, Handlers: handlers, Method: "GET"})
}

func Post(relativePath string, handler gin.HandlerFunc, handlers ...gin.HandlerFunc) {
	handlers = append(handlers, handler)
	routerMap = append(routerMap, data.RouterMapStruct{RelativePath: relativePath, Handlers: handlers, Method: "POST"})
}

func Delete(relativePath string, handler gin.HandlerFunc, handlers ...gin.HandlerFunc) {
	handlers = append(handlers, handler)
	routerMap = append(routerMap, data.RouterMapStruct{RelativePath: relativePath, Handlers: handlers, Method: "DELETE"})
}

//处理资源控制器
func Rest(relativePath string, restController controller.RestInterface, handlers ...gin.HandlerFunc) {
	if relativePath == "/" {
		routerMap = append(routerMap, data.RouterMapStruct{Handlers: []gin.HandlerFunc{restController.Get}, Method: "GET"})
	} else {

		var (
			handlerGet    = []gin.HandlerFunc{}
			handlerShow   = []gin.HandlerFunc{}
			handlerEdit   = []gin.HandlerFunc{}
			handlerStore  = []gin.HandlerFunc{}
			handlerDelete = []gin.HandlerFunc{}
		)

		for _, v := range handlers {
			handlerGet = append(handlerGet, v)
			handlerShow = append(handlerShow, v)
			handlerEdit = append(handlerEdit, v)
			handlerStore = append(handlerStore, v)
			handlerDelete = append(handlerDelete, v)
		}

		handlerGet = append(handlerGet, restController.Get)
		routerMap = append(routerMap, data.RouterMapStruct{RelativePath: relativePath, Handlers: handlerGet, Method: "GET"}) //查询全部
		rPath := strings.TrimRight(relativePath, "/") + "/"

		handlerShow = append(handlerShow, restController.Show)
		routerMap = append(routerMap, data.RouterMapStruct{RelativePath: rPath + ":id", Handlers: handlerShow, Method: "GET"}) //查询一条

		handlerEdit = append(handlerEdit, restController.Edit)
		routerMap = append(routerMap, data.RouterMapStruct{RelativePath: rPath + ":id", Handlers: handlerEdit, Method: "POST"}) //编辑一条

		handlerStore = append(handlerStore, restController.Store)
		routerMap = append(routerMap, data.RouterMapStruct{RelativePath: relativePath, Handlers: handlerStore, Method: "POST"}) //新增一条

		handlerDelete = append(handlerDelete, restController.Delete)
		routerMap = append(routerMap, data.RouterMapStruct{RelativePath: rPath + ":id", Handlers: handlerDelete, Method: "DELETE"}) //删除一条

		routerMap = append(routerMap, data.RouterMapStruct{RelativePath: rPath, Handlers: handlerDelete, Method: "DELETE"}) //删除全部
	}
}

// 路由组
func Group(relativePath string, handlers ...gin.HandlerFunc) GroupMapStruct {
	return GroupMapStruct{
		RelativePath: relativePath,
		Handlers:     handlers,
	}
}

func (group *GroupMapStruct) Rest(relativePath string, restController controller.RestInterface, handlers ...gin.HandlerFunc) {
	// 临时赋值
	rPath := ""
	if group.RelativePath != "" {
		rPath = strings.TrimRight(group.RelativePath, "")
	}
	handlers = append(group.Handlers, handlers...)
	Rest(rPath+relativePath, restController, handlers...)
}

func (group *GroupMapStruct) Get(relativePath string, handler gin.HandlerFunc, handlers ...gin.HandlerFunc) {
	rPath := ""
	if group.RelativePath != "" {
		rPath = strings.TrimRight("/"+group.RelativePath, "/") + "/"
	}

	handlers = append(group.Handlers, handlers...)
	Get(rPath+relativePath, handler, handlers...)
}

func (group *GroupMapStruct) Post(relativePath string, handler gin.HandlerFunc, handlers ...gin.HandlerFunc) {
	rPath := ""
	if group.RelativePath != "" {
		rPath = strings.TrimRight("/"+group.RelativePath, "/") + "/"
	}
	handlers = append(group.Handlers, handlers...)
	Post(rPath+relativePath, handler, handlers...)
}

func (group *GroupMapStruct) Delete(relativePath string, handler gin.HandlerFunc, handlers ...gin.HandlerFunc) {
	rPath := ""
	if group.RelativePath != "" {
		rPath = strings.TrimRight("/"+group.RelativePath, "/") + "/"
	}
	handlers = append(group.Handlers, handlers...)
	Delete(rPath+relativePath, handler, handlers...)
}

func LoadTemplate() {
	//加载全部主题路径
	if TemplateMap.Theme != "" && TemplateMap.ThemePath != "" {
		theme = strings.TrimRight(TemplateMap.Theme, "/") + "/"
		path = strings.TrimRight(TemplateMap.ThemePath, "/") + "/"
		themePath = path + theme
		files := scanFiles(themePath)
		engine.LoadHTMLFiles(files...)
	}
}

type themeDirStruct struct {
	name string
	path string
}

func scanThemeDir(path string) []themeDirStruct {
	dirs, _ := ioutil.ReadDir(path)
	var dirList []themeDirStruct

	for _, dir := range dirs {
		if dir.IsDir() {
			dirList = append(dirList, themeDirStruct{dir.Name(), strings.TrimRight(path+dir.Name(), "/") + "/"})
		}
	}
	return dirList
}

func scanDir(path string) ([]string, []string) {
	dirs, _ := ioutil.ReadDir(path)
	var dirList, fileList []string
	for _, dir := range dirs {
		if dir.IsDir() {
			dirList = append(dirList, strings.TrimRight(path+dir.Name(), "/")+"/")
		} else {
			fileList = append(fileList, path+dir.Name())
		}
	}
	return dirList, fileList
}

// 递归扫描目录
func scanFiles(dirName string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
	}
	var fileList []string
	for _, file := range files {
		if file.IsDir() {
			subList := scanFiles(strings.TrimRight(dirName+file.Name(), "/") + "/")
			fileList = append(fileList, subList...)

		} else {
			suffix := strings.Split(file.Name(), ".")
			if len(suffix) > 1 && suffix[1] == "html" {
				fileList = append(fileList, dirName+file.Name())
			}
		}
	}
	return fileList
}
