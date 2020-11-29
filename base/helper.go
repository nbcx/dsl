package base

import (
	"errors"
	"fmt"
	"github.com/nbcx/gcs/model"
	"github.com/spf13/viper"
	"net"
	"strings"
	"time"
)

// 获取服务器Ip
func getServerIp() (ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
			}
		}
	}

	return
}

func GetOrderIdTime() (orderId string) {

	currentTime := time.Now().Nanosecond()
	orderId = fmt.Sprintf("%d", currentTime)

	return
}

//对称加密IP和端口，当做clientId
func GenClientId() string {
	raw := []byte(LocalIp + ":" + viper.GetString("server.port"))
	str, err := Encrypt(raw, []byte(viper.GetString("server.secret")))
	if err != nil {
		panic(err)
	}

	return str
}

//获取client key地址信息
func GetAddrInfoAndIsLocal(clientId string) (addr string, host string, port string, isLocal bool, err error) {
	//解密ClientId
	addr, err = Decrypt(clientId, []byte(viper.GetString("server.secret")))
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
	addr, err := Decrypt(fd, []byte(viper.GetString("server.secret")))
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
	return host == LocalIp && port == viper.GetString("server.port")
}

func LocalServer() (server *model.Server) {
	server = &model.Server{
		Ip:   LocalIp,
		Port: viper.GetString("server.port"),
	}
	return
}

func AddrToServer(addr string) (server *model.Server) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		panic("addr parameter wrong")
	}
	ip := list[0]
	if len(ip) < 1 {
		ip = LocalIp
	}
	server = &model.Server{
		Ip:   ip,
		Port: list[1],
	}
	return
}
