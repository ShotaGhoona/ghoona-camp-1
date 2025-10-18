package controller

import (
	"github.com/gin-gonic/gin"
)

// TitleController 称号管理コントローラー
type TitleController struct {
	// TODO: UseCase依存注入（UseCase実装後）
	// titleUseCase application.TitleUseCase
}

// NewTitleController TitleControllerのコンストラクタ
func NewTitleController( /* TODO: UseCase追加 */ ) *TitleController {
	return &TitleController{
		// TODO: UseCase注入
	}
}

// GetTitles 全称号一覧を取得（8レベル）
// GET /titles
func (tc *TitleController) GetTitles(c *gin.Context) {
	// TODO: UseCase呼び出し
	// titles, err := tc.titleUseCase.GetAllTitles(c.Request.Context())
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"titles": titles})

	SendInternalError(c)
}

// GetTitleDetail 指定称号の詳細情報を取得
// GET /titles/{titleId}
func (tc *TitleController) GetTitleDetail(c *gin.Context) {
	titleID := c.Param("titleId")
	if titleID == "" {
		SendValidationError(c, "Title ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// title, err := tc.titleUseCase.GetTitleDetail(c.Request.Context(), titleID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrTitleNotFound) {
	//         SendNotFoundError(c, "Title")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"title": title})

	SendInternalError(c)
}
