package title

import (
	"ghoona-camp-backend/internal/domain/common"
)

// UserTitleProgressResponse はユーザーの称号進捗レスポンスDTO
type UserTitleProgressResponse struct {
	UserID           common.UUID          `json:"userId"`
	CurrentTitle     *AchievementResponse `json:"currentTitle"`
	NextTitle        *TitleResponse       `json:"nextTitle"`
	CurrentDays      int                  `json:"currentDays"`
	RequiredDays     int                  `json:"requiredDays"`
	RemainingDays    int                  `json:"remainingDays"`
	ProgressPercent  float64              `json:"progressPercent"`
	HighestLevel     int                  `json:"highestLevel"`
	EligibleTitles   []TitleResponse      `json:"eligibleTitles"`
}

// BuildUserTitleProgressResponse は称号進捗レスポンスを構築する
func BuildUserTitleProgressResponse(
	userID common.UUID,
	currentTitle *AchievementResponse,
	nextTitle *TitleResponse,
	currentDays int,
	highestLevel int,
	eligibleTitles []TitleResponse,
) *UserTitleProgressResponse {
	var requiredDays int
	var remainingDays int
	var progressPercent float64

	if nextTitle != nil {
		requiredDays = nextTitle.RequiredDays
		remainingDays = requiredDays - currentDays
		if remainingDays < 0 {
			remainingDays = 0
		}
		if requiredDays > 0 {
			progressPercent = float64(currentDays) / float64(requiredDays) * 100
			if progressPercent > 100 {
				progressPercent = 100
			}
		}
	}

	return &UserTitleProgressResponse{
		UserID:           userID,
		CurrentTitle:     currentTitle,
		NextTitle:        nextTitle,
		CurrentDays:      currentDays,
		RequiredDays:     requiredDays,
		RemainingDays:    remainingDays,
		ProgressPercent:  progressPercent,
		HighestLevel:     highestLevel,
		EligibleTitles:   eligibleTitles,
	}
}