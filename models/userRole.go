package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type UserRole struct {
	GormModel
	UserID uint `gorm:"primaryKey" json:"user_id" form:"user_id" valid:"required~data is required"`
	RoleID uint `gorm:"primaryKey" json:"role_id" form:"role_id" valid:"required~data is required"`
}

func (r *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(r)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
