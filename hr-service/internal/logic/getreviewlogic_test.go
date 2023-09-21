package logic

import (
	hr_service "HR_Go/hr-service/pb/hr-service"
	"fmt"
	"testing"
)

func TestNewGetReviewLogic(t *testing.T) {
	ctx, svcCtx := GetCtxAndSvcCtxForTest()
	logic := NewGetReviewLogic(ctx, svcCtx)

	applicantId := int64(351)
	review, err := logic.GetReview(&hr_service.ApplicantIdReq{ApplicantId: applicantId})
	if err != nil {
		return
	}

	fmt.Println(review.Text)
}
