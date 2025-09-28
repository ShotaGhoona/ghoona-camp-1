package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userDto "ghoona-camp-backend/internal/application/dto/user"
	userUsecase "ghoona-camp-backend/internal/application/usecase/user"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/interface/controller/common"
)

// UserRivalController はユーザーライバル操作のコントローラー
type UserRivalController struct {
	userRivalUseCase userUsecase.UserRivalUseCase
	userRepo         repository.UserRepository
}

// NewUserRivalController は新しいUserRivalControllerを作成
func NewUserRivalController(
	userRivalUseCase userUsecase.UserRivalUseCase,
	userRepo repository.UserRepository,
) *UserRivalController {
	return &UserRivalController{
		userRivalUseCase: userRivalUseCase,
		userRepo:         userRepo,
	}
}

// GetUserRivals ユーザーのライバル一覧を取得
// GET /users/{userId}/rivals
func (c *UserRivalController) GetUserRivals(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	response, err := c.userRivalUseCase.GetUserRivals(ctx.Request.Context(), userID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// AddRival ライバルを追加
// POST /users/{userId}/rivals
func (c *UserRivalController) AddRival(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	var req userDto.AddRivalRequest
	if !common.BindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userRivalUseCase.AddRival(ctx.Request.Context(), userID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusCreated, response)
}

// RemoveRival ライバルを削除
// DELETE /users/{userId}/rivals/{rivalId}
func (c *UserRivalController) RemoveRival(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	rivalID, err := common.ParseUUIDParam(ctx, "rivalId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "rivalId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	err = c.userRivalUseCase.RemoveRival(ctx.Request.Context(), rivalID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	ctx.Status(http.StatusNoContent)
}