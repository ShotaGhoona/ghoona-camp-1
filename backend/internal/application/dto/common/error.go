package common

import "time"

// ErrorResponse 標準APIエラーレスポンス形式
type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Timestamp string `json:"timestamp"`
}

// NewErrorResponse 新しいエラーレスポンスインスタンスを作成
func NewErrorResponse(code, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    code,
			Message: message,
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

// 共通エラーコード
const (
	ErrorCodeUnauthorized          = "UNAUTHORIZED"
	ErrorCodeAccessDenied          = "ACCESS_DENIED"
	ErrorCodeValidationError       = "VALIDATION_ERROR"
	ErrorCodeUserNotFound          = "USER_NOT_FOUND"
	ErrorCodeGoalNotFound          = "GOAL_NOT_FOUND"
	ErrorCodeEventNotFound         = "EVENT_NOT_FOUND"
	ErrorCodeTitleNotFound         = "TITLE_NOT_FOUND"
	ErrorCodeSocialLinkNotFound    = "SOCIAL_LINK_NOT_FOUND"
	ErrorCodeRivalNotFound         = "RIVAL_NOT_FOUND"
	ErrorCodeAchievementNotFound   = "ACHIEVEMENT_NOT_FOUND"
	ErrorCodeEventFull             = "EVENT_FULL"
	ErrorCodeAlreadyRegistered     = "ALREADY_REGISTERED"
	ErrorCodeRivalLimitExceeded    = "RIVAL_LIMIT_EXCEEDED"
	ErrorCodeRivalAlreadyExists    = "RIVAL_ALREADY_EXISTS"
	ErrorCodeTitleNotAchieved      = "TITLE_NOT_ACHIEVED"
	ErrorCodeInternalServerError   = "INTERNAL_SERVER_ERROR"
)