package gcs

import (
	"github.com/nbcx/dsl/distributed/component"
	"github.com/nbcx/dsl/server"
)

var (
	Manager    = server.GetManager() // 管理者
	iComponent component.IComponent
	secret     string
	local      = make(map[int]string)
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
