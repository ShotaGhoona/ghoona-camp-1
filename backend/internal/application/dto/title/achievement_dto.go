package title

import (
	"time"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title/entity"
)

// AchievementResponse は称号獲得記録のレスポンスDTO
type AchievementResponse struct {
	ID          common.UUID   `json:"id"`
	UserID      common.UUID   `json:"userId"`
	Title       TitleResponse `json:"title"`
	AchievedAt  time.Time     `json:"achievedAt"`
	IsCurrent   bool          `json:"isCurrent"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

// AchievementListResponse はユーザー称号獲得履歴のレスポンスDTO
type AchievementListResponse struct {
	UserID       common.UUID           `json:"userId"`
	Achievements []AchievementResponse `json:"achievements"`
	CurrentTitle *AchievementResponse  `json:"currentTitle"`
	Total        int                   `json:"total"`
}

// SetCurrentTitleRequest は現在称号変更のリクエストDTO
type SetCurrentTitleRequest struct {
	IsCurrent bool `json:"isCurrent" binding:"required"`
}

// AchievementResponseFromEntity はエンティティからレスポンスDTOに変換する
func AchievementResponseFromEntity(achievement *entity.TitleAchievement, title *entity.Title) *AchievementResponse {
	if achievement == nil || title == nil {
		return nil
	}

	return &AchievementResponse{
		ID:          achievement.ID,
		UserID:      achievement.UserID,
		Title:       *TitleResponseFromEntity(title),
		AchievedAt:  achievement.AchievedAt,
		IsCurrent:   achievement.IsCurrent.Bool(),
		CreatedAt:   achievement.CreatedAt,
		UpdatedAt:   achievement.UpdatedAt,
	}
}

// AchievementListFromEntities はエンティティリストからレスポンスDTOリストに変換する
// titlesMapはtitleIDをキーとした称号エンティティのマップ
func AchievementListFromEntities(achievements []*entity.TitleAchievement, titlesMap map[common.UUID]*entity.Title) []AchievementResponse {
	if achievements == nil {
		return []AchievementResponse{}
	}

	responses := make([]AchievementResponse, 0, len(achievements))
	for _, achievement := range achievements {
		if title, exists := titlesMap[achievement.TitleID]; exists {
			if response := AchievementResponseFromEntity(achievement, title); response != nil {
				responses = append(responses, *response)
			}
		}
	}
	return responses
}

// BuildAchievementListResponse は獲得記録リストと称号マップからレスポンスを構築する
func BuildAchievementListResponse(userID common.UUID, achievements []*entity.TitleAchievement, titlesMap map[common.UUID]*entity.Title) *AchievementListResponse {
	achievementResponses := AchievementListFromEntities(achievements, titlesMap)
	
	var currentTitle *AchievementResponse
	for i := range achievementResponses {
		if achievementResponses[i].IsCurrent {
			currentTitle = &achievementResponses[i]
			break
		}
	}

	return &AchievementListResponse{
		UserID:       userID,
		Achievements: achievementResponses,
		CurrentTitle: currentTitle,
		Total:        len(achievementResponses),
	}
}