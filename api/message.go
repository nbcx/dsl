package api

import (
	"fmt"
	"github.com/nbcx/gcs"
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
