package handler

import (
	"errors"
	"fim_server/common/response"
	"fim_server/fim_file/file_api/internal/logic"
	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"fim_server/utils/common"
	comparelist "fim_server/utils/compare_list"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		//获取上传图片，以及图片属性
		file, fileHead, err := r.FormFile("images")
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		imageType := r.FormValue("imageType")
		switch imageType {
		case "avatar", "group_avatar", "chat":
		default:
			response.Response(r, w, nil, errors.New("imageType只能为avatar，group_avatar，chat"))
			return
		}
		//上传文件大小限制
		mSize := float64(fileHead.Size) / float64(1024) / float64(1024)

		if mSize > svcCtx.Config.FileSize {
			response.Response(r, w, nil, fmt.Errorf("上传图片大小超过限制，只能上传%.2f大小的图片", svcCtx.Config.FileSize))
			return
		}
		//文件后缀白名单
		nameList := strings.Split(fileHead.Filename, ".")
		var suffix string
		if len(nameList) > 1 {
			suffix = nameList[len(nameList)-1]
		}
		if !comparelist.ContainsString(svcCtx.Config.WhiteList, suffix) {
			response.Response(r, w, nil, errors.New("图片属于非法后缀"))
			return
		}
		fileName := fileHead.Filename
		filePath := path.Join(svcCtx.Config.UploadPath, imageType, fileName)
		imageData, _ := io.ReadAll(file)
		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)
		resp.Url = "/" + filePath
		//文件重名
		// 在保存用户上传的文件之前，尝试去读文件列表，如果用重名的文件，计算两个文件的哈希值，文件一样的情况就不需要重新写了
		// 如果哈希值不一样，就把最新的文件重新命名一下
		dirPath := path.Join(svcCtx.Config.UploadPath, imageType)
		dir, err := os.ReadDir(dirPath)
		if err != nil {
			os.Mkdir(dirPath, 0666)
		}
		if InDir(dir, fileHead.Filename) {
			//如果有重名文件
			//todo: 计算哈希值，如果哈希值相同，就不需要重新写文件了
			byteData, _ := os.ReadFile(filePath)
			oldFileHash := common.MD5(byteData)
			newFileHash := common.MD5(imageData)
			if oldFileHash == newFileHash {
				logx.Info("两个图片是一样的")
				response.Response(r, w, resp, nil)
				return
			}
			//两个文件不一样，进行改名保存
			var prefix = common.GetFilePrefix(fileHead.Filename)
			newPath := fmt.Sprintf("%s_%s.%s", prefix, common.RandStr(4), suffix)
			filePath = path.Join(dirPath, newPath)
		}
		err = os.WriteFile(filePath, imageData, 0666)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		resp.Url = "/" + filePath
		response.Response(r, w, resp, err)
	}
}

func InDir(dir []os.DirEntry, filename string) (ok bool) {
	for _, entry := range dir {
		if entry.Name() == filename && !entry.IsDir() {
			//存在相同文件，返回true
			return true
		}
	}
	//不存在相同的文件，返回false
	return false
}
