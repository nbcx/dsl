package api

import (
	"github.com/nbcx/gcs"
	gserver "github.com/nbcx/gcs/server"
)

func Login(fd string, userId string) {

	//如果是本机
	gserver.Login(fd, userId)
	//如果非本机
	//Remote.Login()
}

// 全体广播
func Broadcast(msg []string) {

}

// 对应用全体广播
func BroadcastApp(appId string, msg []byte) {
	gserver.GetManager().SendWithApp(appId, msg)
	//Remote.SendMsgAll()
}

// 对组进行广播
func BroadcastGroup(appId, gid string, msg []byte) {

}

// 对登陆用户进行广播
func BroadcastLogin(appId string, msg []byte) {

}

// 对用户进行广播
func BroadcastUid(appId, uid string, msg []byte) {
	c := gserver.GetManager().FindWithUid(appId, uid)
	if c != nil {
		c.Write(msg)
	} else {
		remote.BroadcastUid(appId, uid, msg)
		//Remote.SendMsg()
	}
}

// 向指定FD广播数据
func BroadcastFd(fd string, msg []byte) {
	c := gserver.GetManager().Find(fd)
	if c != nil {
		c.Write(msg)
	} else {
		s, isLocal, _ := gcs.GetServerAndIsLocal(fd)
		if isLocal {
			return
		}
		remote.BroadcastFd(s, fd, msg)
	}
}
