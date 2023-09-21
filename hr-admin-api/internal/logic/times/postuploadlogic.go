package times

import (
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"
	"github.com/samber/lo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostUploadLogic {
	return &PostUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostUploadLogic) PostUpload(req *types.PostUploadReq) error {
	_, err := l.svcCtx.AdminService.UploadTimes(l.ctx, &hr_admin_service.UploadTimesReq{
		Groups: req.Groups,
		Data: lo.Map(req.Data, func(item map[string]string, index int) *hr_admin_service.TimeItem {
			return &hr_admin_service.TimeItem{
				Data: item,
			}
		}),
	})
	if err != nil {
		return err
	}

	return nil
}
