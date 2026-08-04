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

	if _, err = s.NewJob(gocron.DurationJob(jobDuration), gocron.NewTask(func() {
		expired := carbon.Now(cfg.Db.TimeZone).SubSeconds(30).ToDateTimeString()
		_, _ = dbEngine.Where("is_online = 1 and updated_at <= ?", expired).Cols("is_online").Update(&model.Device{
			IsOnline: false,
		})
	})); err != nil {
		return errcode.Errorf(errcode.ERRB006.Code, errcode.ERRB006.Message)
	}

	s.Start()
	return nil
}
