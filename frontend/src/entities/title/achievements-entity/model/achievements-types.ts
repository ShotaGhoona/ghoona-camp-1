/**
 * Achievements Types
 * 
 * ユーザー称号実績関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * achievements-entity/index.ts からエクスポートされた型を使用してください。
 */

// === API Response Types（API仕様準拠） ===

// 称号実績（achievements-entity用）- API仕様準拠
interface AchievementResponse {
  id: string;
  title: {
    id: string;
    level: number;
    name_jp: string;
    name_en: string;
    color_theme: string;
    image_url: string;
  };
  achieved_at: string;
  is_current: boolean;
  achievement_rank: number;
  total_achievers_at_time: number;
}

interface CurrentTitleResponse {
  id: string;
  level: number;
  name_jp: string;
  name_en: string;
  color_theme: string;
  image_url: string;
  achieved_at: string;
  days_held: number;
}

// === GET /users/{userId}/achievements レスポンス（API仕様準拠） ===
export interface UserAchievementsResponse {
  data: {
    user: {
      id: string;
      display_name: string;
      username: string;
      avatar_url: string;
    };
    current_title: CurrentTitleResponse;
    achievements: AchievementResponse[];
    next_title_progress: {
      title: {
        id: string;
        level: number;
        name_jp: string;
        required_days: number;
      };
      progress: {
        current_days: number;
        remaining_days: number;
        progress_rate: number;
        estimated_achievement_date: string;
      };
    };
    statistics: {
      total_achieved: number;
      max_level_achieved: number;
      consistency_score: number;
      total_attendance_days: number;
    };
  };
  message: string;
  timestamp: string;
}

// === PUT /users/{userId}/achievements/{titleId} レスポンス（API仕様準拠） ===
export interface SetCurrentTitleResponse {
  data: {
    id: string;
    title: {
      id: string;
      level: number;
      name_jp: string;
      name_en: string;
      color_theme: string;
      image_url: string;
    };
    is_current: boolean;
    updated_at: string;
    previous_title: {
      id: string;
      level: number;
      name_jp: string;
    };
  };
  message: string;
  timestamp: string;
}

// === Client Types（camelCase - フロントエンド用） ===
export interface Achievement {
  id: string;
  title: {
    id: string;
    level: number;
    nameJp: string;
    nameEn: string;
    colorTheme: string;
    imageUrl: string;
  };
  achievedAt: Date;
  isCurrent: boolean;
  achievementRank: number;
  totalAchieversAtTime: number;
}

export interface CurrentTitle {
  id: string;
  level: number;
  nameJp: string;
  nameEn: string;
  colorTheme: string;
  imageUrl: string;
  achievedAt: Date;
  daysHeld: number;
}

export interface NextTitleProgress {
  title: {
    id: string;
    level: number;
    nameJp: string;
    requiredDays: number;
  };
  progress: {
    currentDays: number;
    remainingDays: number;
    progressRate: number;
    estimatedAchievementDate: Date;
  };
}

export interface UserInfo {
  id: string;
  displayName: string;
  username: string;
  avatarUrl: string;
}

export interface AchievementStatistics {
  totalAchieved: number;
  maxLevelAchieved: number;
  consistencyScore: number;
  totalAttendanceDays: number;
}

export interface UserAchievements {
  user: UserInfo;
  currentTitle: CurrentTitle;
  achievements: Achievement[];
  nextTitleProgress: NextTitleProgress;
  statistics: AchievementStatistics;
}

// === DTO Types（API仕様準拠） ===
export interface SetCurrentTitleDto {
  is_current: boolean;
}

// === Query Parameters Types（API仕様準拠） ===
export interface AchievementsQueryParams extends Record<string, string | number | boolean | undefined> {
  include_progress?: boolean;
  sort_by?: 'achieved_at' | 'level';
  order?: 'asc' | 'desc';
}

