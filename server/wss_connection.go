package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"gosh/base"
	"runtime/debug"
	"time"
)

type WssConnection struct {
	*BaseConnection
	Socket   *websocket.Conn // 用户连接
	wsServer *WsServer
}

// 初始化
func NewWssConnection(appId, addr string, socket *websocket.Conn, ws *WsServer) (client *WssConnection) {
	currentTime := uint64(time.Now().Unix())
	client = &WssConnection{
		BaseConnection: &BaseConnection{
			appId:         appId,
			fd:            base.GenClientId(),
			addr:          addr,
			firstTime:     currentTime,
			heartbeatTime: currentTime,
		},
		Socket:   socket,
		wsServer: ws,
	}
	return
}

// 用户登录
func (c *WssConnection) Login(userId string, loginTime uint64) {
	c.uid = userId
	c.loginTime = loginTime
	c.Heartbeat(loginTime) // 登录成功=心跳一次
	clientManager.login(c)
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
			log.Warn("read clent data error,", c.addr, err)
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
func (c *WssConnection) Write(message []byte) {
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
func (c *WssConnection) WriteString(message string) {
	c.Write([]byte(message))
}
