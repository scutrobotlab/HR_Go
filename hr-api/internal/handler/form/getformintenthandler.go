package form

import (
	"net/http"

	"HR_Go/hr-api/internal/logic/form"
	"HR_Go/hr-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFormIntentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := form.NewGetFormIntentLogic(r.Context(), svcCtx)
		resp, err := l.GetFormIntent()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
