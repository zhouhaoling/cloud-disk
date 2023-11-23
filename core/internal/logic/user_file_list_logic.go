package logic

import (
	"cloud/core/define"
	"cloud/core/models"
	"context"
	"time"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserFileListReply)
	//分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext, user_repository.name,"+
			"repository_pool.size, repository_pool.path").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? or user_repository.deleted_at is null", time.Time{}.Format(define.DataTime)).
		Limit(size, offset).Find(&uf)
	if err != nil {
		return
	}
	//查询用户文件总数
	count, err := l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count(new(models.UserRepository))
	if err != nil {
		return
	}
	resp.List = uf
	resp.Count = count

	return
}
