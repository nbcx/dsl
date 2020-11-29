package component

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/nbcx/gcs/model"
	"github.com/nbcx/gcs/util"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
	"strconv"
	"time"
)

type RedisComponent struct {
	server     *model.Server
	HashKey    string // 在redids里存储全部的服务器的key
	ExpireTime uint64 // hashKey过期时间
	Timeout    uint64 // server地址过期时间，超过设置时间，则server地址不可用
	Store      *redis.Client
}

func (s *RedisComponent) Start() {
	//注册本机信息到redis
	util.Timer(2*time.Second, 60*time.Second, s.autoRenew, "", s.del, "")
}

// 获取集群所有server地址
func (s *RedisComponent) GetAllServer() (servers []*model.Server, err error) {

	currentTime := uint64(time.Now().Unix())
	servers = make([]*model.Server, 0)
	key := s.HashKey

	val, err := s.Store.Do("hGetAll", key).Result()

	valByte, _ := json.Marshal(val)
	fmt.Println("GetServerAll", key, string(valByte))

	serverMap, err := s.Store.HGetAll(key).Result()
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
		if valueUint64+s.Timeout <= currentTime {
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
func (s *RedisComponent) Register(server *model.Server) (err error) {
	s.server = server
	return s.set(server)
}

// 更新服务注册信息
func (s *RedisComponent) autoRenew(param interface{}) (result bool) {
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
func (s *RedisComponent) set(server *model.Server) (err error) {
	currentTime := uint64(time.Now().Unix())

	value := fmt.Sprintf("%d", currentTime)

	number, err := s.Store.Do("hSet", s.HashKey, server.String(), value).Int()
	if err != nil {
		fmt.Println("SetServerInfo", s.HashKey, number, err)

		return
	}

	if number != 1 {

		return
	}

	s.Store.Do("Expire", s.HashKey, s.ExpireTime)

	return
}

// 服务下线
func (s *RedisComponent) del(param interface{}) (result bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Info("服务下线 stop", r, string(debug.Stack()))
		}
	}()
	number, err := s.Store.Do("hDel", s.HashKey, s.server.String()).Int()
	if err != nil {
		fmt.Println("DelServerInfo", s.HashKey, number, err)
		return
	}

	if number != 1 {
		return
	}

	s.Store.Do("Expire", s.HashKey, s.ExpireTime)
	result = true
	return
}
