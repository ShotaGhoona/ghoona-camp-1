package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userDto "ghoona-camp-backend/internal/application/dto/user"
	userUsecase "ghoona-camp-backend/internal/application/usecase/user"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/interface/controller/common"
)

// UserSocialController はユーザーソーシャルリンク操作のコントローラー
type UserSocialController struct {
	userSocialUseCase userUsecase.UserSocialUseCase
	userRepo          repository.UserRepository
}

// NewUserSocialController は新しいUserSocialControllerを作成
func NewUserSocialController(
	userSocialUseCase userUsecase.UserSocialUseCase,
	userRepo repository.UserRepository,
) *UserSocialController {
	return &UserSocialController{
		userSocialUseCase: userSocialUseCase,
		userRepo:          userRepo,
	}
}

// GetUserSocialLinks ユーザーのソーシャルリンク一覧を取得
// GET /users/{userId}/social-links
func (c *UserSocialController) GetUserSocialLinks(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	response, err := c.userSocialUseCase.GetUserSocialLinks(ctx.Request.Context(), userID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// CreateSocialLink ソーシャルリンクを追加
// POST /users/{userId}/social-links
func (c *UserSocialController) CreateSocialLink(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	var req userDto.CreateSocialLinkRequest
	if !common.BindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userSocialUseCase.CreateSocialLink(ctx.Request.Context(), userID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusCreated, response)
}

// UpdateSocialLink ソーシャルリンクを更新
// PUT /users/{userId}/social-links/{linkId}
func (c *UserSocialController) UpdateSocialLink(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	linkID, err := common.ParseUUIDParam(ctx, "linkId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "linkId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	var req userDto.UpdateSocialLinkRequest
	if !common.BindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userSocialUseCase.UpdateSocialLink(ctx.Request.Context(), linkID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// DeleteSocialLink ソーシャルリンクを削除
// DELETE /users/{userId}/social-links/{linkId}
func (c *UserSocialController) DeleteSocialLink(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	linkID, err := common.ParseUUIDParam(ctx, "linkId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "linkId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	err = c.userSocialUseCase.DeleteSocialLink(ctx.Request.Context(), linkID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	ctx.Status(http.StatusNoContent)
}