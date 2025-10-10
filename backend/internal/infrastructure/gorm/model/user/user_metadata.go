package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// UserMetadata ユーザー詳細情報テーブル
type UserMetadata struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	DisplayName *string        `gorm:"type:varchar(100)" json:"display_name"`
	Tagline     *string        `gorm:"type:varchar(150)" json:"tagline"`
	Bio         *string        `gorm:"type:text" json:"bio"`
	Skills      pq.StringArray `gorm:"type:text[]" json:"skills"`
	Interests   pq.StringArray `gorm:"type:text[]" json:"interests"`
	CreatedAt   time.Time      `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:now()" json:"updated_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (um *UserMetadata) BeforeCreate(tx *gorm.DB) error {
	if um.ID == uuid.Nil {
		um.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (UserMetadata) TableName() string {
	return "user_metadata"
}
