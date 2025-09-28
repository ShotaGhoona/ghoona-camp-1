package user

import (
	"time"

	"github.com/google/uuid"

	"ghoona-camp-backend/internal/domain/user/entity"
)

// AddRivalRequest はライバル追加のリクエストDTO
type AddRivalRequest struct {
	RivalUserID uuid.UUID `json:"rival_user_id" binding:"required"`
}

// RivalResponse はライバルのレスポンスDTO
type RivalResponse struct {
	ID          uuid.UUID    `json:"id"`
	UserID      uuid.UUID    `json:"user_id"`
	RivalUser   UserResponse `json:"rival_user"`
	CreatedAt   time.Time    `json:"created_at"`
}

// RivalListResponse はライバル一覧のレスポンスDTO
type RivalListResponse struct {
	Rivals    []RivalResponse `json:"rivals"`
	Total     int             `json:"total"`
	MaxRivals int             `json:"max_rivals"`
}

// RivalResponseFromEntity はエンティティからレスポンスDTOに変換する
func RivalResponseFromEntity(rival *entity.UserRival, rivalUser *entity.User) *RivalResponse {
	if rival == nil || rivalUser == nil {
		return nil
	}

	rivalUserResponse := UserResponseFromEntity(rivalUser)
	if rivalUserResponse == nil {
		return nil
	}

	return &RivalResponse{
		ID:        rival.ID,
		UserID:    rival.UserID,
		RivalUser: *rivalUserResponse,
		CreatedAt: rival.CreatedAt,
	}
}

// RivalListFromEntities はエンティティリストからレスポンスDTOリストに変換する
func RivalListFromEntities(rivals []*entity.UserRival, rivalUsers []*entity.User) []RivalResponse {
	if rivals == nil || rivalUsers == nil {
		return []RivalResponse{}
	}

	// ライバルユーザーをマップ化（効率的な検索のため）
	rivalUserMap := make(map[uuid.UUID]*entity.User)
	for _, user := range rivalUsers {
		if user != nil {
			rivalUserMap[user.ID] = user
		}
	}

	responses := make([]RivalResponse, 0, len(rivals))
	for _, rival := range rivals {
		if rival != nil {
			if rivalUser, exists := rivalUserMap[rival.RivalUserID]; exists {
				if response := RivalResponseFromEntity(rival, rivalUser); response != nil {
					responses = append(responses, *response)
				}
			}
		}
	}

	return responses
}