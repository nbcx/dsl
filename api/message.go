package api

import (
	"fmt"
	"github.com/nbcx/gcs"
	gserver "github.com/nbcx/gcs/server"
)

// 给全体用户发消息
func SendUserMessageAll(appId, userId string, msgId, cmd, message string) (sendResults bool, err error) {
	sendResults = true

	servers, err := component.GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}

	for _, server := range servers {
		if gcs.IsLocal(server) {
			gcs.Manager.SendWithUser(appId, []byte(message))
		} else {
			remote.SendMsgAll(server, msgId, userId, cmd, message)
		}
	}

	return
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
