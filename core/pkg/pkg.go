package pkg

import (
	"crypto/md5"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GetUUID() string {
	return uuid.NewV4().String()
}
