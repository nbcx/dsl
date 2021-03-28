package server

import "time"

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

// 用户连接
type BaseConnection struct {
	fd            string   // 客户端ID
	addr          string   // 客户端地址
	uid           string   // 用户Id，用户登录以后才有
	appId         string   // 应用ID
	group         []string // 连接加入的所有分组
	firstTime     uint64   // 首次连接事件
	heartbeatTime uint64   // 用户上次心跳时间
	loginTime     uint64   // 登录时间 登录以后才有
}

func NewBaseConnection(appId, fd, addr string) *BaseConnection {
	currentTime := uint64(time.Now().Unix())
	return &BaseConnection{
		appId:         appId,
		fd:            fd,
		addr:          addr,
		firstTime:     currentTime,
		heartbeatTime: currentTime,
	}
}
func (c *BaseConnection) SetLoginTime(time uint64) {
	c.loginTime = time
}

func (c *BaseConnection) GetLoginTime() uint64 {
	return c.loginTime
}

func (c *BaseConnection) GetAppId() string {
	return c.appId
}

func (c *BaseConnection) GetUid() string {
	return c.uid
}

func (c *BaseConnection) GetGroup() []string {
	return c.group
}

func (c *BaseConnection) GetFd() string {
	return c.fd
}

func (c *BaseConnection) GetAddr() string {
	return c.addr
}

func (c *BaseConnection) SetAddr(addr string) {
	c.addr = addr
}

func (c *BaseConnection) SetAppId(appId string) {
	c.appId = appId
}

func (c *BaseConnection) SetFd(fd string) {
	c.fd = fd
}
func (c *BaseConnection) SetUid(uid string) {
	c.uid = uid
}

func (c *BaseConnection) Close() {

}

func (c *BaseConnection) Write(message string) {
}

func (c *BaseConnection) WriteByte(message []byte) {
}

// 用户心跳
func (c *BaseConnection) Heartbeat(currentTime uint64) {
	c.heartbeatTime = currentTime
	return
}

// 心跳超时
func (c *BaseConnection) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.heartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

// 是否登录了
func (c *BaseConnection) IsLogin() (isLogin bool) {
	// 用户登录了
	if c.uid != "" {
		isLogin = true
		return
	}
	return
}

// 加入组
func (c *BaseConnection) JoinGroup(gids ...string) (err error) {
	for _, gid := range gids {
		manager.joinGroup(c, gid)
	}
	return
}

// 加入组
func (c *BaseConnection) ExitGroup(gids ...string) (err error) {
	for _, gid := range gids {
		manager.exitGroup(c, gid)
	}
	return
}
