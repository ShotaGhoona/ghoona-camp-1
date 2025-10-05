package title

import (
	"context"
	"sort"
	"strings"

	"ghoona-camp-backend/internal/application/dto/title"
	"ghoona-camp-backend/internal/application/transaction"
	"ghoona-camp-backend/internal/domain/common"
	domainTitle "ghoona-camp-backend/internal/domain/title"
	domainUser "ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/title/entity"
	titleRepository "ghoona-camp-backend/internal/domain/title/repository"
	userRepository "ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/domain/title/service"
)

// TitleAchievementUseCase 称号獲得記録操作のユースケース
type TitleAchievementUseCase interface {
	GetUserAchievements(ctx context.Context, userID common.UUID, includeProgress bool) (*title.AchievementListResponse, error)
	GetUserAchievementsWithUserInfo(ctx context.Context, userID common.UUID, includeProgress bool, sortBy string, order string) (*title.UserAchievementsResponse, error)
	SetCurrentTitle(ctx context.Context, userID common.UUID, titleID common.UUID) (*title.AchievementResponse, error)
}

type titleAchievementUseCase struct {
	userRepo          userRepository.UserRepository
	userMetadataRepo  userRepository.UserMetadataRepository
	titleRepo         titleRepository.TitleRepository
	achievementRepo   titleRepository.TitleAchievementRepository
	titleService      *service.TitleService
	validationService *service.TitleValidationService
	txManager         transaction.Manager
}

// NewTitleAchievementUseCase 新しいTitleAchievementUseCaseを作成
func NewTitleAchievementUseCase(
	userRepo userRepository.UserRepository,
	userMetadataRepo userRepository.UserMetadataRepository,
	titleRepo titleRepository.TitleRepository,
	achievementRepo titleRepository.TitleAchievementRepository,
	titleService *service.TitleService,
	validationService *service.TitleValidationService,
	txManager transaction.Manager,
) TitleAchievementUseCase {
	return &titleAchievementUseCase{
		userRepo:          userRepo,
		userMetadataRepo:  userMetadataRepo,
		titleRepo:         titleRepo,
		achievementRepo:   achievementRepo,
		titleService:      titleService,
		validationService: validationService,
		txManager:         txManager,
	}
}

// GetUserAchievements ユーザーの称号獲得履歴を取得
func (t *titleAchievementUseCase) GetUserAchievements(ctx context.Context, userID common.UUID, includeProgress bool) (*title.AchievementListResponse, error) {
	// ユーザーの存在確認
	userEntity, err := t.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	// ユーザーの獲得記録を取得
	achievements, err := t.achievementRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 称号IDリストを収集
	titleIDs := make([]common.UUID, 0, len(achievements))
	titleIDMap := make(map[common.UUID]bool)
	for _, achievement := range achievements {
		if !titleIDMap[achievement.TitleID] {
			titleIDs = append(titleIDs, achievement.TitleID)
			titleIDMap[achievement.TitleID] = true
		}
	}

	// 称号情報をバッチ取得
	titlesMap := make(map[common.UUID]*entity.Title)
	for _, titleID := range titleIDs {
		titleEntity, err := t.titleRepo.GetByID(ctx, titleID)
		if err != nil {
			return nil, err
		}
		if titleEntity != nil {
			titlesMap[titleID] = titleEntity
		}
	}

	// レスポンスを構築
	response := title.BuildAchievementListResponse(userID, achievements, titlesMap)

	return response, nil
}

