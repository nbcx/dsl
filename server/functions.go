package server

import (
	"github.com/nbcx/gcs/model"
)

func IsLocal(server *model.Server) (isLocal bool) {
	if server.Ip == serverIp && server.Port == serverPort {
		isLocal = true
	}

	return
}

// 用户登录
func Login(fd string, userId string) {

	return
}

func GetManager() *ClientManager {
	return clientManager
}
