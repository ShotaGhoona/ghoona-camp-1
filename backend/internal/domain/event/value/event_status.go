// Package value イベントドメインの値オブジェクトを定義する
package value

// EventStatus イベント自体の進行状態を表す値オブジェクト
// 時間の経過に伴う状態変化を管理し、フィルタリングや表示制御に使用
// upcoming: 開催前, ongoing: 開催中, completed: 完了, cancelled: 中止
type EventStatus string

const (
	EventStatusUpcoming  EventStatus = "upcoming"
	EventStatusOngoing   EventStatus = "ongoing"
	EventStatusCompleted EventStatus = "completed"
	EventStatusCancelled EventStatus = "cancelled"
)

// IsValid イベントステータスが有効かどうかを判定する
func (s EventStatus) IsValid() bool {
	switch s {
	case EventStatusUpcoming, EventStatusOngoing, EventStatusCompleted, EventStatusCancelled:
		return true
	default:
		return false
	}
}

// String イベントステータスを文字列として返す
func (s EventStatus) String() string {
	return string(s)
}
