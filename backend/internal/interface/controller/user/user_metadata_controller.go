package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userDto "ghoona-camp-backend/internal/application/dto/user"
	userUsecase "ghoona-camp-backend/internal/application/usecase/user"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/interface/controller/common"
)

// UserMetadataController はユーザーメタデータ操作のコントローラー
type UserMetadataController struct {
	userMetadataUseCase userUsecase.UserMetadataUseCase
	userRepo            repository.UserRepository
}

// NewUserMetadataController は新しいUserMetadataControllerを作成
func NewUserMetadataController(
	userMetadataUseCase userUsecase.UserMetadataUseCase,
	userRepo repository.UserRepository,
) *UserMetadataController {
	return &UserMetadataController{
		userMetadataUseCase: userMetadataUseCase,
		userRepo:            userRepo,
	}
}

// GetUserMetadata ユーザーメタデータを取得
// GET /users/{userId}/metadata
func (c *UserMetadataController) GetUserMetadata(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	response, err := c.userMetadataUseCase.GetUserMetadata(ctx.Request.Context(), userID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// CreateUserMetadata ユーザーメタデータを作成
// POST /users/{userId}/metadata
func (c *UserMetadataController) CreateUserMetadata(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	var req userDto.CreateUserMetadataRequest
	if !common.BindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userMetadataUseCase.CreateUserMetadata(ctx.Request.Context(), userID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusCreated, response)
}

// UpdateUserMetadata ユーザーメタデータを更新
// PUT /users/{userId}/metadata
func (c *UserMetadataController) UpdateUserMetadata(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	var req userDto.UpdateUserMetadataRequest
	if !common.BindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userMetadataUseCase.UpdateUserMetadata(ctx.Request.Context(), userID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}