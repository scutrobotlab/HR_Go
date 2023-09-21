package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"
	"context"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostApplyLogic {
	return &PostApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostApplyLogic) PostApply(in *hr_service.PostApplyReq) (*hr_service.PostApplyResp, error) {
	a := l.svcCtx.Query.Applicant
	i := l.svcCtx.Query.Intent
	at := l.svcCtx.Query.ApplicantTime

	// 更新申请人信息
	_, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.ApplicantId)).Updates(&model.Applicant{
		Name:   in.Name,
		Gender: common.GenderStringToInt(in.Gender),
		Phone:  in.Phone,
		Form:   in.Form,
	})
	if err != nil {
		return nil, err
	}

	// 更新意向信息
	_, _ = i.WithContext(l.ctx).Where(i.ApplicantID.Eq(in.ApplicantId)).Delete()
	for index, intent := range in.Intents {
		if intent == "" {
			continue
		}
		_ = i.WithContext(l.ctx).Save(&model.Intent{
			Rank:        int32(index),
			Group_:      intent,
			Parallel:    l.svcCtx.Common.Intents.Parallel,
			ApplicantID: in.ApplicantId,
		})
	}

	// 删除已经失效的面试时间
	applicantTimes, _ := at.WithContext(l.ctx).Where(at.ApplicantID.Eq(in.ApplicantId)).Find()
	if applicantTimes != nil {
		for _, time := range applicantTimes {
			_, ok := lo.Find(in.Intents, func(item string) bool {
				return time.Group_ == item
			})
			if !ok {
				// 执行删除已经失效的面试时间
				_, _ = at.WithContext(l.ctx).Where(at.ID.Eq(time.ID)).Delete()
			}
		}
	}

	return &hr_service.PostApplyResp{
		Id: in.ApplicantId,
	}, nil
}
