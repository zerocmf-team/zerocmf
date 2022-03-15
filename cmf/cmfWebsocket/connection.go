/**
** @创建时间: 2021/3/14 1:53 下午
** @作者　　: return
** @描述　　: 定义连接池
 */
package cmfWebsocket

import (
	"errors"
	"fmt"
	"github.com/gincmf/cmf/controller"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
}

// 初始化连接
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {

	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	// 启动读携程
	go conn.readLoop()
	// 启动写携程
	go conn.writeLoop()

	return
}

// 读取消息
func (conn *Connection) ReadMessage() (data []byte, err error) {

	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("连接已断开！")
	}
	return
}

// 发送消息
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("连接已断开！")
	}
	return
}

// 关闭连接
func (conn *Connection) Close() {
	// 线程安全，可重入的
	conn.wsConn.Close()

	conn.mutex.Lock()
	if conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}

		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			goto ERR
		}

	}
ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err := conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}

	}
ERR:
	conn.Close()

}

func (conn *Connection) Success(msg string, data interface{}) error {
	msg = new(controller.Rest).JsonSuccess(msg, data)

	if err := conn.WriteMessage([]byte(msg)); err != nil {
		fmt.Println("err", err)
		conn.Close()
		return errors.New("链接已关闭！")
	}
	return nil
}

func (conn *Connection) Error(msg string, data interface{}) {
	msg = new(controller.Rest).JsonError(msg, nil)
	if err := conn.WriteMessage([]byte(msg)); err != nil {
		conn.Close()
	}
	time.Sleep(time.Microsecond * 100)
	return
}
