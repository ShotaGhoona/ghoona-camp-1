package title

import (
	"net/http"

	"github.com/gin-gonic/gin"

	titleUsecase "ghoona-camp-backend/internal/application/usecase/title"
	"ghoona-camp-backend/internal/interface/controller/common"
)

// TitleController は称号操作のコントローラー
type TitleController struct {
	titleUseCase titleUsecase.TitleUseCase
}

// NewTitleController は新しいTitleControllerを作成
func NewTitleController(titleUseCase titleUsecase.TitleUseCase) *TitleController {
	return &TitleController{
		titleUseCase: titleUseCase,
	}
}

// GetTitles 称号一覧を取得
// GET /titles?include_inactive=true&user_id={userId}
func (c *TitleController) GetTitles(ctx *gin.Context) {
	// クエリパラメータ解析
	includeInactive := false
	if inactive := ctx.Query("include_inactive"); inactive == "true" {
		includeInactive = true
	}
	
	response, err := c.titleUseCase.GetAllTitles(ctx.Request.Context(), includeInactive)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	// API仕様に合わせて配列を直接返却
	common.RespondWithSuccess(ctx, http.StatusOK, response.Titles)
}

// GetTitle 称号詳細を取得
// GET /titles/{titleId}
func (c *TitleController) GetTitle(ctx *gin.Context) {
	titleID, err := common.ParseUUIDParam(ctx, "titleId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "titleId", err)
		return
	}
	
	response, err := c.titleUseCase.GetTitleByID(ctx.Request.Context(), titleID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}
	
	common.RespondWithSuccess(ctx, http.StatusOK, response)
}