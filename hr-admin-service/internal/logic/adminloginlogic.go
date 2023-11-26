package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoginLogic) AdminLogin(in *hr_admin_service.AdminLoginReq) (*hr_admin_service.AdminLoginResp, error) {
	if l.svcCtx.Common.Debug {
		return &hr_admin_service.AdminLoginResp{
			Admin: &hr_admin_service.Admin{
				Id:   1,
				Name: "虎王",
				Groups: lo.Map(l.svcCtx.Common.Groups, func(item string, index int) *hr_admin_service.AdminGroup {
					return &hr_admin_service.AdminGroup{
						Id:    int64(index + 1),
						Name:  item,
						Power: 2,
					}
				}),
				DefaultStandardId: 0,
				SmsEnabled:        true,
			},
		}, nil
	}

	oAuth := l.svcCtx.Common.OAuth
	param1 := fmt.Sprintf("grant_type=authorization_code&client_id=%s&client_secret=%s&redirect_uri=%s&code=%s", oAuth.ClientId, oAuth.ClientSecret, oAuth.RedirectUrl, in.Code)
	req1, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/oauth/token?%s", l.svcCtx.Common.Url.AdminUrl, param1), nil)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}
	resp1, err := http.DefaultClient.Do(req1)
	if err != nil {
		return nil, err
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != http.StatusOK {
		return nil, common.GrpcErrorInvalidArgument(fmt.Errorf("invalid code"))
	}
	bytes1, err := io.ReadAll(resp1.Body)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	type Resp1 struct {
		AccessToken string `json:"access_token"`
	}
	resp1Body := &Resp1{}
	err = json.Unmarshal(bytes1, resp1Body)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	param2 := fmt.Sprintf("access_token=%s", resp1Body.AccessToken)
	req2, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/oauth/userinfo?%s", l.svcCtx.Common.Url.AdminUrl, param2), nil)
	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		return nil, common.GrpcErrorInternal(fmt.Errorf("HTTP %d", resp2.StatusCode))
	}
	bytes2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	resp2Body := make(map[string]any)
	err = json.Unmarshal(bytes2, &resp2Body)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}
	type Group struct {
		Id    int64
		Name  string
		Power int32
	}
	groups := lo.Map(resp2Body["groups"].([]any), func(item any, index int) *Group {
		it := item.(map[string]any)
		return &Group{
			Id:    int64(it["id"].(float64)),
			Name:  it["name"].(string),
			Power: int32(it["pivot"].(map[string]any)["power"].(float64)),
		}
	})
	resp2Body["groups"] = groups
	profileJson, err := json.Marshal(resp2Body)
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	adminId := int64(resp2Body["id"].(float64))
	adminName := resp2Body["name"].(string)
	a := l.svcCtx.Query.Admin
	smsEnabled := false
	if adminId == 1 {
		smsEnabled = true
	}
	admin := model.Admin{
		ID:                adminId,
		Name:              adminName,
		DefaultStandardID: 0,
		Profile:           "{}",
		SmsEnabled:        smsEnabled,
	}
	count, _ := a.WithContext(l.ctx).Where(a.ID.Eq(adminId)).Count()
	if count == 0 {
		err = a.WithContext(l.ctx).Create(&admin)
		if err != nil {
			return nil, common.GrpcErrorInternal(err)
		}
	}
	_, err = a.WithContext(l.ctx).Where(a.ID.Eq(adminId)).UpdateColumns(model.Admin{
		Name:    adminName,
		Profile: string(profileJson),
	})
	if err != nil {
		return nil, common.GrpcErrorInternal(err)
	}

	return &hr_admin_service.AdminLoginResp{
		Admin: &hr_admin_service.Admin{
			Id:     adminId,
			Name:   adminName,
			Uuid:   resp2Body["uuid"].(string),
			Email:  resp2Body["email"].(string),
			Avatar: resp2Body["avatar"].(string),
			Groups: lo.Map(groups, func(item *Group, index int) *hr_admin_service.AdminGroup {
				return &hr_admin_service.AdminGroup{
					Id:    item.Id,
					Name:  item.Name,
					Power: item.Power,
				}
			}),
			SmsEnabled: false,
		},
	}, nil
}
