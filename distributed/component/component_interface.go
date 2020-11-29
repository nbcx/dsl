package component

import "github.com/nbcx/gcs/model"

type IComponent interface {
	Start()
	GetAllServer() (servers []*model.Server, err error)
	Register(server *model.Server) (err error)
}
