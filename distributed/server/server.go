package server

import (
	"context"
	"fmt"
	"github.com/nbcx/gcs/distributed/protobuf"
	log "github.com/sirupsen/logrus"
)

// 处理用户登陆
func (s *Server) Login(c context.Context, req *protobuf.LoginReq) (rsp *protobuf.LoginRsp, err error) {

	rsp = &protobuf.LoginRsp{}
	//websocket.LoginController()
	return
}

// 查询用户是否在线
func (s *Server) QueryUsersOnline(c context.Context, req *protobuf.QueryUsersOnlineReq) (rsp *protobuf.QueryUsersOnlineRsp, err error) {

	fmt.Println("grpc_request 查询用户是否在线", req.String())

	rsp = &protobuf.QueryUsersOnlineRsp{}

	//online := client.CheckUserOnline(req.GetUserId())

	//setErr(req, base.OK, "")
	//rsp.Online = online

	return rsp, nil
}

// 给本机用户发消息
func (s *Server) SendMsg(c context.Context, req *protobuf.SendMsgReq) (rsp *protobuf.SendMsgRsp, err error) {

	log.Info("hello s server:" + req.Cms)
	//fmt.Println("grpc_request 给本机用户发消息", req.String())
	rsp = &protobuf.SendMsgRsp{}
	//if req.GetIsLocal() {
	//
	//	// 不支持
	//	setErr(rsp, base.ParameterIllegal, "")
	//
	//	return
	//}
	//
	//data := model.GetMsgData(req.GetUserId(), req.GetSeq(), req.GetCms(), req.GetMsg())
	//sendResults, err := distributed.SendUserMessageLocal(req.GetAppId(), req.GetUserId(), data)
	//if err != nil {
	//	fmt.Println("系统错误", err)
	//	setErr(rsp, base.ServerError, "")
	//
	//	return rsp, nil
	//}
	//
	//if !sendResults {
	//	fmt.Println("发送失败", err)
	//	setErr(rsp, base.OperationFailure, "")
	//
	//	return rsp, nil
	//}
	//
	rsp.RetCode = 200
	rsp.SendMsgId = "ddddd"
	//
	//fmt.Println("grpc_response 给本机用户发消息", rsp.String())
	return
}

// 给本机全体用户发消息
func (s *Server) SendMsgAll(c context.Context, req *protobuf.SendMsgAllReq) (rsp *protobuf.SendMsgAllRsp, err error) {

	//fmt.Println("grpc_request 给本机全体用户发消息", req.String())
	//
	//rsp = &protobuf.SendMsgAllRsp{}
	//
	//data := model.GetMsgData(req.GetUserId(), req.GetSeq(), req.GetCms(), req.GetMsg())
	//gserver.AllSendMessages(req.GetUserId(), data)
	//
	//setErr(rsp, base.OK, "")
	//
	//fmt.Println("grpc_response 给本机全体用户发消息:", rsp.String())

	return
}

// 获取本机用户列表
func (s *Server) GetUserList(c context.Context, req *protobuf.GetUserListReq) (rsp *protobuf.GetUserListRsp, err error) {

	//fmt.Println("grpc_request 获取本机用户列表", req.String())
	//
	//appId := req.GetAppId()
	//rsp = &protobuf.GetUserListRsp{}
	//
	//// 本机
	//userList := gserver.GetUserList(appId)
	//
	//setErr(rsp, base.OK, "")
	//rsp.UserId = userList
	//
	//fmt.Println("grpc_response 获取本机用户列表:", rsp.String())

	return
}