// GetUserAchievementsWithUserInfo ユーザー情報と称号獲得履歴を取得（API仕様準拠）
func (t *titleAchievementUseCase) GetUserAchievementsWithUserInfo(ctx context.Context, userID common.UUID, includeProgress bool, sortBy string, order string) (*title.UserAchievementsResponse, error) {
	// ユーザーの存在確認
	userEntity, err := t.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	// ユーザーメタデータを取得
	userMetadata, err := t.userMetadataRepo.GetByUserID(ctx, userID)
	if err != nil {
		// メタデータが存在しない場合はエラーにしない
		userMetadata = nil
	}

	// ユーザーの獲得記録を取得
	achievements, err := t.achievementRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 称号IDリストを収集
	titleIDs := make([]common.UUID, 0, len(achievements))
	titleIDMap := make(map[common.UUID]bool)
	for _, achievement := range achievements {
		if !titleIDMap[achievement.TitleID] {
			titleIDs = append(titleIDs, achievement.TitleID)
			titleIDMap[achievement.TitleID] = true
		}
	}

	// 称号情報をバッチ取得
	titlesMap := make(map[common.UUID]*entity.Title)
	for _, titleID := range titleIDs {
		titleEntity, err := t.titleRepo.GetByID(ctx, titleID)
		if err != nil {
			return nil, err
		}
		if titleEntity != nil {
			titlesMap[titleID] = titleEntity
		}
	}

	// ソート処理
	sortedAchievements := make([]*entity.TitleAchievement, len(achievements))
	copy(sortedAchievements, achievements)

	switch strings.ToLower(sortBy) {
	case "achieved_at":
		if strings.ToLower(order) == "desc" {
			sort.Slice(sortedAchievements, func(i, j int) bool {
				return sortedAchievements[i].AchievedAt.After(sortedAchievements[j].AchievedAt)
			})
		} else {
			sort.Slice(sortedAchievements, func(i, j int) bool {
				return sortedAchievements[i].AchievedAt.Before(sortedAchievements[j].AchievedAt)
			})
		}
	case "level":
		if strings.ToLower(order) == "desc" {
			sort.Slice(sortedAchievements, func(i, j int) bool {
				titleI := titlesMap[sortedAchievements[i].TitleID]
				titleJ := titlesMap[sortedAchievements[j].TitleID]
				if titleI != nil && titleJ != nil {
					return titleI.Level > titleJ.Level
				}
				return false
			})
		} else {
			sort.Slice(sortedAchievements, func(i, j int) bool {
				titleI := titlesMap[sortedAchievements[i].TitleID]
				titleJ := titlesMap[sortedAchievements[j].TitleID]
				if titleI != nil && titleJ != nil {
					return titleI.Level < titleJ.Level
				}
				return false
			})
		}
	}

	// レスポンスDTO構築
	achievementResponses := title.AchievementListFromEntities(sortedAchievements, titlesMap)

	// ユーザー情報を構築
	displayName := ""
	if userMetadata != nil && userMetadata.DisplayName != nil {
		displayName = *userMetadata.DisplayName
	} else if userEntity.Username != nil {
		displayName = *userEntity.Username
	}

	userInfo := title.UserInfo{
		ID:          userEntity.ID,
		DisplayName: displayName,
		Username:    userEntity.Username,
		AvatarURL:   userEntity.AvatarURL,
	}

	response := &title.UserAchievementsResponse{
		User:         userInfo,
		Achievements: achievementResponses,
	}

	return response, nil
}

// SetCurrentTitle 現在表示称号を設定
func (t *titleAchievementUseCase) SetCurrentTitle(ctx context.Context, userID common.UUID, titleID common.UUID) (*title.AchievementResponse, error) {
	var response *title.AchievementResponse

	err := t.txManager.ExecuteInTx(ctx, func(txCtx context.Context) error {
		// ユーザーの存在確認
		userEntity, err := t.userRepo.GetByID(txCtx, userID)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return domainUser.ErrUserNotFound
		}

		// 称号の存在確認
		titleEntity, err := t.titleRepo.GetByID(txCtx, titleID)
		if err != nil {
			return err
		}
		if titleEntity == nil {
			return domainTitle.ErrTitleNotFound
		}

		// 称号がアクティブかチェック
		if !titleEntity.IsActive.Bool() {
			return domainTitle.ErrTitleInactive
		}

		// ユーザーが該当称号を獲得しているかチェック
		achievement, err := t.achievementRepo.GetByUserIDAndTitleID(txCtx, userID, titleID)
		if err != nil {
			return err
		}
		if achievement == nil {
			return domainTitle.ErrTitleNotAchieved
		}

		// 既に現在称号の場合はエラー
		if achievement.IsCurrent.Bool() {
			return domainTitle.ErrTitleAlreadyCurrent
		}

		// 現在称号をアトミックに変更（排他制御付き）
		if err := t.achievementRepo.SetCurrent(txCtx, userID, titleID); err != nil {
			return err
		}

		// 更新後の獲得記録を取得
		updatedAchievement, err := t.achievementRepo.GetByUserIDAndTitleID(txCtx, userID, titleID)
		if err != nil {
			return err
		}

		// DTOに変換
		response = title.AchievementResponseFromEntity(updatedAchievement, titleEntity)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}