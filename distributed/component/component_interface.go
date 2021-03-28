package component

import (
	"github.com/nbcx/dsl/util"
)

type IComponent interface {
	Start()
	GetAllServer() (servers []*util.Server, err error)
	Register(server *util.Server) (err error)
}
