package cmd

import (
	"cli/client"

	"github.com/go-openapi/runtime"
)

var (
	Api  *client.Deplyf
	Auth runtime.ClientAuthInfoWriter
)
