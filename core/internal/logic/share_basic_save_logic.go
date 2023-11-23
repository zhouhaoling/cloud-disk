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

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {
	rp := new(models.RepositoryPool)
	get, err := l.svcCtx.Engine.Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("资源不存在")
	}
	ur := &models.UserRepository{
		ParentId:           req.ParentId,
		Identity:           pkg.GetUUID(),
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return
	}
	resp = &types.ShareBasicSaveReply{
		Identity: ur.Identity,
	}
	return
}
