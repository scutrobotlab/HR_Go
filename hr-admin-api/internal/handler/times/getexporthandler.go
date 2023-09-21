package times

import (
	"io"
	"net/http"
	"os"

	"HR_Go/hr-admin-api/internal/logic/times"
	"HR_Go/hr-admin-api/internal/svc"
	"HR_Go/hr-admin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetExportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetExportReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := times.NewGetExportLogic(r.Context(), svcCtx)
		bytes, err := l.GetExport(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			tempDir, err := os.MkdirTemp("", "export")
			if err != nil {
				return
			}
			file, err := os.CreateTemp(tempDir, "export.csv")
			if err != nil {
				return
			}
			defer os.Remove(file.Name())
			_, err = io.WriteString(file, string(bytes))
			if err != nil {
				return
			}
			http.ServeFile(w, r, file.Name())
			return
		}
	}
}
