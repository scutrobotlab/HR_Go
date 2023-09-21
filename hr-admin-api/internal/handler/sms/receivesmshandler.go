package sms

import (
	"net/http"
	"strconv"

	"HR_Go/hr-admin-api/internal/logic/sms"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReceiveSmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReceiveSmsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sms.NewReceiveSmsLogic(r.Context(), svcCtx)
		resp, err := l.ReceiveSms(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			_, err := w.Write([]byte(strconv.Itoa(int(resp.Code))))
			if err != nil {
				return
			}
		}
	}
}
