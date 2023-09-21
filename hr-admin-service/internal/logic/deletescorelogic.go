package logic

import (
	"HR_Go/util"
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteScoreLogic {
	return &DeleteScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteScoreLogic) DeleteScore(in *hr_admin_service.DeleteScoreReq) (*hr_admin_service.StatusResp, error) {
	s := l.svcCtx.Query.Score
	_, err := s.WithContext(l.ctx).Where(s.ApplicantID.Eq(in.ApplicantId), s.AdminID.Eq(in.AdminId), s.Group_.Eq(in.Group)).Delete()
	if err != nil {
		return &hr_admin_service.StatusResp{
			Ok: false,
		}, util.GrpcErrorNotFound(err)
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
