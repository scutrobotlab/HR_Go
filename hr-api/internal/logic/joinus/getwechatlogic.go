package joinus

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWechatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWechatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWechatLogic {
	return &GetWechatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWechatLogic) GetWechat() (resp *types.GetWechatResp, err error) {
	userInfo := common.GetUserInfo(l.ctx)
	infoResp, err := l.svcCtx.HrService.GetWechatInfo(l.ctx, &hr_service.ApplicantIdReq{
		ApplicantId: userInfo.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetWechatResp{
		OpenId:     infoResp.Openid,
		Nickname:   infoResp.Nickname,
		Sex:        infoResp.Sex,
		Language:   infoResp.Language,
		City:       infoResp.City,
		Province:   infoResp.Province,
		Country:    infoResp.Country,
		HeadImgURL: infoResp.Headimgurl,
		Privilege:  infoResp.Privilege,
	}, nil
}
