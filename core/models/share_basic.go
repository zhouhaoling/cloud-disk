package models

import "time"

type ShareBasic struct {
	Id                     int
	ExpiredTime            int
	ClickNum               int
	Identity               string
	UserIdentity           string
	UserRepositoryIdentity string
	RepositoryIdentity     string
	CreatedAt              time.Time `xorm:"created"`
	UpdatedAt              time.Time `xorm:"updated"`
	DeletedAt              time.Time `xorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
