package handler

import (
	"cloud/core/models"
	"cloud/core/pkg"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud/core/internal/logic"
	"cloud/core/internal/svc"
	"cloud/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		//判断文件是否存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		get, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if get {
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity, Name: rp.Name, Ext: rp.Ext})
			return
		}
		//向cos中存储文件
		cosPath, err := pkg.CosUpload(r)
		if err != nil {
			return
		}
		//往 logic中传递request
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
