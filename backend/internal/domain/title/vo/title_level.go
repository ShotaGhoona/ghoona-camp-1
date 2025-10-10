package vo

import "errors"

// TitleLevel 称号のレベル（1-8の固定値）を表すValue Object
type TitleLevel int

const (
	TitleLevel1 TitleLevel = 1 // まどろみ見習い（Sleeper） - 1日
	TitleLevel2 TitleLevel = 2 // 早起き候補生（Early Bird Candidate） - 7日
	TitleLevel3 TitleLevel = 3 // 朝活探検家（Morning Explorer） - 30日
	TitleLevel4 TitleLevel = 4 // 朝の住人（Morning Resident） - 100日
	TitleLevel5 TitleLevel = 5 // 夜明けの戦士（Dawn Warrior） - 200日
	TitleLevel6 TitleLevel = 6 // 朝光の使者（Morning Light Messenger） - 365日
	TitleLevel7 TitleLevel = 7 // 暁の守護者（Dawn Guardian） - 500日
	TitleLevel8 TitleLevel = 8 // 朝活の伝説（Morning Legend） - 1000日
)

const (
	MinTitleLevel = 1
	MaxTitleLevel = 8
)

// NewTitleLevel TitleLevelを作成する
func NewTitleLevel(value int) (TitleLevel, error) {
	if value < MinTitleLevel || value > MaxTitleLevel {
		return 0, errors.New("称号レベルは1-8の範囲で指定してください")
	}
	return TitleLevel(value), nil
}

// Value 内部の整数値を返す
func (l TitleLevel) Value() int {
	return int(l)
}

// IsValid 有効なレベル値かどうかを検証する
func (l TitleLevel) IsValid() bool {
	return int(l) >= MinTitleLevel && int(l) <= MaxTitleLevel
}