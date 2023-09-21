package evaluation

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/evaluation"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteEvaluationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteEvaluationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := evaluation.NewDeleteEvaluationLogic(r.Context(), svcCtx)
		err := l.DeleteEvaluation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
