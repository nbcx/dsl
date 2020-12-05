package client

import (
	"context"
	"fmt"
	"github.com/nbcx/gcs/distributed/protobuf"
	"github.com/nbcx/gcs/util"
	"google.golang.org/grpc"
	"time"
)

func ConnectionOffine(server *util.Server, fd string) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewConnectionServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.ConnectionReq{
		Fd:  fd,
		Uid: "",
	}
	rsp, err := c.Offline(ctx, &req)
	if err != nil {
		fmt.Println("发送消息", err)
		return
	}

	code := rsp.GetCode()
	fmt.Println("发送消息 成功:", code)

	return
}
