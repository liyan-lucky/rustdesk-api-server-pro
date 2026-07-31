package main

import (
	"fmt"
	"os"
	"rustdesk-api-server-pro/cmd"
)

var appVersion = "latest"

func main() {
	if appVersion != "latest" && appVersion != "" {
		os.Setenv("APP_VERSION", appVersion)
	}
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
