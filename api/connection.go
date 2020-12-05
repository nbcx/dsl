package api

import (
	"github.com/nbcx/gcs"
	"github.com/nbcx/gcs/distributed/client"
)

func ConnectionOffine(fd string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		c.Close()
		return
	}

	server, isLocal, _ := gcs.GetServerAndIsLocal(fd)
	if isLocal {
		return
	}
	client.ConnectionOffine(server, fd)
	return
}
