package main

import (
	"fmt"
	"os"
	"rustdesk-api-server-pro/cmd"
	"rustdesk-api-server-pro/internal/service"
)

var appVersion = "latest"
var buildTime = ""

func main() {
	service.SetCompatServerVersion(appVersion)
	if appVersion != "latest" && appVersion != "" {
		os.Setenv("APP_VERSION", appVersion)
	}
	if buildTime != "" {
		os.Setenv("BUILD_TIME", buildTime)
	}
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
