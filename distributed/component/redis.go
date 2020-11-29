package component

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"gosh/base"
	"gosh/model"
	"runtime/debug"
	"strconv"
	"time"
)

type Service struct {
	server     *model.Server
	hashKey    string // 在redids里存储全部的服务器的key
	expireTime uint64 // hashKey过期时间
	timeout    uint64 // server地址过期时间，超过设置时间，则server地址不可用
	store      *redis.Client
}

func NewRedisComponent(h base.H) *Service {
	r := h.Get("redis")
	if r == nil {
		panic("redis parameter missing or wrong")
	}
	return &Service{
		hashKey:    h.GetString("hashKey", "gosh:hash:servers"),
		expireTime: h.GetUint64("expireTime", 2*60*60),
		timeout:    h.GetUint64("timeout", 3*60),
		store:      r.(*redis.Client),
	}
}

func (s *Service) Start() {
	//注册本机信息到redis
	base.Timer(2*time.Second, 60*time.Second, s.autoRenew, "", s.del, "")
}

// 获取集群所有server地址
func (s *Service) GetAllServer() (servers []*model.Server, err error) {

	currentTime := uint64(time.Now().Unix())
	servers = make([]*model.Server, 0)
	key := s.hashKey

	val, err := s.store.Do("hGetAll", key).Result()

	valByte, _ := json.Marshal(val)
	fmt.Println("GetServerAll", key, string(valByte))

	serverMap, err := s.store.HGetAll(key).Result()
	if err != nil {
		fmt.Println("SetServerInfo", key, err)
		return
	}

	for key, value := range serverMap {
		valueUint64, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			fmt.Println("GetServerAll", key, err)

			return nil, err
		}
		// 超时
		if valueUint64+s.timeout <= currentTime {
			continue
		}
		server, err := model.StringToServer(key)
		if err != nil {
			fmt.Println("GetServerAll", key, err)

			return nil, err
		}
		servers = append(servers, server)
	}

	return
}

// 设置服务器信息
func (s *Service) Register(server *model.Server) (err error) {
	s.server = server
	return s.set(server)
}

// 更新服务注册信息
func (s *Service) autoRenew(param interface{}) (result bool) {
	result = true

	defer func() {
		if r := recover(); r != nil {
			log.Error("服务注册 stop", r, param, string(debug.Stack()))
		}
	}()
	//log.Info("定时任务，服务注册", param, server, currentTime)
	s.set(s.server)
	return
}

// 将服务信息保存到redis
func (s *Service) set(server *model.Server) (err error) {
	currentTime := uint64(time.Now().Unix())

	value := fmt.Sprintf("%d", currentTime)

	number, err := s.store.Do("hSet", s.hashKey, server.String(), value).Int()
	if err != nil {
		fmt.Println("SetServerInfo", s.hashKey, number, err)

		return
	}

	if number != 1 {

		return
	}

	s.store.Do("Expire", s.hashKey, s.expireTime)

	return
}

// 服务下线
func (s *Service) del(param interface{}) (result bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Info("服务下线 stop", r, string(debug.Stack()))
		}
	}()
	number, err := s.store.Do("hDel", s.hashKey, s.server.String()).Int()
	if err != nil {
		fmt.Println("DelServerInfo", s.hashKey, number, err)
		return
	}

	if number != 1 {
		return
	}

	s.store.Do("Expire", s.hashKey, s.expireTime)
	result = true
	return
}
