package client

import (
	"context"
	"fmt"
	"github.com/nbcx/gcs/distributed/protobuf"
	"github.com/nbcx/gcs/util"
	"google.golang.org/grpc"
	"time"
)

func GroupJoin(server *util.Server, fd string, gid ...string) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewGroupServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.GroupReq{
		Fd:  fd,
		Gid: gid[0],
	}
	rsp, err := c.Join(ctx, &req)
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

	code := rsp.GetCode()
	fmt.Println("发送消息 成功:", code)

	return

}

func GroupQuit(server *util.Server, fd string, gid ...string) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewGroupServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.GroupReq{
		Fd:  fd,
		Gid: gid[0],
	}
	rsp, err := c.Quit(ctx, &req)
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

	code := rsp.GetCode()
	fmt.Println("发送消息 成功:", code)

	return

}
