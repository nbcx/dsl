package gcs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nbcx/dsl/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Message func(client *WssConnection, message []byte)
type Open func(client *WssConnection, r *http.Request)
type Close func(client *WssConnection)
type Upgrade func(res http.ResponseWriter, req *http.Request) (aid string, conn *websocket.Conn, err error)

type open struct {
	connection *WssConnection
	request    *http.Request
}

type wsServer struct {
	Open          chan *open          // 连接处理
	Close         chan *WssConnection // 断开连接处理程序
	Message       chan []byte         // 广播向全部成员发送数据
	trigerUpgrade Upgrade
	trigerOpen    Open
	trigerMessage Message
	trigerClose   Close
	ip            string
	port          int
	server        *util.Server
	path          string
}

func NewWsServer(ip string, port int, path string) (wsSer *wsServer) {
	wsSer = &wsServer{
		Open:    make(chan *open, 1000),
		Close:   make(chan *WssConnection, 1000),
		Message: make(chan []byte, 1000),
		ip:      ip,
		port:    port,
		server:  util.NewServer(ip, port),
		path:    path,
	}
	return
}

func (ws *wsServer) EventUpgrade(u Upgrade) {
	ws.trigerUpgrade = u
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
func (ws *wsServer) eventRegister(o *open) {
	Manager.Add(o.connection)
	fmt.Println("EventRegister 用户建立连接", o.connection.GetAddr())
	if ws.trigerOpen != nil {
		ws.trigerOpen(o.connection, o.request)
	}
}

// 用户断开连接
func (ws *wsServer) eventUnregister(c *WssConnection) {
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
			ws.eventRegister(conn)
		case conn := <-ws.Close:
			// 断开连接事件
			ws.eventUnregister(conn)
		case message := <-ws.Message:
			// 广播事件
			for _, conn := range Manager.Connections {
				conn.WriteByte(message)
			}
		}
	}
}

func (ws *wsServer) upgrade(res http.ResponseWriter, req *http.Request) {
	var aid string
	var conn *websocket.Conn
	var err error
	if ws.trigerUpgrade != nil {
		aid, conn, err = ws.trigerUpgrade(res, req)
	} else {
		conn, err = (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { // 升级协议
			log.Info("upgrader ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
			aid = "hello"
			return true
		}}).Upgrade(res, req, nil)
	}
	if err != nil {
		res.Write(util.Str2bytes(err.Error()))
		http.NotFound(res, req)
		return
	}
	log.Infof("%s connetion", conn.RemoteAddr().String())
	client := newWssConnection(aid, conn.RemoteAddr().String(), conn, ws)
	go client.read()
	// 用户连接事件
	ws.Open <- &open{
		connection: client,
		request:    req,
	}
}

// Websocket 服务启动
func (ws *wsServer) Start() {
	Manager.Start()
	go ws.event() // 添加事件处理程序,管道处理程序
	log.Infof("websocket server startup in %s:%d", ws.ip, ws.port)
	http.HandleFunc(ws.path, ws.upgrade)
	http.ListenAndServe(fmt.Sprintf("%s:%d", ws.ip, ws.port), nil)
}
