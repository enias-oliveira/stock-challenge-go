package domain

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Email        string `gorm:"type:varchar(100);unique;uniqueIndex;"`
	Password     string `gorm:"-"`
	PasswordHash []byte `gorm:"type:binary(32);"`
	Role         string `gorm:"type:enum('admin','user');default:'user'"`
}
