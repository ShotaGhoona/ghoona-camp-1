package value

// AttendanceStatistics は称号判定に必要な出席統計情報を表す
// 他ドメインへの依存を避けるための値オブジェクト
type AttendanceStatistics struct {
	TotalAttendanceDays int // 総参加日数
}

// NewAttendanceStatistics は新しいAttendanceStatisticsを作成する
func NewAttendanceStatistics(totalDays int) AttendanceStatistics {
	if totalDays < 0 {
		totalDays = 0
	}
	return AttendanceStatistics{
		TotalAttendanceDays: totalDays,
	}
}

// IsValid は統計情報が有効かどうかを確認する
func (a AttendanceStatistics) IsValid() bool {
	return a.TotalAttendanceDays >= 0
}