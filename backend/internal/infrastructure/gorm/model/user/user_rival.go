package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRival ライバル関係テーブル
type UserRival struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index:idx_user_rival_unique" json:"user_id"`
	RivalUserID uuid.UUID `gorm:"type:uuid;not null;index:idx_user_rival_unique" json:"rival_user_id"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now()" json:"updated_at"`

	// Relationships
	User      User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	RivalUser User `gorm:"foreignKey:RivalUserID;constraint:OnDelete:CASCADE" json:"rival_user,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (ur *UserRival) BeforeCreate(tx *gorm.DB) error {
	if ur.ID == uuid.Nil {
		ur.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (UserRival) TableName() string {
	return "user_rivals"
}
