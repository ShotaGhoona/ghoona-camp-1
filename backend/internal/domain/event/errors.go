// Package event イベントドメインで使用するエラーを定義する
// ドメイン固有のエラーによりビジネスルール違反を明確に表現する
package event

import "errors"

// イベント関連エラー
var (
	ErrEventNotFound      = errors.New("イベントが見つかりません")
	ErrEventAlreadyExists = errors.New("イベントが既に存在します")
	ErrEventCancelled     = errors.New("イベントはキャンセルされています")
	ErrEventCompleted     = errors.New("イベントは既に完了しています")
	ErrEventInactive      = errors.New("イベントは無効です")
)

// イベント作成・更新関連エラー
var (
	ErrInvalidTitle       = errors.New("タイトルは必須です")
	ErrTitleTooLong       = errors.New("タイトルは200文字以内で入力してください")
	ErrDescriptionTooLong = errors.New("説明は2000文字以内で入力してください")
	ErrInvalidEventType   = errors.New("無効なイベントタイプです")
	ErrInvalidTimeSlot    = errors.New("終了時間は開始時間より後である必要があります")
	ErrInvalidCapacity    = errors.New("最大参加者数は1-100の範囲で設定してください")
)

// 参加者関連エラー
var (
	ErrParticipantNotFound        = errors.New("参加者が見つかりません")
	ErrParticipantAlreadyExists   = errors.New("既に参加登録済みです")
	ErrInvalidParticipantStatus   = errors.New("無効な参加ステータスです")
	ErrEventFull                  = errors.New("イベントが満員です")
	ErrRegistrationClosed         = errors.New("参加登録は締め切られています")
	ErrCannotReduceCapacity       = errors.New("現在の参加者数より少ない定員には変更できません")
)

// 権限関連エラー
var (
	ErrUnauthorized    = errors.New("権限がありません")
	ErrNotEventCreator = errors.New("この操作はイベント作成者のみ実行できます")
)