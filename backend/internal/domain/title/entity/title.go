package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"ghoona-camp-backend/internal/domain/title/vo"
	titlevo "ghoona-camp-backend/internal/domain/title/vo"
)

// Title 称号エンティティ
type Title struct {
	id           uuid.UUID           // 称号ID
	level        vo.TitleLevel       // レベル (1-8)
	nameJP       string              // 日本語名
	nameEN       string              // 英語名
	description  string              // 称号の説明・ストーリー
	requiredDays int                 // 獲得に必要な参加日数
	imageURL     string              // 称号カード画像URL
	colorTheme   string              // テーマカラー
	isActive     titlevo.ActiveFlag   // 称号の有効状態
	createdAt    time.Time           // 作成日時
	updatedAt    time.Time           // 更新日時
}

// NewTitle Titleエンティティを作成する
func NewTitle(
	level vo.TitleLevel,
	nameJP string,
	nameEN string,
	description string,
	requiredDays int,
	imageURL string,
	colorTheme string,
	isActive titlevo.ActiveFlag,
) (*Title, error) {
	// バリデーション
	if nameJP == "" {
		return nil, errors.New("日本語名は必須です")
	}
	if nameEN == "" {
		return nil, errors.New("英語名は必須です")
	}
	if description == "" {
		return nil, errors.New("説明は必須です")
	}
	if requiredDays < 0 {
		return nil, errors.New("必要日数は0以上で指定してください")
	}
	
	// レベル・必要日数妥当性チェック
	if !level.IsValid() {
		return nil, errors.New("称号レベルが不正です")
	}

	now := time.Now()
	return &Title{
		id:           uuid.New(),
		level:        level,
		nameJP:       nameJP,
		nameEN:       nameEN,
		description:  description,
		requiredDays: requiredDays,
		imageURL:     imageURL,
		colorTheme:   colorTheme,
		isActive:     isActive,
		createdAt:    now,
		updatedAt:    now,
	}, nil
}

// Getters
func (t *Title) ID() uuid.UUID              { return t.id }
func (t *Title) Level() vo.TitleLevel       { return t.level }
func (t *Title) NameJP() string             { return t.nameJP }
func (t *Title) NameEN() string             { return t.nameEN }
func (t *Title) Description() string        { return t.description }
func (t *Title) RequiredDays() int          { return t.requiredDays }
func (t *Title) ImageURL() string           { return t.imageURL }
func (t *Title) ColorTheme() string         { return t.colorTheme }
func (t *Title) IsActive() titlevo.ActiveFlag { return t.isActive }
func (t *Title) CreatedAt() time.Time       { return t.createdAt }
func (t *Title) UpdatedAt() time.Time       { return t.updatedAt }