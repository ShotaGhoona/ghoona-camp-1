-- development.sql
-- 開発環境用シードデータ
-- 作成日: 2025-09-28
-- 用途: 開発・テスト・デモ用の充実したデータセット

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
(8, '朝活の伝説', 'Legend of Morning', '朝活界の伝説的存在。あなたの継続力は多くの人に勇気を与えています。', 365, '#FFD700', true)
ON CONFLICT (level) DO NOTHING;

-- =============================================================================
-- 開発用テストユーザー
-- =============================================================================
INSERT INTO users (clerk_id, email, username, avatar_url, is_active) VALUES
('clerk_dev_user_001', 'developer@ghoona.camp', 'Ghoona Developer', 'https://avatars.githubusercontent.com/u/1?v=4', true),
('clerk_dev_user_002', 'tester@ghoona.camp', 'Test Expert', 'https://avatars.githubusercontent.com/u/2?v=4', true),
('clerk_dev_user_003', 'designer@ghoona.camp', 'UI Designer', 'https://avatars.githubusercontent.com/u/3?v=4', true),
('clerk_dev_user_004', 'pm@ghoona.camp', 'Product Manager', 'https://avatars.githubusercontent.com/u/4?v=4', true),
('clerk_test_user_1', 'test1@ghoona.camp', 'テストユーザー1', 'https://example.com/avatar1.jpg', true),
('clerk_test_user_2', 'test2@ghoona.camp', 'テストユーザー2', 'https://example.com/avatar2.jpg', true)
ON CONFLICT (clerk_id) DO NOTHING;

-- =============================================================================
-- ユーザーメタデータ
-- =============================================================================
INSERT INTO user_metadata (user_id, display_name, tagline, bio, vision, vision_public, timezone, skills, interests) 
SELECT 
    u.id,
    'Ghoona Developer',
    '朝活で人生を変える開発者',
    'Ghoona Campの開発を通じて、朝活コミュニティの価値を最大化することを目指しています。毎朝6時から開発作業を行い、生産性の高い一日をスタートしています。',
    '朝活を通じて、多くの人が充実した人生を送れる世界を作りたい。テクノロジーの力で朝活習慣を支援し、継続できる仕組みを構築する。',
    true,
    'Asia/Tokyo',
    ARRAY['Go', 'React', 'TypeScript', 'Docker', 'AWS', 'Clean Architecture'],
    ARRAY['朝活', 'プログラミング', 'コミュニティ', '健康', 'ライフハック']
FROM users u WHERE u.email = 'developer@ghoona.camp'
ON CONFLICT (user_id) DO NOTHING;

INSERT INTO user_metadata (user_id, display_name, tagline, bio, vision, vision_public, timezone, skills, interests) 
SELECT 
    u.id,
    'Quality Guardian',
    'テスト専門家・品質の守護者',
    'アプリケーションのテストと品質保証を担当。朝の静寂な時間を活用してテストケースの設計と自動化スクリプトの作成を行っています。',
    'バグのない完璧なアプリケーションを提供し、ユーザーに最高の体験を届けたい。',
    true,
    'Asia/Tokyo',
    ARRAY['Testing', 'QA', 'Automation', 'Selenium', 'Jest', 'Postman'],
    ARRAY['品質管理', 'テスト', 'デバッグ', 'ユーザー体験', '自動化']
FROM users u WHERE u.email = 'tester@ghoona.camp'
ON CONFLICT (user_id) DO NOTHING;

INSERT INTO user_metadata (user_id, display_name, tagline, bio, vision, vision_public, timezone, skills, interests) 
SELECT 
    u.id,
    '朝活マスター',
    '早起きは三文の得！',
    '毎朝5時起きで頑張っています。朝活を通じて自分を成長させることで、充実した一日を過ごしています。朝の静寂な時間が一番集中できます。',
    '朝活を通じて規則正しい生活を送り、目標達成に向けて努力する。健康的なライフスタイルを維持したい。',
    true,
    'Asia/Tokyo',
    ARRAY['早起き', '継続力', '目標設定', '時間管理'],
    ARRAY['朝活', '読書', '運動', '瞑想', 'ヨガ']
FROM users u WHERE u.email = 'test1@ghoona.camp'
ON CONFLICT (user_id) DO NOTHING;

-- =============================================================================
-- ソーシャルリンク
-- =============================================================================
INSERT INTO user_social_links (user_id, platform, url, title, is_public)
SELECT 
    u.id,
    'github',
    'https://github.com/ghoona-developer',
    'Ghoona Camp Repository',
    true
FROM users u WHERE u.email = 'developer@ghoona.camp'
ON CONFLICT (user_id, platform) DO NOTHING;

