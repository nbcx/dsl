package gcs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nbcx/gcs/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Message func(client *WssConnection, message []byte)
type Open func(client *WssConnection, r *http.Request)
type Close func(client *WssConnection)

type open struct {
	connection *WssConnection
	request    *http.Request
}

type wsServer struct {
	Open          chan *open          // 连接处理
	Close         chan *WssConnection // 断开连接处理程序
	Message       chan []byte         // 广播向全部成员发送数据
	trigerOpen    Open
	trigerMessage Message
	trigerClose   Close
	addr          string
	server        *util.Server
	path          string
}

func NewWsServer(addr string) (wsSer *wsServer) {
	wsSer = &wsServer{
		Open:    make(chan *open, 1000),
		Close:   make(chan *WssConnection, 1000),
		Message: make(chan []byte, 1000),
		addr:    addr,
		server:  util.AddrToServer(addr),
		path:    "/",
	}
	return
}

func (ws *wsServer) SetPath(path string) {
	ws.path = path
}

func (ws *wsServer) EventOpen(o Open) {
	ws.trigerOpen = o
}

func (ws *wsServer) EventClose(c Close) {
	ws.trigerClose = c
}

func (ws *wsServer) EventMessage(m Message) {
	ws.trigerMessage = m
}

// 用户建立连接事件
func (ws *wsServer) EventRegister(o *open) {
	Manager.Add(o.connection)
	fmt.Println("EventRegister 用户建立连接", o.connection.GetAddr())
	if ws.trigerOpen != nil {
		ws.trigerOpen(o.connection, o.request)
	}
}

// 用户断开连接
func (ws *wsServer) EventUnregister(c *WssConnection) {
	Manager.Del(c)
	fmt.Println("EventUnregister 用户断开连接", c.GetAddr(), c.GetUid())
	if ws.trigerClose != nil {
		ws.trigerClose(c)
	}
}

func (ws *wsServer) event() {
	for {
		select {
		case conn := <-ws.Open:
			// 建立连接事件
			ws.EventRegister(conn)
		case conn := <-ws.Close:
			// 断开连接事件
			ws.EventUnregister(conn)
		case message := <-ws.Message:
			// 广播事件
			for _, conn := range Manager.Connections {
				conn.Write(message)
			}
		}
	}
}

func (ws *wsServer) upgrade(w http.ResponseWriter, req *http.Request) {
	var request *http.Request
	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		log.Info("upgrader ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		request = r
		return true
	}}).Upgrade(w, req, nil)

	if err != nil {
		log.Info(err)
		http.NotFound(w, req)
		return
	}

	log.Infof("%s connetion", conn.RemoteAddr().String())
	client := newWssConnection("hello", conn.RemoteAddr().String(), conn, ws)
	go client.read()

	// 用户连接事件
	ws.Open <- &open{
		connection: client,
		request:    request,
	}
}

// Websocket 服务启动
func (ws *wsServer) Start() {
	Manager.Start()
	go ws.event() // 添加事件处理程序,管道处理程序

	log.Infof("websocket server startup in %s:%s", util.LocalIp, ws.addr)
	http.HandleFunc(ws.path, ws.upgrade)
	http.ListenAndServe(ws.addr, nil)
}
