package api

import (
	"fmt"
	"github.com/nbcx/gcs"
	"github.com/nbcx/gcs/distributed/client"
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

func BroadcastGroup(appId string, message string, gids ...string) {
	gcs.Manager.SendWithGroup(appId, gids, []byte(message))
	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}
	for _, server := range servers {
		if gcs.IsLocal(server) {
			continue
		}
		client.BroadcastGroup(server, appId, gids, message)
	}
}

func BroadcastUser(appId string, message string) {
	gcs.Manager.SendWithUser(appId, []byte(message))
	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}
	for _, server := range servers {
		if gcs.IsLocal(server) {
			continue
		}
		client.BroadcastUser(server, appId, message)
	}
}

// 向指定应用所有连接发送消息
func BroadcastApp(appId, message string) (sendResults bool, err error) {
	gcs.Manager.SendWithApp(appId, []byte(message))
	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}
	for _, server := range servers {
		if gcs.IsLocal(server) {
			continue
		}
		client.BroadcastApp(server, appId, message)
	}
	return
}

func BroadcastAll(message string) {
	gcs.Manager.Send([]byte(message))
	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("给全体用户发消息", err)
		return
	}
	for _, server := range servers {
		if gcs.IsLocal(server) {
			continue
		}
		client.BroadcastAll(server, message)
	}
	return
}
