package client

import (
	"context"
	"fmt"
	"github.com/nbcx/gcs/distributed/protobuf"
	"github.com/nbcx/gcs/util"
	"google.golang.org/grpc"
	"time"
)

func BroadcastFd(server *util.Server, fd, message string) (err error) {

	return
}

func BroadcastUid(server *util.Server, appId, uid, message string) (err error) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewBroadcastServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.BroadcastReq{}

	rsp, err := c.Uid(ctx, &req)
	if err != nil {
		fmt.Println("发送消息", err)
		return
	}

	code := rsp.GetRetCode()
	fmt.Println("发送消息 成功:", code)
	return
}

func BroadcastUser(server *util.Server, appId string, message string) (err error) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewBroadcastServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.BroadcastReq{}

	rsp, err := c.User(ctx, &req)
	if err != nil {
		fmt.Println("发送消息", err)
		return
	}

	code := rsp.GetRetCode()
	fmt.Println("发送消息 成功:", code)
	return
}

func BroadcastGroup(server *util.Server, appId string, gid []string, message string) (err error) {

	return
}

func BroadcastApp(server *util.Server, appId string, message string) (err error) {

	return
}

func BroadcastAll(server *util.Server, message string) (err error) {
	return
}
