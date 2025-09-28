package common

import (
	"time"

	"github.com/google/uuid"
)

// UUID は共通UUID型
// 使用予定: 全エンティティのID (User, Attendance, Goal, Event, Title, Notification)
type UUID = uuid.UUID

// NewUUID は新しいUUIDを生成
// 使用予定: 全エンティティ作成時 (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func NewUUID() UUID {
	return uuid.New()
}

// ParseUUID は文字列からUUIDを解析
// 使用予定: API リクエストパラメータからUUID変換 (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func ParseUUID(s string) (UUID, error) {
	return uuid.Parse(s)
}

// BaseEntity は基本エンティティ
// 使用予定: 全ドメインエンティティの基底構造 (User, Attendance, Goal, Event, Title, Notification)
type BaseEntity struct {
	ID        UUID      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewBaseEntity は基本エンティティを作成
// 使用予定: 全エンティティのコンストラクタ内 (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func NewBaseEntity() BaseEntity {
	now := time.Now()
	return BaseEntity{
		ID:        NewUUID(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateTimestamp は更新タイムスタンプを設定
// 使用予定: エンティティ更新処理 (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func (e *BaseEntity) UpdateTimestamp() {
	e.UpdatedAt = time.Now()
}

// Pagination はページネーション情報
// 使用予定: 
//   - 通知一覧API (BE-09-notify-04)
//   - 目標一覧API (BE-06-goal-04) 
//   - イベント一覧API (BE-07-event-04)
//   - ユーザー一覧API (BE-03-user-04)
//   - 称号一覧API (BE-08-title-04)
//   - 出席ログ一覧API (BE-04-attend-04)
type Pagination struct {
	CurrentPage int  `json:"current_page"`
	TotalPages  int  `json:"total_pages"`
	TotalCount  int  `json:"total_count"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
}

// NewPagination はページネーション情報を作成
// 使用予定: 全リスト取得ユースケース (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func NewPagination(page, limit, total int) *Pagination {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}

	totalPages := (total + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 1
	}

	return &Pagination{
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalCount:  total,
		HasNext:     page < totalPages,
		HasPrev:     page > 1,
	}
}

// GetOffset はオフセット値を取得
// 使用予定: GORM クエリでのOFFSET計算 (BE-03-user-*, BE-04-attend-*, BE-06-goal-*, BE-07-event-*, BE-08-title-*, BE-09-notify-*)
func (p *Pagination) GetOffset() int {
	return (p.CurrentPage - 1) * 20 // TODO: Limitフィールド削除のため固定値、将来的に改善
}