/**
** @创建时间: 2020/12/8 12:01 下午
** @作者　　: return
** @描述　　:
 */
package bootstrap

import "github.com/gorilla/websocket"

type Connection struct {
	WsConn *websocket.Conn
	IsStart bool
}

func InitConn(wsConn *websocket.Conn) (*Connection,error)  {
	return &Connection{
		WsConn: wsConn,
		IsStart:false,
	},nil
}
