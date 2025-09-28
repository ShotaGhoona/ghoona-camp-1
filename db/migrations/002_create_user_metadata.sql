-- 002_create_user_metadata.sql
-- ユーザー詳細情報テーブル作成
-- 作成日: 2025-09-28
-- BE-03-user: ユーザー管理機能

-- =============================================================================
-- User Metadata Table（ユーザー詳細情報）
-- =============================================================================
CREATE TABLE IF NOT EXISTS user_metadata (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    display_name VARCHAR(100),
    profile_image_url TEXT,
    tagline VARCHAR(150),
    bio TEXT,
    vision TEXT,
    vision_public BOOLEAN DEFAULT false,
    timezone VARCHAR(50) DEFAULT 'Asia/Tokyo',
    skills TEXT[],
    interests TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id)
);

-- Indexes for user_metadata
CREATE INDEX IF NOT EXISTS idx_user_metadata_user_id ON user_metadata(user_id);
CREATE INDEX IF NOT EXISTS idx_user_metadata_vision_public ON user_metadata(vision_public) WHERE vision_public = true;

-- Comments
COMMENT ON TABLE user_metadata IS 'ユーザー詳細情報テーブル';
COMMENT ON COLUMN user_metadata.user_id IS 'ユーザーID（外部キー）';
COMMENT ON COLUMN user_metadata.display_name IS '表示名（ユーザー名とは別の表示用名前）';
COMMENT ON COLUMN user_metadata.profile_image_url IS 'プロフィール画像URL';
COMMENT ON COLUMN user_metadata.tagline IS '一言プロフィール';
COMMENT ON COLUMN user_metadata.bio IS '自己紹介文';
COMMENT ON COLUMN user_metadata.vision IS 'ビジョン・目標';
COMMENT ON COLUMN user_metadata.vision_public IS 'ビジョン公開設定';
COMMENT ON COLUMN user_metadata.timezone IS 'タイムゾーン';
COMMENT ON COLUMN user_metadata.skills IS 'スキル一覧';
COMMENT ON COLUMN user_metadata.interests IS '興味・関心一覧';