package scores

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteScoreLogic {
	return &DeleteScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteScoreLogic) DeleteScore(req *types.DeleteScoreReq) (resp *types.Score, err error) {
	_, err = l.svcCtx.AdminService.DeleteScore(l.ctx, &hr_admin_service.DeleteScoreReq{
		ApplicantId: req.Id,
		Group:       req.Group,
		AdminId:     common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.Score{}, nil
}
