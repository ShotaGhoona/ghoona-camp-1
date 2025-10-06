package model

import (
	"time"

	"github.com/google/uuid"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/value"
)

// Title は称号基本情報のGORMモデル
type Title struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Level        int        `gorm:"unique;not null;check:level >= 1 AND level <= 8" json:"level"`
	NameJP       string     `gorm:"not null;size:100" json:"nameJp"`
	NameEN       string     `gorm:"not null;size:100" json:"nameEn"`
	Description  string     `gorm:"not null;type:text" json:"description"`
	RequiredDays int        `gorm:"not null;check:required_days >= 1" json:"requiredDays"`
	ImageURL     *string    `gorm:"type:text" json:"imageUrl"`
	ColorTheme   *string    `gorm:"size:50" json:"colorTheme"`
	IsActive     bool       `gorm:"default:true" json:"isActive"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

// TitleAchievement はユーザー称号獲得実績のGORMモデル
type TitleAchievement struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`
	TitleID    uuid.UUID `gorm:"type:uuid;not null;index" json:"titleId"`
	AchievedAt time.Time `gorm:"not null" json:"achievedAt"`
	IsCurrent  bool      `gorm:"default:false;index" json:"isCurrent"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	
	// 外部キー関係（リレーション定義）
	Title Title `gorm:"foreignKey:TitleID" json:"title"`
	User  User  `gorm:"foreignKey:UserID" json:"user"`
}

// TableName メソッドでテーブル名を指定
func (Title) TableName() string {
	return "titles"
}

func (TitleAchievement) TableName() string {
	return "title_achievements"
}

// ToEntity はGORMモデルからドメインエンティティへ変換します
func (t *Title) ToEntity() (*entity.Title, error) {
	activeFlag := value.ActiveTrue
	if !t.IsActive {
		activeFlag = value.ActiveFalse
	}

	return &entity.Title{
		BaseEntity: common.BaseEntity{
			ID:        common.UUID(t.ID),
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		},
		Level:        t.Level,
		NameJP:       t.NameJP,
		NameEN:       t.NameEN,
		Description:  t.Description,
		RequiredDays: t.RequiredDays,
		ImageURL:     t.ImageURL,
		ColorTheme:   t.ColorTheme,
		IsActive:     activeFlag,
	}, nil
}

// FromEntity はドメインエンティティからGORMモデルへ変換します
func FromEntityTitle(domainTitle *entity.Title) *Title {
	return &Title{
		ID:           uuid.UUID(domainTitle.ID),
		Level:        domainTitle.Level,
		NameJP:       domainTitle.NameJP,
		NameEN:       domainTitle.NameEN,
		Description:  domainTitle.Description,
		RequiredDays: domainTitle.RequiredDays,
		ImageURL:     domainTitle.ImageURL,
		ColorTheme:   domainTitle.ColorTheme,
		IsActive:     domainTitle.IsActive.Bool(),
		CreatedAt:    domainTitle.CreatedAt,
		UpdatedAt:    domainTitle.UpdatedAt,
	}
}

// ToEntity はGORMモデルからドメインエンティティへ変換します
func (ta *TitleAchievement) ToEntity() (*entity.TitleAchievement, error) {
	currentFlag := value.CurrentFalse
	if ta.IsCurrent {
		currentFlag = value.CurrentTrue
	}

	return &entity.TitleAchievement{
		BaseEntity: common.BaseEntity{
			ID:        common.UUID(ta.ID),
			CreatedAt: ta.CreatedAt,
			UpdatedAt: ta.UpdatedAt,
		},
		UserID:     common.UUID(ta.UserID),
		TitleID:    common.UUID(ta.TitleID),
		AchievedAt: ta.AchievedAt,
		IsCurrent:  currentFlag,
	}, nil
}

// FromEntityTitleAchievement はドメインエンティティからGORMモデルへ変換します
func FromEntityTitleAchievement(domainAchievement *entity.TitleAchievement) *TitleAchievement {
	return &TitleAchievement{
		ID:         uuid.UUID(domainAchievement.ID),
		UserID:     uuid.UUID(domainAchievement.UserID),
		TitleID:    uuid.UUID(domainAchievement.TitleID),
		AchievedAt: domainAchievement.AchievedAt,
		IsCurrent:  domainAchievement.IsCurrent.Bool(),
		CreatedAt:  domainAchievement.CreatedAt,
		UpdatedAt:  domainAchievement.UpdatedAt,
	}
}