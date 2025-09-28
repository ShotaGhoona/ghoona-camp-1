package clerk

import (
	"context"
	"errors"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"ghoona-camp-backend/internal/infrastructure/config"
)

// AuthService はClerk認証サービス
type AuthService struct {
	secretKey string
	enabled   bool
}

// NewAuthService は新しいClerk認証サービスを作成
func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		secretKey: cfg.ClerkSecretKey,
		enabled:   cfg.ClerkSecretKey != "",
	}
}

// ClerkUser はClerkユーザー情報
type ClerkUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	// TODO: 詳細フィールドは後で追加
}

// VerifyToken はJWTトークンを検証
func (s *AuthService) VerifyToken(ctx context.Context, token string) (*ClerkUser, error) {
	if s.secretKey == "" {
		return nil, errors.New("Clerk secret key が設定されていません")
	}

	if token == "" {
		return nil, errors.New("トークンが空です")
	}

	// 開発環境用のモック処理
	if !s.enabled {
		if token == "mock-clerk-token" {
			return &ClerkUser{
				ID:       "user_mock123",
				Email:    "test@example.com",
				Username: "testuser",
			}, nil
		}
		return nil, errors.New("無効なモックトークンです")
	}

	// Clerk SDK v2使用
	clerk.SetKey(s.secretKey)

	claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	return &ClerkUser{
		ID:       claims.Subject,
		Email:    "clerk-user@example.com", // TODO: 実際のユーザー情報取得APIを使用
		Username: "clerk-user",            // TODO: 実際のユーザー情報取得APIを使用
	}, nil
}

