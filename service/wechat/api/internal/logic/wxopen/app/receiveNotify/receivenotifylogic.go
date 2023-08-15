package receiveNotify

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"zerocmf/common/bootstrap/data"
	"zerocmf/service/tenant/rpc/tenantclient"

	wechatUtil "github.com/zerocmf/wechatEasySdk/util"

	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewReceiveNotifyLogic(header *http.Request, svcCtx *svc.ServiceContext) *ReceiveNotifyLogic {
	ctx := header.Context()
	return &ReceiveNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveNotifyLogic) ReceiveNotify(req *types.ReceiveNotifyReq) (resp data.Rest) {

	c := l.svcCtx
	r := l.header
	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Error("请求失败！", err.Error())
		return
	}

	// 关闭请求Body
	defer r.Body.Close()

	// 检查目录是否存在
	file, err := os.OpenFile("logs/receive.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法打开日志文件:", err)
	}
	defer file.Close()

	log.SetOutput(file)

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

	authAppId := req.AppId

	//var reply *tenantclient.ShowMpData
	_, err = c.TenantRpc.ShowMp(l.ctx, &tenantclient.ShowMpData{
		AppId: authAppId,
	})

	if err != nil {
		resp.Error("授权小程序不一致！", nil)
		return
	}

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

	log.Println(string(body))
	log.Println(string(result))
	log.Println("--------------------------------------------")

	type agent struct {
		Name      string `xml:"name"`
		Phone     string `xml:"phone"`
		ReachTime string `xml:"reach_time"`
	}

	var message struct {
		ToUserName   string `xml:"ToUserName"`
		FromUserName string `xml:"FromUserName"`
		CreateTime   int    `xml:"CreateTime"`
		MsgType      string `xml:"MsgType"`
		Event        string `xml:"Event"`
		Shopid       string `xml:"shopid"`
		ShopOrderId  string `xml:"shop_order_id"`
		WaybillId    string `xml:"waybill_id"`
		ActionTime   int    `xml:"action_time"`
		OrderStatus  int    `xml:"order_status"`
		ActionMsg    string `xml:"action_msg"`
		ShopNo       string `xml:"shop_no"`
		SuccTime     int    `xml:"SuccTime"`
		FailTime     int    `xml:"FailTime"`
		DelayTime    int    `xml:"DelayTime"`
		Reason       string `xml:"Reason"`
		ScreenShot   string `xml:"ScreenShot"`
		Agent        agent  `xml:"agent"`
	}

	err = xml.Unmarshal(result, &message)
	if err != nil {
		resp.Error("非法请求！", nil)
		return
	}

	switch message.Event {

	}

	resp.String("success")
	return
}
