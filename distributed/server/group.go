package server

import (
	"context"
	"github.com/nbcx/dsl/distributed/protobuf"
	"github.com/nbcx/dsl/server"
)

type Group struct {
}

// 将用户加入指定分组
func (s *Group) Join(cox context.Context, req *protobuf.GroupFdReq) (rsp *protobuf.GroupRsp, err error) {
	gid := req.Gid
	fd := req.Fd
	c := server.GetManager().Find(fd)
	if c == nil {
		rsp.Code = 500
		return
	}
	c.JoinGroup(gid)
	return
}

// 将用户从指定分组移除
func (s *Group) Quit(ctx context.Context, req *protobuf.GroupFdReq) (rsp *protobuf.GroupRsp, err error) {
	gid := req.Gid
	fd := req.Fd
	c := server.GetManager().Find(fd)
	if c == nil {
		rsp.Code = 500
		return
	}
	c.ExitGroup(gid)
	return
}

func (s *Group) Del(ctx context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {
	aid := req.Aid
	gid := req.Gid
	c := server.GetManager().DelGroup(aid, gid)
	if c == false {
		rsp.Code = 500
		return
	}
	return
}
