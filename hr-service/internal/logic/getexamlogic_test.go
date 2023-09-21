package logic

import (
	"HR_Go/dal/model"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGetExamLogic_GetExam(t *testing.T) {
	ctx, svcCtx := GetCtxAndSvcCtxForTest()
	logic := NewGetExamLogic(ctx, svcCtx)

	applicantID := rand.Int63()

	q := &model.Question{
		Question: "为什么",
		Group_:   "测试",
	}
	_ = svcCtx.Query.Question.WithContext(ctx).Create(q)

	aq := &model.ApplicantQuestion{
		ApplicantID: applicantID,
		QuestionID:  q.ID,
		Answer:      "A",
	}
	_ = svcCtx.Query.ApplicantQuestion.WithContext(ctx).Create(aq)

	resp, err := logic.GetExam(&hr_service.GetExamReq{
		ApplicantId: applicantID,
		Group:       "测试",
	})
	if err != nil {
		t.Fail()
		println(err)
	}

	assert.Equal(t, "为什么", resp.Questions[0].Question)
	assert.Equal(t, "A", resp.Questions[0].Answer)
	fmt.Printf("%+v\n", resp)

	_, _ = svcCtx.Query.Question.WithContext(ctx).Delete(q)
	_, _ = svcCtx.Query.ApplicantQuestion.WithContext(ctx).Delete(aq)
}
