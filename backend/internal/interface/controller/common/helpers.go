package common

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	domainEvent "ghoona-camp-backend/internal/domain/event"
	domainTitle "ghoona-camp-backend/internal/domain/title"
	domainUser "ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/infrastructure/clerk"
)

// ParseUUIDParam パスパラメータからUUIDを解析
func ParseUUIDParam(ctx *gin.Context, paramName string) (uuid.UUID, error) {
	paramStr := ctx.Param(paramName)
	if paramStr == "" {
		return uuid.Nil, errors.New("missing parameter: " + paramName)
	}
	
	id, err := uuid.Parse(paramStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid UUID format: " + paramName)
	}
	
	return id, nil
}

// GetCurrentClerkID 認証されたユーザーのClerkIDを取得
func GetCurrentClerkID(ctx *gin.Context) (string, error) {
	// Clerk認証ミドルウェアからユーザー情報を取得
	userInterface, exists := ctx.Get("user")
	if !exists {
		return "", errors.New("user not found in context")
	}
	
	// ClerkUserからIDを取得
	user, ok := userInterface.(*clerk.ClerkUser)
	if !ok {
		return "", errors.New("invalid user format in context")
	}
	
	return user.ID, nil
}

// BindJSON JSONリクエストボディをバインド
func BindJSON(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBindJSON(req); err != nil {
		RespondWithError(ctx, err)
		return false
	}
	return true
}

