package role

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"zerocmf/common/bootstrap/util"
	"zerocmf/service/lowcode/api/internal/svc"
	"zerocmf/service/lowcode/api/internal/types"
	"zerocmf/service/lowcode/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreLogic) Store(req *types.UserAdminRoleSave) (resp types.Response) {
	c := l.svcCtx
	resp = save(c, req)
	return
}

func save(c *svc.ServiceContext, req *types.UserAdminRoleSave) (resp types.Response) {

	siteId, _ := c.Get("siteId")
	// 选择租户表
	db, err := c.MongoDB(siteId.(string))
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	collection := db.Collection("role")
	role := model.Role{
		Status: 1,
	}

	access := req.Access

	var e *casbin.Enforcer
	//访问权限控制
	e, err = db.NewEnforcer()
	if err != nil {
		resp.Error("enforcer err", err.Error())
		return
	}

	roleShow := model.Role{}
	err = db.FindOne(collection, bson.M{
		"name":      req.Name,
		"deletedAt": time.Now().Unix(),
	}, &roleShow)
	if err != nil && err != mongo.ErrNoDocuments {
		resp.Error(err.Error(), nil)
		return
	}

	if req.Id == "" {

		if !roleShow.Id.IsZero() {
			resp.Error("该角色已存在！", nil)
			return
		}

		copier.Copy(&role, &req)
		role.CreateAt = time.Now().Unix()

		var one *mongo.InsertOneResult
		one, err = db.InsertOne(collection, role)
		if err != nil {
			resp.Error(err.Error(), nil)
			return
		}
		role.Id = one.InsertedID.(primitive.ObjectID)

		id := role.Id.Hex()
		for _, v := range access {
			e.AddPolicy(id, v, "*")
		}
	} else {

		var objectId primitive.ObjectID
		objectId, err = primitive.ObjectIDFromHex(req.Id)

		if err != nil {
			resp.Error("系统异常", err.Error())
			return
		}

		filter := bson.M{
			"_id": objectId,
		}

		err = db.FindOne(collection, filter, &role)

		// 不是同一条字段
		if !roleShow.Id.IsZero() && role.Id != roleShow.Id {
			resp.Error("该角色名称已存在", nil)
			return
		}

		if err != nil {
			if err == mongo.ErrNoDocuments {
				resp.Error("改角色不存在或已被删除！", err.Error())
				return
			}
			resp.Error("系统异常", err.Error())
			return
		}

		copier.Copy(&role, &req)
		role.UpdateAt = time.Now().Unix()
		update := bson.M{
			"$set": role,
		}

		_, err = db.UpdateOne(collection, filter, update)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				resp.Error("系统错误", err.Error())
				return
			}
		}

		id := objectId.Hex()

		// 获取全部角色策略
		roles := e.GetFilteredPolicy(0, id)
		existAccess := make([]string, 0)
		for _, v := range roles {
			if len(v) > 1 {
				existAccess = append(existAccess, v[1])
			}
		}

		/*
		*  新增：[1,2,3]
		*  原有：[3,4,5]
		*  筛选去除的规则：[4，5]
		*  筛选新增的规则：[1，2]

		* 新增：[1,2,3]
		* 原有：[]
		* 筛选去除的规则：[]
		* 筛选新增的规则：[1，2，3]

		* 新增：[]
		* 原有：[1，2，3]
		* 筛选去除的规则：[1，2，3]
		* 筛选新增的规则：[]
		 */

		alreadyDel := make([]string, 0)
		// 判断是否需要被删除
		for _, v := range existAccess {
			if util.ToLowerInArray(v, access) == false {
				alreadyDel = append(alreadyDel, v)
			}
		}
		// 如果新增为空，则全部删除
		if len(access) == 0 {
			alreadyDel = existAccess
		}

		alreadyAdd := make([]string, 0)
		for _, v := range access {
			if util.ToLowerInArray(v, existAccess) == false {
				alreadyAdd = append(alreadyAdd, v)
			}
		}

		// 如果数据库不存在，则为新增
		if len(existAccess) == 0 {
			alreadyAdd = access
		}

		// 开始删除策略
		rules := make([][]string, 0)
		for _, v := range alreadyDel {
			rules = append(rules, []string{id, v, "*"})
		}
		if len(rules) > 0 {
			e.RemovePolicies(rules)
		}

		// 开始新增策略
		rules = make([][]string, 0)
		for _, v := range alreadyAdd {
			rules = append(rules, []string{id, v, "*"})
		}
		if len(rules) > 0 {
			e.AddPolicies(rules)
		}
	}

	resp.Success("操作成功！", role)
	return
}
