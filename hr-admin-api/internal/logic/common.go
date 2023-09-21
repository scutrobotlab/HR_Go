package logic

import (
	"HR_Go/hr-admin-api/internal/types"
	"HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/util"
	"encoding/json"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorIsNotFound checks if an error is a grpc not found error
func ErrorIsNotFound(err error) bool {
	s, ok := status.FromError(err)
	return ok && s.Code() == codes.NotFound
}

// IntentsGrpcToApi converts a slice of *hr_admin_service.Intent to a slice of types.Intent
func IntentsGrpcToApi(in []*hr_admin_service.Intent) []types.Intent {
	return lo.Map(in, func(item *hr_admin_service.Intent, index int) types.Intent {
		return types.Intent{
			Group:       item.Group,
			ApplicantId: item.ApplicantId,
			Rank:        item.Rank,
			Parallel:    item.Parallel,
		}
	})
}

// TimesGrpcToApi converts a slice of *hr_admin_service.Time to a slice of types.SavedTime
func TimesGrpcToApi(in []*hr_admin_service.Time) []types.SavedTime {
	return lo.Map(in, func(item *hr_admin_service.Time, index int) types.SavedTime {
		extend := make(map[string]interface{})
		_ = json.Unmarshal([]byte(item.Extend), &extend)
		return types.SavedTime{
			Id:       item.Id,
			Group:    item.Group,
			Date:     item.Date,
			Time:     item.Time,
			Rank:     item.Rank,
			Location: item.Location,
			TotalCnt: item.TotalCnt,
			Grade:    item.Grade,
			Campus:   item.Campus,
			Pivot: types.Pivot{
				ApplicantId: item.ApplicantId,
				TimeId:      item.Id,
			},
			ApplicantId:     0,
			Extend:          extend,
			ApplicantTimeId: item.ApplicantTimeId,
		}
	})
}

// AdmitsGrpcToApi converts a slice of *hr_admin_service.Admit to a slice of types.Admit
func AdmitsGrpcToApi(in []*hr_admin_service.Admit) []types.Admit {
	return lo.Map(in, func(item *hr_admin_service.Admit, index int) types.Admit {
		return types.Admit{
			AdminId:     item.AdminId,
			ApplicantId: item.ApplicantId,
			Group:       item.Group,
			Admin: types.Admin{
				Id:                item.Admin.Id,
				Name:              item.Admin.Name,
				DefaultStandardId: item.Admin.DefaultStandardId,
			},
			UpdatedAt: item.UpdatedAt,
		}
	})
}

// ScoresGrpcToApi converts a slice of *hr_admin_service.Score to a slice of types.Score
func ScoresGrpcToApi(in []*hr_admin_service.Score) []types.ScoreItem {
	return lo.Map(in, func(item *hr_admin_service.Score, index int) types.ScoreItem {
		return types.ScoreItem{
			AdminId:           item.AdminId,
			ApplicantId:       item.ApplicantId,
			Score:             float32(item.Score),
			Group:             item.Group,
			StandardId:        item.StandardId,
			EvaluationDetails: item.EvaluationDetail,
			Admin: types.Admin{
				Id:                item.Admin.Id,
				Name:              item.Admin.Name,
				DefaultStandardId: item.Admin.DefaultStandardId,
			},
		}
	})
}

// NameListGrpcToApi converts a slice of *hr_admin_service.NameList to a slice of types.NameList
func NameListGrpcToApi(in []*hr_admin_service.NameListItem) []types.NameListItem {
	return lo.Map(in, func(item *hr_admin_service.NameListItem, index int) types.NameListItem {
		return types.NameListItem{
			Id:   item.Id,
			Name: item.Name,
		}
	})
}

// StandardsGrpcToApi converts a slice of *hr_admin_service.Standard to a slice of types.Standard
func StandardsGrpcToApi(in []*hr_admin_service.Standard) []types.Standard {
	return lo.Map(in, func(item *hr_admin_service.Standard, index int) types.Standard {
		return types.Standard{
			Name:   item.Name,
			Weight: item.Weight,
		}
	})
}

// EvaluationsGrpcToApi converts a slice of *hr_admin_service.Evaluation to a slice of types.Evaluation
func EvaluationsGrpcToApi(in []*hr_admin_service.Evaluation) []types.EvaluationStandard {
	return lo.Map(in, func(item *hr_admin_service.Evaluation, index int) types.EvaluationStandard {
		return types.EvaluationStandard{
			Id:              item.Id,
			Name:            item.Name,
			ScoresCount:     item.ScoresCount,
			LastEditAdminId: item.UpdatedById,
			LastEditAdmin: types.EvaluationAdmin{
				Id:                item.UpdatedBy.Id,
				Name:              item.UpdatedBy.Name,
				DefaultStandardId: item.UpdatedBy.DefaultStandardId,
			},
			Standard:  StandardsGrpcToApi(item.Standard),
			UpdatedAt: item.UpdatedAt,
		}
	})
}

// StandardsApiToGrpc converts a slice of types.Standard to a slice of *hr_admin_service.Standard
func StandardsApiToGrpc(in []types.Standard) []*hr_admin_service.Standard {
	return lo.Map(in, func(item types.Standard, index int) *hr_admin_service.Standard {
		return &hr_admin_service.Standard{
			Name:   item.Name,
			Weight: item.Weight,
		}
	})
}

// QuestionsGrpcToApi converts a slice of *hr_admin_service.Question to a slice of types.Question
func QuestionsGrpcToApi(in []*hr_admin_service.Question) []types.Question {
	return lo.Map(in, func(item *hr_admin_service.Question, index int) types.Question {
		return types.Question{
			Id:       item.Id,
			Group:    item.Group,
			Question: item.Question,
		}
	})
}

// ScoresGrpcToApi2 converts a slice of *hr_admin_service.Score to a slice of types.Score
func ScoresGrpcToApi2(in []*hr_admin_service.Score) []types.Score {
	return lo.Map(in, func(item *hr_admin_service.Score, index int) types.Score {
		EvaluationDetails := make([]types.EvaluationDetail, 0)
		_ = json.Unmarshal([]byte(item.EvaluationDetail), &EvaluationDetails)
		return types.Score{
			AdminId:           item.AdminId,
			ApplicantId:       item.ApplicantId,
			Score:             float32(item.Score),
			Group:             item.Group,
			StandardId:        item.StandardId,
			EvaluationDetails: EvaluationDetails,
		}
	})
}

// ClassGrpcToApi converts a slice of *hr_admin_service.StatisticsClassify to a slice of types.ClassItem
func ClassGrpcToApi(in []*hr_admin_service.StatisticsClassify) []types.ClassItem {
	return lo.Map(in, func(item *hr_admin_service.StatisticsClassify, index int) types.ClassItem {
		return types.ClassItem{
			Key:   item.Key,
			Count: item.Count,
		}
	})
}

// TimeConfigApiToGrpc converts a slice of types.TimeConfigItem to a slice of *hr_admin_service.TimeConfig
func TimeConfigApiToGrpc(in []types.TimeConfigItem) []*hr_admin_service.TimeConfig {
	return lo.Map(in, func(item types.TimeConfigItem, index int) *hr_admin_service.TimeConfig {
		return &hr_admin_service.TimeConfig{
			Key:   item.Key,
			Value: item.Value,
		}
	})
}

// TimeConfigGrpcToApi converts a slice of *hr_admin_service.TimeConfig to a slice of types.TimeConfigItem
func TimeConfigGrpcToApi(in []*hr_admin_service.TimeConfig) []types.TimeConfigItem {
	return lo.Map(in, func(item *hr_admin_service.TimeConfig, index int) types.TimeConfigItem {
		return types.TimeConfigItem{
			Key:   item.Key,
			Value: item.Value,
		}
	})
}

// TimeStatisticsGrpcToApi converts a slice of *hr_admin_service.TimeStatistics to a slice of types.StatisticsItem
func TimeStatisticsGrpcToApi(in []*hr_admin_service.TimeStatistics) []types.StatisticsItem {
	return lo.Map(in, func(item *hr_admin_service.TimeStatistics, index int) types.StatisticsItem {
		return types.StatisticsItem{
			Group: item.Group,
			Count: item.Count,
		}
	})
}

// QuestionAndAnswerGrpcToApi converts a slice of *hr_admin_service.QuestionAndAnswer to a slice of types.QuestionAndAnswer
func QuestionAndAnswerGrpcToApi(in []*hr_admin_service.QuestionAndAnswer) []types.QuestionAndAnswer {
	return lo.Map(in, func(item *hr_admin_service.QuestionAndAnswer, index int) types.QuestionAndAnswer {
		return types.QuestionAndAnswer{
			QuestionId: item.Id,
			Group:      item.Group,
			Question:   item.Question,
			Answer:     item.Answer,
		}
	})
}

// RemarkGrpcToApi converts a slice of *hr_admin_service.Remark to a slice of types.RemarkItem
func RemarkGrpcToApi(in []*hr_admin_service.Remark) []types.RemarkItem {
	return lo.Map(in, func(item *hr_admin_service.Remark, index int) types.RemarkItem {
		return types.RemarkItem{
			AdminId:     item.AdminId,
			ApplicantId: item.ApplicantId,
			Remark:      item.Remark,
			AdminName:   item.Admin.Name,
			UpdatedAt:   item.UpdatedAt,
		}
	})
}

// ApplicantSmsGrpcToApi converts a slice of *hr_admin_service.ApplicantSms to a slice of types.ApplicantSms
func ApplicantSmsGrpcToApi(in []*hr_admin_service.ApplicantSms) []types.ApplicantSms {
	return lo.Map(in, func(item *hr_admin_service.ApplicantSms, index int) types.ApplicantSms {
		return types.ApplicantSms{
			ApplicantId: item.ApplicantId,
			Name:        item.Name,
			Phone:       item.Phone,
			Status:      item.Status,
			Time:        item.Time,
			Args:        item.Args,
			Content:     item.Content,
			ReplyCnt:    item.ReplyCnt,
		}
	})
}

// HistorySmsGrpcToApi converts a slice of *hr_admin_service.HistorySms to a slice of types.HistorySms
func HistorySmsGrpcToApi(in []*hr_admin_service.HistorySms) []types.HistorySms {
	return lo.Map(in, func(item *hr_admin_service.HistorySms, index int) types.HistorySms {
		return types.HistorySms{
			SenderName:    item.SenderName,
			SenderPhone:   item.SenderPhone,
			ReceiverName:  item.ReceiverName,
			ReceiverPhone: item.ReceiverPhone,
			Content:       item.Content,
			Time:          item.Time,
			Right:         item.Right,
		}
	})
}

// RoomGrpcToApi converts a slice of *hr_admin_service.Room to a slice of types.Room
func RoomGrpcToApi(in []*hr_admin_service.Room) []types.Room {
	return lo.Map(in, func(item *hr_admin_service.Room, index int) types.Room {
		return types.Room{
			Id:                 item.Id,
			Name:               item.Name,
			Location:           item.Location,
			Status:             item.Status,
			Admins:             util.NotNullList(item.Admins),
			Applicant:          item.Applicant,
			Group:              item.Group,
			Time:               item.Time,
			UpdatedAt:          item.UpdatedAt,
			ApplicantId:        item.ApplicantId,
			UpdatedBy:          item.UpdatedBy,
			InterviewerComment: item.InterviewerComment,
			ReceiverComment:    item.ReceiverComment,
			GroupLabel:         item.GroupLabel,
		}
	})
}

// PushableApplicantGrpcToApi converts a slice of *hr_admin_service.PushableApplicant to a slice of types.PushableApplicant
func PushableApplicantGrpcToApi(in []*hr_admin_service.PushableApplicant) []types.PushableApplicant {
	return lo.Map(in, func(item *hr_admin_service.PushableApplicant, index int) types.PushableApplicant {
		return types.PushableApplicant{
			ApplicantId:     item.ApplicantId,
			Name:            item.Name,
			Group:           item.Group,
			Time:            item.Time,
			ApplicantTimeId: item.ApplicantTimeId,
		}
	})
}

// TimeTableItemGrpcToApi converts a slice of *hr_admin_service.TimeTableItem to a slice of types.TimeTableItem
func TimeTableItemGrpcToApi(in []*hr_admin_service.TimeTableItem) []types.TimeTableItem {
	return lo.Map(in, func(item *hr_admin_service.TimeTableItem, index int) types.TimeTableItem {
		return types.TimeTableItem{
			Id:       item.Id,
			Name:     item.Name,
			Start:    item.Start,
			End:      item.End,
			Color:    item.Color,
			Category: item.Category,
		}
	})
}
