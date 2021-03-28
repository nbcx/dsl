package server

type IConnection interface {
	GetAddr() string
	GetAppId() string
	GetFd() string
	GetUid() string
	GetLoginTime() uint64

	SetLoginTime(time uint64)
	SetAddr(string)
	SetAppId(string)
	SetFd(string)
	SetUid(string)

	JoinGroup(gids ...string) (err error) //将链接加入指定分组
	ExitGroup(gids ...string) (err error) //将链接从指定分组移除
	GetGroup() []string                   //获取链接已经加入的分组

	IsHeartbeatTimeout(currentTime uint64) (timeout bool)
	Heartbeat(currentTime uint64)
	Close() // 关闭链接

	Write(message string)     //向链接写入字符串
	WriteByte(message []byte) //向链接写入字字节

}
