package user

import "errors"

// User domain errors
var (
	// User errors
	ErrUserNotFound        = errors.New("ユーザーが見つかりません")
	ErrInvalidUsername     = errors.New("ユーザー名は3文字以上50文字以内で入力してください")
	ErrDuplicateEmail      = errors.New("このメールアドレスは既に使用されています")
	ErrDuplicateClerkID    = errors.New("このClerk IDは既に使用されています")
	ErrInvalidEmail        = errors.New("無効なメールアドレスです")
	ErrInvalidClerkID      = errors.New("無効なClerk IDです")

	// User metadata errors
	ErrUserMetadataNotFound    = errors.New("ユーザーメタデータが見つかりません")
	ErrUserMetadataAlreadyExists = errors.New("ユーザーメタデータが既に存在します")

	// Social link errors
	ErrUserSocialLinkNotFound = errors.New("ソーシャルリンクが見つかりません")
	ErrInvalidURL          = errors.New("無効なURL形式です")
	ErrInvalidURLScheme    = errors.New("URLはhttp またはhttpsスキームを使用してください")
	ErrInvalidTwitterURL   = errors.New("無効なTwitter URLです")
	ErrInvalidGitHubURL    = errors.New("無効なGitHub URLです")
	ErrInvalidLinkedInURL  = errors.New("無効なLinkedIn URLです")
	ErrInvalidPlatform     = errors.New("無効なプラットフォームです")
	ErrInvalidTitle        = errors.New("タイトルは100文字以内で入力してください")
	ErrDuplicatePlatform   = errors.New("同一プラットフォームのリンクは1つまでです")

	// Rival errors
	ErrCannotRivalSelf     = errors.New("自分自身をライバルに設定できません")
	ErrRivalLimitExceeded  = errors.New("ライバルは最大3人まで設定できます")
	ErrDuplicateRival      = errors.New("既にライバル関係が存在します")

	// Metadata errors
	ErrInvalidDisplayName  = errors.New("表示名は100文字以内で入力してください")
	ErrInvalidTagline      = errors.New("一言プロフィールは150文字以内で入力してください")
	ErrInvalidBio          = errors.New("自己紹介は1000文字以内で入力してください")
	ErrInvalidVision       = errors.New("ビジョンは2000文字以内で入力してください")
	ErrTooManySkills       = errors.New("スキルは最大20個まで設定できます")
	ErrTooManyInterests    = errors.New("興味・関心は最大20個まで設定できます")
)