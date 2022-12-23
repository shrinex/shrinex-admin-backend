package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/internal/svc"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpLogic) SignUp(input *pb.SignUpInput) (output *pb.SignUpOutput, err error) {
	fmt.Println(uuid.NewString())
	return &pb.SignUpOutput{}, nil
}
