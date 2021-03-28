package api

import (
	"fmt"
	"github.com/nbcx/dsl"
	"github.com/nbcx/dsl/distributed/client"
	log "github.com/sirupsen/logrus"
)

// 向指定连接ID发送数据
func BroadcastFd(fd, message string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		fmt.Println("isLocal Write")
		c.Write(message)
		return
	}
	server, isLocal, err := gcs.GetServerAndIsLocal(fd)
	if err != nil {
		fmt.Println("BroadcastFd  GetServerAndIsLocal:", err)
		return
	}
	if isLocal {
		return
	}
	client.BroadcastFd(server, fd, message)
}

func BroadcastUid(aid, uid, message string) {
	c := gcs.Manager.FindWithUid(aid, uid)
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
		client.BroadcastUid(server, aid, uid, message)
	}
}

func BroadcastGroup(aid, gid, message string) {
	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}
	for _, server := range servers {
		if gcs.IsLocal(server) {
			log.Info("local server")
			for _, c := range gcs.Manager.Connections {
				c.Write(message)
			}
		} else {
			log.Info("remote server:", " ", server.Ip, ":", server.Port)
			client.BroadcastGroup(server, aid, gid, message)
		}
	}
}

func BroadcastUser() {

}

// 向指定应用所有连接发送消息
func BroadcastApp(appId, message string) (sendResults bool, err error) {
	sendResults = true

	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}

	for _, server := range servers {
		if gcs.IsLocal(server) {
			gcs.Manager.SendWithUser(appId, message)
		} else {
			// remote.SendMsgAll(server, msgId, userId, cmd, message)
		}
	}

	return
}

// 向所有链接广播消息
func BroadcastAll(message string) {
	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}
	for _, server := range servers {
		if gcs.IsLocal(server) {
			log.Info("local server")
			for _, c := range gcs.Manager.Connections {
				c.Write(message)
			}
		} else {
			log.Info("remote server:", " ", server.Ip, ":", server.Port)
			client.BroadcastAll(server, message)
			//remote.Send(server, "msgId", "userId", "cmd", "message", "ddd")
		}
	}
}
