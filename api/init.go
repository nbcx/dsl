package api

import (
	"github.com/nbcx/gcs"
	"github.com/nbcx/gcs/distributed/client"
)

var (
	remote    = &client.Remote{}
	component = gcs.GetComponent()
)
