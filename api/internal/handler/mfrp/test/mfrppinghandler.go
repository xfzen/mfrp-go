package test

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mfrp/api/internal/logic/mfrp/test"
	"mfrp/api/internal/svc"
)

func MfrpPingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewMfrpPingLogic(r.Context(), svcCtx)
		err := l.MfrpPing()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
