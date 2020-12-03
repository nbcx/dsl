package api

import (
	"fmt"
	"github.com/nbcx/gcs"
	"github.com/nbcx/gcs/distributed/client"
	log "github.com/sirupsen/logrus"
)

// 向指定连接ID发送数据
func BroadcastFd(fd, message string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		c.Write(message)
		return
	}
	server, isLocal, _ := gcs.GetServerAndIsLocal(fd)
	if isLocal {
		return
	}
	client.BroadcastFd(server, fd, message)
}

func BroadcastUid(appId, uid, message string) {
	c := gcs.Manager.FindWithUid(appId, uid)
	if c != nil {
		c.Write(message)
		return
	}

	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}

	for _, server := range servers {
		if gcs.IsLocal(server) {
			continue
		}
		client.BroadcastUid(server, appId, uid, message)
	}
}

func BroadcastGroup() {

}

func BroadcastUser() {

}

// 向指定应用所有连接发送消息
func BroadcastApp(appId, userId string, msgId, cmd, message string) (sendResults bool, err error) {
	sendResults = true

	servers, err := gcs.GetComponent().GetAllServer()
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

func BroadcastAll() {
	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}
	for _, server := range servers {
		if gcs.IsLocal(server) {
			log.Info("local server")
		} else {
			log.Info("remote server")
			remote.Send(server, "msgId", "userId", "cmd", "message", "ddd")
		}
	}
}
