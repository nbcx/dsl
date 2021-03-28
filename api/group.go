package api

import (
	"fmt"
	"github.com/nbcx/dsl"
	"github.com/nbcx/dsl/distributed/client"
)

func GroupJoin(aid, fd string, gid ...string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		fmt.Println("isLocal Write")
		c.JoinGroup(gid...)
		return
	}
	server, isLocal, err := gcs.GetServerAndIsLocal(fd)
	if err != nil {
		fmt.Println("GroupJoin  GetServerAndIsLocal:", err)
		return
	}
	if isLocal {
		return
	}
	client.GroupJoin(server, fd, gid)
}

func GroupQuit(aid, fd string, gid ...string) {
	c := gcs.Manager.Find(fd)
	if c != nil {
		fmt.Println("isLocal Write")
		c.ExitGroup(gid...)
		return
	}
	server, isLocal, err := gcs.GetServerAndIsLocal(fd)
	if err != nil {
		fmt.Println("GroupJoin  GetServerAndIsLocal:", err)
		return
	}
	if isLocal {
		return
	}
	client.GroupQuit(server, fd, gid)
}

func GroupDel(aid string, gid ...string) {
	servers, err := gcs.GetComponent().GetAllServer()
	if err != nil {
		fmt.Println("获取服务集群异常：", err)
		return
	}
	for _, server := range servers {
		if gcs.IsLocal(server) {
			gcs.Manager.DelGroup(aid, gid...)
		} else {
			client.GroupDel(server, aid, gid)
		}
	}
}
