package room

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetMyRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMyRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMyRoomLogic {
	return &SetMyRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMyRoomLogic) SetMyRoom(req *types.SetMyRoomReq) error {
	_, err := l.svcCtx.AdminService.SetMyRoom(l.ctx, &hr_admin_service.SetMyRoomReq{
		AdminId: common.GetUserInfo(l.ctx).Id,
		RoomId:  req.RoomId,
	})
	if err != nil {
		return err
	}

	return nil
}
