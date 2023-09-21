package exam

import (
	"net/http"

	"HR_Go/hr-api/internal/logic/exam"
	"HR_Go/hr-api/internal/svc"
	"HR_Go/hr-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetGuideHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetGuideReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := exam.NewGetGuideLogic(r.Context(), svcCtx)
		resp, err := l.GetGuide(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
