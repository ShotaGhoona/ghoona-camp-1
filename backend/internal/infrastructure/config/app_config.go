package config

import (
	"os"
)

// Config はアプリケーション設定
// 使用予定: BE-02-arch-02でDB設定追加、BE-02-arch-03でClerk設定追加
type Config struct {
	Port      string
	JWTSecret string
	Env       string
}

// LoadConfig は設定を読み込み
// 使用予定: main.goでアプリケーション起動時
func LoadConfig() (*Config, error) {
	config := &Config{
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "default-secret-key"),
		Env:       getEnv("APP_ENV", "development"),
	}

	return config, nil
}

// getEnv は環境変数を取得、なければデフォルト値を返す
// 使用予定: 全環境変数取得時
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// IsProduction は本番環境かどうかを判定
// 使用予定: Ginモード設定、ログレベル制御
func (c *Config) IsProduction() bool {
	return c.Env == "production"
}

// IsDevelopment は開発環境かどうかを判定  
// 使用予定: デバッグ機能有効化判定
func (c *Config) IsDevelopment() bool {
	return c.Env == "development"
}