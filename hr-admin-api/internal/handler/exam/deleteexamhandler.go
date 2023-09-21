package exam

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/exam"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteExamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteExamReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := exam.NewDeleteExamLogic(r.Context(), svcCtx)
		resp, err := l.DeleteExam(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
