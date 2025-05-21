package handler

import (
	"fim_server/common/response"
	"fim_server/fim_auth/api/internal/logic"
	"fim_server/fim_auth/api/internal/svc"
	"fmt"
	"net/http"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		fmt.Println("token is " + token)
		resp, err := l.Logout(token)
		response.Response(r, w, resp, err)
	}
}
