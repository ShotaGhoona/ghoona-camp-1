package user

import (
	"time"

	"github.com/google/uuid"

	"ghoona-camp-backend/internal/domain/user/entity"
)

// CreateUserRequest はユーザー作成のリクエストDTO
type CreateUserRequest struct {
	ClerkID   string  `json:"clerk_id" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Username  *string `json:"username" binding:"omitempty,min=3,max=50"`
	AvatarURL *string `json:"avatar_url" binding:"omitempty"`
	DiscordID *string `json:"discord_id" binding:"omitempty"`
}

// UpdateUserRequest はユーザー更新のリクエストDTO
type UpdateUserRequest struct {
	Username  *string `json:"username" binding:"omitempty,min=3,max=50"`
	AvatarURL *string `json:"avatar_url" binding:"omitempty"`
	DiscordID *string `json:"discord_id" binding:"omitempty"`
}

// UserResponse はユーザーのレスポンスDTO
type UserResponse struct {
	ID        uuid.UUID             `json:"id"`
	ClerkID   string                `json:"clerk_id"`
	Email     string                `json:"email"`
	Username  *string               `json:"username"`
	AvatarURL *string               `json:"avatar_url"`
	DiscordID *string               `json:"discord_id"`
	Status    string                `json:"status"`
	Metadata  *UserMetadataResponse `json:"metadata,omitempty"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

// UserListResponse はユーザー一覧のレスポンスDTO
type UserListResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
}

// UserResponseFromEntity はエンティティからレスポンスDTOに変換する
func UserResponseFromEntity(user *entity.User) *UserResponse {
	if user == nil {
		return nil
	}

	return &UserResponse{
		ID:        user.ID,
		ClerkID:   user.ClerkID,
		Email:     user.Email,
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
		DiscordID: user.DiscordID,
		Status:    user.Status.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// UserResponseFromEntityWithMetadata はメタデータ付きでエンティティからレスポンスDTOに変換する
func UserResponseFromEntityWithMetadata(user *entity.User, metadata *entity.UserMetadata) *UserResponse {
	response := UserResponseFromEntity(user)
	if response != nil && metadata != nil {
		response.Metadata = UserMetadataResponseFromEntity(metadata)
	}
	return response
}

// UserListFromEntities はエンティティリストからレスポンスDTOリストに変換する
func UserListFromEntities(users []*entity.User) []UserResponse {
	if users == nil {
		return []UserResponse{}
	}

	responses := make([]UserResponse, len(users))
	for i, user := range users {
		if response := UserResponseFromEntity(user); response != nil {
			responses[i] = *response
		}
	}
	return responses
}