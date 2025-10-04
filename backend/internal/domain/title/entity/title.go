package entity

import (
	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title/value"
)

// Title は称号の基本情報を表す
type Title struct {
	common.BaseEntity            // 共通フィールド (ID, CreatedAt, UpdatedAt)
	Level        int             // レベル (1-8)
	NameJP       string          // 日本語名
	NameEN       string          // 英語名
	Description  string          // 称号の説明・ストーリー
	RequiredDays int             // 獲得に必要な参加日数
	ImageURL     *string         // 称号カード画像URL（オプショナル）
	ColorTheme   *string         // テーマカラー（オプショナル）
	IsActive     value.ActiveFlag // 称号の有効状態
}

// NewTitle は新しいTitleエンティティを作成する
func NewTitle(level int, nameJP, nameEN, description string, requiredDays int) *Title {
	return &Title{
		BaseEntity:   common.NewBaseEntity(),
		Level:        level,
		NameJP:       nameJP,
		NameEN:       nameEN,
		Description:  description,
		RequiredDays: requiredDays,
		IsActive:     value.ActiveTrue,
	}
}

// IsEligibleFor は指定された参加日数で獲得可能かどうかを判定する
func (t *Title) IsEligibleFor(attendanceDays int) bool {
	return t.IsActive.Bool() && attendanceDays >= t.RequiredDays
}

// IsValidLevel は称号レベルが有効範囲かどうかを確認する
func (t *Title) IsValidLevel() bool {
	return t.Level >= 1 && t.Level <= 8
}