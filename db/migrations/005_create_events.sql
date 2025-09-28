-- 005_create_events.sql
-- イベント関連テーブル作成
-- 作成日: 2025-09-28
-- BE-05-event: イベント管理機能（将来実装予定）

-- =============================================================================
-- Events Table（朝活イベント）
-- =============================================================================
CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    creator_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
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
CREATE INDEX IF NOT EXISTS idx_events_start_time ON events(start_time);

-- Comments
COMMENT ON TABLE events IS '朝活イベントテーブル';
COMMENT ON COLUMN events.creator_id IS 'イベント作成者ID';
COMMENT ON COLUMN events.title IS 'イベントタイトル';
COMMENT ON COLUMN events.description IS 'イベント説明';
COMMENT ON COLUMN events.event_type IS 'イベント種別（general, development, wellness等）';
COMMENT ON COLUMN events.scheduled_date IS '開催予定日';
COMMENT ON COLUMN events.start_time IS '開始時刻';
COMMENT ON COLUMN events.end_time IS '終了時刻';
COMMENT ON COLUMN events.max_participants IS '最大参加者数';
COMMENT ON COLUMN events.is_recurring IS '定期開催フラグ';
COMMENT ON COLUMN events.recurrence_pattern IS '繰り返しパターン';
COMMENT ON COLUMN events.discord_channel_id IS 'Discord連携チャンネルID';
COMMENT ON COLUMN events.is_active IS 'アクティブ状態フラグ';