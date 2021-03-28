package client

import (
	"context"
	"fmt"
	"github.com/nbcx/dsl/distributed/protobuf"
	"github.com/nbcx/dsl/util"
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

	c := protobuf.NewGroupClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.GroupFdReq{
		Fd:  fd,
		Gid: gid[0],
	}
	rsp, err := c.Join(ctx, &req)
	if err != nil {
		fmt.Println("发送消息", err)
		return
	}

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

	c := protobuf.NewGroupClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.GroupFdReq{
		Fd:  fd,
		Gid: gid[0],
	}
	rsp, err := c.Quit(ctx, &req)
	if err != nil {
		fmt.Println("发送消息", err)
		return
	}

	code := rsp.GetCode()
	fmt.Println("发送消息 成功:", code)

	return

}

func GroupDel(server *util.Server, aid string, gid []string) {
	conn, err := grpc.Dial(server.String(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败", server.String())
		return
	}
	defer conn.Close()

	c := protobuf.NewGroupClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := protobuf.GroupReq{
		Aid: aid,
		Gid: gid[0],
	}
	rsp, err := c.Del(ctx, &req)
	if err != nil {
		fmt.Println("删除分组 失败:", err)
		return
	}

	code := rsp.GetCode()
	fmt.Println("删除分组 成功:", code)

	return

}
