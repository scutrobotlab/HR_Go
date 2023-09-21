package room

import (
	"HR_Go/common"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoomCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetRoomCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoomCommentLogic {
	return &SetRoomCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetRoomCommentLogic) SetRoomComment(req *types.SetRoomCommentReq) error {
	_, err := l.svcCtx.AdminService.SetRoomComment(l.ctx, &hr_admin_service.SetRoomCommentReq{
		RoomId:  req.RoomId,
		Comment: req.Comment,
		AdminId: common.GetUserInfo(l.ctx).Id,
		Type:    req.Type,
	})
	if err != nil {
		return err
	}

	return nil
}
