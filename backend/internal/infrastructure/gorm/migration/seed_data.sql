-- seed_data.sql
-- Ghoona Camp 開発用初期データ
-- 作成日: 2025-01-28
-- 説明: 開発・テスト用の初期データを投入

-- =============================================================================
-- 称号マスターデータ（8段階）
-- =============================================================================
INSERT INTO titles (level, name_jp, name_en, description, required_days, color_theme, is_active) VALUES
(1, 'まどろみ見習い', 'Sleeper', '朝活の世界への第一歩。まだ眠りの世界から抜け出せないあなたに贈る称号です。', 1, '#8B5A3C', true),
(2, '夜明け前の戦士', 'Dawn Warrior', '暗闇の中でも立ち上がる勇気を見せたあなた。夜明け前の静寂を味方にする戦士です。', 7, '#4A5568', true),
(3, '朝日の使者', 'Sunrise Messenger', '朝日と共に動き出すあなたは、新しい一日の使者として認められました。', 14, '#F6AD55', true),
(4, '早起きの達人', 'Early Bird Master', '早起きのコツを掴み、朝の時間を有効活用できるようになった達人です。', 30, '#48BB78', true),
(5, '朝活の戦士', 'Morning Warrior', '継続的な朝活により、真の朝活戦士として認定されました。', 60, '#4299E1', true),
(6, '黎明の守護者', 'Dawn Guardian', '朝活コミュニティを支える存在として、黎明を守護する者です。', 120, '#9F7AEA', true),
(7, '朝の賢者', 'Morning Sage', '長年の朝活経験により培われた智慧を持つ、朝の賢者です。', 200, '#F56565', true),
(8, '朝活の伝説', 'Legend of Morning', '朝活界の伝説的存在。あなたの継続力は多くの人に勇気を与えています。', 365, '#FFD700', true);

-- =============================================================================
-- 開発用テストユーザー
-- =============================================================================
-- 注意: 実際のClerk IDは実際の値に置き換えてください
INSERT INTO users (clerk_id, email, username, avatar_url, is_active) VALUES
('test_clerk_user_001', 'developer@ghoona-camp.com', 'developer', 'https://via.placeholder.com/150', true),
('test_clerk_user_002', 'tester@ghoona-camp.com', 'tester', 'https://via.placeholder.com/150', true);

-- =============================================================================
-- テストユーザーの詳細情報
-- =============================================================================
INSERT INTO user_metadata (user_id, display_name, tagline, bio, vision, vision_public, skills, interests) 
SELECT 
    u.id,
    'Ghoona Developer',
    '朝活で人生を変える開発者',
    'Ghoona Campの開発を通じて、朝活コミュニティの価値を最大化することを目指しています。',
    '朝活を通じて、多くの人が充実した人生を送れる世界を作りたい。',
    true,
    ARRAY['Go', 'React', 'TypeScript', 'Docker'],
    ARRAY['朝活', 'プログラミング', 'コミュニティ', '健康']
FROM users u WHERE u.email = 'developer@ghoona-camp.com';

INSERT INTO user_metadata (user_id, display_name, tagline, bio, vision, vision_public, skills, interests) 
SELECT 
    u.id,
    'Test User',
    'テスト専門家',
    'アプリケーションのテストと品質保証を担当しています。',
    'バグのない完璧なアプリケーションを提供したい。',
    false,
    ARRAY['Testing', 'QA', 'Bug Detection'],
    ARRAY['品質管理', 'テスト', 'デバッグ']
FROM users u WHERE u.email = 'tester@ghoona-camp.com';

-- =============================================================================
-- 開発用サンプルイベント
-- =============================================================================
INSERT INTO events (creator_id, title, description, event_type, scheduled_date, start_time, end_time, max_participants, is_active)
SELECT 
    u.id,
    '朝活開発セッション',
    'Ghoona Campの開発を進める朝活セッション。一緒にコードを書きましょう！',
    'development',
    CURRENT_DATE + INTERVAL '1 day',
    '06:00:00',
    '07:00:00',
    10,
    true
FROM users u WHERE u.email = 'developer@ghoona-camp.com';

INSERT INTO events (creator_id, title, description, event_type, scheduled_date, start_time, end_time, max_participants, is_active)
SELECT 
    u.id,
    '朝のヨガ＆瞑想',
    '心と体を整える朝のヨガと瞑想の時間です。',
    'wellness',
    CURRENT_DATE + INTERVAL '2 days',
    '06:30:00',
    '07:30:00',
    15,
    true
FROM users u WHERE u.email = 'tester@ghoona-camp.com';

-- =============================================================================
-- テスト用出席統計初期化
-- =============================================================================
INSERT INTO attendance_statistics (user_id, total_attendance_days, current_streak_days, max_streak_days, first_attendance_date, last_attendance_date, total_duration_minutes)
SELECT 
    u.id,
    5,
    3,
    5,
    CURRENT_DATE - INTERVAL '7 days',
    CURRENT_DATE - INTERVAL '1 day',
    300
FROM users u WHERE u.email = 'developer@ghoona-camp.com';

INSERT INTO attendance_statistics (user_id, total_attendance_days, current_streak_days, max_streak_days, first_attendance_date, last_attendance_date, total_duration_minutes)
SELECT 
    u.id,
    2,
    2,
    2,
    CURRENT_DATE - INTERVAL '3 days',
    CURRENT_DATE - INTERVAL '1 day',
    120
FROM users u WHERE u.email = 'tester@ghoona-camp.com';

-- =============================================================================
-- テスト用称号付与
-- =============================================================================
-- 開発者に「夜明け前の戦士」称号を付与
INSERT INTO title_achievements (user_id, title_id, is_current)
SELECT 
    u.id,
    t.id,
    true
FROM users u, titles t 
WHERE u.email = 'developer@ghoona-camp.com' 
AND t.level = 2;

-- テスターに「まどろみ見習い」称号を付与
INSERT INTO title_achievements (user_id, title_id, is_current)
SELECT 
    u.id,
    t.id,
    true
FROM users u, titles t 
WHERE u.email = 'tester@ghoona-camp.com' 
AND t.level = 1;

-- =============================================================================
-- Comments
-- =============================================================================

-- データ投入完了メッセージをログに出力
DO $$
BEGIN
    RAISE NOTICE '✅ Ghoona Camp シードデータの投入が完了しました';
    RAISE NOTICE '📊 投入データ: 称号8件、テストユーザー2件、イベント2件、統計2件、称号実績2件';
    RAISE NOTICE '🎯 開発環境での動作確認が可能になりました';
END $$;