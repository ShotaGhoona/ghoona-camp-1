package user

import "errors"

// User domain specific errors

// Rival management errors
var (
	ErrCannotAddSelfAsRival    = errors.New("自分自身をライバルに設定することはできません")
	ErrMaxRivalsLimitExceeded  = errors.New("ライバルは最大3人までしか設定できません")
	ErrRivalAlreadyExists      = errors.New("このユーザーは既にライバルに設定されています")
)

// User validation errors
var (
	ErrUserClerkIDRequired     = errors.New("Clerk IDは必須です")
	ErrUserEmailRequired       = errors.New("メールアドレスは必須です")
	ErrUserDisplayNameRequired = errors.New("表示名は必須です")
)

// User metadata validation errors
var (
	ErrUserMetadataUserIDRequired = errors.New("ユーザーIDは必須です")
)

// User rival validation errors
var (
	ErrUserRivalUserIDRequired  = errors.New("ユーザーIDは必須です")
	ErrUserRivalRivalIDRequired = errors.New("ライバルユーザーIDは必須です")
	ErrUserRivalSameUser        = errors.New("自分自身をライバルに設定することはできません")
)

// User vision validation errors
var (
	ErrUserVisionUserIDRequired = errors.New("ユーザーIDは必須です")
)

// User social link validation errors
var (
	ErrUserSocialLinkUserIDRequired = errors.New("ユーザーIDは必須です")
)