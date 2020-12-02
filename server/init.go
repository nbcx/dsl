package server

import (
	"github.com/nbcx/gcs/util"
)

var (
	manager    = newManager() // 管理者
	serverIp   = util.LocalIp
	serverPort string
)
