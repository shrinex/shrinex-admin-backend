package session

import (
	"github.com/shrinex/shrinex-core-backend/result"
	"net/http"

	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/logic/session"
	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/svc"
	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignUpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignUpReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Compute(r, w, nil, err)
			return
		}

		l := session.NewSignUpLogic(r.Context(), svcCtx)
		resp, err := l.SignUp(&req)
		result.Compute(r, w, resp, err)
	}
}
