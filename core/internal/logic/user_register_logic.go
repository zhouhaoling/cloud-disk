package logic

import (
	"cloud/core/models"
	"cloud/core/pkg"
	"context"
	"errors"
	"log"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	//判断验证码是否错误
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取到该邮箱验证码")
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		return
	}
	//判断用户是否存在
	count, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("用户已存在")
		return
	}
	//注册用户入库
	user := &models.UserBasic{
		Identity: pkg.GetUUID(),
		Name:     req.Name,
		Password: pkg.Md5(req.Password),
		Email:    req.Email,
	}
	insert, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row:", insert)
	return
}
