package main

import (
	"fmt"
	"os"
	"rustdesk-api-server-pro/cmd"
)

var appVersion = "latest"
var buildTime = ""

func main() {
	if appVersion != "latest" && appVersion != "" {
		os.Setenv("APP_VERSION", appVersion)
	}
	if buildTime != "" {
		os.Setenv("BUILD_TIME", buildTime)
	}
	if err := cmd.RootCmdBExecute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
