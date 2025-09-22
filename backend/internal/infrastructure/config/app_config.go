package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config はアプリケーション設定
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Clerk    ClerkConfig
	Discord  DiscordConfig
	JWT      JWTConfig
}

// ServerConfig はサーバー設定
type ServerConfig struct {
	Port        string
	Environment string
	Debug       bool
}

// DatabaseConfig はデータベース設定
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

// ClerkConfig はClerk認証設定
type ClerkConfig struct {
	SecretKey   string
	WebhookSecret string
}

// DiscordConfig はDiscord連携設定
type DiscordConfig struct {
	BotToken  string
	GuildID   string
	WebhookURL string
}

// JWTConfig はJWT設定
type JWTConfig struct {
	SecretKey string
	ExpiryHours int
}

// LoadConfig は設定を読み込み
func LoadConfig() (*Config, error) {
	// .envファイルを読み込み（存在する場合）
	if err := godotenv.Load(); err != nil {
		// .envファイルが存在しない場合はエラーとしない
		fmt.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		Server: ServerConfig{
			Port:        getEnv("PORT", "8080"),
			Environment: getEnv("ENVIRONMENT", "development"),
			Debug:       getEnvBool("DEBUG", true),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5433),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Database: getEnv("DB_NAME", "ghoona_camp"),
			SSLMode:  getEnv("DB_SSL_MODE", "require"),
		},
		Clerk: ClerkConfig{
			SecretKey:     getEnv("CLERK_SECRET_KEY", ""),
			WebhookSecret: getEnv("CLERK_WEBHOOK_SECRET", ""),
		},
		Discord: DiscordConfig{
			BotToken:   getEnv("DISCORD_BOT_TOKEN", ""),
			GuildID:    getEnv("DISCORD_GUILD_ID", ""),
			WebhookURL: getEnv("DISCORD_WEBHOOK_URL", ""),
		},
		JWT: JWTConfig{
			SecretKey:   getEnv("JWT_SECRET", "default-secret-key"),
			ExpiryHours: getEnvInt("JWT_EXPIRY_HOURS", 24),
		},
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}

// Validate は設定をバリデート
func (c *Config) Validate() error {
	// 必須設定のチェック
	if c.Database.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	if c.Clerk.SecretKey == "" {
		return fmt.Errorf("CLERK_SECRET_KEY is required")
	}

	// ポート番号の妥当性チェック
	if c.Database.Port <= 0 || c.Database.Port > 65535 {
		return fmt.Errorf("invalid database port: %d", c.Database.Port)
	}

	return nil
}

// GetDSN はデータベース接続文字列を取得
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.SSLMode)
}

// IsProduction は本番環境かどうかを判定
func (c *Config) IsProduction() bool {
	return strings.ToLower(c.Server.Environment) == "production"
}

// IsDevelopment は開発環境かどうかを判定
func (c *Config) IsDevelopment() bool {
	return strings.ToLower(c.Server.Environment) == "development"
}

// getEnv は環境変数を取得、なければデフォルト値を返す
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvInt は環境変数を整数として取得
func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvBool は環境変数をブール値として取得
func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}