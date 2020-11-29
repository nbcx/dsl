package gcs

import (
	"github.com/go-redis/redis"
	"github.com/nbcx/gcs/distributed/component"
	"github.com/nbcx/gcs/distributed/protobuf"
	"github.com/nbcx/gcs/distributed/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

func newRedisComponent(h H) *component.RedisComponent {
	r := h.Get("redis")
	if r == nil {
		panic("redis parameter missing or wrong")
	}
	return &component.RedisComponent{
		HashKey:    h.GetString("hashKey", "gosh:hash:servers"),
		ExpireTime: h.GetUint64("expireTime", 2*60*60),
		Timeout:    h.GetUint64("timeout", 3*60),
		Store:      r.(*redis.Client),
	}
}

// 注入服务发现组件
func (c *cluster) Use(i component.IComponent) {
	c.component = i
}

// 使用内置的服务发现组件
func (c *cluster) UseBuiltIn(name string, option H) {
	switch name {
	case "redis":
		c.useRedis(option)
	case "etcd":
		c.useEtcd(option)
	default:
		panic("the corresponding built-in component could not be found")
	}
}

func (c *cluster) useRedis(option H) {
	c.component = newRedisComponent(option)
}

func (c *cluster) useEtcd(option H) {
	c.component = newRedisComponent(option)
}

func (c *cluster) getComponent() (i component.IComponent) {
	return c.component
}

// 开始集群服务
func (c *cluster) Start() (err error) {
	ser := AddrToServer(c.addr)

	c.component.Register(ser)
	c.component.Start()

	iComponent = c.component

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
