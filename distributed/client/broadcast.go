package client

import (
	"context"
	"fmt"
	"github.com/nbcx/dsl/distributed/protobuf"
	"github.com/nbcx/dsl/util"
	"google.golang.org/grpc"
	"time"
)

func BroadcastFd(server *util.Server, fd, message string) (err error) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewBroadcastClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.BroadcastFdReq{
		Fd:      fd,
		Aid:     "",
		Message: message,
	}

	rsp, err := c.Fd(ctx, &req)
	if err != nil {
		fmt.Println("发送消息 失败", server.String(), err)
		return
	}

	code := rsp.GetCode()
	fmt.Println("发送消息 成功:", code)
	return
}

func BroadcastUid(server *util.Server, appId, uid, message string) (err error) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewBroadcastClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.BroadcastUidReq{}

	rsp, err := c.Uid(ctx, &req)
	if err != nil {
		fmt.Println("发送消息", err)
		return
	}

	code := rsp.GetCode()
	fmt.Println("发送消息 成功:", code)
	return
}

func BroadcastUser(server util.Server, req *protobuf.BroadcastReq) (err error) {

	return
}

func BroadcastGroup(server *util.Server, aid, gid, message string) (err error) {

	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewBroadcastClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.BroadcastGroupReq{
		Gid:     gid,
		Aid:     "",
		Message: message,
	}
	rsp, err := c.Group(ctx, &req)
	if err != nil {
		fmt.Println("发送消息 失败", server.String(), err)
		return
	}

	code := rsp.GetCode()
	fmt.Println("发送消息 成功:", code)
	return
}

func BroadcastApp(server util.Server, req *protobuf.BroadcastReq) (err error) {

	return
}

func BroadcastAll(server *util.Server, message string) (err error) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewBroadcastClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.BroadcastReq{
		Seq:     "sssss",
		Message: message,
	}

	rsp, err := c.All(ctx, &req)
	if err != nil {
		fmt.Println("发送消息通信 失败:", err)
		return
	}

	code := rsp.GetCode()
	fmt.Println("发送消息 成功:", code, ",Message:", rsp.Message, "Seq:", rsp.Seq)
	return
}
