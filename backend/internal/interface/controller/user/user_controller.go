package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userDto "ghoona-camp-backend/internal/application/dto/user"
	userUsecase "ghoona-camp-backend/internal/application/usecase/user"
	"ghoona-camp-backend/internal/interface/controller/common"
)

// UserController はユーザー基本操作のコントローラー
type UserController struct {
	userUseCase userUsecase.UserUseCase
}

// NewUserController は新しいUserControllerを作成
func NewUserController(userUseCase userUsecase.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

// GetCurrentUser 現在のユーザー情報を取得
// GET /auth/me
func (c *UserController) GetCurrentUser(ctx *gin.Context) {
	clerkID, err := common.GetCurrentClerkID(ctx)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	response, err := c.userUseCase.GetUserByClerkID(ctx.Request.Context(), clerkID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// GetUser ユーザー詳細を取得
// GET /users/{userId}
func (c *UserController) GetUser(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	response, err := c.userUseCase.GetUserByID(ctx.Request.Context(), userID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// UpdateUser ユーザー基本情報を更新
// PUT /users/{userId}
func (c *UserController) UpdateUser(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userUseCase.GetUserRepo()) {
		return
	}
	
	var req userDto.UpdateUserRequest
	if !common.BindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.UpdateUser(ctx.Request.Context(), userID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// GetUsers 全ユーザー一覧を取得
// GET /users
func (c *UserController) GetUsers(ctx *gin.Context) {
	response, err := c.userUseCase.GetUsers(ctx.Request.Context())
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// CreateUser 新しいユーザーを作成
// POST /users
func (c *UserController) CreateUser(ctx *gin.Context) {
	var req userDto.CreateUserRequest
	if !common.BindJSON(ctx, &req) {
		return
	}
	
	response, err := c.userUseCase.CreateUser(ctx.Request.Context(), &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusCreated, response)
}