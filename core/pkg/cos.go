package pkg

import (
	"bytes"
	"cloud/core/define"
	"context"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// CosUpload 上传文件到腾讯云
func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + GetUUID() + path.Ext(fileHeader.Filename)
	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		return "", err
	}
	return define.CosBucket + "/" + key, nil
}

// CosInitPart 分片上传初始化
func CosInitPart(ext string) (string, string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	key := "cloud-disk/" + GetUUID() + ext
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	UploadID := v.UploadID
	return key, UploadID, nil
}

// CosPartUpload 分片上传
func CosPartUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	key := r.PostForm.Get("key")
	UploadID := r.PostForm.Get("upload_id")
	partNumber, err := strconv.Atoi(r.PostForm.Get("part_number"))
	if err != nil {
		return "", err
	}
	f, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)

	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, partNumber, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}
	return strings.Trim(resp.Header.Get("ETag"), "\""), nil
}

// CosFileUploadComplete 分片上传完成
func CosFileUploadComplete(key, uploadId string, cs []cos.Object) error {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cs...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	if err != nil {
		return err
	}
	return nil
}
