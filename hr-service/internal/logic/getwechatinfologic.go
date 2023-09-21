package logic

import (
	"HR_Go/util"
	"context"
	"encoding/json"

	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWechatInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWechatInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWechatInfoLogic {
	return &GetWechatInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWechatInfoLogic) GetWechatInfo(in *hr_service.ApplicantIdReq) (*hr_service.GetWechatInfoResp, error) {
	a := l.svcCtx.Query.Applicant
	applicant, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.ApplicantId)).First()
	if err != nil {
		return nil, util.GrpcErrorNotFound(err)
	}

	profileStr := applicant.Profile
	profile := &Profile{}
	_ = json.Unmarshal([]byte(profileStr), profile)
	return &hr_service.GetWechatInfoResp{
		Openid:     profile.OpenId,
		Nickname:   profile.Nickname,
		Sex:        int32(profile.Sex),
		Language:   profile.Language,
		City:       profile.City,
		Province:   profile.Province,
		Country:    profile.Country,
		Headimgurl: profile.HeadImgURL,
		Privilege:  profile.Privilege,
	}, nil
}
