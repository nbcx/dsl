package util

import (
	"errors"
	"fmt"
	"strings"
)

type Server struct {
	Ip   string `json:"ip"`   // ip
	Port string `json:"port"` // 端口
}

func NewServer(ip string, port string) *Server {
	return &Server{Ip: ip, Port: port}
}

func AddrToServer(addr string) (server *Server) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		panic("addr parameter wrong")
	}
	ip := list[0]
	if len(ip) < 1 {
		ip = LocalIp
	}
	server = &Server{
		Ip:   ip,
		Port: list[1],
	}
	return
}

func (s *Server) String() (str string) {
	if s == nil {
		return
	}

	str = fmt.Sprintf("%s:%s", s.Ip, s.Port)

	return
}

func (s *Server) IsLocal() bool {
	if s.Ip != LocalIp {
		return false
	}

	return true
}

func StringToServer(str string) (server *Server, err error) {
	list := strings.Split(str, ":")
	if len(list) != 2 {
		return nil, errors.New("err")
	}

	server = &Server{
		Ip:   list[0],
		Port: list[1],
	}
	return
}
