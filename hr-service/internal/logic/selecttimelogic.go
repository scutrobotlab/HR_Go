package logic

import (
	"HR_Go/dal/model"
	"HR_Go/dal/query"
	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type SelectTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectTimeLogic {
	return &SelectTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectTimeLogic) SelectTime(in *hr_service.SelectTimeReq) (*hr_service.SelectTimeResp, error) {
	err := l.svcCtx.Query.Transaction(func(tx *query.Query) error {
		at := tx.ApplicantTime
		t := tx.Time

		// 删除原有的时间段
		_, _ = at.WithContext(l.ctx).
			Where(at.ApplicantID.Eq(in.ApplicantId), at.Group_.Eq(in.Group)).
			Delete()

		if in.TimeId > 0 {
			// 检查人数是否已满
			theTime, err := t.WithContext(l.ctx).Where(t.ID.Eq(in.TimeId)).First()
			if err != nil {
				return err
			}
			count, _ := at.WithContext(l.ctx).Where(at.TimeID.Eq(in.TimeId)).Count()
			if int32(count) >= theTime.TotalCnt {
				// 人数已满
				return fmt.Errorf("人数已满")
			}

			// 选择新的时间段
			_ = at.WithContext(l.ctx).Save(&model.ApplicantTime{
				ApplicantID: in.ApplicantId,
				Group_:      in.Group,
				TimeID:      in.TimeId,
				Extend:      "{}",
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &hr_service.SelectTimeResp{
		ApplicantId: in.ApplicantId,
		Group:       in.Group,
		TimeId:      in.TimeId,
	}, nil
}
