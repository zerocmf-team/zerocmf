package mp

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"html"
	"io/ioutil"
	"time"
	"zerocmf/service/wechat/api/internal/svc"
	"zerocmf/service/wechat/api/internal/types"
)

type GatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GatewayLogic {
	return &GatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	SUBSCRIBE = "subscribe"
	SCAN      = "SCAN"
)

type GatewayReq struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Event        string
	EventKey     string
	Ticket       string
}

func (l *GatewayLogic) Gateway() (resp types.Response) {
	c := l.svcCtx
	r := c.Request

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.Error("系统错误", err.Error())
		return
	}

	fmt.Println(string(body))

	gateway := GatewayReq{}
	err = xml.Unmarshal(body, &gateway)
	if err != nil {
		resp.Error("系统错误", err.Error())
		return
	}

	switch gateway.Event {
	case SUBSCRIBE:
		resp = l.subscribe(gateway)
	case SCAN:
		resp = l.subscribe(gateway)
	default:
		resp.Success("暂未处理的事件类型", "")
	}
	return
}

type ReplyText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

func cDATA(value string) (str string) {
	str = fmt.Sprintf("<![CDATA[%s]]>", value)
	return
}

func (l *GatewayLogic) subscribe(gateway GatewayReq) (resp types.Response) {
	eventKey := gateway.EventKey

	fmt.Println("额外加密消息eventKey：", eventKey)

	c := l.svcCtx
	redis := c.Redis
	redis.HMSet(eventKey, map[string]interface{}{"openId": gateway.FromUserName, "appId": gateway.ToUserName})
	redis.Expire(eventKey, time.Second*120)

	reply := ReplyText{
		ToUserName:   cDATA(gateway.FromUserName),
		FromUserName: cDATA(gateway.ToUserName),
		CreateTime:   time.Now().Unix(),
		MsgType:      cDATA("text"),
		Content:      cDATA("欢迎您登录，你的openid为" + gateway.FromUserName),
	}

	res, _ := xml.Marshal(reply)
	resStr := html.UnescapeString(string(res))

	resp.Success("获取成功", resStr)
	return
}
