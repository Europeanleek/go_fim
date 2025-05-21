package handler

import (
	"fim_server/common/response"
	"fim_server/fim_auth/api/internal/logic"
	"fim_server/fim_auth/api/internal/svc"
	"net/http"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication(token)
		response.Response(r, w, resp, err)

	}
}
