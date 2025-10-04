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
	
	// Achievement errors  
	ErrAchievementNotFound      = common.ErrNotFound
	ErrTitleNotAchieved        = common.NewDomainError("TITLE_NOT_ACHIEVED", "未獲得の称号です", nil)
	ErrTitleAlreadyCurrent     = common.NewDomainError("TITLE_ALREADY_CURRENT", "既に現在の称号に設定済み", nil)
	ErrAchievementAlreadyExists = common.ErrAlreadyExists
)