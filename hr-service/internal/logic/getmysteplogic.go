package logic

import (
	"HR_Go/common"
	"context"
	"time"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyStepLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	PostApply  = 1 // 提交报名表
	SelectTime = 2 // 选择面试时间
	Done       = 3 // 已完成
)

func NewGetMyStepLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyStepLogic {
	return &GetMyStepLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyStepLogic) GetMyStep(in *hr_service.ApplicantIdReq) (*hr_service.GetMyStepResp, error) {
	publishResult := common.GetTimeConfig(l.svcCtx.Query, common.PublishResultStart)
	if !publishResult.IsZero() && time.Now().After(publishResult) {
		// 已经发布结果
		return &hr_service.GetMyStepResp{Step: Done}, nil
	}

	a := l.svcCtx.Query.Applicant
	applicant, _ := a.WithContext(l.ctx).Where(a.ID.Eq(in.ApplicantId)).First()
	if applicant.Phone == "" || applicant.Form == "" || applicant.Form == "{}" {
		// 还没有提交报名表
		return &hr_service.GetMyStepResp{Step: PostApply}, nil
	}

	i := l.svcCtx.Query.Intent
	at := l.svcCtx.Query.ApplicantTime
	intents, _ := i.WithContext(l.ctx).Where(i.ApplicantID.Eq(in.ApplicantId)).Find()
	applicantTimes, _ := at.WithContext(l.ctx).Where(at.Where(at.ApplicantID.Eq(in.ApplicantId))).Find()
	if len(intents) != len(applicantTimes) {
		// 存在组别没有选择面试时间
		return &hr_service.GetMyStepResp{Step: SelectTime}, nil
	}

	// 已完成
	return &hr_service.GetMyStepResp{Step: Done}, nil
}
