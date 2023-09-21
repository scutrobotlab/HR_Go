package sms

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/sms"
	"HR_Go/hr-admin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSmsConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sms.NewGetSmsConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetSmsConfig()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
