package cronjob

import (
	"HR_Go/common/sms"
	"HR_Go/dal/model"
	"HR_Go/dal/query"
	"HR_Go/hr-admin-service/internal/logic"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	SmsSendLockKey = "SmsSendLock"
)

type SendSmsJob struct {
	BaseJob
}

func (t *SendSmsJob) Run() {
	ctx := context.Background()
	logx.Infof("cron job %s start", t.Name)

	smsEnabled := logic.GetConfigOrDefault(ctx, t.SvcCtx, logic.SmsServiceEnabled, "0")
	if smsEnabled != "1" {
		logx.Infof("sms service is not enabled, cron job %s end", t.Name)
		return
	}

	// 获取锁
	locked, tx := t.GetLock(ctx)
	if locked {
		logx.Infof("lock is already locked, cron job %s end", t.Name)
		return
	}
	logx.Infof("get lock success, cron job %s continue", t.Name)

	getApplicantSmsLogic := logic.NewGetApplicantSmsLogic(ctx, t.SvcCtx)
	sendSmsLogic := logic.NewSendSmsLogic(ctx, t.SvcCtx)

	for _, smsTyp := range sms.Types {
		// 获取待发送短信
		smsResp, err := getApplicantSmsLogic.GetApplicantSms(&hr_admin_service.GetApplicantSmsReq{
			Typ:      smsTyp,
			PageSize: -1,
		})
		if err != nil {
			continue
		}
		// 获取待发送短信的申请人ID
		ids := make([]int64, 0)
		for _, applicantSms := range smsResp.ApplicantSms {
			if applicantSms.Status == sms.SendRightNow {
				ids = append(ids, applicantSms.ApplicantId)
			}
		}
		logx.Infof("send ids: %v", ids)
		// 发送短信
		_, _ = sendSmsLogic.SendSms(&hr_admin_service.SendSmsReq{
			AdminId:      1, // 自动发送以管理员1的身份发送
			Typ:          smsTyp,
			ApplicantIds: ids,
		})
	}

	// 释放锁
	_ = t.ReleaseLock(ctx, tx)
	logx.Infof("release lock success, cron job %s end", t.Name)
}

func (t *SendSmsJob) GetLock(ctx context.Context) (locked bool, tx *query.QueryTx) {
	// 开启事务
	tx = t.SvcCtx.Query.Begin()
	c := tx.Configuration
	// 获取锁
	_, err := c.WithContext(ctx).Clauses(clause.Locking{
		Strength: "UPDATE",
		Options:  "NOWAIT",
	}).Where(c.Key.Eq(SmsSendLockKey)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在则创建
			_ = c.WithContext(ctx).Create(&model.Configuration{Key: SmsSendLockKey, Value: "0"})
			_ = tx.Commit()
			return true, nil
		}
		_ = tx.Rollback()
		// 锁已被占用
		return true, nil
	}
	// 锁未被占用
	return false, tx
}

func (t *SendSmsJob) ReleaseLock(ctx context.Context, tx *query.QueryTx) error {
	c := tx.Configuration
	_, err := c.WithContext(ctx).Where(c.Key.Eq(SmsSendLockKey)).Update(c.Value, "0")
	if err != nil {
		return tx.Rollback()
	}
	// 释放锁
	return tx.Commit()
}
