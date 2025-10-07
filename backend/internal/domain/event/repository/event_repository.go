// Package repository イベントドメインのリポジトリインターフェースを定義する
// インターフェースをドメイン層で定義し、実装をインフラ層で行う（依存関係逆転の原則）
package repository

import (
	"context"

	"ghoona-camp-backend/internal/domain/event/entity"
	"ghoona-camp-backend/internal/domain/event/value"

	"github.com/google/uuid"
)

// EventRepository イベントの永続化を抽象化するインターフェース
type EventRepository interface {
	// 詳細取得 - イベント詳細表示、更新、削除処理で使用
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Event, error)
	
	// 一覧取得 - イベント一覧API、検索機能で使用
	FindAll(ctx context.Context, filter EventFilter) ([]*entity.Event, error)
	
	// 作成 - イベント作成APIで使用
	Create(ctx context.Context, event *entity.Event) error
	
	// 更新 - イベント更新APIで使用
	Update(ctx context.Context, event *entity.Event) error
	
	// 削除 - イベント削除APIで使用
	Delete(ctx context.Context, id uuid.UUID) error

	// 作成者別取得 - 自分が作成したイベント一覧で使用
	FindByCreatorID(ctx context.Context, creatorID uuid.UUID) ([]*entity.Event, error)
	
	// 日付範囲取得 - 開催予定イベント一覧で使用
	FindByDateRange(ctx context.Context, startDate, endDate string) ([]*entity.Event, error)
	
	// 参加者数カウント - 定員チェック、一覧表示のparticipant_countで使用
	CountParticipants(ctx context.Context, eventID uuid.UUID) (int, error)
}

// EventFilter イベント一覧取得時のフィルタ条件
type EventFilter struct {
	DateFrom   *string
	DateTo     *string
	EventType  *value.EventType
	Search     *string
	Page       int
	Limit      int
}