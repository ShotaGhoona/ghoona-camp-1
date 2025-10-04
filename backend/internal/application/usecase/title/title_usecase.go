package title

import (
	"context"

	"ghoona-camp-backend/internal/application/dto/title"
	"ghoona-camp-backend/internal/domain/common"
	domainTitle "ghoona-camp-backend/internal/domain/title"
	"ghoona-camp-backend/internal/domain/title/entity"
	"ghoona-camp-backend/internal/domain/title/repository"
	"ghoona-camp-backend/internal/domain/title/service"
)

// TitleUseCase 称号基本操作のユースケース
type TitleUseCase interface {
	GetAllTitles(ctx context.Context, includeInactive bool) (*title.TitleListResponse, error)
	GetTitleByID(ctx context.Context, titleID common.UUID) (*title.TitleResponse, error)
}

type titleUseCase struct {
	titleRepo         repository.TitleRepository
	validationService *service.TitleValidationService
}

// NewTitleUseCase 新しいTitleUseCaseを作成
func NewTitleUseCase(
	titleRepo repository.TitleRepository,
	validationService *service.TitleValidationService,
) TitleUseCase {
	return &titleUseCase{
		titleRepo:         titleRepo,
		validationService: validationService,
	}
}

// GetAllTitles 全称号を取得
func (t *titleUseCase) GetAllTitles(ctx context.Context, includeInactive bool) (*title.TitleListResponse, error) {
	var titles []*entity.Title
	var err error

	if includeInactive {
		titles, err = t.titleRepo.GetAll(ctx)
	} else {
		titles, err = t.titleRepo.GetActiveTitles(ctx)
	}

	if err != nil {
		return nil, err
	}

	titleResponses := title.TitleListFromEntities(titles)

	return &title.TitleListResponse{
		Titles: titleResponses,
		Total:  len(titleResponses),
	}, nil
}

// GetTitleByID IDで称号を取得
func (t *titleUseCase) GetTitleByID(ctx context.Context, titleID common.UUID) (*title.TitleResponse, error) {
	titleEntity, err := t.titleRepo.GetByID(ctx, titleID)
	if err != nil {
		return nil, err
	}
	if titleEntity == nil {
		return nil, domainTitle.ErrTitleNotFound
	}

	return title.TitleResponseFromEntity(titleEntity), nil
}