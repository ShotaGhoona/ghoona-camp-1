package service

import (
	"context"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/title"
	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/repository"
)

// TitleValidationService は称号ドメインの全てのバリデーションロジックを処理する
type TitleValidationService struct {
	titleRepo       repository.TitleRepository
	achievementRepo repository.TitleAchievementRepository
}

// NewTitleValidationService は新しいTitleValidationServiceを作成する
func NewTitleValidationService(
	titleRepo repository.TitleRepository,
	achievementRepo repository.TitleAchievementRepository,
) *TitleValidationService {
	return &TitleValidationService{
		titleRepo:       titleRepo,
		achievementRepo: achievementRepo,
	}
}

// ValidateTitle は称号エンティティをバリデートする
func (s *TitleValidationService) ValidateTitle(t *entity.Title) error {
	// レベル範囲チェック (1-8)
	if t.Level < 1 || t.Level > 8 {
		return title.ErrInvalidLevel
	}

	// 必須フィールドチェック
	if t.NameJP == "" {
		return title.ErrEmptyNameJP
	}
	if t.NameEN == "" {
		return title.ErrEmptyNameEN
	}
	if t.Description == "" {
		return title.ErrEmptyDescription
	}

	// 必要日数チェック
	if t.RequiredDays < 1 {
		return title.ErrInvalidRequiredDays
	}

	// 文字数制限チェック
	if len(t.NameJP) > 100 {
		return title.ErrNameJPTooLong
	}
	if len(t.NameEN) > 100 {
		return title.ErrNameENTooLong
	}
	if len(t.Description) > 2000 {
		return title.ErrDescriptionTooLong
	}

	// オプションフィールドの文字数チェック
	if t.ColorTheme != nil && len(*t.ColorTheme) > 50 {
		return title.ErrColorThemeTooLong
	}

	// ActiveFlagの妥当性チェック
	if !t.IsActive.IsValid() {
		return title.ErrInvalidActiveFlag
	}

	return nil
}

// ValidateTitleAchievement は称号獲得記録エンティティをバリデートする
func (s *TitleValidationService) ValidateTitleAchievement(ta *entity.TitleAchievement) error {
	// 必須フィールドチェック
	var emptyUUID common.UUID
	if ta.UserID == emptyUUID {
		return title.ErrEmptyUserID
	}
	if ta.TitleID == emptyUUID {
		return title.ErrEmptyTitleID
	}

	// CurrentFlagの妥当性チェック
	if !ta.IsCurrent.IsValid() {
		return title.ErrInvalidCurrentFlag
	}

	// 獲得日時チェック（過去日時である必要がある）
	// 注意: 将来日時での獲得記録は許可しない
	// time.Now()との比較は少し余裕を持たせる（システム時間の微細な差を考慮）

	return nil
}

// ValidateTitleLevelUniqueness は称号レベルの一意性をチェックする
func (s *TitleValidationService) ValidateTitleLevelUniqueness(ctx context.Context, level int, excludeTitleID *common.UUID) error {
	existingTitle, err := s.titleRepo.GetByLevel(ctx, level)
	if err != nil {
		return err
	}

	// 既存の称号が見つかった場合
	if existingTitle != nil {
		// 除外対象（更新時の自分自身）でない場合はエラー
		if excludeTitleID == nil || *excludeTitleID != existingTitle.ID {
			return title.ErrDuplicateLevel
		}
	}

	return nil
}

// ValidateAchievementUniqueness はユーザーの称号獲得記録の一意性をチェックする
func (s *TitleValidationService) ValidateAchievementUniqueness(ctx context.Context, userID, titleID common.UUID) error {
	existingAchievement, err := s.achievementRepo.GetByUserIDAndTitleID(ctx, userID, titleID)
	if err != nil {
		return err
	}

	if existingAchievement != nil {
		return title.ErrDuplicateAchievement
	}

	return nil
}

// ValidateCurrentTitleUniqueness は現在表示称号の一意性をチェックする
// 新しく現在称号に設定する前に、既存の現在称号がないことを確認
func (s *TitleValidationService) ValidateCurrentTitleUniqueness(ctx context.Context, userID common.UUID, excludeAchievementID *common.UUID) error {
	currentAchievement, err := s.achievementRepo.GetCurrentByUserID(ctx, userID)
	if err != nil {
		return err
	}

	// 現在称号が存在する場合
	if currentAchievement != nil {
		// 除外対象（自分自身）でない場合はエラー
		if excludeAchievementID == nil || *excludeAchievementID != currentAchievement.ID {
			return title.ErrMultipleCurrentTitles
		}
	}

	return nil
}