package app

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/errcode"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/golang-module/carbon/v2"
	"xorm.io/xorm"
)

func StartJobs(cfg *config.ServerConfig, dbEngine *xorm.Engine) error {
	if dbEngine == nil {
		return errcode.New(errcode.ERRB003.Code, errcode.ERRB003.Message)
	}

	s, err := gocron.NewScheduler()
	if err != nil {
		return errcode.Errorf(errcode.ERRB005.Code, errcode.ERRB005.Message)
	}

	jobDuration := time.Duration(cfg.JobsConfig.DeviceCheckJob.Duration) * time.Second
	if jobDuration <= 0 {
		return errcode.New(errcode.ERRB004.Code, errcode.ERRB004.Message)
	}
	offlineAfterSeconds := cfg.JobsConfig.DeviceCheckJob.OfflineAfterSeconds
	if offlineAfterSeconds <= 0 {
		offlineAfterSeconds = 30
	}

	if _, err = s.NewJob(gocron.DurationJob(jobDuration), gocron.NewTask(func() {
		expired := carbon.Now(cfg.Db.TimeZone).SubSeconds(offlineAfterSeconds).ToStdTime().UTC().Format("2006-01-02 15:04:05")
		_, _ = markExpiredDevicesOffline(dbEngine, expired)
	})); err != nil {
		return errcode.Errorf(errcode.ERRB006.Code, errcode.ERRB006.Message)
	}

	s.Start()
	return nil
}

// markExpiredDevicesOffline 将最后心跳早于截止时间的在线设备标记为离线。
func markExpiredDevicesOffline(dbEngine *xorm.Engine, expired string) (int64, error) {
	return dbEngine.Where("is_online = 1 and updated_at <= ?", expired).Cols("is_online").Update(&model.Device{
		IsOnline: false,
	})
}
