package joinus

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostApplyLogic {
	return &PostApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostApplyLogic) PostApply(req *types.PostApplyReq) (resp *types.PostApplyResp, err error) {
	apply, err := l.svcCtx.HrService.PostApply(l.ctx, &hr_service.PostApplyReq{
		ApplicantId: common.GetUserInfo(l.ctx).Id,
		Name:        req.Name,
		Gender:      req.Gender,
		Phone:       req.Phone,
		Intents:     req.Intents,
		Parallel:    req.Parallel,
		Form:        req.Form,
	})
	if err != nil {
		return nil, err
	}

	return &types.PostApplyResp{
		Id: apply.Id,
	}, nil
}
