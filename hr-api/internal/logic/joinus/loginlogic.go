package joinus

import (
	"HR_Go/common"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"HR_Go/util"
	"context"

	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.HrService.Login(l.ctx, &hr_service.LoginReq{Token: req.Token})
	if err != nil {
		return nil, err
	}

	userInfo := common.UserInfo{
		Id:         loginResp.ApplicantId,
		Permission: common.Applicant,
	}
	// 是否 Debug 模式不影响
	token, err := util.GetLoginJwtToken(l.svcCtx.Config.Auth, userInfo)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Nickname:   loginResp.Nickname,
		OpenId:     loginResp.Openid,
		HeadImgURL: loginResp.Headimgurl,
		Token:      token,
	}, nil
}
