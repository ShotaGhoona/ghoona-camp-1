package title

import (
	"net/http"

	"github.com/gin-gonic/gin"

	titleDto "ghoona-camp-backend/internal/application/dto/title"
	titleUsecase "ghoona-camp-backend/internal/application/usecase/title"
	"ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/interface/controller/common"
)

// TitleAchievementController は称号獲得記録操作のコントローラー
type TitleAchievementController struct {
	titleAchievementUseCase titleUsecase.TitleAchievementUseCase
	userRepo               repository.UserRepository
}

// NewTitleAchievementController は新しいTitleAchievementControllerを作成
func NewTitleAchievementController(
	titleAchievementUseCase titleUsecase.TitleAchievementUseCase,
	userRepo repository.UserRepository,
) *TitleAchievementController {
	return &TitleAchievementController{
		titleAchievementUseCase: titleAchievementUseCase,
		userRepo:               userRepo,
	}
}

// GetUserAchievements ユーザーの称号獲得履歴を取得
// GET /users/{userId}/achievements?include_progress=true&sort_by=level&order=asc
func (c *TitleAchievementController) GetUserAchievements(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	// クエリパラメータ解析
	includeProgress := false
	if progress := ctx.Query("include_progress"); progress == "true" {
		includeProgress = true
	}
	
	sortBy := ctx.DefaultQuery("sort_by", "level")
	order := ctx.DefaultQuery("order", "asc")
	
	// ユーザー情報と実績を含む複合レスポンス取得
	response, err := c.titleAchievementUseCase.GetUserAchievementsWithUserInfo(
		ctx.Request.Context(), 
		userID, 
		includeProgress,
		sortBy,
		order,
	)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	// API仕様準拠: userとachievementsを含むオブジェクト構造で返却
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}

// SetCurrentTitle 現在表示称号を変更
// PUT /users/{userId}/achievements/{titleId}
func (c *TitleAchievementController) SetCurrentTitle(ctx *gin.Context) {
	userID, err := common.ParseUUIDParam(ctx, "userId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "userId", err)
		return
	}
	
	titleID, err := common.ParseUUIDParam(ctx, "titleId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "titleId", err)
		return
	}
	
	// 本人のみアクセス可能
	if !common.RequireSelfAccess(ctx, userID, c.userRepo) {
		return
	}
	
	var req titleDto.SetCurrentTitleRequest
	if !common.BindJSON(ctx, &req) {
		return
	}
	
	response, err := c.titleAchievementUseCase.SetCurrentTitle(
		ctx.Request.Context(), 
		userID, 
		titleID,
	)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}