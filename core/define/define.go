package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

// MailPassword 邮箱密钥
var MailPassword = "MyMailPassword"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 过期时间
var CodeExpire = 400

// TencentSecretID 存储桶id
var TencentSecretID = "MySecretID"

// TencentSecretKey 存储桶key
var TencentSecretKey = "MySecretKey"

// CosBucket 上传地址
var CosBucket = "https://zhl-1937468-1322349519.cos.ap-guangzhou.myqcloud.com/"

// PageSize 分页默认参数
var PageSize = 20

// DataTime 时期
var DataTime = "2006-01-02 15:04:05"

// TokenExpire token过期时间
var TokenExpire = 3600

// RefreshTokenExpire refreshToken过期时间
var RefreshTokenExpire = 7200
