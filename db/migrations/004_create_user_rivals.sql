-- 004_create_user_rivals.sql
-- ユーザーライバル関係テーブル作成
-- 作成日: 2025-09-28
-- BE-03-user: ユーザー管理機能

-- =============================================================================
-- User Rivals Table（ユーザーライバル関係）
-- =============================================================================
CREATE TABLE IF NOT EXISTS user_rivals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rival_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, rival_user_id),
    CHECK (user_id != rival_user_id)
);

-- Indexes for user_rivals
CREATE INDEX IF NOT EXISTS idx_user_rivals_user_id ON user_rivals(user_id);
CREATE INDEX IF NOT EXISTS idx_user_rivals_rival_user_id ON user_rivals(rival_user_id);

-- Comments
COMMENT ON TABLE user_rivals IS 'ユーザーライバル関係テーブル';
COMMENT ON COLUMN user_rivals.user_id IS 'ユーザーID（外部キー）';
COMMENT ON COLUMN user_rivals.rival_user_id IS 'ライバルユーザーID（外部キー）';
COMMENT ON CONSTRAINT user_rivals_user_id_rival_user_id_key ON user_rivals IS '同一ユーザー間の重複ライバル関係防止';
COMMENT ON CONSTRAINT user_rivals_check ON user_rivals IS '自分自身をライバルに設定することを防止';