package announce

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnnounceStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAnnounceStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnnounceStatusLogic {
	return &GetAnnounceStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAnnounceStatusLogic) GetAnnounceStatus(req *types.GetAnnounceStatusReq) (resp *types.GetAnnounceStatusResp, err error) {
	status, err := l.svcCtx.HrService.GetAnnounceStatus(l.ctx, &hr_service.GetAnnounceStatusReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetAnnounceStatusResp{
		Id:     status.Id,
		Status: status.Status,
		Body:   status.Body,
	}, nil
}
