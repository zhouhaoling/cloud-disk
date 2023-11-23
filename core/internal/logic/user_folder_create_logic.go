package logic

import (
	"cloud/core/internal/svc"
	"cloud/core/internal/types"
	"cloud/core/models"
	"cloud/core/pkg"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	count, err := l.svcCtx.Engine.Where("name = ? AND parent_id = ?", req.Name, req.ParentId).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("文件夹名称已存在")
	}
	folder := &models.UserRepository{
		Identity:     pkg.GetUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(folder)
	if err != nil {
		return
	}
	return &types.UserFolderCreateReply{
		Identity: folder.Identity,
	}, nil
}
