package user

import (
	"ghoona-camp-backend/internal/domain/common"
)

// User domain errors
var (
	// User errors
	ErrUserNotFound        = common.ErrNotFound
	ErrInvalidUsername     = common.NewValidationError("username", "ユーザー名は3文字以上50文字以内で入力してください")
	ErrDuplicateEmail      = common.ErrDuplicateEntry
	ErrDuplicateClerkID    = common.ErrDuplicateEntry
	ErrInvalidEmail        = common.NewValidationError("email", "無効なメールアドレスです")
	ErrInvalidClerkID      = common.NewValidationError("clerkId", "無効なClerk IDです")

	// User metadata errors
	ErrUserMetadataNotFound      = common.ErrNotFound
	ErrUserMetadataAlreadyExists = common.ErrAlreadyExists

	// Social link errors
	ErrUserSocialLinkNotFound = common.ErrNotFound
	ErrInvalidURL             = common.NewValidationError("url", "無効なURL形式です")
	ErrInvalidURLScheme       = common.NewValidationError("url", "URLはhttp またはhttpsスキームを使用してください")
	ErrInvalidTwitterURL      = common.NewValidationError("url", "無効なTwitter URLです")
	ErrInvalidGitHubURL       = common.NewValidationError("url", "無効なGitHub URLです")
	ErrInvalidLinkedInURL     = common.NewValidationError("url", "無効なLinkedIn URLです")
	ErrInvalidPlatform        = common.NewValidationError("platform", "無効なプラットフォームです")
	ErrInvalidTitle           = common.NewValidationError("title", "タイトルは100文字以内で入力してください")
	ErrDuplicatePlatform      = common.ErrDuplicateEntry

	// Rival errors
	ErrCannotRivalSelf    = common.NewDomainError("CANNOT_RIVAL_SELF", "自分自身をライバルに設定できません", nil)
	ErrRivalLimitExceeded = common.NewDomainError("RIVAL_LIMIT_EXCEEDED", "ライバルは最大3人まで設定できます", nil)
	ErrDuplicateRival     = common.ErrDuplicateEntry

	// Metadata errors
	ErrInvalidDisplayName = common.NewValidationError("displayName", "表示名は100文字以内で入力してください")
	ErrInvalidTagline     = common.NewValidationError("tagline", "一言プロフィールは150文字以内で入力してください")
	ErrInvalidBio         = common.NewValidationError("bio", "自己紹介は1000文字以内で入力してください")
	ErrInvalidVision      = common.NewValidationError("vision", "ビジョンは2000文字以内で入力してください")
	ErrTooManySkills      = common.NewValidationError("skills", "スキルは最大20個まで設定できます")
	ErrTooManyInterests   = common.NewValidationError("interests", "興味・関心は最大20個まで設定できます")
)