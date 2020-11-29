package gcs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nbcx/gcs/server"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

type wssConnection struct {
	*server.BaseConnection
	Socket   *websocket.Conn // 用户连接
	wsServer *wsServer
}

// 初始化
func newWssConnection(appId, addr string, socket *websocket.Conn, ws *wsServer) (client *wssConnection) {
	client = &wssConnection{
		BaseConnection: server.NewBaseConnection(appId, GenClientId(), addr),
		Socket:         socket,
		wsServer:       ws,
	}
	return
}

// 读取客户端数据
func (c *wssConnection) read() {
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
func (c *wssConnection) Close() {
	if c == nil {
		return
	}

	c.wsServer.Close <- c
	c.Socket.Close()
}

// 直接向客户端写数据
func (c *wssConnection) Write(message []byte) {
	if c == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			c.Close()
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()
	c.Socket.WriteMessage(websocket.TextMessage, message)
}

// 直接向客户端写数据
func (c *wssConnection) WriteString(message string) {
	c.Write([]byte(message))
}
