package logic

import (
	"context"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGuideLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGuideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGuideLogic {
	return &GetGuideLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGuideLogic) GetGuide(in *hr_service.GetGuideReq) (*hr_service.GetGuideResp, error) {
	g := l.svcCtx.Query.Guide

	guide, err := g.WithContext(l.ctx).Where(g.Group_.Eq(in.Group)).First()
	if err != nil {
		return nil, err
	}

	return &hr_service.GetGuideResp{
		Guide: guide.Guide,
	}, nil
}
