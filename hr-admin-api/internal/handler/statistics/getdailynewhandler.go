package statistics

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/statistics"
	"HR_Go/hr-admin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetDailyNewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := statistics.NewGetDailyNewLogic(r.Context(), svcCtx)
		resp, err := l.GetDailyNew()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
