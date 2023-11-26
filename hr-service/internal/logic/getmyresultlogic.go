package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"context"
	"github.com/samber/lo"
	"time"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyResultLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyResultLogic {
	return &GetMyResultLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyResultLogic) GetMyResult(in *hr_service.ApplicantIdReq) (*hr_service.GetMyResultResp, error) {
	tc := l.svcCtx.Query.TimeConfig
	a := l.svcCtx.Query.Applicant
	admit := l.svcCtx.Query.Admit

	timeConfigs, err := tc.WithContext(l.ctx).Find()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}
	applicant, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.ApplicantId)).First()
	if err != nil {
		return nil, common.GrpcErrorNotFound(err)
	}

	tcMap := lo.Associate(timeConfigs, func(item *model.TimeConfig) (string, time.Time) {
		return item.Key, item.Value
	})

	if startTime, ok := tcMap[common.PublishResultStart]; ok && time.Now().After(startTime) {
		admits, _ := admit.WithContext(l.ctx).Where(admit.ApplicantID.Eq(in.ApplicantId)).Find()
		if len(admits) > 0 {
			return &hr_service.GetMyResultResp{
				Status: "passed",
				Groups: lo.Map(admits, func(item *model.Admit, index int) *hr_service.AdmitGroup {
					return &hr_service.AdmitGroup{
						Group: item.Group_,
						Hint:  "TODO",
					}
				}),
			}, nil
		} else {
			return &hr_service.GetMyResultResp{
				Status: "failed",
			}, nil
		}
	}

	if applicant != nil {
		return &hr_service.GetMyResultResp{
			Status: common.HaveNotPublished,
			Times:  GetTimeAndApplicantTime(l.ctx, l.svcCtx, in.ApplicantId),
		}, nil
	} else {
		return &hr_service.GetMyResultResp{
			Status: common.HaveNotAppliedForm,
		}, nil
	}
}
