package server

import (
	"context"
	"fmt"
	"github.com/nbcx/gcs/distributed/protobuf"
	"github.com/nbcx/gcs/server"
)

type Connection struct {
}

// 使连接断开
func (s *Connection) Offline(ctx context.Context, req *protobuf.ConnectionReq) (rsp *protobuf.ConnectionRsp, err error) {

	fmt.Println("grpc_request 查询用户是否在线", req.String())

	rsp = &protobuf.ConnectionRsp{}
	fd := req.Fd

	c := server.GetManager().Find(fd)
	if c == nil {
		rsp.Code = 500
		return
	}
	c.Close()
	return rsp, nil
}

// 获取连接信息
func (s *Connection) Info(ctx context.Context, req *protobuf.ConnectionReq) (rsp *protobuf.ConnectionRsp, err error) {

	fmt.Println("grpc_request 查询用户是否在线", req.String())

	rsp = &protobuf.ConnectionRsp{}

	//online := client.CheckUserOnline(req.GetUserId())

	//setErr(req, base.OK, "")
	//rsp.Online = online

	return rsp, nil
}
