package timeconfig

import (
	"net/http"

	"HR_Go/hr-api/internal/logic/timeconfig"
	"HR_Go/hr-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTimeConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := timeconfig.NewGetTimeConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetTimeConfig()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
