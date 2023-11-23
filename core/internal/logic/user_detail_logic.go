package logic

import (
	"cloud/core/models"
	"context"
	"errors"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserDetailReply{}
	ub := new(models.UserBasic)
	get, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(ub)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("用户不存在")
	}
	resp.Name = ub.Name
	resp.Email = ub.Email
	return
}
