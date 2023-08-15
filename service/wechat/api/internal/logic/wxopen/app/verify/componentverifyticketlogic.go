package verify

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/wechat/api/internal/svc"

	wechatUtil "github.com/zerocmf/wechatEasySdk/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComponentVerifyTicketLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewComponentVerifyTicketLogic(header *http.Request, svcCtx *svc.ServiceContext) *ComponentVerifyTicketLogic {
	ctx := header.Context()
	return &ComponentVerifyTicketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ComponentVerifyTicketLogic) ComponentVerifyTicket() (resp data.Rest) {
	r := l.header
	c := l.svcCtx
	redis := c.Redis

	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Error("请求失败！", err.Error())
		return
	}

	// 关闭请求Body
	defer r.Body.Close()

	// 检查目录是否存在
	dirPath := "logs"
	if _, err = os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，创建它
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	err = os.WriteFile("logs/test.log", body, 0644)
	if err != nil {
		resp.Error("请求失败！", err.Error())
		return
	}

	var form struct {
		AppId   string `xml:"AppId"`
		Encrypt string `xml:"Encrypt"`
	}

	err = xml.Unmarshal(body, &form)
	if err != nil {
		resp.Error("请求失败！", err.Error())
		return
	}

	appId := strings.TrimSpace(form.AppId)
	encrypt := strings.TrimSpace(form.Encrypt)

	encryptBytes, _ := base64.StdEncoding.DecodeString(encrypt)

	aesKey := c.Config.Wechat.WxOpen.Aeskey
	if aesKey == "" {
		resp.Error("请求失败！aesKey为空", nil)
		return
	}

	var result []byte
	result, err = wechatUtil.AesDecrypt(encryptBytes, []byte(aesKey))
	if err != nil {
		resp.Error("请求失败!", err.Error())
		return
	}

	end := strings.LastIndex(string(result), appId)

	if end == -1 {
		resp.Error("非法签名！", nil)
		return
	}

	result = result[20:end]

	var data struct {
		AppId                 string
		CreateTime            int64
		InfoType              string
		ComponentVerifyTicket string
	}

	// 解析XML数据到Map
	err = xml.Unmarshal(result, &data)
	if err != nil {
		resp.Error("请求失败！", err.Error())
		return
	}

	fmt.Println(data)

	switch data.InfoType {
	case "component_verify_ticket":
		tickets := strings.TrimSpace(data.ComponentVerifyTicket)
		if tickets != "" {
			c.Config.Wechat.WxOpen.ComponentVerifyTicket = tickets
			redis.Set("componentVerifyTicket", tickets, time.Hour*12)
		}
	default:
		resp.Error("err", data)
		return
	}

	// 输出解析结果
	resp.String("success")

	return
}
