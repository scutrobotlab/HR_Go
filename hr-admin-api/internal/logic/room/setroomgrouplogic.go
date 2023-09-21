package room

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoomGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoomGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoomGroupLogic {
	return &SetRoomGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoomGroupLogic) SetRoomGroup(req *types.SetRoomGroupReq) error {
	_, err := l.svcCtx.AdminService.SetRoomGroup(l.ctx, &hr_admin_service.SetRoomGroupReq{
		RoomId:  req.RoomId,
		Group:   req.Group,
		AdminId: common.GetUserInfo(l.ctx).Id,
	})
	if err != nil {
		return err
	}

	return nil
}
