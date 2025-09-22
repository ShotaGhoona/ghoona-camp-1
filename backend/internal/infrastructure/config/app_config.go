package config

import (
	"os"
	"strconv"
	"time"
)

// Config はアプリケーション設定
type Config struct {
	Port               string
	ClerkSecretKey     string // Clerk JWT検証用秘密鍵
	ClerkWebhookSecret string // Clerk Webhook検証用秘密鍵
	Env                string
	
	// Database 設定 (BE-02-arch-02で追加)
	Database DatabaseConfig
}

// DatabaseConfig はデータベース関連の設定
type DatabaseConfig struct {
	URL             string        // データベース接続文字列
	MaxOpenConns    int           // 最大オープン接続数
	MaxIdleConns    int           // 最大アイドル接続数
	ConnMaxLifetime time.Duration // 接続最大生存時間
	LogLevel        string        // ログレベル
}

// LoadConfig は設定を読み込み
// 使用予定: main.goでアプリケーション起動時
func LoadConfig() (*Config, error) {
	config := &Config{
		Port:               getEnv("PORT", "8080"),
		ClerkSecretKey:     getEnv("CLERK_SECRET_KEY", ""),
		ClerkWebhookSecret: getEnv("CLERK_WEBHOOK_SECRET", ""),
		Env:                getEnv("APP_ENV", "development"),
		Database: DatabaseConfig{
			URL:             getEnv("DATABASE_URL", ""),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", "30m"),
			LogLevel:        getEnv("DB_LOG_LEVEL", "warn"),
		},
	}

	return config, nil
}

// getEnv は環境変数を取得、なければデフォルト値を返す
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt は環境変数を整数として取得
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsDuration は環境変数を時間として取得
func getEnvAsDuration(key, defaultValue string) time.Duration {
	value := getEnv(key, defaultValue)
	if duration, err := time.ParseDuration(value); err == nil {
		return duration
	}
	// パースに失敗した場合はデフォルト値をパース
	if duration, err := time.ParseDuration(defaultValue); err == nil {
		return duration
	}
	return 30 * time.Minute // 最終フォールバック
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