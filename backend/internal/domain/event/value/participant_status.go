// Package value イベントドメインの値オブジェクトを定義する
package value

// ParticipantStatus イベント参加者の状態を表す値オブジェクト
// 参加登録とキャンセルの2状態のみをサポートし、履歴保持を可能にする
// DELETEではなくステータス変更により、再参加や分析データの保持を実現
type ParticipantStatus string

const (
	ParticipantStatusRegistered ParticipantStatus = "registered"
	ParticipantStatusCancelled  ParticipantStatus = "cancelled"
)

// IsValid 参加ステータスが有効かどうかを判定する
func (s ParticipantStatus) IsValid() bool {
	switch s {
	case ParticipantStatusRegistered, ParticipantStatusCancelled:
		return true
	default:
		return false
	}
}

// String 参加ステータスを文字列として返す
func (s ParticipantStatus) String() string {
	return string(s)
}