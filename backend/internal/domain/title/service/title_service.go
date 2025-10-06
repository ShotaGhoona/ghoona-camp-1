package service

import (
	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title"
	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/value"
)

// TitleService は称号関連のビジネスロジックを処理する
type TitleService struct{}

// NewTitleService は新しいTitleServiceを作成する
func NewTitleService() *TitleService {
	return &TitleService{}
}

// CheckEligibleTitles は参加統計に基づいて獲得可能な称号を判定する
func (s *TitleService) CheckEligibleTitles(stats value.AttendanceStatistics, allTitles []*entity.Title) []*entity.Title {
	var eligibleTitles []*entity.Title
	
	for _, title := range allTitles {
		if title.IsEligibleFor(stats.TotalAttendanceDays) {
			eligibleTitles = append(eligibleTitles, title)
		}
	}
	
	return eligibleTitles
}

// GetNextTitle は現在のレベルから次に獲得可能な称号を取得する
func (s *TitleService) GetNextTitle(currentLevel int, allTitles []*entity.Title) *entity.Title {
	nextLevel := currentLevel + 1
	
	for _, title := range allTitles {
		if title.Level == nextLevel && title.IsActive.Bool() {
			return title
		}
	}
	
	return nil
}

// CalculateProgress は次の称号への進捗を計算する
func (s *TitleService) CalculateProgress(stats value.AttendanceStatistics, nextTitle *entity.Title) int {
	if nextTitle == nil {
		return 0
	}
	
	if stats.TotalAttendanceDays >= nextTitle.RequiredDays {
		return 100 // 既に達成済み
	}
	
	progress := (stats.TotalAttendanceDays * 100) / nextTitle.RequiredDays
	if progress > 100 {
		progress = 100
	}
	
	return progress
}

// GetHighestAchievedLevel は獲得済み称号の最高レベルを取得する
func (s *TitleService) GetHighestAchievedLevel(userAchievements []*entity.TitleAchievement, allTitles []*entity.Title) int {
	highestLevel := 0
	
	// 称号IDからレベルを取得するためのマップを作成
	titleLevelMap := make(map[common.UUID]int)
	for _, title := range allTitles {
		titleLevelMap[title.ID] = title.Level
	}
	
	// 獲得済み称号の最高レベルを検索
	for _, achievement := range userAchievements {
		if level, exists := titleLevelMap[achievement.TitleID]; exists {
			if level > highestLevel {
				highestLevel = level
			}
		}
	}
	
	return highestLevel
}

// ValidateCurrentTitleChange は現在表示称号変更の妥当性をチェックする
func (s *TitleService) ValidateCurrentTitleChange(userAchievements []*entity.TitleAchievement, targetTitleID common.UUID) error {
	// 獲得済み称号かどうかをチェック
	var targetAchievement *entity.TitleAchievement
	for _, achievement := range userAchievements {
		if achievement.TitleID == targetTitleID {
			targetAchievement = achievement
			break
		}
	}
	
	if targetAchievement == nil {
		return title.ErrTitleNotAchieved
	}
	
	// 既に現在の称号として設定済みかチェック
	if targetAchievement.IsCurrentTitle() {
		return title.ErrTitleAlreadyCurrent
	}
	
	return nil
}

// GetRemainingDaysToNextTitle は次の称号獲得まで残り何日かを計算する
func (s *TitleService) GetRemainingDaysToNextTitle(stats value.AttendanceStatistics, nextTitle *entity.Title) int {
	if nextTitle == nil {
		return 0
	}
	
	remaining := nextTitle.RequiredDays - stats.TotalAttendanceDays
	if remaining < 0 {
		remaining = 0
	}
	
	return remaining
}