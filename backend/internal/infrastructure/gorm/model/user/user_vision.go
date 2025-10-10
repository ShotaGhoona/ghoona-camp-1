package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserVision ユーザービジョンテーブル
type UserVision struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	Vision    *string   `gorm:"type:text" json:"vision"`
	IsPublic  bool      `gorm:"default:false" json:"is_public"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (uv *UserVision) BeforeCreate(tx *gorm.DB) error {
	if uv.ID == uuid.Nil {
		uv.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (UserVision) TableName() string {
	return "user_visions"
}
