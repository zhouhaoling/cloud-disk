package models

import "time"

type UserRepository struct {
	Id                 int
	ParentId           int64
	Identity           string
	UserIdentity       string
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
