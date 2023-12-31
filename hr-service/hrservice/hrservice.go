// Code generated by goctl. DO NOT EDIT.
// Source: hr.proto

package hrservice

import (
	"context"

	"HR_Go/hr-service/pb/hr-service"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AdmitGroup            = hr_service.AdmitGroup
	Answer                = hr_service.Answer
	ApplicantIdReq        = hr_service.ApplicantIdReq
	FormItem              = hr_service.FormItem
	GetAnnounceStatusReq  = hr_service.GetAnnounceStatusReq
	GetAnnounceStatusResp = hr_service.GetAnnounceStatusResp
	GetCanJoinReq         = hr_service.GetCanJoinReq
	GetCanJoinResp        = hr_service.GetCanJoinResp
	GetExamReq            = hr_service.GetExamReq
	GetExamResp           = hr_service.GetExamResp
	GetFormFormatResp     = hr_service.GetFormFormatResp
	GetFormGroupsResp     = hr_service.GetFormGroupsResp
	GetFormIntentResp     = hr_service.GetFormIntentResp
	GetGuideReq           = hr_service.GetGuideReq
	GetGuideResp          = hr_service.GetGuideResp
	GetIntentListResp     = hr_service.GetIntentListResp
	GetMyFormResp         = hr_service.GetMyFormResp
	GetMyResultResp       = hr_service.GetMyResultResp
	GetMyStepResp         = hr_service.GetMyStepResp
	GetReviewResp         = hr_service.GetReviewResp
	GetTimeConfigResp     = hr_service.GetTimeConfigResp
	GetTimeReq            = hr_service.GetTimeReq
	GetTimeResp           = hr_service.GetTimeResp
	GetWechatInfoResp     = hr_service.GetWechatInfoResp
	IntentItem            = hr_service.IntentItem
	LoginReq              = hr_service.LoginReq
	LoginResp             = hr_service.LoginResp
	PingRequest           = hr_service.PingRequest
	PingResponse          = hr_service.PingResponse
	PostApplyReq          = hr_service.PostApplyReq
	PostApplyResp         = hr_service.PostApplyResp
	PostExamReq           = hr_service.PostExamReq
	PostExamResp          = hr_service.PostExamResp
	Question              = hr_service.Question
	SelectTimeReq         = hr_service.SelectTimeReq
	SelectTimeResp        = hr_service.SelectTimeResp
	TimeConfig            = hr_service.TimeConfig
	TimeItem              = hr_service.TimeItem

	HrService interface {
		Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
		// announce
		GetAnnounceStatus(ctx context.Context, in *GetAnnounceStatusReq, opts ...grpc.CallOption) (*GetAnnounceStatusResp, error)
		// exam
		GetExam(ctx context.Context, in *GetExamReq, opts ...grpc.CallOption) (*GetExamResp, error)
		PostExam(ctx context.Context, in *PostExamReq, opts ...grpc.CallOption) (*PostExamResp, error)
		GetGuide(ctx context.Context, in *GetGuideReq, opts ...grpc.CallOption) (*GetGuideResp, error)
		// form
		GetFormFormat(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetFormFormatResp, error)
		GetFormGroups(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetFormGroupsResp, error)
		GetFormIntent(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetFormIntentResp, error)
		// join-us
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		GetWechatInfo(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetWechatInfoResp, error)
		PostApply(ctx context.Context, in *PostApplyReq, opts ...grpc.CallOption) (*PostApplyResp, error)
		GetCanJoin(ctx context.Context, in *GetCanJoinReq, opts ...grpc.CallOption) (*GetCanJoinResp, error)
		GetMyForm(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetMyFormResp, error)
		GetMyResult(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetMyResultResp, error)
		GetMyStep(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetMyStepResp, error)
		GetIntentList(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetIntentListResp, error)
		GetTime(ctx context.Context, in *GetTimeReq, opts ...grpc.CallOption) (*GetTimeResp, error)
		SelectTime(ctx context.Context, in *SelectTimeReq, opts ...grpc.CallOption) (*SelectTimeResp, error)
		GetReview(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetReviewResp, error)
		// time-config
		GetTimeConfig(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetTimeConfigResp, error)
	}

	defaultHrService struct {
		cli zrpc.Client
	}
)

func NewHrService(cli zrpc.Client) HrService {
	return &defaultHrService{
		cli: cli,
	}
}

func (m *defaultHrService) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

// announce
func (m *defaultHrService) GetAnnounceStatus(ctx context.Context, in *GetAnnounceStatusReq, opts ...grpc.CallOption) (*GetAnnounceStatusResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetAnnounceStatus(ctx, in, opts...)
}

// exam
func (m *defaultHrService) GetExam(ctx context.Context, in *GetExamReq, opts ...grpc.CallOption) (*GetExamResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetExam(ctx, in, opts...)
}

func (m *defaultHrService) PostExam(ctx context.Context, in *PostExamReq, opts ...grpc.CallOption) (*PostExamResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.PostExam(ctx, in, opts...)
}

func (m *defaultHrService) GetGuide(ctx context.Context, in *GetGuideReq, opts ...grpc.CallOption) (*GetGuideResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetGuide(ctx, in, opts...)
}

// form
func (m *defaultHrService) GetFormFormat(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetFormFormatResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetFormFormat(ctx, in, opts...)
}

func (m *defaultHrService) GetFormGroups(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetFormGroupsResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetFormGroups(ctx, in, opts...)
}

func (m *defaultHrService) GetFormIntent(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetFormIntentResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetFormIntent(ctx, in, opts...)
}

// join-us
func (m *defaultHrService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultHrService) GetWechatInfo(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetWechatInfoResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetWechatInfo(ctx, in, opts...)
}

func (m *defaultHrService) PostApply(ctx context.Context, in *PostApplyReq, opts ...grpc.CallOption) (*PostApplyResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.PostApply(ctx, in, opts...)
}

func (m *defaultHrService) GetCanJoin(ctx context.Context, in *GetCanJoinReq, opts ...grpc.CallOption) (*GetCanJoinResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetCanJoin(ctx, in, opts...)
}

func (m *defaultHrService) GetMyForm(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetMyFormResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetMyForm(ctx, in, opts...)
}

func (m *defaultHrService) GetMyResult(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetMyResultResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetMyResult(ctx, in, opts...)
}

func (m *defaultHrService) GetMyStep(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetMyStepResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetMyStep(ctx, in, opts...)
}

func (m *defaultHrService) GetIntentList(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetIntentListResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetIntentList(ctx, in, opts...)
}

func (m *defaultHrService) GetTime(ctx context.Context, in *GetTimeReq, opts ...grpc.CallOption) (*GetTimeResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetTime(ctx, in, opts...)
}

func (m *defaultHrService) SelectTime(ctx context.Context, in *SelectTimeReq, opts ...grpc.CallOption) (*SelectTimeResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.SelectTime(ctx, in, opts...)
}

func (m *defaultHrService) GetReview(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetReviewResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetReview(ctx, in, opts...)
}

// time-config
func (m *defaultHrService) GetTimeConfig(ctx context.Context, in *ApplicantIdReq, opts ...grpc.CallOption) (*GetTimeConfigResp, error) {
	client := hr_service.NewHrServiceClient(m.cli.Conn())
	return client.GetTimeConfig(ctx, in, opts...)
}
