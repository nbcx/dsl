package server

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

func (c *BaseConnection) AppId() string {
	return c.appId
}

func (c *BaseConnection) UID() string {
	return c.uid
}

func (c *BaseConnection) Group() []string {
	return c.group
}

func (c *BaseConnection) Fd() string {
	return c.fd
}

func (c *BaseConnection) Addr() string {
	return c.addr
}

func (c *BaseConnection) Close() {
}

func (c *BaseConnection) Write(message []byte) {
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
func (c *BaseConnection) JoinGroup(groupId string) (result bool) {
	clientManager.joinGroup(c, groupId)
	// 用户登录了
	if c.uid != "" {
		result = true
		return
	}
	return
}
