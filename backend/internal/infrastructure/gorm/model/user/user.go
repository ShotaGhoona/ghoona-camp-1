package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User ユーザー基本情報テーブル
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ClerkID   string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"clerk_id"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Username  *string   `gorm:"type:varchar(100)" json:"username"`
	AvatarURL *string   `gorm:"type:text" json:"avatar_url"`
	DiscordID *string   `gorm:"type:varchar(255);uniqueIndex" json:"discord_id"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`

	// Relationships (同一ドメイン内のみ定義、他ドメインは循環参照回避のため省略)
	Metadata    *UserMetadata    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"metadata,omitempty"`
	Vision      *UserVision      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"vision,omitempty"`
	SocialLinks []UserSocialLink `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"social_links,omitempty"`
	Rivals      []UserRival      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"rivals,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (User) TableName() string {
	return "users"
}