package server

import (
	"context"
	"github.com/nbcx/gcs/distributed/protobuf"
)

type Broadcast struct {
}

func (s *Broadcast) Fd(c context.Context, req *protobuf.BroadcastReq) (rsp *protobuf.BroadcastRsp, err error) {

	return
}

func (s *Broadcast) Uid(c context.Context, req *protobuf.BroadcastReq) (rsp *protobuf.BroadcastRsp, err error) {

	return
}

func (s *Broadcast) User(c context.Context, req *protobuf.BroadcastReq) (rsp *protobuf.BroadcastRsp, err error) {

	return
}

func (s *Broadcast) Group(c context.Context, req *protobuf.BroadcastReq) (rsp *protobuf.BroadcastRsp, err error) {

	return
}

func (s *Broadcast) App(c context.Context, req *protobuf.BroadcastReq) (rsp *protobuf.BroadcastRsp, err error) {

	return
}

func (s *Broadcast) All(c context.Context, req *protobuf.BroadcastReq) (rsp *protobuf.BroadcastRsp, err error) {
	return
}
