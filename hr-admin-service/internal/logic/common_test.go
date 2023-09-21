package logic

import (
	"HR_Go/hr-admin-service/internal/config"
	"HR_Go/hr-admin-service/internal/svc"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func GetCtxAndSvcCtxForTest() (context.Context, *svc.ServiceContext) {
	c := config.Config{
		DataSource: "root:123456@(localhost:3306)/hr?charset=utf8mb4&parseTime=True&loc=Local",
	}
	return context.Background(), svc.NewServiceContext4Test(c)
}

func TestExportApplicants(t *testing.T) {
	ctx, svcCtx := GetCtxAndSvcCtxForTest()

	a := svcCtx.Query.Applicant
	find, err := a.WithContext(ctx).Find()
	if err != nil {
		return
	}
	for _, applicant := range find {
		form := applicant.Form
		var formMap map[string]interface{}
		_ = json.Unmarshal([]byte(form), &formMap)
		var intro = ""
		if formMap["self_intro"] != nil {
			intro = formMap["self_intro"].(string)
			intro = strings.ReplaceAll(intro, "\n", "")
		}
		fmt.Printf("%d\t%s\t%s\n", applicant.ID, applicant.Name, intro)
	}
}
