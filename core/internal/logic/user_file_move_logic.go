package logic

import (
	"cloud/core/models"
	"context"
	"errors"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	parentDate := new(models.UserRepository)
	get, err := l.svcCtx.Engine.Where("identity = ? and user_identity = ?", req.ParentIdentity, userIdentity).Get(parentDate)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("文件夹不存在")
	}

	//更新记录的parentid
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(models.UserRepository{
		ParentId: int64(parentDate.Id),
	})
	if err != nil {
		return nil, err
	}
	return
}
