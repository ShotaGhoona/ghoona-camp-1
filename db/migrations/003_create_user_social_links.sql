-- 003_create_user_social_links.sql
-- ユーザーソーシャルリンクテーブル作成
-- 作成日: 2025-09-28
-- BE-03-user: ユーザー管理機能

-- =============================================================================
-- User Social Links Table（ユーザーソーシャルリンク）
-- =============================================================================
CREATE TABLE IF NOT EXISTS user_social_links (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    platform VARCHAR(50) NOT NULL,
    url TEXT NOT NULL,
    title VARCHAR(100),
    is_public BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, platform)
);

-- Indexes for user_social_links
CREATE INDEX IF NOT EXISTS idx_user_social_links_user_id ON user_social_links(user_id);
CREATE INDEX IF NOT EXISTS idx_user_social_links_platform ON user_social_links(platform);
CREATE INDEX IF NOT EXISTS idx_user_social_links_is_public ON user_social_links(is_public) WHERE is_public = true;

-- Comments
COMMENT ON TABLE user_social_links IS 'ユーザーソーシャルリンクテーブル';
COMMENT ON COLUMN user_social_links.user_id IS 'ユーザーID（外部キー）';
COMMENT ON COLUMN user_social_links.platform IS 'プラットフォーム（twitter, github, linkedin等）';
COMMENT ON COLUMN user_social_links.url IS 'ソーシャルリンクURL';
COMMENT ON COLUMN user_social_links.title IS 'リンクタイトル（オプション）';
COMMENT ON COLUMN user_social_links.is_public IS '公開設定';