package api

import gserver "github.com/nbcx/gcs/server"

func Login(fd string, userId string) {

	//如果是本机
	gserver.Login(fd, userId)
	//如果非本机
	//Remote.Login()
}
