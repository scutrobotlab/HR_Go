package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"context"
	"encoding/json"
	"github.com/samber/lo"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEvaluationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEvaluationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEvaluationsLogic {
	return &GetEvaluationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEvaluationsLogic) GetEvaluations(in *hr_admin_service.AdminIdReq) (*hr_admin_service.GetEvaluationsResp, error) {
	es := l.svcCtx.Query.EvaluationStandard
	s := l.svcCtx.Query.Score

	evaluationStandards, err := es.WithContext(l.ctx).Find()
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	return &hr_admin_service.GetEvaluationsResp{
		Evaluations: lo.Map(evaluationStandards, func(item *model.EvaluationStandard, index int) *hr_admin_service.Evaluation {
			updatedBy, err := getAbstractAdmin(l.ctx, l.svcCtx, item.LastEditAdminID)
			if err != nil {
				return nil
			}
			count, _ := s.WithContext(l.ctx).Where(s.StandardID.Eq(item.ID)).Count()
			standard := make([]*hr_admin_service.Standard, 0)
			err = json.Unmarshal([]byte(item.Standard), &standard)
			if err != nil {
				return nil
			}
			return &hr_admin_service.Evaluation{
				Id:          item.ID,
				Name:        item.Name,
				ScoresCount: int32(count),
				UpdatedById: item.LastEditAdminID,
				UpdatedBy:   updatedBy,
				Standard:    standard,
				UpdatedAt:   item.UpdatedAt.Format(time.DateTime),
			}
		}),
	}, nil
}
