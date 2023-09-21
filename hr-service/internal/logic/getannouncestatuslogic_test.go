package logic

import (
	"HR_Go/common"
	"HR_Go/dal/model"
	hr_service "HR_Go/hr-service/pb/hr-service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAnnounceStatusLogic_GetAnnounceStatus(t *testing.T) {
	ctx, svcCtx := GetCtxAndSvcCtxForTest()
	logic := NewGetAnnounceStatusLogic(ctx, svcCtx)

	ac := svcCtx.Query.AnnounceConfig
	m := &model.AnnounceConfig{
		Status: common.HaveNotAppliedForm,
		Body:   "您还没有提交报名表。",
	}
	err := ac.WithContext(ctx).Create(m)
	if err != nil {
		t.Fail()
	}

	status, err := logic.GetAnnounceStatus(&hr_service.GetAnnounceStatusReq{
		Status: common.HaveNotAppliedForm,
	})
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "您还没有提交报名表。", status.Body)

	_, err = ac.WithContext(ctx).Delete(m)
}
