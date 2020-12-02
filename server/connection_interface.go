package server

type IConnection interface {
	GetAddr() string
	GetAppId() string
	GetFd() string
	GetUid() string
	GetGroup() []string
	GetLoginTime() uint64

	SetLoginTime(time uint64)
	SetAddr(string)
	SetAppId(string)
	SetFd(string)
	SetUid(string)
	SetGroup([]string)

	IsHeartbeatTimeout(currentTime uint64) (timeout bool)
	Heartbeat(currentTime uint64)
	Close()
	JoinGroup(groupId string) (result bool)
	Write(message string)
	WriteByte(message []byte)
}
