package logic

import (
	"cloud/core/define"
	"cloud/core/models"
	"cloud/core/pkg"
	"context"
	"errors"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	user := new(models.UserBasic)
	get, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, pkg.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("用户名或密码错误")
	}
	token, err := pkg.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	//生成refreshToken
	refreshToken, err := pkg.GenerateToken(user.Id, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
