package clerk

import (
	"context"
	"errors"
	"fmt"
	"ghoona-camp-backend/internal/infrastructure/config"
)

// TODO: BE-02-arch-02で実際のClerk連携を実装
// 現在は基盤のみ実装

// AuthService はClerk認証サービス
type AuthService struct {
	secretKey string
}

// NewAuthService は新しいClerk認証サービスを作成
func NewAuthService(cfg *config.ClerkConfig) *AuthService {
	return &AuthService{
		secretKey: cfg.SecretKey,
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
	// TODO: 実際のClerk JWT検証ロジックを実装
	return nil, errors.New("not implemented: Clerk JWT verification")
}

// GetUser はユーザー情報を取得
func (s *AuthService) GetUser(ctx context.Context, userID string) (*ClerkUser, error) {
	// TODO: Clerk APIからユーザー情報を取得
	return nil, errors.New("not implemented: Clerk user fetch")
}

// CreateUser はユーザーを作成
func (s *AuthService) CreateUser(ctx context.Context, email, password string) (*ClerkUser, error) {
	// TODO: Clerk APIでユーザーを作成
	return nil, errors.New("not implemented: Clerk user creation")
}

// UpdateUser はユーザー情報を更新
func (s *AuthService) UpdateUser(ctx context.Context, userID string, updates map[string]interface{}) (*ClerkUser, error) {
	// TODO: Clerk APIでユーザー情報を更新
	return nil, errors.New("not implemented: Clerk user update")
}

// DeleteUser はユーザーを削除
func (s *AuthService) DeleteUser(ctx context.Context, userID string) error {
	// TODO: Clerk APIでユーザーを削除
	return errors.New("not implemented: Clerk user deletion")
}

// ParseWebhook はClerk Webhookを解析
func (s *AuthService) ParseWebhook(payload []byte, signature string) (map[string]interface{}, error) {
	// TODO: Clerk Webhook検証と解析を実装
	return nil, errors.New("not implemented: Clerk webhook parsing")
}

// validateTokenFormat はトークンフォーマットを検証（ヘルパー）
func (s *AuthService) validateTokenFormat(token string) error {
	if token == "" {
		return fmt.Errorf("token is empty")
	}
	// TODO: より詳細な検証ロジック
	return nil
}