package room

import (
	"net/http"

	"HR_Go/hr-admin-api/internal/logic/room"
	"HR_Go/hr-admin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllRoomHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := room.NewGetAllRoomLogic(r.Context(), svcCtx)
		resp, err := l.GetAllRoom()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
