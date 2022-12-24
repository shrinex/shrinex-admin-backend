package logic

import (
	"context"
	"errors"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/dao"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/internal/svc"
	"github.com/shrinex/shrinex-admin-backend/ums-rpc/pb"
	"github.com/shrinex/shrinex-core-backend/errx"
	"golang.org/x/crypto/bcrypt"

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

func (l *SignUpLogic) SignUp(input *pb.SignUpInput) (*pb.SignUpOutput, error) {
	_, err := l.svcCtx.AdminUserDao.FindOneByNickname(l.ctx, input.GetUsername())
	if err == nil {
		return nil, errx.New(1024, "用户已存在")
	}
	if !errors.Is(err, dao.ErrNotFound) {
		l.Logger.Errorf("查询用户失败: %v", err)
		return nil, errx.New(errx.DataAccess, errx.DataAccessDesc)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		l.Logger.Errorf("run bcrypt failed: %v", err)
		return nil, errx.New(errx.Regular, "加密失败")
	}

	user := &dao.UmsAdminUser{
		Nickname: input.GetUsername(),
		Password: string(hashedPassword),
	}
	_, err = l.svcCtx.AdminUserDao.Insert(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("新增用户失败: %v", err)
		return nil, errx.New(errx.DataAccess, errx.DataAccessDesc)
	}

	return &pb.SignUpOutput{}, nil
}
