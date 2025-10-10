package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// UserRival ユーザーライバル関係エンティティ
type UserRival struct {
	id          uuid.UUID // 関係ID
	userID      uuid.UUID // ユーザーID
	rivalUserID uuid.UUID // ライバルのユーザーID
	createdAt   time.Time // 設定日時
	updatedAt   time.Time // 更新日時
}

// NewUserRival UserRivalエンティティを作成する
func NewUserRival(
	userID uuid.UUID,
	rivalUserID uuid.UUID,
) (*UserRival, error) {
	// 自分自身不可チェック
	if userID == rivalUserID {
		return nil, errors.New("自分自身をライバルに設定することはできません")
	}

	now := time.Now()
	return &UserRival{
		id:          uuid.New(),
		userID:      userID,
		rivalUserID: rivalUserID,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// Getters
func (r *UserRival) ID() uuid.UUID          { return r.id }
func (r *UserRival) UserID() uuid.UUID      { return r.userID }
func (r *UserRival) RivalUserID() uuid.UUID { return r.rivalUserID }
func (r *UserRival) CreatedAt() time.Time   { return r.createdAt }
func (r *UserRival) UpdatedAt() time.Time   { return r.updatedAt }