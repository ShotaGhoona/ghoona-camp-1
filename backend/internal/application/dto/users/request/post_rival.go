package request

// PostRivalRequest POST /users/{userId}/rivals のリクエストボディ
type PostRivalRequest struct {
	RivalUserID string `json:"rival_user_id" validate:"required,uuid"`
}