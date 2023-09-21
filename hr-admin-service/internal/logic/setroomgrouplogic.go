package logic

import (
	"context"
	"time"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoomGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRoomGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoomGroupLogic {
	return &SetRoomGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRoomGroupLogic) SetRoomGroup(in *hr_admin_service.SetRoomGroupReq) (*hr_admin_service.StatusResp, error) {
	r := l.svcCtx.Query.Room
	_, err := r.WithContext(l.ctx).Where(r.ID.Eq(in.RoomId)).Updates(map[string]interface{}{
		"group_label":       in.Group,
		"status_updated_at": time.Now(),
		"updated_by":        in.AdminId,
	})
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
