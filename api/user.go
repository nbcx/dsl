package api

import (
	"github.com/nbcx/dsl"
	"github.com/nbcx/dsl/server"
)

func Login(fd string, userId string) {

	//如果是本机
	server.Login(fd, userId)
	//如果非本机
	//Remote.Login()
}

func Logout(fd string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		gcs.Manager.Logout(c)
		return
	}
}
