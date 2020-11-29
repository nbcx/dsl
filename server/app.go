package server

import (
	"sync"
)

type App struct {
	Connections    map[string]IConnection // 全部的连接
	ConnectionLock sync.RWMutex           // 读写锁
	Users          map[string]IConnection // 登录的用户 // appId+uuid
	UserLock       sync.RWMutex           // 读写锁
	Groups         map[string][]string    // 分组连接
	GroupLock      sync.RWMutex           // 读写锁
}

// 加入组
func (a *App) JoinGroup(groupId string, c IConnection) {
	a.GroupLock.Lock()
	defer a.GroupLock.Unlock()
	a.Groups[groupId] = append(a.Groups[groupId], c.Fd())
}

// 退出组
func (a *App) ExitGroup(groupId string, c IConnection) {
	a.GroupLock.Lock()
	defer a.GroupLock.Unlock()

	for index, fd := range a.Groups[groupId] {
		if fd == c.Fd() {
			a.Groups[groupId] = append(a.Groups[groupId][:index], a.Groups[groupId][index+1:]...)
		}
	}
}

// 获取组数据
func (a *App) ListGroup(groupId string) []string {
	a.GroupLock.Lock()
	defer a.GroupLock.Unlock()

	return a.Groups[groupId]
}

// 添加一个客户端连接
func (a *App) Login(c IConnection) {
	a.UserLock.Lock()
	defer a.UserLock.Unlock()
	a.Users[c.UID()] = c
}

func (a *App) Logout(c IConnection) {
	a.UserLock.Lock()
	defer a.UserLock.Unlock()
	delete(a.Connections, c.UID())
}

func (a *App) getUserConnection(uid string) (c IConnection) {
	a.UserLock.Lock()
	defer a.UserLock.Unlock()
	c = a.Users[uid]
	return
}

// 添加一个客户端连接
func (a *App) add(c IConnection) {
	a.ConnectionLock.Lock()
	defer a.ConnectionLock.Unlock()
	a.Connections[c.Fd()] = c
}

// 客户端数量
func (a *App) Count() int {
	a.ConnectionLock.RLock()
	defer a.ConnectionLock.RUnlock()
	return len(a.Connections)
}

// 删除客户端
func (a *App) del(c IConnection) {
	a.ConnectionLock.RLock()
	defer a.ConnectionLock.RUnlock()

	delete(a.Connections, c.Fd())

	// 删除所在的分组
	if len(c.Group()) > 0 {
		for _, groupName := range c.Group() {
			a.ExitGroup(groupName, c)
		}
	}

	// 删除系统里的客户端
	if len(c.UID()) > 0 {
		a.Logout(c)
	}
}
