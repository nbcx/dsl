package gcs

import (
	"github.com/nbcx/gcs/distributed/component"
	"github.com/nbcx/gcs/server"
)

var (
	Manager    = server.GetManager() // 管理者
	iComponent component.IComponent
	secret     string
	local      = make(map[string]string)
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