// RespondWithError エラーレスポンスを返す
func RespondWithError(ctx *gin.Context, err error) {
	var statusCode int
	var errorCode string
	
	// ドメインエラーをHTTPステータスコードにマッピング
	switch {
	// Title Domain Errors
	case errors.Is(err, domainTitle.ErrTitleNotFound):
		statusCode = http.StatusNotFound
		errorCode = "TITLE_NOT_FOUND"
	case errors.Is(err, domainTitle.ErrTitleInactive):
		statusCode = http.StatusUnprocessableEntity
		errorCode = "TITLE_INACTIVE"
	case errors.Is(err, domainTitle.ErrTitleNotAchieved):
		statusCode = http.StatusForbidden
		errorCode = "TITLE_NOT_ACHIEVED"
	case errors.Is(err, domainTitle.ErrTitleAlreadyCurrent):
		statusCode = http.StatusConflict
		errorCode = "TITLE_ALREADY_CURRENT"
	case errors.Is(err, domainTitle.ErrAchievementNotFound):
		statusCode = http.StatusNotFound
		errorCode = "ACHIEVEMENT_NOT_FOUND"
	case errors.Is(err, domainTitle.ErrAchievementAlreadyExists):
		statusCode = http.StatusConflict
		errorCode = "DUPLICATE_ACHIEVEMENT"
	case errors.Is(err, domainTitle.ErrMultipleCurrentTitles):
		statusCode = http.StatusConflict
		errorCode = "MULTIPLE_CURRENT_TITLES"
	case errors.Is(err, domainTitle.ErrInvalidTitleLevel):
		statusCode = http.StatusUnprocessableEntity
		errorCode = "INVALID_TITLE_LEVEL"
	// User Domain Errors
	case errors.Is(err, domainUser.ErrUserNotFound):
		statusCode = http.StatusNotFound
		errorCode = "USER_NOT_FOUND"
	case errors.Is(err, domainUser.ErrUserMetadataNotFound):
		statusCode = http.StatusNotFound
		errorCode = "USER_METADATA_NOT_FOUND"
	case errors.Is(err, domainUser.ErrUserMetadataAlreadyExists):
		statusCode = http.StatusConflict
		errorCode = "USER_METADATA_ALREADY_EXISTS"
	case errors.Is(err, domainUser.ErrUserSocialLinkNotFound):
		statusCode = http.StatusNotFound
		errorCode = "SOCIAL_LINK_NOT_FOUND"
	case errors.Is(err, domainUser.ErrDuplicateEmail):
		statusCode = http.StatusConflict
		errorCode = "DUPLICATE_EMAIL"
	case errors.Is(err, domainUser.ErrDuplicateClerkID):
		statusCode = http.StatusConflict
		errorCode = "DUPLICATE_CLERK_ID"
	case errors.Is(err, domainUser.ErrDuplicatePlatform):
		statusCode = http.StatusConflict
		errorCode = "DUPLICATE_PLATFORM"
	case errors.Is(err, domainUser.ErrRivalLimitExceeded):
		statusCode = http.StatusUnprocessableEntity
		errorCode = "RIVAL_LIMIT_EXCEEDED"
	case errors.Is(err, domainUser.ErrCannotRivalSelf):
		statusCode = http.StatusUnprocessableEntity
		errorCode = "CANNOT_RIVAL_SELF"
	case errors.Is(err, domainUser.ErrInvalidEmail):
		statusCode = http.StatusUnprocessableEntity
		errorCode = "INVALID_EMAIL"
	case errors.Is(err, domainUser.ErrInvalidURL):
		statusCode = http.StatusUnprocessableEntity
		errorCode = "INVALID_URL"
	case errors.Is(err, domainUser.ErrInvalidPlatform):
		statusCode = http.StatusUnprocessableEntity
		errorCode = "INVALID_PLATFORM"
	// Event Domain Errors
	case errors.Is(err, domainEvent.ErrEventNotFound):
		statusCode = http.StatusNotFound
		errorCode = "EVENT_NOT_FOUND"
	case errors.Is(err, domainEvent.ErrUnauthorized):
		statusCode = http.StatusForbidden
		errorCode = "UNAUTHORIZED"
	case errors.Is(err, domainEvent.ErrParticipantAlreadyExists):
		statusCode = http.StatusConflict
		errorCode = "PARTICIPANT_ALREADY_EXISTS"
	case errors.Is(err, domainEvent.ErrEventFull):
		statusCode = http.StatusConflict
		errorCode = "EVENT_FULL"
	case errors.Is(err, domainEvent.ErrParticipantNotFound):
		statusCode = http.StatusNotFound
		errorCode = "PARTICIPANT_NOT_FOUND"
	case errors.Is(err, domainEvent.ErrCannotReduceCapacity):
		statusCode = http.StatusBadRequest
		errorCode = "CANNOT_REDUCE_CAPACITY"
	case errors.Is(err, domainEvent.ErrInvalidTimeSlot):
		statusCode = http.StatusBadRequest
		errorCode = "INVALID_TIME_SLOT"
	case errors.Is(err, domainEvent.ErrEventAlreadyExists):
		statusCode = http.StatusConflict
		errorCode = "EVENT_ALREADY_EXISTS"
	case errors.Is(err, domainEvent.ErrInvalidTitle):
		statusCode = http.StatusBadRequest
		errorCode = "INVALID_TITLE"
	case errors.Is(err, domainEvent.ErrInvalidEventType):
		statusCode = http.StatusBadRequest
		errorCode = "INVALID_EVENT_TYPE"
	case errors.Is(err, domainEvent.ErrInvalidCapacity):
		statusCode = http.StatusBadRequest
		errorCode = "INVALID_CAPACITY"
	default:
		statusCode = http.StatusInternalServerError
		errorCode = "INTERNAL_SERVER_ERROR"
	}
	
	ctx.JSON(statusCode, gin.H{
		"error": gin.H{
			"code":    errorCode,
			"message": err.Error(),
		},
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// RespondWithSuccess 成功レスポンスを返す
func RespondWithSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"data":      data,
		"message":   "success",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// RequireSelfAccess 本人のみアクセス可能かチェック（ClerkIDベース）
func RequireSelfAccess(ctx *gin.Context, targetUserID uuid.UUID, userRepo repository.UserRepository) bool {
	clerkID, err := GetCurrentClerkID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": gin.H{
				"code":    "UNAUTHORIZED",
				"message": "Authentication required",
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return false
	}
	
	// ClerkIDから内部UUIDへの変換
	currentUser, err := userRepo.GetByClerkID(ctx.Request.Context(), clerkID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": gin.H{
				"code":    "USER_NOT_FOUND",
				"message": "Current user not found",
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return false
	}
	
	// 本人確認
	if currentUser.ID != targetUserID {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": gin.H{
				"code":    "FORBIDDEN",
				"message": "Access denied: can only access own resources",
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return false
	}
	
	return true
}

// HandleUUIDParamError UUIDパラメータエラーを処理
func HandleUUIDParamError(ctx *gin.Context, paramName string, err error) {
	var errorCode string
	switch paramName {
	case "userId":
		errorCode = "INVALID_USER_ID"
	case "titleId":
		errorCode = "INVALID_TITLE_ID"
	case "linkId":
		errorCode = "INVALID_LINK_ID"
	case "rivalId":
		errorCode = "INVALID_RIVAL_ID"
	case "eventId":
		errorCode = "INVALID_EVENT_ID"
	default:
		errorCode = "INVALID_PARAMETER"
	}
	
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": gin.H{
			"code":    errorCode,
			"message": err.Error(),
		},
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}