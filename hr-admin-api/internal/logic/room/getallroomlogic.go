package room

import (
	"HR_Go/common"
	"HR_Go/hr-admin-api/internal/logic"
	"HR_Go/hr-admin-service/hrservice"
	"context"

	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllRoomLogic {
	return &GetAllRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllRoomLogic) GetAllRoom() (resp *types.GetRoomResp, err error) {
	allRoom, err := l.svcCtx.AdminService.GetAllRoom(l.ctx, &hrservice.AdminIdReq{AdminId: common.GetUserInfo(l.ctx).Id})
	if err != nil {
		return nil, err
	}

	return &types.GetRoomResp{
		Rooms:      logic.RoomGrpcToApi(allRoom.Rooms),
		MyRoomName: allRoom.MyRoomName,
		MyRoomId:   allRoom.MyRoomId,
	}, nil
}
