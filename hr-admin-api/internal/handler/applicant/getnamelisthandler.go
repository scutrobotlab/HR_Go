package applicant

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/applicant"
	"HR_Go/hr-admin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetNameListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := applicant.NewGetNameListLogic(r.Context(), svcCtx)
		resp, err := l.GetNameList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