INSERT INTO user_social_links (user_id, platform, url, title, is_public)
SELECT 
    u.id,
    'twitter',
    'https://twitter.com/test_user_1',
    '朝活アカウント',
    true
FROM users u WHERE u.email = 'test1@ghoona.camp'
ON CONFLICT (user_id, platform) DO NOTHING;

-- =============================================================================
-- ライバル関係
-- =============================================================================
INSERT INTO user_rivals (user_id, rival_user_id)
SELECT 
    u1.id,
    u2.id
FROM users u1, users u2 
WHERE u1.email = 'test1@ghoona.camp' 
AND u2.email = 'test2@ghoona.camp'
ON CONFLICT (user_id, rival_user_id) DO NOTHING;

-- =============================================================================
-- 開発用サンプルイベント
-- =============================================================================
INSERT INTO events (creator_id, title, description, event_type, scheduled_date, start_time, end_time, max_participants, is_active)
SELECT 
    u.id,
    '朝活開発セッション',
    'Ghoona Campの開発を進める朝活セッション。一緒にコードを書きましょう！新機能の実装やバグ修正、アーキテクチャの改善について議論します。',
    'development',
    CURRENT_DATE + INTERVAL '1 day',
    '06:00:00',
    '07:30:00',
    10,
    true
FROM users u WHERE u.email = 'developer@ghoona.camp'
ON CONFLICT DO NOTHING;

INSERT INTO events (creator_id, title, description, event_type, scheduled_date, start_time, end_time, max_participants, is_active)
SELECT 
    u.id,
    '朝のヨガ＆瞑想',
    '心と体を整える朝のヨガと瞑想の時間です。一日を穏やかにスタートしましょう。',
    'wellness',
    CURRENT_DATE + INTERVAL '2 days',
    '06:30:00',
    '07:30:00',
    15,
    true
FROM users u WHERE u.email = 'test1@ghoona.camp'
ON CONFLICT DO NOTHING;

-- =============================================================================
-- テスト用出席統計
-- =============================================================================
INSERT INTO attendance_statistics (user_id, total_attendance_days, current_streak_days, max_streak_days, first_attendance_date, last_attendance_date, total_duration_minutes)
SELECT 
    u.id,
    15,
    7,
    10,
    CURRENT_DATE - INTERVAL '20 days',
    CURRENT_DATE - INTERVAL '1 day',
    900
FROM users u WHERE u.email = 'developer@ghoona.camp'
ON CONFLICT (user_id) DO NOTHING;

INSERT INTO attendance_statistics (user_id, total_attendance_days, current_streak_days, max_streak_days, first_attendance_date, last_attendance_date, total_duration_minutes)
SELECT 
    u.id,
    5,
    3,
    5,
    CURRENT_DATE - INTERVAL '7 days',
    CURRENT_DATE - INTERVAL '1 day',
    300
FROM users u WHERE u.email = 'test1@ghoona.camp'
ON CONFLICT (user_id) DO NOTHING;

-- =============================================================================
-- テスト用称号付与
-- =============================================================================
-- 開発者に「朝日の使者」称号を付与
INSERT INTO title_achievements (user_id, title_id, is_current)
SELECT 
    u.id,
    t.id,
    true
FROM users u, titles t 
WHERE u.email = 'developer@ghoona.camp' 
AND t.level = 3
ON CONFLICT (user_id, title_id) DO NOTHING;

-- テスト1ユーザーに「夜明け前の戦士」称号を付与
INSERT INTO title_achievements (user_id, title_id, is_current)
SELECT 
    u.id,
    t.id,
    true
FROM users u, titles t 
WHERE u.email = 'test1@ghoona.camp' 
AND t.level = 2
ON CONFLICT (user_id, title_id) DO NOTHING;

-- =============================================================================
-- メッセージ
-- =============================================================================
DO $$
BEGIN
    RAISE NOTICE '✅ 開発環境シードデータの投入が完了しました';
    RAISE NOTICE '📊 投入データ:';
    RAISE NOTICE '  - 称号: 8件';
    RAISE NOTICE '  - ユーザー: 6件（開発・テスト用）';
    RAISE NOTICE '  - メタデータ: 3件';
    RAISE NOTICE '  - ソーシャルリンク: 2件';
    RAISE NOTICE '  - ライバル関係: 1件';
    RAISE NOTICE '  - イベント: 2件';
    RAISE NOTICE '  - 出席統計: 2件';
    RAISE NOTICE '  - 称号実績: 2件';
    RAISE NOTICE '🎯 開発環境での動作確認・デモが可能になりました';
END $$;