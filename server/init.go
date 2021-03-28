package server

import (
	"github.com/nbcx/dsl/util"
)

var (
	manager    = newManager() // 管理者
	serverIp   = util.LocalIp
	serverPort int
)
