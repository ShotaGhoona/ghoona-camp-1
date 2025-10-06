package title

import (
	"ghoona-camp-backend/internal/domain/common"
)

// Title domain errors
var (
	// Title errors
	ErrTitleNotFound        = common.ErrNotFound
	ErrTitleInactive        = common.NewDomainError("TITLE_INACTIVE", "非アクティブな称号です", nil)
	ErrInvalidTitleLevel    = common.NewDomainError("INVALID_TITLE_LEVEL", "無効な称号レベルです", nil)
	
	// Title validation errors
	ErrInvalidLevel          = common.NewDomainError("INVALID_LEVEL", "レベルは1-8の範囲で設定してください", nil)
	ErrEmptyNameJP          = common.NewDomainError("EMPTY_NAME_JP", "日本語名は必須です", nil)
	ErrEmptyNameEN          = common.NewDomainError("EMPTY_NAME_EN", "英語名は必須です", nil)
	ErrEmptyDescription     = common.NewDomainError("EMPTY_DESCRIPTION", "説明は必須です", nil)
	ErrInvalidRequiredDays  = common.NewDomainError("INVALID_REQUIRED_DAYS", "必要日数は1以上で設定してください", nil)
	ErrNameJPTooLong        = common.NewDomainError("NAME_JP_TOO_LONG", "日本語名は100文字以内で設定してください", nil)
	ErrNameENTooLong        = common.NewDomainError("NAME_EN_TOO_LONG", "英語名は100文字以内で設定してください", nil)
	ErrDescriptionTooLong   = common.NewDomainError("DESCRIPTION_TOO_LONG", "説明は2000文字以内で設定してください", nil)
	ErrColorThemeTooLong    = common.NewDomainError("COLOR_THEME_TOO_LONG", "カラーテーマは50文字以内で設定してください", nil)
	ErrInvalidActiveFlag    = common.NewDomainError("INVALID_ACTIVE_FLAG", "無効なアクティブフラグです", nil)
	ErrDuplicateLevel       = common.NewDomainError("DUPLICATE_LEVEL", "同じレベルの称号が既に存在します", nil)
	
	// Achievement errors  
	ErrAchievementNotFound      = common.ErrNotFound
	ErrTitleNotAchieved        = common.NewDomainError("TITLE_NOT_ACHIEVED", "未獲得の称号です", nil)
	ErrTitleAlreadyCurrent     = common.NewDomainError("TITLE_ALREADY_CURRENT", "既に現在の称号に設定済み", nil)
	ErrAchievementAlreadyExists = common.ErrAlreadyExists
	
	// Achievement validation errors
	ErrEmptyUserID           = common.NewDomainError("EMPTY_USER_ID", "ユーザーIDは必須です", nil)
	ErrEmptyTitleID          = common.NewDomainError("EMPTY_TITLE_ID", "称号IDは必須です", nil)
	ErrInvalidCurrentFlag    = common.NewDomainError("INVALID_CURRENT_FLAG", "無効な現在フラグです", nil)
	ErrDuplicateAchievement  = common.NewDomainError("DUPLICATE_ACHIEVEMENT", "既に獲得済みの称号です", nil)
	ErrMultipleCurrentTitles = common.NewDomainError("MULTIPLE_CURRENT_TITLES", "現在称号は1つまでしか設定できません", nil)
)