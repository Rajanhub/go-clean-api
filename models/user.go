package models

import (
	"time"

	"github.com/Rajanhub/goapi/lib"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	lib.ModelBase
	Name       string        `json:"name" form:"name" binding:"required"`
	Email      string        `json:"email" form:"email" binding:"required" gorm:"unique"`
	ProfilePic lib.SignedURL `json:"profile_pic"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (t *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = lib.BinaryUUID(id)
	return err
}
