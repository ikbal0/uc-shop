package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Role struct {
	GormModel
	Name        string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required~description is required"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(r)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
