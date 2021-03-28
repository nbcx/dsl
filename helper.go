package gcs

import (
	"errors"
	"fmt"
	"github.com/nbcx/dsl/distributed/component"
	"github.com/nbcx/dsl/util"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

//对称加密IP和端口，当做clientId
func GenClientId(ip string, port int) string {
	raw := util.Str2bytes(fmt.Sprintf("%s:%d", ip, port))
	str, err := util.Encrypt(raw, util.Str2bytes(getSecret()))
	if err != nil {
		panic(err)
	}
	return str
}

func GenFd() string {
	cluster := GetCluster()
	fmt.Println("GenFd:", cluster.ip, ":", cluster.port)
	return GenClientId(cluster.ip, cluster.port)
}

//获取client key地址信息
func GetAddrInfoAndIsLocal(fd string) (addr string, host string, port int, isLocal bool, err error) {
	//解密ClientId
	addr, err = util.Decrypt(fd, util.Str2bytes(getSecret()))
	if err != nil {
		return
	}

	host, port, err = ParseRedisAddrValue(addr)
	if err != nil {
		return
	}

	server := util.NewServer(host, port)
	isLocal = IsLocal(server)
	return
}

func GetServerAndIsLocal(fd string) (server *util.Server, isLocal bool, err error) {
	//解密ClientId
	addr, err := util.Decrypt(fd, util.Str2bytes(getSecret()))
	if err != nil {
		fmt.Println("GetServerAndIsLocal Err:", err)
		return
	}
	host, port, err := ParseRedisAddrValue(addr)
	if err != nil {
		return
	}
	server = util.NewServer(host, port)
	isLocal = IsLocal(server)
	return
}

//解析addr的地址格式
func ParseRedisAddrValue(addr string) (host string, port int, err error) {
	if addr == "" {
		err = errors.New("解析地址错误,不能为空")
		return
	}
	addrs := strings.Split(addr, ":")
	if len(addrs) != 2 {
		err = errors.New(fmt.Sprintf("解析地址错误，格式错误:%s", addr))
		return
	}
	port, _ = strconv.Atoi(addrs[1])
	host = addrs[0]
	return
}

//判断地址是否为本机
func IsLocalWithValue(server *util.Server) (is bool, value string) {
	if server.Ip != util.LocalIp {
		is = false
		return
	}

	if v, ok := local[server.Port]; ok {
		is = true
		value = v
		return
	}
	return
}

func IsLocal(server *util.Server) bool {
	log.Info("IsLocal:", server.Ip, "-", util.LocalIp)
	if server.Ip != util.LocalIp {
		return false
	}

	if _, ok := local[server.Port]; ok {
		return true
	}
	return false
}

func GetComponent() component.IComponent {
	return iComponent
}
