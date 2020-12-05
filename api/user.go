package api

import (
	"github.com/nbcx/gcs"
	"github.com/nbcx/gcs/distributed/client"
)

func UserLogin(fd, uid string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		// todo
		gcs.Manager.Login(c, uid, 12323223232)
		return
	}

	server, isLocal, _ := gcs.GetServerAndIsLocal(fd)
	if isLocal {
		return
	}
	client.UserLogin(server, fd, uid)
}

func Logout(fd string) {

	c := gcs.Manager.Find(fd)
	if c != nil {
		gcs.Manager.Logout(c)
		return
	}

	server, isLocal, _ := gcs.GetServerAndIsLocal(fd)
	if isLocal {
		return
	}
	client.UserLogout(server, fd)

}
