package vo

import "errors"

// ParticipantStatus イベント参加者の状態を表すValue Object
type ParticipantStatus string

const (
	ParticipantStatusRegistered ParticipantStatus = "registered" // 参加登録済み
	ParticipantStatusCancelled  ParticipantStatus = "cancelled"  // キャンセル
)

// NewParticipantStatus ParticipantStatusを作成する
func NewParticipantStatus(value string) (ParticipantStatus, error) {
	status := ParticipantStatus(value)
	if !status.isValid() {
		return "", errors.New("不正な参加者ステータスです")
	}
	return status, nil
}

// isValid 有効な参加者ステータス値かどうかを検証する
func (s ParticipantStatus) isValid() bool {
	switch s {
	case ParticipantStatusRegistered, ParticipantStatusCancelled:
		return true
	default:
		return false
	}
}

// IsRegistered 参加登録済みかどうかを判定する
func (s ParticipantStatus) IsRegistered() bool {
	return s == ParticipantStatusRegistered
}

// String 文字列表現を返す
func (s ParticipantStatus) String() string {
	return string(s)
}

// Value 内部の文字列値を返す
func (s ParticipantStatus) Value() string {
	return string(s)
}