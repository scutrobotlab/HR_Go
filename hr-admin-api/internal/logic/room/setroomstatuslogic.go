package room

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoomStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoomStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoomStatusLogic {
	return &SetRoomStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoomStatusLogic) SetRoomStatus(req *types.SetRoomStatusReq) error {
	_, err := l.svcCtx.AdminService.SetRoomStatus(l.ctx, &hr_admin_service.SetRoomStatusReq{
		RoomId:  req.RoomId,
		Status:  req.Status,
		AdminId: common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return err
	}

	return nil
}
