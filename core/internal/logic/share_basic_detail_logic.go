package logic

import (
	"context"
	"errors"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	_, err = l.svcCtx.Engine.Exec("update share_basic set	click_num = click_num + 1 where identity = ?", req.Identity)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	resp = new(types.ShareBasicDetailReply)
	_, err = l.svcCtx.Engine.Table("share_basic").
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("left", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Join("left", "user_repository", "user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ?", req.Identity).Get(resp)
	if err != nil {
		return nil, errors.New("获取资源详情失败")
	}
	return resp, nil
}
