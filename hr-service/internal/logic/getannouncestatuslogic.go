package logic

import (
	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"
	"HR_Go/util"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnnounceStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnnounceStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnnounceStatusLogic {
	return &GetAnnounceStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnnounceStatusLogic) GetAnnounceStatus(in *hr_service.GetAnnounceStatusReq) (*hr_service.GetAnnounceStatusResp, error) {
	ac := l.svcCtx.Query.AnnounceConfig
	first, err := ac.WithContext(l.ctx).Where(ac.Status.Eq(in.Status)).First()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	return &hr_service.GetAnnounceStatusResp{
		Id:     first.ID,
		Status: first.Status,
		Body:   first.Body,
	}, nil
}
