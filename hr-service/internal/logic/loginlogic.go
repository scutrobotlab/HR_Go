package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *hr_service.LoginReq) (*hr_service.LoginResp, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/oauth/profile/%s", l.svcCtx.Common.Url.WechatUrl, in.Token))
	defer resp.Body.Close()
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}
	respBody, err := io.ReadAll(resp.Body)
	respStr := string(respBody)
	if err != nil {
		return nil, common.GrpcErrorInvalidArgument(err)
	}

	profile := &Profile{}
	_ = json.Unmarshal(respBody, profile)

	applicant := &model.Applicant{
		WechatID: profile.OpenId,
		Avatar:   profile.HeadImgURL,
		Profile:  respStr,
		Form:     "{}",
	}
	a := l.svcCtx.Query.Applicant
	l.svcCtx.Db.WithContext(l.ctx).Where(model.Applicant{WechatID: profile.OpenId}).FirstOrCreate(&applicant)
	_, err = a.WithContext(l.ctx).Where(a.ID.Eq(applicant.ID)).Update(a.Profile, respStr)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	return &hr_service.LoginResp{
		Nickname:    profile.Nickname,
		Openid:      profile.OpenId,
		Headimgurl:  profile.HeadImgURL,
		ApplicantId: applicant.ID,
	}, nil
}
