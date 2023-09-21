package scores

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/scores"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRankReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := scores.NewGetRankLogic(r.Context(), svcCtx)
		resp, err := l.GetRank(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
