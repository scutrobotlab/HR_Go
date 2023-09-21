package evaluation

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/evaluation"
	"HR_Go/hr-admin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetEvaluationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := evaluation.NewGetEvaluationListLogic(r.Context(), svcCtx)
		resp, err := l.GetEvaluationList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
