package handler

import (
	"errors"
	"fim_server/common/response"
	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"net/http"
	"os"
	"path"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		filePath := path.Join("uploads", req.ImageType, req.ImageName)
		byteData, err := os.ReadFile(filePath)
		if err != nil {
			response.Response(r, w, nil, errors.New("不存在该头像文件"))
			return
		}
		w.Write(byteData)
	}
}
