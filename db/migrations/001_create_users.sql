-- 001_create_users.sql
-- ユーザー基本情報テーブル作成
-- 作成日: 2025-09-28
-- BE-03-user: ユーザー管理機能

-- Extension: UUID生成用
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- =============================================================================
-- Users Table（ユーザー基本情報）
-- =============================================================================
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    clerk_id VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100),
    avatar_url TEXT,
    discord_id VARCHAR(255) UNIQUE,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Indexes for users
CREATE INDEX IF NOT EXISTS idx_users_clerk_id ON users(clerk_id);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_discord_id ON users(discord_id) WHERE discord_id IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_users_is_active ON users(is_active);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at) WHERE deleted_at IS NULL;

-- Comments
COMMENT ON TABLE users IS 'ユーザー基本情報テーブル';
COMMENT ON COLUMN users.clerk_id IS 'Clerk認証システムのユーザーID';
COMMENT ON COLUMN users.email IS 'メールアドレス（ログイン用）';
COMMENT ON COLUMN users.username IS '表示名';
COMMENT ON COLUMN users.avatar_url IS 'プロフィール画像URL';
COMMENT ON COLUMN users.discord_id IS 'Discord連携用ID';
COMMENT ON COLUMN users.is_active IS 'アクティブ状態フラグ';
COMMENT ON COLUMN users.deleted_at IS 'ソフトデリート用タイムスタンプ';