package repository

import "rustdesk-api-server-pro/internal/core"

type CompatRepository interface {
	ApplyDevicesCli(cmd core.CompatDevicesCliCommand) error
}
