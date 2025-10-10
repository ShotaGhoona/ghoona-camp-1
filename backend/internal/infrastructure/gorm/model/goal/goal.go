package goal

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Goal 目標管理テーブル
type Goal struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Title       string     `gorm:"type:varchar(200);not null" json:"title"`
	Description *string    `gorm:"type:text" json:"description"`
	StartedAt   time.Time  `gorm:"type:date;default:CURRENT_DATE" json:"started_at"`
	EndedAt     *time.Time `gorm:"type:date" json:"ended_at"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`
	IsPublic    bool       `gorm:"default:false" json:"is_public"`
	CreatedAt   time.Time  `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:now()" json:"updated_at"`
}

// BeforeCreate GORM hook - IDを自動生成
func (g *Goal) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (Goal) TableName() string {
	return "goals"
}