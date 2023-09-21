package form

import (
	"net/http"

	"HR_Go/hr-api/internal/logic/form"
	"HR_Go/hr-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFormGroupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := form.NewGetFormGroupsLogic(r.Context(), svcCtx)
		resp, err := l.GetFormGroups()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
