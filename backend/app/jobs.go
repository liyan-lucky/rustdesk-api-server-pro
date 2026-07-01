package app

import (
	"errors"
	"fmt"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/golang-module/carbon/v2"
	"xorm.io/xorm"
)

func StartJobs(cfg *config.ServerConfig, dbEngine *xorm.Engine) error {
	if dbEngine == nil {
		return errors.New("job db engine is nil")
	}

	s, err := gocron.NewScheduler()
	if err != nil {
		return fmt.Errorf("create scheduler: %w", err)
	}

	jobDuration := time.Duration(cfg.JobsConfig.DeviceCheckJob.Duration) * time.Second
	if jobDuration <= 0 {
		return fmt.Errorf("invalid device check job duration: %s", jobDuration)
	}

	if _, err = s.NewJob(gocron.DurationJob(jobDuration), gocron.NewTask(func() {
		expired := carbon.Now(cfg.Db.TimeZone).SubSeconds(30).ToDateTimeString()
		_, _ = dbEngine.Where("is_online = 1 and updated_at <= ?", expired).Cols("is_online").Update(&model.Device{
			IsOnline: false,
		})
	})); err != nil {
		return fmt.Errorf("create device check job: %w", err)
	}

	s.Start()
	return nil
}
