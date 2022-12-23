package session

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/svc"
	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/types"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/pb"
	"github.com/shrinex/shrinex-core-backend/errx"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	input := &pb.SignUpInput{}
	_ = copier.Copy(input, req)

	output, err := l.svcCtx.UmsClient.SignUp(l.ctx, input)
	if err != nil {
		return nil, fmt.Errorf("%w: call UmsClient#SignUp() failed, err: %v",
			errx.New(4001, "注册失败"), err)
	}

	_ = copier.Copy(resp, output)
	return
}
