package server

import (
	"gosh/base"
)

var (
	clientManager = NewClientManager() // 管理者
	serverIp      = base.LocalIp
	serverPort    string
)
