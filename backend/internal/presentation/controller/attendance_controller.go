package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AttendanceController 参加記録管理コントローラー
type AttendanceController struct {
	// TODO: UseCase依存注入（UseCase実装後）
	// attendanceUseCase application.AttendanceUseCase
}

// NewAttendanceController AttendanceControllerのコンストラクタ
func NewAttendanceController( /* TODO: UseCase追加 */ ) *AttendanceController {
	return &AttendanceController{
		// TODO: UseCase注入
	}
}

// GetRankingMonthly 月次参加ランキングを取得
// GET /ranking/monthly?month=&limit=&offset=
func (ac *AttendanceController) GetRankingMonthly(c *gin.Context) {
	// クエリパラメータ取得
	var query struct {
		PaginationQuery
		Month int `form:"month" binding:"omitempty,min=1,max=12"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// ranking, total, err := ac.attendanceUseCase.GetMonthlyRanking(c.Request.Context(), query.Month, query.Limit, query.Offset)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "ranking": ranking.Users,
	//     "pagination": PaginationMeta{
	//         Total:   total,
	//         Limit:   query.Limit,
	//         Offset:  query.Offset,
	//         HasMore: query.Offset+query.Limit < total,
	//     },
	//     "current_user_rank": ranking.CurrentUserRank,
	//     "month": query.Month,
	// })

	SendInternalError(c)
}

// GetRankingTotal 総合参加ランキングを取得
// GET /ranking/total?limit=&offset=
func (ac *AttendanceController) GetRankingTotal(c *gin.Context) {
	// クエリパラメータ取得
	var query PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// ranking, total, err := ac.attendanceUseCase.GetTotalRanking(c.Request.Context(), query.Limit, query.Offset)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "ranking": ranking.Users,
	//     "pagination": PaginationMeta{
	//         Total:   total,
	//         Limit:   query.Limit,
	//         Offset:  query.Offset,
	//         HasMore: query.Offset+query.Limit < total,
	//     },
	//     "current_user_rank": ranking.CurrentUserRank,
	// })

	SendInternalError(c)
}

// GetRankingStreak 連続参加ランキングを取得
// GET /ranking/streak?limit=&offset=
func (ac *AttendanceController) GetRankingStreak(c *gin.Context) {
	// クエリパラメータ取得
	var query PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// ranking, total, err := ac.attendanceUseCase.GetStreakRanking(c.Request.Context(), query.Limit, query.Offset)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "ranking": ranking.Users,
	//     "pagination": PaginationMeta{
	//         Total:   total,
	//         Limit:   query.Limit,
	//         Offset:  query.Offset,
	//         HasMore: query.Offset+query.Limit < total,
	//     },
	//     "current_user_rank": ranking.CurrentUserRank,
	// })

	SendInternalError(c)
}