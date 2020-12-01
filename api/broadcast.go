package api

import (
	"fmt"
	"github.com/nbcx/gcs"
	log "github.com/sirupsen/logrus"
)

func Send() {
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

// 给全体用户发消息
func SendUserMessageAll(appId, userId string, msgId, cmd, message string) (sendResults bool, err error) {
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
