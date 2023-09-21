package config

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/config"
	"HR_Go/hr-admin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSmsServiceConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := config.NewGetSmsServiceConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetSmsServiceConfig()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
