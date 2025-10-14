package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GoalController 目標管理コントローラー
type GoalController struct {
	// TODO: UseCase依存注入（UseCase実装後）
	// goalUseCase application.GoalUseCase
}

// NewGoalController GoalControllerのコンストラクタ
func NewGoalController( /* TODO: UseCase追加 */ ) *GoalController {
	return &GoalController{
		// TODO: UseCase注入
	}
}

// GetGoalsMe 現在のユーザーの目標を取得（プライベート + パブリック）
// GET /goals/me?search=&is_public=&started_from=&started_to=&ended_from=&ended_to=&limit=&offset=
func (gc *GoalController) GetGoalsMe(c *gin.Context) {
	// ユーザーコンテキストから取得
	userID := c.GetString("user_id")
	if userID == "" {
		SendUnauthorizedError(c)
		return
	}

	// クエリパラメータ取得
	var query struct {
		SearchQuery
		PaginationQuery
		IsPublic    *bool  `form:"is_public"`
		StartedFrom string `form:"started_from" binding:"omitempty,datetime=2006-01-02"`
		StartedTo   string `form:"started_to" binding:"omitempty,datetime=2006-01-02"`
		EndedFrom   string `form:"ended_from" binding:"omitempty,datetime=2006-01-02"`
		EndedTo     string `form:"ended_to" binding:"omitempty,datetime=2006-01-02"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// goals, total, err := gc.goalUseCase.GetGoalsByUser(c.Request.Context(), userID, query)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "goals": goals,
	//     "pagination": PaginationMeta{
	//         Total:   total,
	//         Limit:   query.Limit,
	//         Offset:  query.Offset,
	//         HasMore: query.Offset+query.Limit < total,
	//     },
	// })

	SendInternalError(c)
}

// GetGoalsPublic 全ユーザーの公開目標を取得
// GET /goals/public?search=&user=&started_from=&started_to=&ended_from=&ended_to=&limit=&offset=
func (gc *GoalController) GetGoalsPublic(c *gin.Context) {
	// クエリパラメータ取得
	var query struct {
		SearchQuery
		PaginationQuery
		User        string `form:"user"`
		StartedFrom string `form:"started_from" binding:"omitempty,datetime=2006-01-02"`
		StartedTo   string `form:"started_to" binding:"omitempty,datetime=2006-01-02"`
		EndedFrom   string `form:"ended_from" binding:"omitempty,datetime=2006-01-02"`
		EndedTo     string `form:"ended_to" binding:"omitempty,datetime=2006-01-02"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// goals, total, err := gc.goalUseCase.GetPublicGoals(c.Request.Context(), query)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "goals": goals,
	//     "pagination": PaginationMeta{
	//         Total:   total,
	//         Limit:   query.Limit,
	//         Offset:  query.Offset,
	//         HasMore: query.Offset+query.Limit < total,
	//     },
	// })

	SendInternalError(c)
}

// PostGoal 新しい目標を作成
// POST /goals
func (gc *GoalController) PostGoal(c *gin.Context) {
	// ユーザーコンテキストから取得
	userID := c.GetString("user_id")
	if userID == "" {
		SendUnauthorizedError(c)
		return
	}

	// リクエストボディ取得
	var req struct {
		Title       string `json:"title" binding:"required,min=1,max=200"`
		Description string `json:"description" binding:"omitempty,max=1000"`
		StartedAt   string `json:"started_at" binding:"omitempty,datetime=2006-01-02"`
		EndedAt     string `json:"ended_at" binding:"omitempty,datetime=2006-01-02"`
		IsPublic    bool   `json:"is_public"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// goalID, err := gc.goalUseCase.CreateGoal(c.Request.Context(), userID, req)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusCreated, gin.H{
	//     "goal_id": goalID,
	//     "message": "Goal created successfully",
	// })

	SendInternalError(c)
}

// PutGoal 目標の内容を更新
// PUT /goals/{goalId}
func (gc *GoalController) PutGoal(c *gin.Context) {
	goalID := c.Param("goalId")
	if goalID == "" {
		SendValidationError(c, "Goal ID is required")
		return
	}

	// リクエストボディ取得
	var req struct {
		Title       string `json:"title" binding:"omitempty,min=1,max=200"`
		Description string `json:"description" binding:"omitempty,max=1000"`
		EndedAt     string `json:"ended_at" binding:"omitempty,datetime=2006-01-02"`
		IsPublic    *bool  `json:"is_public"`
		IsActive    *bool  `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// err := gc.goalUseCase.UpdateGoal(c.Request.Context(), goalID, req)
	// if err != nil {
	//     if errors.Is(err, domain.ErrGoalNotFound) {
	//         SendNotFoundError(c, "Goal")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"message": "Goal updated successfully"})

	SendInternalError(c)
}

// DeleteGoal 目標を削除
// DELETE /goals/{goalId}
func (gc *GoalController) DeleteGoal(c *gin.Context) {
	goalID := c.Param("goalId")
	if goalID == "" {
		SendValidationError(c, "Goal ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// err := gc.goalUseCase.DeleteGoal(c.Request.Context(), goalID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrGoalNotFound) {
	//         SendNotFoundError(c, "Goal")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }

	c.Status(http.StatusNoContent)
}