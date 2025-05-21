package handler

import (
	"fim_server/common/response"
	"fim_server/fim_auth/api/internal/logic"
	"fim_server/fim_auth/api/internal/svc"
	"net/http"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication()
		response.Response(r, w, resp, err)

	}
}
