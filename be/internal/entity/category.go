package entity

import "github.com/google/uuid"

type Category struct {
	Base

	Name        string  `gorm:"type:varchar(128); not null" validate:"required,max=128" json:"name"`
	Img 		string `gorm:"type:varchar(10485760); not null" json:"img"`
	UserId 		uuid.UUID  `gorm:"type:uuid; not null" json:"userId"`
}