// Code generated by goctl. DO NOT EDIT.
package types

type FileUploadChunkCompleteRequest struct {
	Key        string      `json:"key"`
	UploadId   string      `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}

type CosObject struct {
	PartNumber int    `json:"part_number"`
	Etag       string `json:"etag"`
}

type FileUploadChunkCompleteReply struct {
}

type FileUploadChunkRequest struct {
}

type FileUploadChunkReply struct {
	Etag string `json:"etag"` //md5
}

type FileUploadPrepareRequest struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

type FileUploadPrepareReply struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key      string `json:"key"`
}

type RefreshAuthorizationRequest struct {
}

type RefreshAuthorizationReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ShareBasicSaveRequest struct {
	ParentId           int64  `json:"parent_id"`
	RepositoryIdentity string `json:"repository_identity"`
}

type ShareBasicSaveReply struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailReply struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type ShareBasicCreateRequest struct {
	ExpiredTime            int    `json:"expired_time"`
	UserRepositoryIdentity string `json:"user_repository_identity"`
}

type ShareBasicCreateReply struct {
	Identity string `json:"identity"`
}

type UserFileMoveRequest struct {
	Identity       string `json:"identity"`
	ParentIdentity string `json:"parent_identity"`
}

type UserFileMoveReply struct {
}

type UserFileDeleteReply struct {
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFolderCreateRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFolderCreateReply struct {
	Identity string `json:"identity"`
}

type UserFileNameUpdateReply struct {
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileListRequest struct {
	Id   int64 `json:"id,optional"`
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type UserFileListReply struct {
	List  []*UserFile `json:"list"`
	Count int64       `json:"count"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	Size               int64  `json:"size"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveReply struct {
	Identity string `json:"identity"`
}

type FileUploadRequest struct {
	Size int64  `json:"size, optional"`
	Hash string `json:"hash, optional"`
	Name string `json:"name, optional"`
	Ext  string `json:"ext, optional"`
	Path string `json:"path, optional"`
}

type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterReply struct {
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserDetailReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendReply struct {
}
