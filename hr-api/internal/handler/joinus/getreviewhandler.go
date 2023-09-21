package joinus

import (
	"net/http"

	"HR_Go/hr-api/internal/logic/joinus"
	"HR_Go/hr-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetReviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := joinus.NewGetReviewLogic(r.Context(), svcCtx)
		resp, err := l.GetReview()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
