/**
 * Titles Types
 * 
 * 称号関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * titles-entity/index.ts からエクスポートされた型を使用してください。
 */

// === API Response Types ===

// 称号基本情報（API仕様準拠）
interface TitleResponse {
  id: string;
  level: number;
  name_jp: string;
  name_en: string;
  description: string;
  required_days: number;
  image_url: string;
  color_theme: string;
  holders_count: number;
  achievement_rate: number;
  user_status: {
    is_achieved: boolean;
    achieved_at?: string;
    is_current: boolean;
    progress?: {
      current_days: number;
      remaining_days: number;
      progress_rate: number;
    };
  };
  is_active: boolean;
  created_at: string;
}

// === GET /titles レスポンス（API仕様準拠） ===
export interface TitlesListResponse {
  data: {
    titles: TitleResponse[];
    user_summary: {
      achieved_count: number;
      total_count: number;
      current_title: {
        id: string;
        level: number;
        name_jp: string;
        color_theme: string;
      };
      next_title: {
        id: string;
        level: number;
        name_jp: string;
        progress_rate: number;
        remaining_days: number;
      };
    };
  };
  message: string;
  timestamp: string;
}

// === GET /titles/{titleId} レスポンス（API仕様準拠） ===
export interface TitleDetailResponse {
  data: {
    id: string;
    level: number;
    name_jp: string;
    name_en: string;
    description: string;
    story: string;
    required_days: number;
    image_url: string;
    color_theme: string;
    badge_design: {
      primary_color: string;
      secondary_color: string;
      icon: string;
      pattern: string;
    };
    achievement_conditions: Array<{
      type: string;
      value: number;
      description: string;
    }>;
    statistics: {
      total_holders: number;
      achievement_rate: number;
      average_achievement_days: number;
      recent_achievers_count: number;
    };
    recent_achievers: Array<{
      user: {
        id: string;
        display_name: string;
        username: string;
        avatar_url: string;
      };
      achieved_at: string;
    }>;
    user_status: {
      is_achieved: boolean;
      progress?: {
        current_days: number;
        remaining_days: number;
        progress_rate: number;
        estimated_achievement_date: string;
      };
    };
    perks: Array<{
      type: string;
      description: string;
    }>;
    next_title?: {
      id: string;
      level: number;
      name_jp: string;
      required_days: number;
      additional_days_needed: number;
    };
    is_active: boolean;
    created_at: string;
    updated_at: string;
  };
  message: string;
  timestamp: string;
}

// === Client Types（camelCase - フロントエンド用） ===
export interface Title {
  id: string;
  level: number;
  nameJp: string;
  nameEn: string;
  description: string;
  requiredDays: number;
  imageUrl: string;
  colorTheme: string;
  holdersCount: number;
  achievementRate: number;
  userStatus: {
    isAchieved: boolean;
    achievedAt?: Date;
    isCurrent: boolean;
    progress?: {
      currentDays: number;
      remainingDays: number;
      progressRate: number;
    };
  };
  isActive: boolean;
  createdAt: Date;
}

export interface TitleDetail extends Title {
  story: string;
  badgeDesign: {
    primaryColor: string;
    secondaryColor: string;
    icon: string;
    pattern: string;
  };
  achievementConditions: Array<{
    type: string;
    value: number;
    description: string;
  }>;
  statistics: {
    totalHolders: number;
    achievementRate: number;
    averageAchievementDays: number;
    recentAchieversCount: number;
  };
  recentAchievers: Array<{
    user: {
      id: string;
      displayName: string;
      username: string;
      avatarUrl: string;
    };
    achievedAt: Date;
  }>;
  userStatus: {
    isAchieved: boolean;
    achievedAt?: Date;
    isCurrent: boolean;
    progress?: {
      currentDays: number;
      remainingDays: number;
      progressRate: number;
      estimatedAchievementDate: Date;
    };
  };
  perks: Array<{
    type: string;
    description: string;
  }>;
  nextTitle?: {
    id: string;
    level: number;
    nameJp: string;
    requiredDays: number;
    additionalDaysNeeded: number;
  };
  updatedAt: Date;
}

export interface UserSummary {
  achievedCount: number;
  totalCount: number;
  currentTitle: {
    id: string;
    level: number;
    nameJp: string;
    colorTheme: string;
  };
  nextTitle: {
    id: string;
    level: number;
    nameJp: string;
    progressRate: number;
    remainingDays: number;
  };
}

// === Title Level Constants ===
export const TITLE_LEVELS = {
  MIN: 1,
  MAX: 8,
} as const;

export type TitleLevel = 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8;

// === Title Names (from API spec) ===
export const TITLE_NAMES = {
  1: { jp: 'まどろみ見習い', en: 'Sleeper' },
  2: { jp: '早起き戦士', en: 'Early Riser' },
  3: { jp: '陽光探求者', en: 'Dawn Seeker' },
  4: { jp: '朝陽の使者', en: 'Sun Messenger' },
  5: { jp: '暁の守護者', en: 'Dawn Guardian' },
  6: { jp: '朝活の覇者', en: 'Morning Master' },
  7: { jp: '夜明けの皇帝', en: 'Dawn Emperor' },
  8: { jp: '永遠の夜明け', en: 'Eternal Dawn' },
} as const;

// === Title Required Days (from API spec) ===
export const TITLE_REQUIRED_DAYS = {
  1: 1,
  2: 5,
  3: 15,
  4: 30,
  5: 60,
  6: 100,
  7: 200,
  8: 365,
} as const;

// === Query Parameters Types（API仕様準拠） ===
export interface TitlesQueryParams extends Record<string, string | number | boolean | undefined> {
  include_inactive?: boolean;
  user_id?: string;
}