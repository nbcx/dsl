package server

import (
	"context"
	"github.com/nbcx/gcs/distributed/protobuf"
	"github.com/nbcx/gcs/server"
)

type Group struct {
}

// 处理用户登陆
func (s *Group) Join(cox context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {
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

func (s *Group) Quit(ctx context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {
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
func (s *Group) Add(c context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {

	return
}
func (s *Group) Del(c context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {

	return
}
