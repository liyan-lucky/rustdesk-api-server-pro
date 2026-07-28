package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"rustdesk-api-server-pro/app"
)

var startCmd = &cobra.Command{
	Use:                   "start",
	Short:                 "Start the api-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		_, err := app.StartServerWithContext(ctx)
		if err != nil && ctx.Err() == nil {
			fmt.Println("api-server failed to start:", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
