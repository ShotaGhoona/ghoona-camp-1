package title

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TitleAchievement 称号実績テーブル
type TitleAchievement struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	TitleID     uuid.UUID `gorm:"type:uuid;not null;index" json:"title_id"`
	AchievedAt  time.Time `gorm:"default:now()" json:"achieved_at"`
	IsCurrent   bool      `gorm:"default:false" json:"is_current"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now()" json:"updated_at"`

	// Relationships
	Title Title `gorm:"foreignKey:TitleID;constraint:OnDelete:CASCADE" json:"title,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (ta *TitleAchievement) BeforeCreate(tx *gorm.DB) error {
	if ta.ID == uuid.Nil {
		ta.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (TitleAchievement) TableName() string {
	return "title_achievements"
}