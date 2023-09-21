package times

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/times"
	"HR_Go/hr-admin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetStatisticsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := times.NewGetStatisticsLogic(r.Context(), svcCtx)
		resp, err := l.GetStatistics()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
