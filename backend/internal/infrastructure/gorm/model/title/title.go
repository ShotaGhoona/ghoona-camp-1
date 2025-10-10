package title

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Title 称号マスターテーブル
type Title struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Level        int       `gorm:"type:integer;uniqueIndex;not null" json:"level"`
	NameJP       string    `gorm:"type:varchar(50);not null" json:"name_jp"`
	NameEN       string    `gorm:"type:varchar(50);not null" json:"name_en"`
	Description  string    `gorm:"type:text;not null" json:"description"`
	RequiredDays int       `gorm:"type:integer;not null" json:"required_days"`
	ImageURL     *string   `gorm:"type:text" json:"image_url"`
	ColorTheme   *string   `gorm:"type:varchar(50)" json:"color_theme"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:now()" json:"updated_at"`

	// Relationships
	Achievements []TitleAchievement `gorm:"foreignKey:TitleID;constraint:OnDelete:CASCADE" json:"achievements,omitempty"`
}

// BeforeCreate GORM hook - IDを自動生成
func (t *Title) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

// TableName テーブル名を指定
func (Title) TableName() string {
	return "titles"
}