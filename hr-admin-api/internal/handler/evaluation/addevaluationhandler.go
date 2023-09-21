package evaluation

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/evaluation"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddEvaluationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddEvaluationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := evaluation.NewAddEvaluationLogic(r.Context(), svcCtx)
		resp, err := l.AddEvaluation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
