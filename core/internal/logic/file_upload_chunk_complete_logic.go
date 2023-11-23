package logic

import (
	"cloud/core/pkg"
	"context"

	"github.com/tencentyun/cos-go-sdk-v5"

	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteRequest) (resp *types.FileUploadChunkCompleteReply, err error) {
	cs := make([]cos.Object, 0)
	for _, v := range req.CosObjects {
		cs = append(cs, cos.Object{
			ETag:       v.Etag,
			PartNumber: v.PartNumber,
		})
	}
	err = pkg.CosFileUploadComplete(req.Key, req.UploadId, cs)
	return
}
