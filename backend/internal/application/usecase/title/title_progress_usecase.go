package title

import (
	"context"

	"ghoona-camp-backend/internal/application/dto/title"
	"ghoona-camp-backend/internal/domain/common"
	domainUser "ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/value"
	titleRepository "ghoona-camp-backend/internal/domain/title/repository"
	userRepository "ghoona-camp-backend/internal/domain/user/repository"
	"ghoona-camp-backend/internal/domain/title/service"
)

// AttendanceStatsProvider は出席統計情報提供のインターフェース
// 将来的に出席ドメインから実装される
type AttendanceStatsProvider interface {
	GetUserAttendanceStats(ctx context.Context, userID common.UUID) (*value.AttendanceStatistics, error)
}

// TitleProgressUseCase 称号進捗追跡のユースケース
type TitleProgressUseCase interface {
	GetUserProgress(ctx context.Context, userID common.UUID) (*title.UserTitleProgressResponse, error)
	CheckEligibleTitles(ctx context.Context, userID common.UUID, attendanceDays int) ([]*entity.Title, error)
}

type titleProgressUseCase struct {
	userRepo            userRepository.UserRepository
	titleRepo           titleRepository.TitleRepository
	achievementRepo     titleRepository.TitleAchievementRepository
	titleService        *service.TitleService
	attendanceProvider  AttendanceStatsProvider // 将来の出席サービス統合用
}

// NewTitleProgressUseCase 新しいTitleProgressUseCaseを作成
func NewTitleProgressUseCase(
	userRepo userRepository.UserRepository,
	titleRepo titleRepository.TitleRepository,
	achievementRepo titleRepository.TitleAchievementRepository,
	titleService *service.TitleService,
	attendanceProvider AttendanceStatsProvider,
) TitleProgressUseCase {
	return &titleProgressUseCase{
		userRepo:           userRepo,
		titleRepo:          titleRepo,
		achievementRepo:    achievementRepo,
		titleService:       titleService,
		attendanceProvider: attendanceProvider,
	}
}

// GetUserProgress ユーザーの称号進捗を取得
func (t *titleProgressUseCase) GetUserProgress(ctx context.Context, userID common.UUID) (*title.UserTitleProgressResponse, error) {
	// ユーザーの存在確認
	userEntity, err := t.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	// 現在の表示称号を取得
	currentAchievement, err := t.achievementRepo.GetCurrentByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var currentTitleResponse *title.AchievementResponse
	if currentAchievement != nil {
		// 現在称号の詳細情報を取得
		currentTitleEntity, err := t.titleRepo.GetByID(ctx, currentAchievement.TitleID)
		if err != nil {
			return nil, err
		}
		if currentTitleEntity != nil {
			currentTitleResponse = title.AchievementResponseFromEntity(currentAchievement, currentTitleEntity)
		}
	}

	// 出席統計を取得（将来実装）
	var currentDays int
	if t.attendanceProvider != nil {
		stats, err := t.attendanceProvider.GetUserAttendanceStats(ctx, userID)
		if err != nil {
			// 出席統計取得エラーは0日として処理を継続
			currentDays = 0
		} else if stats != nil {
			currentDays = stats.TotalAttendanceDays
		}
	}

	// 全称号を取得
	allTitles, err := t.titleRepo.GetActiveTitles(ctx)
	if err != nil {
		return nil, err
	}

	// ユーザーの獲得記録を取得
	achievements, err := t.achievementRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 獲得済み称号IDのマップを作成
	achievedTitleIDs := make(map[common.UUID]bool)
	for _, achievement := range achievements {
		achievedTitleIDs[achievement.TitleID] = true
	}

	// 次の称号と獲得可能称号を計算
	var nextTitle *title.TitleResponse
	var highestLevel int
	var eligibleTitles []title.TitleResponse

	for _, titleEntity := range allTitles {
		if titleEntity.Level > highestLevel {
			if achievedTitleIDs[titleEntity.ID] {
				highestLevel = titleEntity.Level
			}
		}

		// 未獲得で条件を満たす称号を獲得可能として追加
		if !achievedTitleIDs[titleEntity.ID] && currentDays >= titleEntity.RequiredDays {
			titleResponse := title.TitleResponseFromEntity(titleEntity)
			if titleResponse != nil {
				eligibleTitles = append(eligibleTitles, *titleResponse)
			}
		}

		// 次の称号を設定（未獲得の中で最も低いレベル）
		if !achievedTitleIDs[titleEntity.ID] {
			if nextTitle == nil || titleEntity.Level < nextTitle.Level {
				nextTitle = title.TitleResponseFromEntity(titleEntity)
			}
		}
	}

	// 進捗レスポンスを構築
	response := title.BuildUserTitleProgressResponse(
		userID,
		currentTitleResponse,
		nextTitle,
		currentDays,
		highestLevel,
		eligibleTitles,
	)

	return response, nil
}

// CheckEligibleTitles 獲得可能な称号をチェック
func (t *titleProgressUseCase) CheckEligibleTitles(ctx context.Context, userID common.UUID, attendanceDays int) ([]*entity.Title, error) {
	// ユーザーの存在確認
	userEntity, err := t.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, domainUser.ErrUserNotFound
	}

	// アクティブな称号を取得
	allTitles, err := t.titleRepo.GetActiveTitles(ctx)
	if err != nil {
		return nil, err
	}

	// ユーザーの獲得記録を取得
	achievements, err := t.achievementRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 獲得済み称号IDのマップを作成
	achievedTitleIDs := make(map[common.UUID]bool)
	for _, achievement := range achievements {
		achievedTitleIDs[achievement.TitleID] = true
	}

	// 獲得可能な称号をフィルタリング
	eligibleTitles := make([]*entity.Title, 0)
	for _, titleEntity := range allTitles {
		// 未獲得で条件を満たす称号
		if !achievedTitleIDs[titleEntity.ID] && attendanceDays >= titleEntity.RequiredDays {
			eligibleTitles = append(eligibleTitles, titleEntity)
		}
	}

	return eligibleTitles, nil
}