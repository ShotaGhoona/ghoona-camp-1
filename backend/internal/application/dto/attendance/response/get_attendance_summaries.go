package response

// GetAttendanceSummariesResponse GET /users/{userId}/attendance/summaries のレスポンス形式
type GetAttendanceSummariesResponse struct {
	Data struct {
		Summaries []GetAttendanceSummariesSummary `json:"summaries"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetAttendanceSummariesSummary GET /users/{userId}/attendance/summaries の個別サマリー情報
type GetAttendanceSummariesSummary struct {
	ID                   string `json:"id"`
	UserID               string `json:"user_id"`
	Date                 string `json:"date"`
	TotalDurationMinutes int    `json:"total_duration_minutes"`
	SessionCount         int    `json:"session_count"`
	FirstJoinTime        string `json:"first_join_time,omitempty"`
	LastLeaveTime        string `json:"last_leave_time,omitempty"`
	IsMorningActive      bool   `json:"is_morning_active"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}