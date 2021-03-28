package server

import (
	"github.com/nbcx/dsl/util"
)

func IsLocal(server *util.Server) (isLocal bool) {
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
	return manager
}
