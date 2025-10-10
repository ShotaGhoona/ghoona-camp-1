package response

// GetAttendanceStatisticsResponse GET /users/{userId}/attendance/statistics のレスポンス形式
type GetAttendanceStatisticsResponse struct {
	Data struct {
		Statistics GetAttendanceStatisticsStatistics `json:"statistics"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetAttendanceStatisticsStatistics GET /users/{userId}/attendance/statistics の統計情報
type GetAttendanceStatisticsStatistics struct {
	ID                   string `json:"id"`
	UserID               string `json:"user_id"`
	TotalAttendanceDays  int    `json:"total_attendance_days"`
	CurrentStreakDays    int    `json:"current_streak_days"`
	MaxStreakDays        int    `json:"max_streak_days"`
	LastAttendanceDate   string `json:"last_attendance_date,omitempty"`
	FirstAttendanceDate  string `json:"first_attendance_date,omitempty"`
	TotalDurationMinutes int    `json:"total_duration_minutes"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}