package logic

import (
	"context"
	"fmt"

	"HR_Go/hr-admin-service/internal/svc"
	"HR_Go/hr-admin-service/pb/hr-admin-service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoomCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRoomCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoomCommentLogic {
	return &SetRoomCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRoomCommentLogic) SetRoomComment(in *hr_admin_service.SetRoomCommentReq) (*hr_admin_service.StatusResp, error) {
	r := l.svcCtx.Query.Room

	var err error
	if in.Type == InterviewerComment {
		_, err = r.WithContext(l.ctx).Where(r.ID.Eq(in.RoomId)).UpdateColumn(r.InterviewerComment, in.Comment)
	} else if in.Type == ReceiverComment {
		_, err = r.WithContext(l.ctx).Where(r.ID.Eq(in.RoomId)).UpdateColumn(r.ReceiverComment, in.Comment)
	} else {
		err = fmt.Errorf("invalid type")
	}
	if err != nil {
		return nil, err
	}

	return &hr_admin_service.StatusResp{
		Ok: true,
	}, nil
}
