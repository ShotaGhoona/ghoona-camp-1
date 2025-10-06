package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/value"
)

// User はユーザー基本情報のGORMモデル
type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ClerkID   string     `gorm:"uniqueIndex;not null;size:255" json:"clerkId"`
	Email     string     `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Username  *string    `gorm:"size:100" json:"username"`
	AvatarURL *string    `gorm:"type:text" json:"avatarUrl"`
	DiscordID *string    `gorm:"uniqueIndex;size:255" json:"discordId"`
	IsActive  bool       `gorm:"default:true" json:"isActive"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

// UserMetadata はユーザー詳細情報のGORMモデル
type UserMetadata struct {
	ID               uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID           uuid.UUID      `gorm:"type:uuid;uniqueIndex;not null" json:"userId"`
	DisplayName      *string        `gorm:"size:100" json:"displayName"`
	ProfileImageURL  *string        `gorm:"type:text" json:"profileImageUrl"`
	Tagline          *string        `gorm:"size:150" json:"tagline"`
	Bio              *string        `gorm:"type:text" json:"bio"`
	Vision           *string        `gorm:"type:text" json:"vision"`
	VisionPublic     bool           `gorm:"default:false" json:"visionPublic"`
	Timezone         string         `gorm:"size:50;default:'Asia/Tokyo'" json:"timezone"`
	Skills           pq.StringArray `gorm:"type:text[]" json:"skills"`
	Interests        pq.StringArray `gorm:"type:text[]" json:"interests"`
	User             User           `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
}

// UserSocialLink はユーザーソーシャルリンクのGORMモデル
type UserSocialLink struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`
	Platform  string    `gorm:"not null;size:50" json:"platform"`
	URL       string    `gorm:"not null;type:text" json:"url"`
	Title     *string   `gorm:"size:100" json:"title"`
	IsPublic  bool      `gorm:"default:true" json:"isPublic"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UserRival はユーザーライバル関係のGORMモデル
type UserRival struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`
	RivalUserID uuid.UUID `gorm:"type:uuid;not null;index" json:"rivalUserId"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	RivalUser   User      `gorm:"foreignKey:RivalUserID" json:"rivalUser"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TableName メソッドでテーブル名を指定
func (User) TableName() string {
	return "users"
}

func (UserMetadata) TableName() string {
	return "user_metadata"
}

func (UserSocialLink) TableName() string {
	return "user_social_links"
}

func (UserRival) TableName() string {
	return "user_rivals"
}

// ToEntity はGORMモデルからドメインエンティティへ変換します
func (u *User) ToEntity() (*entity.User, error) {
	status := value.UserStatusActive
	if !u.IsActive {
		status = value.UserStatusInactive
	}

	return &entity.User{
		BaseEntity: common.BaseEntity{
			ID:        common.UUID(u.ID),
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		ClerkID:   u.ClerkID,
		Email:     u.Email,
		Username:  u.Username,
		AvatarURL: u.AvatarURL,
		DiscordID: u.DiscordID,
		Status:    status,
	}, nil
}

// FromEntity はドメインエンティティからGORMモデルへ変換します
func FromEntity(domainUser *entity.User) *User {
	return &User{
		ID:        uuid.UUID(domainUser.ID),
		ClerkID:   domainUser.ClerkID,
		Email:     domainUser.Email,
		Username:  domainUser.Username,
		AvatarURL: domainUser.AvatarURL,
		DiscordID: domainUser.DiscordID,
		IsActive:  domainUser.Status.IsActiveState(),
		CreatedAt: domainUser.CreatedAt,
		UpdatedAt: domainUser.UpdatedAt,
	}
}

// ToEntity はGORMモデルからドメインエンティティへ変換します
func (um *UserMetadata) ToEntity() (*entity.UserMetadata, error) {
	visionPublic := value.PublicFalse
	if um.VisionPublic {
		visionPublic = value.PublicTrue
	}

	return &entity.UserMetadata{
		BaseEntity: common.BaseEntity{
			ID:        common.UUID(um.ID),
			CreatedAt: um.CreatedAt,
			UpdatedAt: um.UpdatedAt,
		},
		UserID:           common.UUID(um.UserID),
		DisplayName:      um.DisplayName,
		ProfileImageURL:  um.ProfileImageURL,
		Tagline:          um.Tagline,
		Bio:              um.Bio,
		Vision:           um.Vision,
		VisionPublic:     visionPublic,
		Timezone:         um.Timezone,
		Skills:           []string(um.Skills),
		Interests:        []string(um.Interests),
	}, nil
}

// FromEntityUserMetadata はドメインエンティティからGORMモデルへ変換します
func FromEntityUserMetadata(domainMetadata *entity.UserMetadata) *UserMetadata {
	return &UserMetadata{
		ID:               uuid.UUID(domainMetadata.ID),
		UserID:           uuid.UUID(domainMetadata.UserID),
		DisplayName:      domainMetadata.DisplayName,
		ProfileImageURL:  domainMetadata.ProfileImageURL,
		Tagline:          domainMetadata.Tagline,
		Bio:              domainMetadata.Bio,
		Vision:           domainMetadata.Vision,
		VisionPublic:     domainMetadata.VisionPublic.Bool(),
		Timezone:         domainMetadata.Timezone,
		Skills:           pq.StringArray(domainMetadata.Skills),
		Interests:        pq.StringArray(domainMetadata.Interests),
		CreatedAt:        domainMetadata.CreatedAt,
		UpdatedAt:        domainMetadata.UpdatedAt,
	}
}

// ToEntity はGORMモデルからドメインエンティティへ変換します
func (usl *UserSocialLink) ToEntity() (*entity.UserSocialLink, error) {
	platform := value.Platform(usl.Platform)
	
	isPublic := value.PublicFalse
	if usl.IsPublic {
		isPublic = value.PublicTrue
	}

	return &entity.UserSocialLink{
		BaseEntity: common.BaseEntity{
			ID:        common.UUID(usl.ID),
			CreatedAt: usl.CreatedAt,
			UpdatedAt: usl.UpdatedAt,
		},
		UserID:    common.UUID(usl.UserID),
		Platform:  platform,
		URL:       usl.URL,
		Title:     usl.Title,
		IsPublic:  isPublic,
	}, nil
}

// FromEntityUserSocialLink はドメインエンティティからGORMモデルへ変換します
func FromEntityUserSocialLink(domainLink *entity.UserSocialLink) *UserSocialLink {
	return &UserSocialLink{
		ID:        uuid.UUID(domainLink.ID),
		UserID:    uuid.UUID(domainLink.UserID),
		Platform:  domainLink.Platform.String(),
		URL:       domainLink.URL,
		Title:     domainLink.Title,
		IsPublic:  domainLink.IsPublic.Bool(),
		CreatedAt: domainLink.CreatedAt,
		UpdatedAt: domainLink.UpdatedAt,
	}
}

// ToEntity はGORMモデルからドメインエンティティへ変換します
func (ur *UserRival) ToEntity() (*entity.UserRival, error) {
	return &entity.UserRival{
		BaseEntity: common.BaseEntity{
			ID:        common.UUID(ur.ID),
			CreatedAt: ur.CreatedAt,
			UpdatedAt: ur.UpdatedAt,
		},
		UserID:      common.UUID(ur.UserID),
		RivalUserID: common.UUID(ur.RivalUserID),
	}, nil
}

// FromEntityUserRival はドメインエンティティからGORMモデルへ変換します
func FromEntityUserRival(domainRival *entity.UserRival) *UserRival {
	return &UserRival{
		ID:          uuid.UUID(domainRival.ID),
		UserID:      uuid.UUID(domainRival.UserID),
		RivalUserID: uuid.UUID(domainRival.RivalUserID),
		CreatedAt:   domainRival.CreatedAt,
		UpdatedAt:   domainRival.UpdatedAt,
	}
}