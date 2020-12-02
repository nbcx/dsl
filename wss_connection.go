package gcs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nbcx/gcs/server"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

type WssConnection struct {
	*server.BaseConnection
	Socket   *websocket.Conn // 用户连接
	wsServer *wsServer
}

// 初始化
func newWssConnection(appId, addr string, socket *websocket.Conn, ws *wsServer) (client *WssConnection) {
	client = &WssConnection{
		BaseConnection: server.NewBaseConnection(appId, GenClientId(), addr),
		Socket:         socket,
		wsServer:       ws,
	}
	return
}

// 读取客户端数据
func (c *WssConnection) read() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		log.Info("connection close", c)
		c.Close()
		//close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			log.Warn("read clent data error,", c.GetAddr(), err)
			return
		}

		log.Info("client data:", string(message))

		//router(c, message)
		c.wsServer.trigerMessage(c, message)
	}
}

// 关闭客户端连接
func (c *WssConnection) Close() {
	if c == nil {
		return
	}

	c.wsServer.Close <- c
	c.Socket.Close()
}

// 直接向客户端写数据
func (c *WssConnection) Write(message string) {
	if c == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			c.Close()
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()
	c.WriteByte([]byte(message))
}

// 直接向客户端写数据
func (c *WssConnection) WriteByte(message []byte) {
	c.Socket.WriteMessage(websocket.TextMessage, message)
}
