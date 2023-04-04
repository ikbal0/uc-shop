package models

import (
	"uc-shop/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FirstName string    `gorm:"not null" json:"first_name" form:"first_name" valid:"required~first name is required"`
	LastName  string    `json:"last_name" form:"last_name"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~email is required, email~email-Invalid email format"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required~Password password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Products  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL;"`
	Roles     []Role    `gorm:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL; many2many:user_roles"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
