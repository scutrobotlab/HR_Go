package joinus

import (
	"net/http"

	"HR_Go/hr-api/internal/logic/joinus"
	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTimeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTimeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := joinus.NewGetTimeLogic(r.Context(), svcCtx)
		resp, err := l.GetTime(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
