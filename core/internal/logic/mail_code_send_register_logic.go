package logic

import (
	"cloud/core/define"
	"cloud/core/models"
	"cloud/core/pkg"
	"context"
	"errors"
	"time"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	count, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}
	code := pkg.RandCode()
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	//发送验证码
	err = pkg.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
