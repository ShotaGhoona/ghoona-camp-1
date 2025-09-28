-- 007_create_titles.sql
-- 称号管理関連テーブル作成
-- 作成日: 2025-09-28
-- BE-07-title: 称号管理機能（将来実装予定）

-- =============================================================================
-- Titles Table（称号マスター）
-- =============================================================================
CREATE TABLE IF NOT EXISTS titles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    level INTEGER UNIQUE NOT NULL,
    name_jp VARCHAR(50) NOT NULL,
    name_en VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    required_days INTEGER NOT NULL,
    image_url TEXT,
    color_theme VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for titles
CREATE INDEX IF NOT EXISTS idx_titles_level ON titles(level);
CREATE INDEX IF NOT EXISTS idx_titles_required_days ON titles(required_days);
CREATE INDEX IF NOT EXISTS idx_titles_is_active ON titles(is_active);

-- =============================================================================
-- Title Achievements Table（称号実績）
-- =============================================================================
CREATE TABLE IF NOT EXISTS title_achievements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title_id UUID NOT NULL REFERENCES titles(id) ON DELETE CASCADE,
    achieved_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    is_current BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, title_id)
);

-- Indexes for title_achievements
CREATE INDEX IF NOT EXISTS idx_title_achievements_user_id ON title_achievements(user_id);
CREATE INDEX IF NOT EXISTS idx_title_achievements_title_id ON title_achievements(title_id);
CREATE INDEX IF NOT EXISTS idx_title_achievements_is_current ON title_achievements(is_current) WHERE is_current = true;
CREATE INDEX IF NOT EXISTS idx_title_achievements_achieved_at ON title_achievements(achieved_at);

-- Comments
COMMENT ON TABLE titles IS '称号マスターテーブル';
COMMENT ON COLUMN titles.level IS '称号レベル（1-8）';
COMMENT ON COLUMN titles.name_jp IS '称号名（日本語）';
COMMENT ON COLUMN titles.name_en IS '称号名（英語）';
COMMENT ON COLUMN titles.description IS '称号説明';
COMMENT ON COLUMN titles.required_days IS '獲得に必要な日数';
COMMENT ON COLUMN titles.image_url IS '称号画像URL';
COMMENT ON COLUMN titles.color_theme IS 'テーマカラー';
COMMENT ON COLUMN titles.is_active IS 'アクティブ状態フラグ';

COMMENT ON TABLE title_achievements IS 'ユーザー称号獲得実績テーブル';
COMMENT ON COLUMN title_achievements.user_id IS 'ユーザーID';
COMMENT ON COLUMN title_achievements.title_id IS '称号ID';
COMMENT ON COLUMN title_achievements.achieved_at IS '獲得日時';
COMMENT ON COLUMN title_achievements.is_current IS '現在の称号フラグ';