package distributed

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gosh/base"
	"gosh/distributed/component"
	"gosh/distributed/protobuf"
	"gosh/distributed/server"
	"net"
)

type cluster struct {
	addr      string
	component component.IComponent
}

func NewCluster(addr string) *cluster {
	return &cluster{
		addr: addr,
	}
}

// 注入服务发现组件
func (c *cluster) Use(i component.IComponent) {
	c.component = i
}

// 使用内置的服务发现组件
func (c *cluster) UseBuiltIn(name string, option base.H) {
	switch name {
	case "redis":
		c.useRedis(option)
	case "etcd":
		c.useEtcd(option)
	default:
		panic("the corresponding built-in component could not be found")
	}
}

func (c *cluster) useRedis(option base.H) {
	c.component = component.NewRedisComponent(option)
}

func (c *cluster) useEtcd(option base.H) {
	c.component = component.NewRedisComponent(option)
}

func (c *cluster) getComponent() (i component.IComponent) {
	return c.component
}

// 开始集群服务
func (c *cluster) Start() (err error) {
	ser := base.AddrToServer(c.addr)

	c.component.Register(ser)
	c.component.Start()

	Remote.Use(c.component)

	//rpcPort := viper.GetString("distributed.port")
	//serverIp := base.LocalIp
	log.Infof("rpc server startup in %s:%s", ser.Ip, ser.Port)

	listen, err := net.Listen("tcp", ":"+ser.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterAccServerServer(s, &server.Server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return
}
