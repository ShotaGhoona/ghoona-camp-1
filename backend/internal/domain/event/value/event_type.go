// Package value イベントドメインの値オブジェクトを定義する
// 値オブジェクトは不変で業務ルールを表現し、エンティティの属性として使用される
package value

// EventType イベントの種類を表す値オブジェクト
// 朝活の多様なニーズに対応するため、勉強・運動・瞑想などのカテゴリを定義
// 無効な値の混入を防ぎ、フロントエンドでの選択肢としても使用される
type EventType string

const (
	EventTypeGeneral    EventType = "general"
	EventTypeStudy      EventType = "study"
	EventTypeExercise   EventType = "exercise"
	EventTypeMeditation EventType = "meditation"
	EventTypeCreative   EventType = "creative"
	EventTypeBusiness   EventType = "business"
)

// IsValid イベントタイプが有効かどうかを判定する
func (t EventType) IsValid() bool {
	switch t {
	case EventTypeGeneral, EventTypeStudy, EventTypeExercise, EventTypeMeditation, EventTypeCreative, EventTypeBusiness:
		return true
	default:
		return false
	}
}

// String イベントタイプを文字列として返す
func (t EventType) String() string {
	return string(t)
}