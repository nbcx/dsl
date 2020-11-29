package gcs

import (
	"errors"
	"fmt"
	"github.com/nbcx/gcs/distributed/component"
	"github.com/nbcx/gcs/model"
	"github.com/nbcx/gcs/util"
	"strings"
	"time"
)

func GetOrderIdTime() (orderId string) {

	currentTime := time.Now().Nanosecond()
	orderId = fmt.Sprintf("%d", currentTime)

	return
}

//对称加密IP和端口，当做clientId
func GenClientId() string {
	raw := []byte(util.LocalIp + ":" + viper.GetString("server.port"))
	str, err := util.Encrypt(raw, []byte(getSecret()))
	if err != nil {
		panic(err)
	}

	return str
}

//获取client key地址信息
func GetAddrInfoAndIsLocal(clientId string) (addr string, host string, port string, isLocal bool, err error) {
	//解密ClientId
	addr, err = util.Decrypt(clientId, []byte(getSecret()))
	if err != nil {
		return
	}

	host, port, err = ParseRedisAddrValue(addr)
	if err != nil {
		return
	}

	isLocal = IsAddrLocal(host, port)
	return
}

func GetServerAndIsLocal(fd string) (server *model.Server, isLocal bool, err error) {
	//解密ClientId
	addr, err := util.Decrypt(fd, []byte(getSecret()))
	if err != nil {
		return
	}

	host, port, err := ParseRedisAddrValue(addr)
	if err != nil {
		return
	}

	isLocal = IsAddrLocal(host, port)
	server = model.NewServer(host, port)
	return
}

//解析redis的地址格式
func ParseRedisAddrValue(redisValue string) (host string, port string, err error) {
	if redisValue == "" {
		err = errors.New("解析地址错误")
		return
	}
	addr := strings.Split(redisValue, ":")
	if len(addr) != 2 {
		err = errors.New("解析地址错误")
		return
	}
	host, port = addr[0], addr[1]

	return
}

//判断地址是否为本机
func IsAddrLocal(host string, port string) bool {
	return host == util.LocalIp && port == viper.GetString("server.port")
}

func AddrToServer(addr string) (server *model.Server) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		panic("addr parameter wrong")
	}
	ip := list[0]
	if len(ip) < 1 {
		ip = util.LocalIp
	}
	server = &model.Server{
		Ip:   ip,
		Port: list[1],
	}
	return
}

func GetComponent() component.IComponent {
	return iComponent
}
