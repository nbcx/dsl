package server

import (
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gosh/base"
	"gosh/distributed/protobuf"
	"net"
)

type Server struct {
}

func setErr(rsp proto.Message, code uint32, message string) {

	message = base.GetErrorMessage(code, message)
	switch v := rsp.(type) {
	case *protobuf.QueryUsersOnlineRsp:
		v.RetCode = code
		v.ErrMsg = message
	case *protobuf.SendMsgRsp:
		v.RetCode = code
		v.ErrMsg = message
	case *protobuf.SendMsgAllRsp:
		v.RetCode = code
		v.ErrMsg = message
	case *protobuf.GetUserListRsp:
		v.RetCode = code
		v.ErrMsg = message
	default:

	}

}

// rpc Server
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
func Start() {

	rpcPort := viper.GetString("distributed.port")
	serverIp := base.LocalIp
	log.Infof("rpc Server startup in %s:%s", serverIp, rpcPort)

	lis, err := net.Listen("tcp", ":"+rpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterAccServerServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
