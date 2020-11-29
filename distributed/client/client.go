package client

import (
	"context"
	"fmt"
	"github.com/nbcx/gcs/distributed/component"
	"github.com/nbcx/gcs/distributed/protobuf"
	"github.com/nbcx/gcs/model"
	"github.com/nbcx/gcs/util"
	"google.golang.org/grpc"
	"time"
)

type Remote struct {
	i component.IComponent
}

func (r *Remote) GetComponent() (i component.IComponent) {
	return r.i
}

// 给全体用户发消息
func (r *Remote) SendAll() (err error) {
	servers, err := r.i.GetAllServer()

	if err != nil {
		return
	}
	for _, server := range servers {
		if r.IsLocal(server) {
			continue
		}
		//r.Send(server)
	}

	return
}

func (r *Remote) IsLocal(server *model.Server) (isLocal bool) {
	if server.Ip == util.LocalIp { //&& server.Port == serverPort
		isLocal = true
	}

	return
}

// Remote Remote
// 发送消息
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func (s *Remote) Send(server *model.Server, seq string, userId string, cmd string, msgType string, message string) (sendMsgId string, err error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.SendMsgReq{
		Seq:     seq,
		UserId:  userId,
		Cms:     cmd,
		Type:    msgType,
		Msg:     message,
		IsLocal: false,
	}
	rsp, err := c.SendMsg(ctx, &req)
	if err != nil {
		fmt.Println("发送消息", err)

		return
	}

	//if rsp.GetRetCode() != base.OK {
	//	fmt.Println("发送消息", rsp.String())
	//	err = errors.New(fmt.Sprintf("发送消息失败 code:%d", rsp.GetRetCode()))
	//
	//	return
	//}

	sendMsgId = rsp.GetSendMsgId()
	fmt.Println("发送消息 成功:", sendMsgId)

	return
}

// Remote Remote
// 给全体用户发送消息
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func (s *Remote) SendMsgAll(server *model.Server, seq string, userId string, cmd string, message string) (sendMsgId string, err error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.SendMsgAllReq{
		Seq:    seq,
		UserId: userId,
		Cms:    cmd,
		Msg:    message,
	}
	rsp, err := c.SendMsgAll(ctx, &req)
	if err != nil {
		fmt.Println("给全体用户发送消息", err)

		return
	}

	//if rsp.GetRetCode() != base.OK {
	//	fmt.Println("给全体用户发送消息", rsp.String())
	//	err = errors.New(fmt.Sprintf("发送消息失败 code:%d", rsp.GetRetCode()))
	//
	//	return
	//}

	sendMsgId = rsp.GetSendMsgId()
	fmt.Println("给全体用户发送消息 成功:", sendMsgId)

	return
}

// 获取用户列表
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func (s *Remote) GetUserList(server *model.Server, appId uint32) (userIds []string, err error) {
	userIds = make([]string, 0)

	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())

		return
	}
	defer conn.Close()

	c := protobuf.NewAccServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.GetUserListReq{
		AppId: appId,
	}
	rsp, err := c.GetUserList(ctx, &req)
	if err != nil {
		fmt.Println("获取用户列表 发送请求错误:", err)

		return
	}

	//if rsp.GetRetCode() != base.OK {
	//	fmt.Println("获取用户列表 返回码错误:", rsp.String())
	//	err = errors.New(fmt.Sprintf("发送消息失败 code:%d", rsp.GetRetCode()))
	//
	//	return
	//}

	userIds = rsp.GetUserId()
	fmt.Println("获取用户列表 成功:", userIds)

	return
}
