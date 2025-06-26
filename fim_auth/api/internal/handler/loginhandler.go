package handler

import (
	"bytes"
	"encoding/json"
	"fim_server/common/response"
	"fim_server/fim_auth/api/internal/logic"
	"fim_server/fim_auth/api/internal/svc"
	"fim_server/fim_auth/api/internal/types"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type ReqData struct {
	UserName string `json:"userName"`
	Passward string `json:"password"`
}

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		fmt.Println(r.ContentLength)
		fmt.Println(r.Method)
		forwardedBody, _ := io.ReadAll(r.Body)
		fmt.Println(forwardedBody)
		r.Body = io.NopCloser(bytes.NewBuffer(forwardedBody))
		var data map[string]interface{}
		if err := json.Unmarshal(forwardedBody, &data); err != nil {
			log.Printf("非法 JSON 数据: %v", err)
			// 返回错误响应
		} else {
			log.Println("请求体是合法 JSON")
		}
		fmt.Println(data)
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		l := logic.NewLoginLogic(r.Context(), svcCtx)
		fmt.Println("login函数执行")
		resp, err := l.Login(&req)
		response.Response(r, w, resp, err)
	}
}
