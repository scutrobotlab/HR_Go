package logic

import (
	"HR_Go/dal/model"
	"context"
	"github.com/samber/lo"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemarkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRemarkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRemarkListLogic {
	return &GetRemarkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRemarkListLogic) GetRemarkList(in *hr_admin_service.GetRemarkListReq) (*hr_admin_service.GetRemarkListResp, error) {
	r := l.svcCtx.Query.Remark
	remarks, err := r.WithContext(l.ctx).Where(r.ApplicantID.Eq(in.ApplicantId)).Find()
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.GetRemarkListResp{
		Remarks: lo.Map(remarks, func(item *model.Remark, index int) *hr_admin_service.Remark {
			admin, err := getAbstractAdmin(l.ctx, l.svcCtx, item.AdminID)
			if err != nil {
				return nil
			}

			return &hr_admin_service.Remark{
				ApplicantId: item.ApplicantID,
				AdminId:     item.AdminID,
				Remark:      item.Remark,
				Admin:       admin,
				UpdatedAt:   item.UpdatedAt.Format(time.DateTime),
			}
		}),
	}, nil
}
