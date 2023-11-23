package logic

import (
	"cloud/core/models"
	"context"
	"errors"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	//该名称是否已存在
	count, err := l.svcCtx.Engine.Where("name = ? AND parent_id = (select parent_id from user_repository ur where ur.identity = ?)", req.Name, req.Identity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("文件名称已存在")
	}
	data := models.UserRepository{Name: req.Name}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return
	}
	return
}
