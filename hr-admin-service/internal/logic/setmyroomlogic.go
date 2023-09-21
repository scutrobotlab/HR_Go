package logic

import (
	"context"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetMyRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetMyRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMyRoomLogic {
	return &SetMyRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetMyRoomLogic) SetMyRoom(in *hr_admin_service.SetMyRoomReq) (*hr_admin_service.StatusResp, error) {
	a := l.svcCtx.Query.Admin
	_, err := a.WithContext(l.ctx).Where(a.ID.Eq(in.AdminId)).UpdateColumn(a.RoomID, in.RoomId)
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
