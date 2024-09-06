package handler

import (
	"net/http"

	"github.com/asjard/benchmark/servers/bench_go_zero/internal/logic"
	"github.com/asjard/benchmark/servers/bench_go_zero/internal/svc"
	"github.com/asjard/benchmark/servers/bench_go_zero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Go_zeroHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGo_zeroLogic(r.Context(), svcCtx)
		resp, err := l.Go_zero(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
