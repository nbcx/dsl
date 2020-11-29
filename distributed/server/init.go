package server

import (
	"github.com/nbcx/gcs/distributed/protobuf"
	"github.com/nbcx/gcs/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
}

// rpc Server
// link::https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
func Start() {

	rpcPort := viper.GetString("distributed.port")
	serverIp := util.LocalIp
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
