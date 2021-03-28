package gcs

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/nbcx/dsl/distributed/component"
	"github.com/nbcx/dsl/distributed/protobuf"
	"github.com/nbcx/dsl/distributed/server"
	"github.com/nbcx/dsl/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type cluster struct {
	ip          string
	port        int
	component   component.IComponent
	supplyServe bool
}

var instance *cluster

func NewCluster(ip string, port int) *cluster {
	if instance == nil {
		instance = &cluster{
			ip:          ip,
			port:        port,
			supplyServe: true,
		} // <--- NOT THREAD SAFE
	}
	return instance
}

func GetCluster() *cluster {
	return instance
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

// 是否在集群内部提供服务
func (c *cluster) SupplyServe(enable bool) {
	c.supplyServe = enable
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
	ser := util.NewServer(c.ip, c.port)

	if c.supplyServe == true {
		c.component.Register(ser)
	}

	c.component.Start()
	iComponent = c.component

	local[ser.Port] = "wss"

	log.Infof("rpc server startup in %s:%d", ser.Ip, ser.Port)

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ser.Ip, ser.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterBroadcastServer(s, &server.Broadcast{})
	protobuf.RegisterGroupServer(s, &server.Group{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return
}
