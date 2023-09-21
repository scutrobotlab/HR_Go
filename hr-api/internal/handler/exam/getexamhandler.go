package exam

import (
	"net/http"

	"HR_Go/hr-api/internal/logic/exam"
	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetExamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetExamReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := exam.NewGetExamLogic(r.Context(), svcCtx)
		resp, err := l.GetExam(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
