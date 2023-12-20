package dao

import (
	"blog/model"

	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *model.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {

}
