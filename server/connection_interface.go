package server

type IConnection interface {
	Write(message []byte)
	Addr() string
	AppId() string
	Fd() string
	UID() string
	Group() []string
	IsHeartbeatTimeout(currentTime uint64) (timeout bool)
	Heartbeat(currentTime uint64)
	Close()
	JoinGroup(groupId string) (result bool)
}
