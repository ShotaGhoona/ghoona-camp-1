package vo

import "errors"

// EventType イベントタイプの分類を表すValue Object
type EventType string

const (
	EventTypeGeneral      EventType = "general"      // 一般
	EventTypeStudy        EventType = "study"        // 勉強
	EventTypeExercise     EventType = "exercise"     // 運動
	EventTypeMeditation   EventType = "meditation"   // 瞑想
	EventTypeDiscussion   EventType = "discussion"   // ディスカッション
	EventTypePresentation EventType = "presentation" // プレゼンテーション
	EventTypeWorkshop     EventType = "workshop"     // ワークショップ
)

// NewEventType EventTypeを作成する
func NewEventType(value string) (EventType, error) {
	eventType := EventType(value)
	if !eventType.isValid() {
		return "", errors.New("不正なイベントタイプです")
	}
	return eventType, nil
}

// isValid 有効なイベントタイプ値かどうかを検証する
func (t EventType) isValid() bool {
	switch t {
	case EventTypeGeneral, EventTypeStudy, EventTypeExercise, EventTypeMeditation,
		EventTypeDiscussion, EventTypePresentation, EventTypeWorkshop:
		return true
	default:
		return false
	}
}

// String 文字列表現を返す
func (t EventType) String() string {
	return string(t)
}

// Value 内部の文字列値を返す
func (t EventType) Value() string {
	return string(t)
}