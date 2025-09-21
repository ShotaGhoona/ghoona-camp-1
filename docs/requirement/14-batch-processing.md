# Ghoona Camp - バッチ処理仕様書

## Overview
Ghoona Campアプリケーションで実行する定期バッチ処理の仕様書です。
Discord参加ログの集計、ランキング更新、称号付与等を自動化します。

## 日次バッチ処理

### 実行スケジュール
- **実行時刻**: 毎朝9:00 (JST)
- **処理対象**: その日分のデータを集計
- **失敗時の対応**: 自動リトライ (最大3回)

### 処理フロー

#### 1. Discord参加ログ集計 (9:00-9:10)
```sql
-- attendance_logsからattendance_summariesへの集計
INSERT INTO attendance_summaries (
  user_id, date, total_duration_minutes, session_count,
  first_join_time, last_leave_time, is_morning_active
)
SELECT 
  user_id,
  DATE(joined_at) as date,
  SUM(duration_minutes) as total_duration_minutes,
  COUNT(*) as session_count,
  MIN(TIME(joined_at)) as first_join_time,
  MAX(TIME(left_at)) as last_leave_time,
  BOOL_OR(TIME(joined_at) BETWEEN '06:00:00' AND '06:30:00') as is_morning_active
FROM attendance_logs
WHERE DATE(joined_at) = CURRENT_DATE
  AND is_valid = true
GROUP BY user_id, DATE(joined_at);
```

#### 2. 参加統計更新 (9:10-9:20)
```sql
-- attendance_statisticsの更新
UPDATE attendance_statistics 
SET 
  total_attendance_days = (
    SELECT COUNT(DISTINCT date) 
    FROM attendance_summaries 
    WHERE user_id = attendance_statistics.user_id
      AND is_morning_active = true
  ),
  last_attendance_date = (
    SELECT MAX(date)
    FROM attendance_summaries
    WHERE user_id = attendance_statistics.user_id
      AND is_morning_active = true
  ),
  total_duration_minutes = (
    SELECT SUM(total_duration_minutes)
    FROM attendance_summaries
    WHERE user_id = attendance_statistics.user_id
  ),
  current_streak_days = calculate_current_streak(user_id),
  max_streak_days = calculate_max_streak(user_id)
WHERE user_id IN (
  SELECT DISTINCT user_id 
  FROM attendance_summaries 
  WHERE date = CURRENT_DATE
);
```

#### 3. 称号自動付与チェック (9:20-9:25)
```sql
-- 新規称号獲得者の自動付与
INSERT INTO title_achievements (user_id, title_id, achieved_at)
SELECT 
  stats.user_id,
  titles.id as title_id,
  NOW() as achieved_at
FROM attendance_statistics stats
CROSS JOIN titles
WHERE stats.total_attendance_days >= titles.required_days
  AND NOT EXISTS (
    SELECT 1 FROM title_achievements ta
    WHERE ta.user_id = stats.user_id 
      AND ta.title_id = titles.id
  );
```

#### 4. 称号獲得通知送信 (9:25-9:30)
```sql
-- 称号獲得通知の作成
INSERT INTO notifications (user_id, type, title, message, data)
SELECT 
  ta.user_id,
  'achievement' as type,
  CONCAT('称号「', t.name_jp, '」を獲得しました！') as title,
  CONCAT('おめでとうございます！', t.required_days, '日間の朝活達成で「', t.name_jp, '」の称号を獲得しました。') as message,
  JSON_BUILD_OBJECT('title_id', t.id, 'title_name', t.name_jp, 'level', t.level) as data
FROM title_achievements ta
JOIN titles t ON ta.title_id = t.id
WHERE DATE(ta.achieved_at) = CURRENT_DATE;
```

## 月次バッチ処理

### 実行スケジュール
- **実行時刻**: 毎月1日 2:00 (JST)
- **処理対象**: 前月分のデータ
- **保存期間**: 過去24ヶ月分保持

### 処理内容

#### 1. 月次ランキング確定
```sql
-- 月次ランキングのスナップショット保存
CREATE TABLE IF NOT EXISTS monthly_rankings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  year_month VARCHAR(7) NOT NULL, -- 'YYYY-MM'
  user_id UUID REFERENCES users(id),
  rank INTEGER NOT NULL,
  attendance_days INTEGER NOT NULL,
  total_duration_minutes INTEGER NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

INSERT INTO monthly_rankings (year_month, user_id, rank, attendance_days, total_duration_minutes)
SELECT 
  TO_CHAR(DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month'), 'YYYY-MM') as year_month,
  user_id,
  ROW_NUMBER() OVER (ORDER BY COUNT(*) DESC, SUM(total_duration_minutes) DESC) as rank,
  COUNT(*) as attendance_days,
  SUM(total_duration_minutes) as total_duration_minutes
FROM attendance_summaries
WHERE date >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
  AND date < DATE_TRUNC('month', CURRENT_DATE)
  AND is_morning_active = true
GROUP BY user_id
ORDER BY attendance_days DESC, total_duration_minutes DESC;
```

