-- 001_initial_tables.sql
-- Ghoona Camp 初期テーブル作成SQL
-- 作成日: 2025-01-28
-- 説明: ユーザー、イベント、出席記録の基本テーブルを作成

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
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for users
CREATE INDEX IF NOT EXISTS idx_users_clerk_id ON users(clerk_id);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_discord_id ON users(discord_id) WHERE discord_id IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_users_is_active ON users(is_active);

-- =============================================================================
-- User Metadata Table（ユーザー詳細情報）
-- =============================================================================
CREATE TABLE IF NOT EXISTS user_metadata (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
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

-- =============================================================================
-- Events Table（朝活イベント）
-- =============================================================================
CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    creator_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    event_type VARCHAR(50) DEFAULT 'general',
    scheduled_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    max_participants INTEGER,
    is_recurring BOOLEAN DEFAULT false,
    recurrence_pattern VARCHAR(50),
    discord_channel_id VARCHAR(255),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for events
CREATE INDEX IF NOT EXISTS idx_events_creator_id ON events(creator_id);
CREATE INDEX IF NOT EXISTS idx_events_scheduled_date ON events(scheduled_date);
CREATE INDEX IF NOT EXISTS idx_events_is_active ON events(is_active);
CREATE INDEX IF NOT EXISTS idx_events_event_type ON events(event_type);

-- =============================================================================
-- Attendance Logs Table（出席ログ）
-- =============================================================================
CREATE TABLE IF NOT EXISTS attendance_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    event_id UUID REFERENCES events(id) ON DELETE SET NULL,
    discord_channel_id VARCHAR(255) NOT NULL,
    joined_at TIMESTAMP WITH TIME ZONE NOT NULL,
    left_at TIMESTAMP WITH TIME ZONE,
    duration_minutes INTEGER,
    is_valid BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for attendance_logs
CREATE INDEX IF NOT EXISTS idx_attendance_logs_user_id ON attendance_logs(user_id);
CREATE INDEX IF NOT EXISTS idx_attendance_logs_event_id ON attendance_logs(event_id);
CREATE INDEX IF NOT EXISTS idx_attendance_logs_joined_at ON attendance_logs(joined_at);
CREATE INDEX IF NOT EXISTS idx_attendance_logs_is_valid ON attendance_logs(is_valid);

-- =============================================================================
-- Attendance Statistics Table（出席統計）
-- =============================================================================
CREATE TABLE IF NOT EXISTS attendance_statistics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    total_attendance_days INTEGER DEFAULT 0,
    current_streak_days INTEGER DEFAULT 0,
    max_streak_days INTEGER DEFAULT 0,
    last_attendance_date DATE,
    first_attendance_date DATE,
    total_duration_minutes INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id)
);

-- Indexes for attendance_statistics
CREATE INDEX IF NOT EXISTS idx_attendance_statistics_user_id ON attendance_statistics(user_id);
CREATE INDEX IF NOT EXISTS idx_attendance_statistics_total_days ON attendance_statistics(total_attendance_days);
CREATE INDEX IF NOT EXISTS idx_attendance_statistics_current_streak ON attendance_statistics(current_streak_days);

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
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title_id UUID REFERENCES titles(id) ON DELETE CASCADE,
    achieved_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    is_current BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, title_id)
);

-- Indexes for title_achievements
CREATE INDEX IF NOT EXISTS idx_title_achievements_user_id ON title_achievements(user_id);
CREATE INDEX IF NOT EXISTS idx_title_achievements_title_id ON title_achievements(title_id);
CREATE INDEX IF NOT EXISTS idx_title_achievements_is_current ON title_achievements(is_current);

-- =============================================================================
-- Comments
-- =============================================================================

COMMENT ON TABLE users IS 'ユーザー基本情報テーブル';
COMMENT ON TABLE user_metadata IS 'ユーザー詳細情報テーブル';
COMMENT ON TABLE events IS '朝活イベントテーブル';
COMMENT ON TABLE attendance_logs IS 'Discord参加ログテーブル';
COMMENT ON TABLE attendance_statistics IS 'ユーザー参加統計テーブル';
COMMENT ON TABLE titles IS '称号マスターテーブル';
COMMENT ON TABLE title_achievements IS 'ユーザー称号獲得実績テーブル';