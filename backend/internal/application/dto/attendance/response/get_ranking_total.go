package response

import "ghoona-camp-backend/internal/application/dto/common"

// GetRankingTotalResponse GET /ranking/total のレスポンス形式
type GetRankingTotalResponse struct {
	Data struct {
		Ranking    []GetRankingTotalRanking `json:"ranking"`
		Pagination *common.Pagination       `json:"pagination,omitempty"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetRankingTotalRanking GET /ranking/total の個別ランキング情報
type GetRankingTotalRanking struct {
	Rank                 int    `json:"rank"`
	UserID               string `json:"user_id"`
	Username             string `json:"username"`
	AvatarURL            string `json:"avatar_url,omitempty"`
	TotalAttendanceDays  int    `json:"total_attendance_days"`
	TotalDurationMinutes int    `json:"total_duration_minutes"`
	FirstAttendanceDate  string `json:"first_attendance_date,omitempty"`
}