package models

import (
	"time"

	"github.com/Rajanhub/goapi/lib"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	lib.ModelBase
	Title       string         `json:"title" form:"title"`
	Description string         `json:"description" form:"description"`
	UserID      lib.BinaryUUID `json:"userId" form:"userId"`
	User        *User          `binding:"-" json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (t *Post) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = lib.BinaryUUID(id)
	return err
}
