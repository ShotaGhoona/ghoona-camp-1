package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	userDto "ghoona-camp-backend/internal/application/dto/user"
	"ghoona-camp-backend/internal/application/usecase"
	domainUser "ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/infrastructure/clerk"
)

// UserController はユーザー管理のコントローラー
type UserController struct {
	userUseCase usecase.UserUseCase
}

// NewUserController は新しいUserControllerを作成
func NewUserController(userUseCase usecase.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

// ヘルパー関数

// parseUUIDParam パスパラメータからUUIDを解析
func parseUUIDParam(ctx *gin.Context, paramName string) (uuid.UUID, error) {
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

// getCurrentClerkID 認証されたユーザーのClerkIDを取得
func getCurrentClerkID(ctx *gin.Context) (string, error) {
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

// bindJSON JSONリクエストボディをバインド
func bindJSON(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBindJSON(req); err != nil {
		respondWithError(ctx, err)
		return false
	}
	return true
}

// respondWithError エラーレスポンスを返す
func respondWithError(ctx *gin.Context, err error) {
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
	case errors.Is(err, domainUser.ErrUserSocialLinkNotFound):
		statusCode = http.StatusNotFound
		errorCode = "SOCIAL_LINK_NOT_FOUND"
	case errors.Is(err, domainUser.ErrDuplicateEmail):
		statusCode = http.StatusConflict
		errorCode = "DUPLICATE_EMAIL"
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

// respondWithSuccess 成功レスポンスを返す
func respondWithSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"data":      data,
		"message":   "success",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// requireSelfAccess 本人のみアクセス可能かチェック（ClerkIDベース）
func requireSelfAccess(ctx *gin.Context, targetUserID uuid.UUID, userRepo repository.UserRepository) bool {
	clerkID, err := getCurrentClerkID(ctx)
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

// コントローラーメソッド

// GetCurrentUser 現在のユーザー情報を取得
// GET /auth/me
func (c *UserController) GetCurrentUser(ctx *gin.Context) {
	clerkID, err := getCurrentClerkID(ctx)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	response, err := c.userUseCase.GetUserByClerkID(ctx.Request.Context(), clerkID)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// GetUser ユーザー詳細を取得
// GET /users/{userId}
func (c *UserController) GetUser(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	response, err := c.userUseCase.GetUserByID(ctx.Request.Context(), userID)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// UpdateUser ユーザー基本情報を更新
// PUT /users/{userId}
func (c *UserController) UpdateUser(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	var req userDto.UpdateUserRequest
	if !bindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.UpdateUser(ctx.Request.Context(), userID, &req)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// GetUserMetadata ユーザーメタデータを取得
// GET /users/{userId}/metadata
func (c *UserController) GetUserMetadata(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	response, err := c.userUseCase.GetUserMetadata(ctx.Request.Context(), userID)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// CreateUserMetadata ユーザーメタデータを作成
// POST /users/{userId}/metadata
func (c *UserController) CreateUserMetadata(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	var req userDto.CreateUserMetadataRequest
	if !bindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.CreateUserMetadata(ctx.Request.Context(), userID, &req)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusCreated, response)
}

// UpdateUserMetadata ユーザーメタデータを更新
// PUT /users/{userId}/metadata
func (c *UserController) UpdateUserMetadata(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	var req userDto.UpdateUserMetadataRequest
	if !bindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.UpdateUserMetadata(ctx.Request.Context(), userID, &req)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// GetUserSocialLinks ユーザーのソーシャルリンク一覧を取得
// GET /users/{userId}/social-links
func (c *UserController) GetUserSocialLinks(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	response, err := c.userUseCase.GetUserSocialLinks(ctx.Request.Context(), userID)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// CreateSocialLink ソーシャルリンクを追加
// POST /users/{userId}/social-links
func (c *UserController) CreateSocialLink(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	var req userDto.CreateSocialLinkRequest
	if !bindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.CreateSocialLink(ctx.Request.Context(), userID, &req)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusCreated, response)
}

// UpdateSocialLink ソーシャルリンクを更新
// PUT /users/{userId}/social-links/{linkId}
func (c *UserController) UpdateSocialLink(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	linkID, err := parseUUIDParam(ctx, "linkId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_LINK_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	var req userDto.UpdateSocialLinkRequest
	if !bindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.UpdateSocialLink(ctx.Request.Context(), linkID, &req)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// DeleteSocialLink ソーシャルリンクを削除
// DELETE /users/{userId}/social-links/{linkId}
func (c *UserController) DeleteSocialLink(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	linkID, err := parseUUIDParam(ctx, "linkId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_LINK_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	err = c.userUseCase.DeleteSocialLink(ctx.Request.Context(), linkID)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	ctx.Status(http.StatusNoContent)
}

// GetUserRivals ユーザーのライバル一覧を取得
// GET /users/{userId}/rivals
func (c *UserController) GetUserRivals(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	response, err := c.userUseCase.GetUserRivals(ctx.Request.Context(), userID)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// AddRival ライバルを追加
// POST /users/{userId}/rivals
func (c *UserController) AddRival(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	var req userDto.AddRivalRequest
	if !bindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.AddRival(ctx.Request.Context(), userID, &req)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusCreated, response)
}

// RemoveRival ライバルを削除
// DELETE /users/{userId}/rivals/{rivalId}
func (c *UserController) RemoveRival(ctx *gin.Context) {
	userID, err := parseUUIDParam(ctx, "userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_USER_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	rivalID, err := parseUUIDParam(ctx, "rivalId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "INVALID_RIVAL_ID",
				"message": err.Error(),
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}
	
	// 本人のみアクセス可能
	if !requireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	err = c.userUseCase.RemoveRival(ctx.Request.Context(), rivalID)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	ctx.Status(http.StatusNoContent)
}

// GetUsers 全ユーザー一覧を取得
// GET /users
func (c *UserController) GetUsers(ctx *gin.Context) {
	response, err := c.userUseCase.GetUsers(ctx.Request.Context())
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusOK, response)
}

// CreateUser 新しいユーザーを作成
// POST /users
func (c *UserController) CreateUser(ctx *gin.Context) {
	var req userDto.CreateUserRequest
	if !bindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.CreateUser(ctx.Request.Context(), &req)
	if err != nil {
		respondWithError(ctx, err)
		return
	}
	
	respondWithSuccess(ctx, http.StatusCreated, response)
}