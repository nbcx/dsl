package server

import (
	"context"
	"github.com/nbcx/gcs/distributed/protobuf"
)

type Group struct {
}

// 处理用户登陆
func (s *Group) Join(c context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {

	//websocket.LoginController()
	return
}

func (s *Group) Quit(c context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {

	//websocket.LoginController()
	return
}
func (s *Group) Add(c context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {

	//websocket.LoginController()
	return
}
func (s *Group) Del(c context.Context, req *protobuf.GroupReq) (rsp *protobuf.GroupRsp, err error) {

	//websocket.LoginController()
	return
}
