package api

import (
	"github.com/nbcx/gcs"
	"github.com/nbcx/gcs/distributed/client"
)

func GroupJoin(fd string, gids ...string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		c.JoinGroup(gids...)
		return
	}

	server, isLocal, _ := gcs.GetServerAndIsLocal(fd)
	if isLocal {
		return
	}
	client.GroupJoin(server, fd, gids...)
}

func GroupQuit(fd string, gids ...string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		c.ExitGroup(gids...)
		return
	}
	server, isLocal, _ := gcs.GetServerAndIsLocal(fd)
	if isLocal {
		return
	}
	client.GroupQuit(server, fd, gids...)
}
