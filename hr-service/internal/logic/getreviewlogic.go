package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	hr_admin_service "HR_Go/hr-admin-service/pb/hr-admin-service"
	"HR_Go/hr-service/internal/svc"
	"HR_Go/hr-service/pb/hr-service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"k8s.io/apimachinery/pkg/api/errors"
	"math"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReviewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetReviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReviewLogic {
	return &GetReviewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetReviewLogic) GetReview(in *hr_service.ApplicantIdReq) (*hr_service.GetReviewResp, error) {
	reviewTime := common.GetTimeConfig(l.svcCtx.Query, common.ReviewThisSeason)
	if time.Now().Before(reviewTime) {
		return nil, fmt.Errorf("没有到赛季回顾时间")
	}

	var text string
	var result []string
	writeResult := func() {
		result = append(result, text)
		text = ""
	}

	aid := in.ApplicantId

	query := l.svcCtx.Query
	ad := query.Admin
	t := query.Time
	admit := query.Admit
	s := query.Score
	a := query.Applicant
	aq := query.ApplicantQuestion
	at := query.ApplicantTime
	i := query.Intent
	r := query.Remark

	getAdmin := func(adminId int64) model.Admin {
		admin, err := ad.WithContext(l.ctx).Where(ad.ID.Eq(adminId)).First()
		if err != nil {
			return model.Admin{}
		}
		return *admin
	}

	applicant, err := a.WithContext(l.ctx).Where(a.ID.Eq(aid)).First()
	if err != nil {
		return nil, err
	}
	applicantCount, err := a.WithContext(l.ctx).Count()
	if err != nil {
		return nil, err
	}
	firstApplicant, err := a.WithContext(l.ctx).First()
	if err != nil {
		return nil, err
	}

	var form map[string]any
	err = json.Unmarshal([]byte(applicant.Form), &form)
	if err != nil {
		return nil, err
	}
	selfIntro := form["self_intro"].(string)

	intents, err := i.WithContext(l.ctx).Where(i.ApplicantID.Eq(aid)).Unscoped().Find()
	if err != nil {
		return nil, err
	}
	intentsStrList := lo.Uniq(lo.Map(intents, func(item *model.Intent, index int) string {
		return item.Group_
	}))
	intentsStr := strings.Join(intentsStrList, "、")
	chosenIntents, err := i.WithContext(l.ctx).Where(i.ApplicantID.Eq(aid)).Find()
	if err != nil {
		return nil, err
	}
	chosenIntentsStr := strings.Join(lo.Map(chosenIntents, func(item *model.Intent, index int) string {
		return item.Group_
	}), "、")

	applicantQuestions, err := aq.WithContext(l.ctx).Where(aq.ApplicantID.Eq(aid)).Find()
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}
	var correctCount, wrongCount, cancelCount int
	for _, it := range applicantQuestions {
		switch it.Answer {
		case "0":
			cancelCount++
		case "1":
			correctCount++
		case "2":
			wrongCount++
		}
	}

	text += fmt.Sprintf("%s同学，你好！\n", applicant.Name)
	text += fmt.Sprintf("%s，你首次进入面试系统，并提交了报名表。\n\n", applicant.CreatedAt.Format("2006年01月02日 15:04:05"))
	text += fmt.Sprintf("在%d位面试者中，你是第%d位进入系统的用户。", applicantCount, applicant.ID-firstApplicant.ID+1)
	writeResult()

	text += fmt.Sprintf("你提交的自我介绍是\n%s\n", selfIntro)
	if len(intentsStrList) == len(chosenIntents) {
		text += fmt.Sprintf("你坚定地选择了%s作为志愿。\n\n", chosenIntentsStr)
	} else {
		text += fmt.Sprintf("你曾经在%s中纠结过，但最终你选择了把%s作为志愿。\n\n", intentsStr, chosenIntentsStr)
	}
	text += fmt.Sprintf("你共完成了面试题库中的%d道题，其中有%d道你认为自己会，%d道你认为自己不会，还有%d道你取消了答案。", len(applicantQuestions), correctCount, wrongCount, cancelCount)
	writeResult()

	for index, intent := range chosenIntents {
		applicantTime, err := at.WithContext(l.ctx).Where(at.ApplicantID.Eq(aid), at.Group_.Eq(intent.Group_)).First()
		if err != nil {
			continue
		}
		oneTime, err := t.WithContext(l.ctx).Where(t.ID.Eq(applicantTime.TimeID)).First()
		if err != nil {
			continue
		}
		dateTime := oneTime.Time.Format("2006年01月02日 15:04")

		scores, err := s.WithContext(l.ctx).Where(s.ApplicantID.Eq(aid), s.Group_.Eq(intent.Group_)).Find()
		if err != nil {
			continue
		}

		text += fmt.Sprintf("你于%s在%s进行了%s的面试。\n", dateTime, oneTime.Location, intent.Group_)

		sumScore := 0.0
		for _, score := range scores {
			admin := getAdmin(score.AdminID)
			details := score.EvaluationDetails
			detailMap := make([]map[string]any, 0)
			_ = json.Unmarshal([]byte(details), &detailMap)
			standards := make([]string, 0)
			thisScores := make([]float64, 0)
			sumWeight := lo.SumBy(detailMap, func(item map[string]any) float64 {
				return item["weight"].(float64)
			})
			for _, d := range detailMap {
				name := d["name"].(string)
				score := d["score"].(float64)
				weight := d["weight"].(float64)
				proportion := weight / sumWeight
				standards = append(standards, fmt.Sprintf("%s占%.0f%%", name, 100*proportion))
				thisScores = append(thisScores, score)
			}
			scoresStrList := lo.Map(thisScores, func(item float64, index int) string {
				return fmt.Sprintf("%.1f", item)
			})
			text += fmt.Sprintf("面试官%s采用了如下评价标准进行打分：%s，你分别获得了五分制%s分。你的加权成绩为%.2f分\n", admin.Name, strings.Join(standards, "、"), strings.Join(scoresStrList, "、"), score.Score)
			sumScore += score.Score
		}
		avgScore := sumScore / float64(len(scores))
		if math.IsNaN(avgScore) {
			avgScore = 0.0
		}

		rankResp, err := l.svcCtx.AdminService.GetRank(l.ctx, &hr_admin_service.GetRankReq{
			ApplicantId: aid,
			Group:       intent.Group_,
		})
		if err != nil {
			continue
		}

		text += fmt.Sprintf("你的最终成绩是五分制%.2f分，在%d位面试者中排名第%d名。\n", avgScore, rankResp.Total, rankResp.Rank)
		if index != (len(chosenIntents) - 1) {
			text += "\n"
		}
	}
	writeResult()

	remarks, err := r.WithContext(l.ctx).Where(r.ApplicantID.Eq(aid)).Find()
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	text += fmt.Sprintf("共有%d位面试官对你做出了评价。\n", len(remarks))
	adminNames := make([]string, 0)
	for _, remark := range remarks {
		admin := getAdmin(remark.AdminID)
		adminNames = append(adminNames, admin.Name)
	}
	if len(adminNames) > 0 {
		text += "分别是："
		text += strings.Join(adminNames, "、")
	}
	text += "\n评价内容不统一展示，你可以礼貌地询问面试官。"
	writeResult()

	admits, err := admit.WithContext(l.ctx).Where(admit.ApplicantID.Eq(aid)).Find()
	if err != nil {
		return nil, err
	}
	for _, admit := range admits {
		admin := getAdmin(admit.AdminID)
		text += fmt.Sprintf("%s，面试官%s录取了你，你成为了华南虎%s组实习生。\n", admit.CreatedAt.Format("2006年01月02日 15:04:05"), admin.Name, admit.Group_)
	}
	writeResult()

	return &hr_service.GetReviewResp{
		Text: result,
	}, nil
}
