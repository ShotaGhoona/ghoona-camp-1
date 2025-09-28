-- 006_create_attendance.sql
-- 出席管理関連テーブル作成
-- 作成日: 2025-09-28
-- BE-04-attend: 出席管理機能（将来実装予定）

-- =============================================================================
-- Attendance Logs Table（出席ログ）
-- =============================================================================
CREATE TABLE IF NOT EXISTS attendance_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
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
CREATE INDEX IF NOT EXISTS idx_attendance_logs_discord_channel ON attendance_logs(discord_channel_id);

-- =============================================================================
-- Attendance Statistics Table（出席統計）
-- =============================================================================
CREATE TABLE IF NOT EXISTS attendance_statistics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
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
CREATE INDEX IF NOT EXISTS idx_attendance_statistics_max_streak ON attendance_statistics(max_streak_days);

-- Comments
COMMENT ON TABLE attendance_logs IS 'Discord参加ログテーブル';
COMMENT ON COLUMN attendance_logs.user_id IS 'ユーザーID';
COMMENT ON COLUMN attendance_logs.event_id IS 'イベントID（オプション）';
COMMENT ON COLUMN attendance_logs.discord_channel_id IS 'Discord参加チャンネルID';
COMMENT ON COLUMN attendance_logs.joined_at IS '参加開始時刻';
COMMENT ON COLUMN attendance_logs.left_at IS '退出時刻';
COMMENT ON COLUMN attendance_logs.duration_minutes IS '参加時間（分）';
COMMENT ON COLUMN attendance_logs.is_valid IS '有効な参加記録かのフラグ';

COMMENT ON TABLE attendance_statistics IS 'ユーザー参加統計テーブル';
COMMENT ON COLUMN attendance_statistics.user_id IS 'ユーザーID';
COMMENT ON COLUMN attendance_statistics.total_attendance_days IS '総参加日数';
COMMENT ON COLUMN attendance_statistics.current_streak_days IS '現在の連続参加日数';
COMMENT ON COLUMN attendance_statistics.max_streak_days IS '最大連続参加日数';
COMMENT ON COLUMN attendance_statistics.last_attendance_date IS '最後の参加日';
COMMENT ON COLUMN attendance_statistics.first_attendance_date IS '初回参加日';
COMMENT ON COLUMN attendance_statistics.total_duration_minutes IS '総参加時間（分）';