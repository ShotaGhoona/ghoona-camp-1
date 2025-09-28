package common

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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
	case "linkId":
		errorCode = "INVALID_LINK_ID"
	case "rivalId":
		errorCode = "INVALID_RIVAL_ID"
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