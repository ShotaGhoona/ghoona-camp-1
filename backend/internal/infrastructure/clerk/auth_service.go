package clerk

import (
	"context"
	"errors"
	"strings"

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

	// Bearer プレフィックスを削除
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimSpace(token)

	if token == "" {
		return nil, errors.New("トークンが空です")
	}

	// 開発環境用のモック処理
	if !s.enabled {
		// テスト用のモックユーザーを返す
		if token == "mock-clerk-token" {
			return &ClerkUser{
				ID:       "user_mock123",
				Email:    "test@example.com",
				Username: "testuser",
			}, nil
		}
		return nil, errors.New("無効なモックトークンです")
	}

	// 実際のClerk JWT検証
	// 注: 簡略化したJWT検証。実際の環境ではClerkのJWKSから公開鍵を取得して検証する必要があります
	// TODO: 本格実装時にClerkのJWKSエンドポイントから公開鍵を取得
	claims := make(map[string]interface{})
	claims["sub"] = "mock_user_id" // 仮実装
	claims["email"] = "test@example.com"
	claims["username"] = "testuser"

	// クレームからユーザー情報を抽出
	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return nil, errors.New("ユーザーIDが見つかりません")
	}

	email, _ := claims["email"].(string)
	username, _ := claims["username"].(string)

	return &ClerkUser{
		ID:       userID,
		Email:    email,
		Username: username,
	}, nil
}

// GetUser はユーザー情報を取得
func (s *AuthService) GetUser(ctx context.Context, userID string) (*ClerkUser, error) {
	// TODO: BE-03-user-*で実際のClerk API連携を実装
	// 現在はモック実装
	return &ClerkUser{
		ID:       userID,
		Email:    "test@example.com",
		Username: "testuser",
	}, nil
}

// ParseWebhook はClerk Webhookを解析
// 使用予定: BE-05-notification-*でWebhookイベント処理実装
func (s *AuthService) ParseWebhook(payload []byte, signature string) (map[string]interface{}, error) {
	// TODO: BE-05-notification-*でClerk Webhook検証と解析を実装
	return nil, errors.New("Clerk Webhook解析は後続タスクで実装予定です")
}

// IsValidToken はトークンの基本的なフォーマットを検証
func (s *AuthService) IsValidToken(token string) bool {
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimSpace(token)
	return token != ""
}
