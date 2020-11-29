package gcs

import (
	"github.com/nbcx/gcs/distributed/component"
	"github.com/nbcx/gcs/model"
	"github.com/nbcx/gcs/server"
	"github.com/nbcx/gcs/util"
)

var (
	Manager     = server.NewClientManager() // 管理者
	iComponent  component.IComponent
	secret      string
	localServer = &model.Server{
		Ip: util.LocalIp,
	}
)

func getSecret() string {
	if len(secret) > 0 {
		return secret
	}
	return "Adba723b7fe06819"
}

func SetSecret(v string) {
	secret = v
}
