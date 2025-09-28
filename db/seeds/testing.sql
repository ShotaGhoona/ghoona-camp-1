-- testing.sql
-- テスト環境用シードデータ
-- 作成日: 2025-09-28
-- 用途: 自動テスト・CI/CD用の最小限のデータセット

-- =============================================================================
-- 称号マスターデータ（最小限）
-- =============================================================================
INSERT INTO titles (level, name_jp, name_en, description, required_days, color_theme, is_active) VALUES
(1, 'まどろみ見習い', 'Sleeper', 'テスト用称号レベル1', 1, '#8B5A3C', true),
(2, '夜明け前の戦士', 'Dawn Warrior', 'テスト用称号レベル2', 7, '#4A5568', true),
(3, '朝日の使者', 'Sunrise Messenger', 'テスト用称号レベル3', 14, '#F6AD55', true)
ON CONFLICT (level) DO NOTHING;

-- =============================================================================
-- テスト用ユーザー（最小限）
-- =============================================================================
INSERT INTO users (clerk_id, email, username, avatar_url, is_active) VALUES
('test_clerk_001', 'test1@example.com', 'Test User 1', 'https://example.com/avatar1.jpg', true),
('test_clerk_002', 'test2@example.com', 'Test User 2', 'https://example.com/avatar2.jpg', true),
('test_clerk_003', 'test3@example.com', 'Test User 3', 'https://example.com/avatar3.jpg', false)
ON CONFLICT (clerk_id) DO NOTHING;

-- =============================================================================
-- テスト用メタデータ（最小限）
-- =============================================================================
INSERT INTO user_metadata (user_id, display_name, tagline, bio, vision, vision_public, timezone, skills, interests) 
SELECT 
    u.id,
    'Test User 1',
    'テスト用ユーザー',
    'API テスト用のユーザーアカウントです。',
    'テスト環境での動作確認',
    true,
    'Asia/Tokyo',
    ARRAY['Testing', 'API'],
    ARRAY['テスト', 'API確認']
FROM users u WHERE u.email = 'test1@example.com'
ON CONFLICT (user_id) DO NOTHING;

-- =============================================================================
-- テスト用ソーシャルリンク（最小限）
-- =============================================================================
INSERT INTO user_social_links (user_id, platform, url, title, is_public)
SELECT 
    u.id,
    'github',
    'https://github.com/test-user-1',
    'Test Repository',
    true
FROM users u WHERE u.email = 'test1@example.com'
ON CONFLICT (user_id, platform) DO NOTHING;

-- =============================================================================
-- テスト用ライバル関係
-- =============================================================================
INSERT INTO user_rivals (user_id, rival_user_id)
SELECT 
    u1.id,
    u2.id
FROM users u1, users u2 
WHERE u1.email = 'test1@example.com' 
AND u2.email = 'test2@example.com'
ON CONFLICT (user_id, rival_user_id) DO NOTHING;

-- =============================================================================
-- テスト用イベント（最小限）
-- =============================================================================
INSERT INTO events (creator_id, title, description, event_type, scheduled_date, start_time, end_time, max_participants, is_active)
SELECT 
    u.id,
    'Test Morning Event',
    'Test event for API verification',
    'test',
    CURRENT_DATE + INTERVAL '1 day',
    '06:00:00',
    '07:00:00',
    5,
    true
FROM users u WHERE u.email = 'test1@example.com'
ON CONFLICT DO NOTHING;

-- =============================================================================
-- テスト用出席統計
-- =============================================================================
INSERT INTO attendance_statistics (user_id, total_attendance_days, current_streak_days, max_streak_days, first_attendance_date, last_attendance_date, total_duration_minutes)
SELECT 
    u.id,
    3,
    2,
    3,
    CURRENT_DATE - INTERVAL '5 days',
    CURRENT_DATE - INTERVAL '1 day',
    180
FROM users u WHERE u.email = 'test1@example.com'
ON CONFLICT (user_id) DO NOTHING;

-- =============================================================================
-- テスト用称号付与
-- =============================================================================
INSERT INTO title_achievements (user_id, title_id, is_current)
SELECT 
    u.id,
    t.id,
    true
FROM users u, titles t 
WHERE u.email = 'test1@example.com' 
AND t.level = 1
ON CONFLICT (user_id, title_id) DO NOTHING;

-- =============================================================================
-- メッセージ
-- =============================================================================
DO $$
BEGIN
    RAISE NOTICE '✅ テスト環境シードデータの投入が完了しました';
    RAISE NOTICE '📊 投入データ:';
    RAISE NOTICE '  - 称号: 3件（最小限）';
    RAISE NOTICE '  - ユーザー: 3件（テスト用）';
    RAISE NOTICE '  - メタデータ: 1件';
    RAISE NOTICE '  - ソーシャルリンク: 1件';
    RAISE NOTICE '  - ライバル関係: 1件';
    RAISE NOTICE '  - イベント: 1件';
    RAISE NOTICE '  - 出席統計: 1件';
    RAISE NOTICE '  - 称号実績: 1件';
    RAISE NOTICE '🧪 自動テスト・CI/CD環境での検証が可能になりました';
END $$;