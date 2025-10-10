package response

import "ghoona-camp-backend/internal/application/dto/common"

// GetRankingStreakResponse GET /ranking/streak のレスポンス形式
type GetRankingStreakResponse struct {
	Data struct {
		Ranking    []GetRankingStreakRanking `json:"ranking"`
		Pagination *common.Pagination        `json:"pagination,omitempty"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetRankingStreakRanking GET /ranking/streak の個別ランキング情報
type GetRankingStreakRanking struct {
	Rank                int    `json:"rank"`
	UserID              string `json:"user_id"`
	Username            string `json:"username"`
	AvatarURL           string `json:"avatar_url,omitempty"`
	CurrentStreakDays   int    `json:"current_streak_days"`
	MaxStreakDays       int    `json:"max_streak_days"`
	LastAttendanceDate  string `json:"last_attendance_date,omitempty"`
}