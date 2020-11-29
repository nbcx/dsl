package server

import (
	"fmt"
	"github.com/nbcx/gcs/util"
	"sync"
	"time"
)

// 连接管理
type ClientManager struct {
	Connections            map[string]IConnection // 全部的连接
	ConnectionLock         sync.RWMutex           // 读写锁
	apps                   map[string]*App        // 针对应用进行连接记录管理
	heartbeatCheckInterval int32                  // 设置每5秒服务器会侦测一次心跳
	heartbeatIdleTime      int32                  // 设置一个TCP连接如果在10秒内未向服务器发送数据则被切断
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Connections: make(map[string]IConnection),
		apps:        make(map[string]*App),
	}
	return
}

func (manager *ClientManager) getApp(appId string) (result *App) {
	if value, ok := manager.apps[appId]; ok {
		result = value
	} else {
		manager.apps[appId] = &App{
			Connections: make(map[string]IConnection),
			Users:       make(map[string]IConnection),
		}
		result = manager.apps[appId]
	}
	return
}

func (manager *ClientManager) joinGroup(client IConnection, groupId string) (result bool) {
	app := manager.getApp(client.GetAppId())
	app.JoinGroup(groupId, client)
	return
}

func (manager *ClientManager) exitGroup(groupId string, client IConnection) (result bool) {
	app := manager.getApp(client.GetAppId())
	app.ExitGroup(groupId, client)
	return
}

// 删除客户端
func (manager *ClientManager) Del(c IConnection) {
	app := manager.getApp(c.GetAppId())
	app.del(c)
	manager.ConnectionLock.RLock()
	defer manager.ConnectionLock.RUnlock()
	delete(manager.Connections, c.GetFd())
}

// 添加客户端
func (manager *ClientManager) Add(c IConnection) {
	app := manager.getApp(c.GetAppId())
	app.add(c)
	manager.ConnectionLock.Lock()
	defer manager.ConnectionLock.Unlock()
	manager.Connections[c.GetFd()] = c
}

// 通过FD找到连接
func (manager *ClientManager) Find(fd string) (c IConnection) {
	manager.ConnectionLock.Lock()
	defer manager.ConnectionLock.Unlock()
	c = manager.Connections[fd]
	return
}

// 通过UID获取连接
func (manager *ClientManager) FindWithUid(appId, uid string) (c IConnection) {
	app := manager.getApp(appId)
	c = app.getUserConnection(uid)
	return
}

// 向服务器上所有连接发送消息
func (manager *ClientManager) Send(msg []byte) {
	for _, c := range manager.Connections {
		c.Write(msg)
	}
}

// 向指定APP发送消息
func (manager *ClientManager) SendWithApp(appId string, msg []byte) {
	app := manager.getApp(appId)

	for _, c := range app.Connections {
		c.Write(msg)
	}
}

// 向指定分组发送消息
func (manager *ClientManager) SendWithGroup(appId, gid string, msg []byte) {
	app := manager.getApp(appId)
	for _, c := range app.Groups[gid] {
		fmt.Println(c)
	}
}

// 向App上所有登录用户发送消息
func (manager *ClientManager) SendWithUser(appId string, msg []byte) {
	app := manager.getApp(appId)
	for _, c := range app.Users {
		fmt.Println(c)
	}
}

func (manager *ClientManager) ClearTimeoutConnections(param interface{}) (result bool) {
	currentTime := uint64(time.Now().Unix())
	for _, client := range manager.Connections {
		if client.IsHeartbeatTimeout(currentTime) {
			//log.Info("心跳时间超时 关闭连接", client.UID(), client.LoginTime, client.HeartbeatTime)
			client.Close()
		}
	}
	return
}

func (manager *ClientManager) Start() {
	// 定时清理不活跃连接
	util.Timer(3*time.Second, 30*time.Second, manager.ClearTimeoutConnections, "", nil, nil)
}

// 用户登录
func (manager *ClientManager) Login(c IConnection, userId string, loginTime uint64) {
	app := manager.getApp(c.GetAppId())
	c.SetUid(userId)
	c.SetLoginTime(loginTime)
	c.Heartbeat(loginTime) // 登录成功=心跳一次
	app.Login(c)
}
