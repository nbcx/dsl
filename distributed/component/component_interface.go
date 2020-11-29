package component

import "gosh/model"

type IComponent interface {
	Start()
	GetAllServer() (servers []*model.Server, err error)
	Register(server *model.Server) (err error)
}
