package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserSocialLink 外部リンクテーブル
type UserSocialLink struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Platform  string    `gorm:"type:varchar(50);not null" json:"platform"`
	URL       string    `gorm:"type:text;not null" json:"url"`
	Title     *string   `gorm:"type:varchar(100)" json:"title"`
	IsPublic  bool      `gorm:"default:true" json:"is_public"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (usl *UserSocialLink) BeforeCreate(tx *gorm.DB) error {
	if usl.ID == uuid.Nil {
		usl.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (UserSocialLink) TableName() string {
	return "user_social_links"
}
