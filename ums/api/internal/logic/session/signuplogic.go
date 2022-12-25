package session

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/svc"
	"github.com/shrinex/shrinex-admin-backend/ums-api/internal/types"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/pb"
	"github.com/shrinex/shrinex-core-backend/errx"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
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

func (l *SignUpLogic) SignUp(req *types.SignUpReq) (*types.SignUpResp, error) {
	err := preCheck(req)
	if err != nil {
		return nil, err
	}

	input := &pb.SignUpInput{}
	_ = copier.Copy(input, req)

	output, err := l.svcCtx.UmsClient.SignUp(l.ctx, input)
	if err != nil {
		// 用户已存在
		if errx.Is(err, 1024) {
			return nil, errors.WithStack(err)
		}
		// 其它错误
		return nil, errx.NewRegular("注册失败", errx.WithCause(err))
	}

	resp := &types.SignUpResp{}
	_ = copier.Copy(resp, output)
	return resp, nil
}

func preCheck(req *types.SignUpReq) error {
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	if len(req.Username) == 0 {
		return errx.NewValidation("用户名不能为空")
	}

	if len(req.Password) < 5 {
		return errx.NewValidation("密码不得低于5位")
	}

	// other constraints balabala...
	return nil
}
