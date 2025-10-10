package response

import "ghoona-camp-backend/internal/application/dto/common"

// GetRankingMonthlyResponse GET /ranking/monthly のレスポンス形式
type GetRankingMonthlyResponse struct {
	Data struct {
		Ranking    []GetRankingMonthlyRanking `json:"ranking"`
		Pagination *common.Pagination         `json:"pagination,omitempty"`
		Month      int                        `json:"month"`
		Year       int                        `json:"year"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetRankingMonthlyRanking GET /ranking/monthly の個別ランキング情報
type GetRankingMonthlyRanking struct {
	Rank                    int    `json:"rank"`
	UserID                  string `json:"user_id"`
	Username                string `json:"username"`
	AvatarURL               string `json:"avatar_url,omitempty"`
	MonthlyAttendanceDays   int    `json:"monthly_attendance_days"`
	TotalDurationMinutes    int    `json:"total_duration_minutes"`
}