#### 2. 古いデータのクリーンアップ
```sql
-- 2年以上前の詳細ログを削除
DELETE FROM attendance_logs 
WHERE created_at < CURRENT_DATE - INTERVAL '2 years';

-- 古い通知を削除（既読・3ヶ月以上前）
DELETE FROM notifications 
WHERE is_read = true 
  AND created_at < CURRENT_DATE - INTERVAL '3 months';
```

## 週次バッチ処理

### 実行スケジュール
- **実行時刻**: 毎週月曜日 1:00 (JST)
- **処理対象**: 前週分のデータ

### 処理内容

#### 1. 週次リマインダー送信
```sql
-- 先週参加がなかったユーザーへのリマインダー
INSERT INTO notifications (user_id, type, title, message, scheduled_at)
SELECT 
  u.id as user_id,
  'reminder' as type,
  '朝活参加のリマインダー' as title,
  '先週は朝活への参加がありませんでした。今週も一緒に頑張りましょう！' as message,
  NOW() + INTERVAL '8 hours' as scheduled_at -- 朝9時配信
FROM users u
LEFT JOIN attendance_summaries as_week ON u.id = as_week.user_id
  AND as_week.date >= DATE_TRUNC('week', CURRENT_DATE - INTERVAL '1 week')
  AND as_week.date < DATE_TRUNC('week', CURRENT_DATE)
  AND as_week.is_morning_active = true
JOIN notification_settings ns ON u.id = ns.user_id
WHERE as_week.user_id IS NULL
  AND u.is_active = true
  AND ns.reminder_enabled = true;
```

#### 2. ライバル更新通知
```sql
-- ライバルの進捗更新通知
INSERT INTO notifications (user_id, type, title, message, data)
SELECT 
  ur.user_id,
  'rival_update' as type,
  CONCAT(rival_meta.display_name, 'さんの進捗更新') as title,
  CONCAT(rival_meta.display_name, 'さんが先週', weekly_count.attendance_days, '日参加しました！') as message,
  JSON_BUILD_OBJECT(
    'rival_user_id', ur.rival_user_id,
    'rival_name', rival_meta.display_name,
    'weekly_attendance', weekly_count.attendance_days
  ) as data
FROM user_rivals ur
JOIN user_metadata rival_meta ON ur.rival_user_id = rival_meta.user_id
JOIN notification_settings ns ON ur.user_id = ns.user_id
JOIN (
  SELECT 
    user_id,
    COUNT(*) as attendance_days
  FROM attendance_summaries
  WHERE date >= DATE_TRUNC('week', CURRENT_DATE - INTERVAL '1 week')
    AND date < DATE_TRUNC('week', CURRENT_DATE)
    AND is_morning_active = true
  GROUP BY user_id
) weekly_count ON ur.rival_user_id = weekly_count.user_id
WHERE ns.rival_update_enabled = true
  AND weekly_count.attendance_days > 0;
```

## エラーハンドリング

### バッチ処理失敗時の対応
1. **自動リトライ**: 各処理で最大3回自動リトライ
2. **エラーログ記録**: 詳細なエラー情報をログに記録
3. **アラート送信**: 管理者へのSlack/メール通知
4. **手動実行**: 管理画面からの手動バッチ実行機能

### データ整合性チェック
```sql
-- 日次チェック：attendance_summariesとlogsの整合性
SELECT 
  'attendance_mismatch' as check_type,
  COUNT(*) as error_count
FROM (
  SELECT user_id, date
  FROM attendance_summaries
  WHERE date = CURRENT_DATE
  EXCEPT
  SELECT user_id, DATE(joined_at) as date
  FROM attendance_logs
  WHERE DATE(joined_at) = CURRENT_DATE
    AND TIME(joined_at) BETWEEN '06:00:00' AND '06:30:00'
    AND is_valid = true
) mismatches;
```

## 監視・アラート

### 処理時間監視
- **正常処理時間**: 各バッチ10分以内
- **アラート条件**: 15分超過または失敗
- **通知先**: 開発者Slackチャンネル

### データ量監視
- **日次処理件数**: 想定ユーザー数の80%以上
- **異常検知**: 前日比50%以下の処理件数
- **対応**: 自動チェック＋管理者通知

## 運用考慮事項

### メンテナンス時の対応
- **定期メンテナンス**: バッチ処理時間を避けた深夜実行
- **緊急メンテナンス**: バッチ停止＋復旧後の差分処理

### スケーリング対応
- **データ量増加**: パーティショニング（月次テーブル分割）
- **処理時間増加**: 並列処理化、インデックス最適化
- **高可用性**: 複数サーバーでの冗長実行