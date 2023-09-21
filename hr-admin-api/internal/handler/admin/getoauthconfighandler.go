package admin

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/admin"
	"HR_Go/hr-admin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOAuthConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewGetOAuthConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetOAuthConfig()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
