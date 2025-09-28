package service

import (
	"net/url"
	"strings"

	"ghoona-camp-backend/internal/domain/user"
	"ghoona-camp-backend/internal/domain/user/entity"
	"ghoona-camp-backend/internal/domain/user/value"
)

// UserService はユーザー関連のビジネスロジックを処理する
type UserService struct{}

// NewUserService は新しいUserServiceを作成する
func NewUserService() *UserService {
	return &UserService{}
}

// ValidateUsername はユーザー名の形式と制約をバリデートする
func (s *UserService) ValidateUsername(username string) error {
	if len(username) < 3 || len(username) > 50 {
		return user.ErrInvalidUsername
	}
	return nil
}

// GenerateDefaultDisplayName はユーザー情報からデフォルトの表示名を生成する
func (s *UserService) GenerateDefaultDisplayName(user *entity.User) string {
	if user.Username != nil && *user.Username != "" {
		return *user.Username
	}
	return strings.Split(user.Email, "@")[0]
}

// ValidateSocialLinkURL はプラットフォームに基づいてURL形式をバリデートする
func (s *UserService) ValidateSocialLinkURL(platform value.Platform, urlStr string) error {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return user.ErrInvalidURL
	}

	if parsedURL.Scheme != "https" && parsedURL.Scheme != "http" {
		return user.ErrInvalidURLScheme
	}

	switch platform {
	case value.PlatformTwitter:
		if !strings.Contains(parsedURL.Host, "twitter.com") && !strings.Contains(parsedURL.Host, "x.com") {
			return user.ErrInvalidTwitterURL
		}
	case value.PlatformGitHub:
		if !strings.Contains(parsedURL.Host, "github.com") {
			return user.ErrInvalidGitHubURL
		}
	case value.PlatformLinkedIn:
		if !strings.Contains(parsedURL.Host, "linkedin.com") {
			return user.ErrInvalidLinkedInURL
		}
	}

	return nil
}

