package server

import (
	"context"
	"fmt"
	"github.com/nbcx/dsl/distributed/protobuf"
	"github.com/nbcx/dsl/server"
)

type Broadcast struct {
}

func (s *Broadcast) Fd(c context.Context, req *protobuf.BroadcastFdReq) (rsp *protobuf.BroadcastRsp, err error) {
	fmt.Println("Fd:", req.Message, ":", req.Seq)
	server.GetManager().Send(req.Fd, req.Message)
	rsp = &protobuf.BroadcastRsp{
		Code:    200,
		Seq:     req.Seq,
		Message: "success",
	}
	return
}

func (s *Broadcast) Uid(c context.Context, req *protobuf.BroadcastUidReq) (rsp *protobuf.BroadcastRsp, err error) {
	fmt.Println("Uid:", req.Message, ":", req.Seq)
	connection := server.GetManager().FindWithUid(req.Aid, req.Uid)
	connection.Write(req.Message)
	return
}

func (s *Broadcast) User(c context.Context, req *protobuf.BroadcastUidReq) (rsp *protobuf.BroadcastRsp, err error) {
	fmt.Println("Uid:", req.Message, ":", req.Seq)
	server.GetManager().SendWithUser(req.Aid, req.Message)
	return
}

func (s *Broadcast) Group(c context.Context, req *protobuf.BroadcastGroupReq) (rsp *protobuf.BroadcastRsp, err error) {
	fmt.Println("Group:", req.Message, ":", req.Seq)
	server.GetManager().SendWithGroup(req.Aid, req.Gid, req.Message)
	return
}

func (s *Broadcast) App(c context.Context, req *protobuf.BroadcastAppReq) (rsp *protobuf.BroadcastRsp, err error) {
	fmt.Println("App:", req.Message, ":", req.Seq)
	server.GetManager().SendWithApp(req.Aid, req.Message)
	return
}

func (s *Broadcast) All(c context.Context, req *protobuf.BroadcastReq) (rsp *protobuf.BroadcastRsp, err error) {
	fmt.Println("All:", req.Message, ":", req.Seq)
	server.GetManager().SendAll(req.Message)
	rsp = &protobuf.BroadcastRsp{
		Code:    200,
		Seq:     req.Seq,
		Message: "success",
	}
	return
}
