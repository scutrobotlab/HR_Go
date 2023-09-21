package cronjob

import (
	"HR_Go/hr-admin-service/internal/svc"
	"github.com/robfig/cron/v3"
)

type BaseJob struct {
	Name   string
	SvcCtx *svc.ServiceContext
}

func NewBaseJob(name string, svcCtx *svc.ServiceContext) *BaseJob {
	return &BaseJob{
		Name:   name,
		SvcCtx: svcCtx,
	}
}

func (t *BaseJob) Run() {
}

func InitCronJob(svcCtx *svc.ServiceContext, c *cron.Cron) {
	_, _ = c.AddJob("@every 1m", &SendSmsJob{BaseJob: *NewBaseJob("SendSms", svcCtx)})
}
