package logic

import (
	"cloud/core/models"
	"cloud/core/pkg"
	"context"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {
	rp := new(models.RepositoryPool)
	resp = new(types.FileUploadPrepareReply)
	get, err := l.svcCtx.Engine.Where("hash = ?", req.Md5).Get(rp)
	if err != nil {
		return
	}
	if get {
		//妙传成功
		resp.Identity = rp.Identity
	} else {
		// TODO: 获取该文件的uploadID，用来进行文件的分片上传
		key, uploadId, err := pkg.CosInitPart(req.Ext)
		if err != nil {
			return nil, err
		}
		resp.Key = key
		resp.UploadId = uploadId
	}
	return
}
