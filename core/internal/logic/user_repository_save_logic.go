package logic

import (
	"cloud/core/internal/svc"
	"cloud/core/internal/types"
	"cloud/core/models"
	"cloud/core/pkg"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	ur := new(models.UserRepository)
	ur = &models.UserRepository{
		ParentId:           req.ParentId,
		Identity:           pkg.GetUUID(),
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}

	return &types.UserRepositorySaveReply{
		Identity: ur.Identity,
	}, nil
}
