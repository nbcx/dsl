package server

import (
	"github.com/nbcx/gcs/util"
)

var (
	clientManager = NewClientManager() // 管理者
	serverIp      = util.LocalIp
	serverPort    string
)
