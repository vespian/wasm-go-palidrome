package main

import (
	onelog "github.com/francoispqt/onelog"
	kubewarden "github.com/kubewarden/policy-sdk-go"
	wapc "github.com/wapc/wapc-guest-tinygo"
)

var (
	logWriter = kubewarden.KubewardenLogWriter{}
	logger    = onelog.New(
		&logWriter,
		onelog.ALL, // shortcut for onelog.DEBUG|onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL
	)
)

func main() {
	wapc.RegisterFunctions(wapc.Functions{
		"validate": validate,
		"validate_settings": func(_ []byte) ([]byte, error) {
			// We are not using any settings
			return kubewarden.AcceptSettings()
		},
	})
}
