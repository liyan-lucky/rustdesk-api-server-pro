package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"rustdesk-api-server-pro/helper/github"
	"rustdesk-api-server-pro/helper/rustdesk"
	"rustdesk-api-server-pro/util"
	"strings"

	"github.com/spf13/cobra"
)

var rustdeskServerCmd = &cobra.Command{
	Use:   "rustdesk [command]",
	Short: "About rustdesk-server command",
}

var rustdeskInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Download and run rustdesk-server",
	Long:  "This command will be download rustdesk-server from https://github.com/rustdesk/rustdesk-server/releases and run it.",
	Run: func(cmd *cobra.Command, args []string) {
		hbbr, hbbs := rustdesk.GetRustdeskServerBin()
		if util.FileExists(hbbr) && util.FileExists(hbbs) {
			fmt.Println("The rustdesk-server has been initialized.")
			os.Exit(0)
		}

		repo := "rustdesk/rustdesk-server"
		var release *github.Release
		var err error
		rustdeskServerVersion := cmd.Flag("version").Value.String()
		if rustdeskServerVersion == "latest" {
			release, err = github.GetLatestRelease(repo)
		} else {
			release, err = github.GetReleaseByTag(repo, rustdeskServerVersion)
		}
		if err != nil {
			fmt.Println("rustdesk-server release lookup error:", err)
			os.Exit(1)
		}

		matchedAsset := github.Asset{}
		arch := runtime.GOARCH
		for _, asset := range release.Assets {
			if runtime.GOOS == "windows" {
				if strings.Contains(asset.Name, "windows") {
					matchedAsset = asset
					arch = "x86_64"
					break
				}
			}
			if runtime.GOOS == "linux" {
				if arch == "arm64" {
					if asset.Name == "rustdesk-server-linux-arm64v8.zip" {
						matchedAsset = asset
						arch = "arm64v8"
						break
					}
				}
				if arch == "amd64" {
					if asset.Name == "rustdesk-server-linux-amd64.zip" {
						matchedAsset = asset
						arch = "amd64"
						break
					}
				}
			}
		}
		if matchedAsset.Name == "" {
			fmt.Println("Your operating system is not supported, only support windows and linux ")
			os.Exit(1)
		}

		// Unzip the rustdesk-server zip if it already exists locally, otherwise download it from github
		if !util.FileExists(matchedAsset.Name) {
			proxyServer := cmd.Flag("proxy").Value.String()
			util.SetHttpProxy(proxyServer)
			err := util.DownloadFile(matchedAsset.BrowserDownloadURL, matchedAsset.Name, true)
			if err != nil {
				fmt.Println("rustdesk-server download error: ", err)
				os.Exit(1)
			}
		}

		fmt.Println("unzipping", matchedAsset.Name)
		err = util.Unzip(matchedAsset.Name, rustdesk.GetRustdeskServerBinDir())
		if err != nil {
			fmt.Println(matchedAsset.Name, "unzipped error: ", err)
			os.Exit(1)
		}
		src := path.Join(rustdesk.GetRustdeskServerBinDir(), arch)
		if err = util.MoveFiles(src, rustdesk.GetRustdeskServerBinDir()); err != nil {
			fmt.Println("rustdesk-server move files error:", err)
			os.Exit(1)
		}
		if err = os.Remove(matchedAsset.Name); err != nil {
			fmt.Println("rustdesk-server cleanup archive error:", err)
			os.Exit(1)
		}
		if err = os.RemoveAll(src); err != nil {
			fmt.Println("rustdesk-server cleanup extracted dir error:", err)
			os.Exit(1)
		}
		fmt.Println("The rustdesk-server has been initialized.")
	},
}

var rustdeskStartCmd = &cobra.Command{
	Use:                   "start",
	Short:                 "Start the rustdesk-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		ret, err := rustdesk.StartServer()
		if ret {
			fmt.Println("rustdesk-server started")
		} else {
			if err != nil {
				fmt.Println("rustdesk-server failed to start:", err.Error())
			} else {
				fmt.Println("rustdesk-server failed to start")
			}
			os.Exit(1)
		}
	},
}

var rustdeskStopCmd = &cobra.Command{
	Use:                   "stop",
	Short:                 "Stop the rustdesk-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		rustdesk.StopServer()
		fmt.Println("rustdesk-server stopped")
	},
}

var rustdeskRestartCmd = &cobra.Command{
	Use:                   "restart",
	Short:                 "Restart the rustdesk-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		rustdesk.StopServer()
		fmt.Println("rustdesk-server stopped")

		ret, err := rustdesk.StartServer()
		if !ret {
			if err != nil {
				fmt.Println("rustdesk-server start failed:", err.Error())
			} else {
				fmt.Println("rustdesk-server start failed")
			}
			os.Exit(1)
		}
		fmt.Println("rustdesk-server started")
	},
}

var rustdeskStatusCmd = &cobra.Command{
	Use:                   "status",
	Short:                 "Show rustdesk-server status",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		hbbrIsRunning, hbbsIsRunning := rustdesk.Status()

		fmt.Println("ServerName\t", "IsRunning")
		fmt.Println("hbbr\t\t", hbbrIsRunning)
		fmt.Println("hbbs\t\t", hbbsIsRunning)
	},
}

var rustdeskKeysCmd = &cobra.Command{
	Use:                   "keys",
	Short:                 "Show rustdesk-server public key",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("public key:")
		fmt.Println(rustdesk.PublicKey())

		showPrivate, _ := cmd.Flags().GetBool("show-private")
		if showPrivate {
			fmt.Println("")
			fmt.Println("private key:")
			fmt.Println(rustdesk.PrivateKey())
		} else {
			fmt.Println("")
			fmt.Println("private key: hidden; use --show-private only when you explicitly need to reveal it")
		}
	},
}

var rustdeskListCmd = &cobra.Command{
	Use:                   "list",
	Short:                 "List the releases rustdesk-server",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		proxyServer := cmd.Flag("proxy").Value.String()
		util.SetHttpProxy(proxyServer)
		releases, err := github.GetReleases("rustdesk/rustdesk-server")
		if err != nil {
			fmt.Println("rustdesk-server release list error:", err)
			os.Exit(1)
		}
		fmt.Printf("%-20s%s\n", "Version", "Published")
		for _, release := range *releases {
			fmt.Printf("%-20s%s\n", release.TagName, release.PublishedAt)
		}
	},
}

func init() {
	rustdeskInstallCmd.Flags().StringP("proxy", "p", "", "Setting up a proxy to download rustdesk-server program (e.g [http|https|socks5]://proxy-host:port)")
	rustdeskInstallCmd.Flags().StringP("version", "v", "latest", "Setting the rustdesk-server program version")
	rustdeskServerCmd.AddCommand(rustdeskInstallCmd)
	rustdeskServerCmd.AddCommand(rustdeskStartCmd)
	rustdeskServerCmd.AddCommand(rustdeskStopCmd)
	rustdeskServerCmd.AddCommand(rustdeskRestartCmd)
	rustdeskServerCmd.AddCommand(rustdeskStatusCmd)
	rustdeskKeysCmd.Flags().Bool("show-private", false, "Print the private key; avoid using this unless required")
	rustdeskServerCmd.AddCommand(rustdeskKeysCmd)

	rustdeskListCmd.Flags().StringP("proxy", "p", "", "Setting up a proxy to download rustdesk-server program (e.g [http|https|socks5]://proxy-host:port)")
	rustdeskServerCmd.AddCommand(rustdeskListCmd)
	RootCmd.AddCommand(rustdeskServerCmd)
}
