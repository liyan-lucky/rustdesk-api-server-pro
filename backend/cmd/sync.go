package cmd

import (
	"fmt"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/util"

	"github.com/spf13/cobra"
)

var dbSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "The api-server database synchronization",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetServerConfig()
		engine, err := db.NewEngine(cfg.Db)
		if err != nil {
			fmt.Println("Db Engine create error:", err)
			return
		}
		models := []interface{}{
			new(model.User),
			new(model.Peer),
			new(model.Tags),
			new(model.AuthToken),
			new(model.Audit),
			new(model.AlarmAudit),
			new(model.OperationAudit),
			new(model.SecurityAudit),
			new(model.CompatAPIAudit),
			new(model.FileTransfer),
			new(model.Device),
			new(model.AddressBook),
			new(model.AddressBookTag),
			new(model.MailLogs),
			new(model.VerifyCode),
			new(model.SystemSettings),
			new(model.MailTemplate),
			new(model.DeviceGroup),
			new(model.DeviceGroupDevice),
			new(model.UserGroup),
			new(model.UserGroupMember),
			new(model.Strategy),
			new(model.StrategyAssignment),
			new(model.OAuthAccount),
		}
		err = engine.Sync(models...)
		if err != nil {
			fmt.Println("Db init error:", err)
			return
		}
		if err = migratePlaintextAuthTokens(engine); err != nil {
			fmt.Println("Auth token migration warning:", err)
		}
		fmt.Println("Database tables sync success")
	},
}

func migratePlaintextAuthTokens(engine db.EngineLike) error {
	legacyTokens := make([]model.AuthToken, 0)
	if err := engine.Where("token != '' and token_hash = ''").Find(&legacyTokens); err != nil {
		return err
	}
	for _, token := range legacyTokens {
		token.TokenHash = util.Sha256Hex(token.Token)
		token.Token = ""
		if _, err := engine.Where("id = ?", token.Id).Cols("token_hash", "token").Update(&token); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	RootCmd.AddCommand(dbSyncCmd)
}
