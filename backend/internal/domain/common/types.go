package common

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// UUID は共通UUID型
type UUID = uuid.UUID

// NullUUID はnull許可UUID型
type NullUUID struct {
	UUID  UUID
	Valid bool
}

// Scan implements the Scanner interface
func (nu *NullUUID) Scan(value interface{}) error {
	if value == nil {
		nu.UUID, nu.Valid = UUID{}, false
		return nil
	}
	nu.Valid = true
	return nu.UUID.Scan(value)
}

// Value implements the driver Valuer interface
func (nu NullUUID) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}
	return nu.UUID.Value()
}

// NewUUID は新しいUUIDを生成
func NewUUID() UUID {
	return uuid.New()
}

// ParseUUID は文字列からUUIDを解析
func ParseUUID(s string) (UUID, error) {
	return uuid.Parse(s)
}

// BaseEntity は基本エンティティ
type BaseEntity struct {
	ID        UUID      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewBaseEntity は基本エンティティを作成
func NewBaseEntity() BaseEntity {
	now := time.Now()
	return BaseEntity{
		ID:        NewUUID(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateTimestamp は更新タイムスタンプを設定
func (e *BaseEntity) UpdateTimestamp() {
	e.UpdatedAt = time.Now()
}

// Pagination はページネーション情報
type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
	HasNext   bool `json:"has_next"`
	HasPrev   bool `json:"has_prev"`
}

// NewPagination はページネーション情報を作成
func NewPagination(page, limit, total int) *Pagination {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}

	totalPage := (total + limit - 1) / limit
	if totalPage == 0 {
		totalPage = 1
	}

	return &Pagination{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: totalPage,
		HasNext:   page < totalPage,
		HasPrev:   page > 1,
	}
}

// GetOffset はオフセット値を取得
func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

// SortOrder はソート順序
type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

// IsValid はソート順序が有効かチェック
func (s SortOrder) IsValid() bool {
	return s == SortOrderAsc || s == SortOrderDesc
}

// String はstring型に変換
func (s SortOrder) String() string {
	return string(s)
}

// FilterOptions は共通フィルターオプション
type FilterOptions struct {
	Page      int       `json:"page"`
	Limit     int       `json:"limit"`
	SortBy    string    `json:"sort_by"`
	SortOrder SortOrder `json:"sort_order"`
	Search    string    `json:"search"`
}

// Validate はフィルターオプションをバリデート
func (f *FilterOptions) Validate() error {
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.Limit <= 0 {
		f.Limit = 20
	}
	if f.Limit > 100 {
		f.Limit = 100
	}
	if f.SortOrder != "" && !f.SortOrder.IsValid() {
		return fmt.Errorf("invalid sort order: %s", f.SortOrder)
	}
	if f.SortOrder == "" {
		f.SortOrder = SortOrderDesc
	}
	return nil
}

// GetOffset はオフセット値を取得
func (f *FilterOptions) GetOffset() int {
	return (f.Page - 1) * f.Limit
}