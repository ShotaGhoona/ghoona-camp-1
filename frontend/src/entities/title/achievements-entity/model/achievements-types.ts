/**
 * Achievements Types
 * 
 * ユーザー称号実績関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * achievements-entity/index.ts からエクスポートされた型を使用してください。
 */

// === API Response Types（API仕様準拠） ===

// 称号実績レスポンス（API仕様準拠）
interface AchievementResponse {
  id: string;
  title: {
    id: string;
    level: number;
    name_jp: string;
    name_en: string;
    description: string;
    required_days: number;
    image_url: string;
    color_theme: string;
    is_active: boolean;
    created_at: string;
    updated_at: string;
  };
  achieved_at: string;
  is_current: boolean;
  created_at: string;
  updated_at: string;
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
    achievements: AchievementResponse[];
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
      description: string;
      required_days: number;
      image_url: string;
      color_theme: string;
      is_active: boolean;
      created_at: string;
      updated_at: string;
    };
    achieved_at: string;
    is_current: boolean;
    updated_at: string;
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
    description: string;
    requiredDays: number;
    imageUrl: string;
    colorTheme: string;
    isActive: boolean;
    createdAt: Date;
    updatedAt: Date;
  };
  achievedAt: Date;
  isCurrent: boolean;
  createdAt: Date;
  updatedAt: Date;
}

export interface UserInfo {
  id: string;
  displayName: string;
  username: string;
  avatarUrl: string;
}

export interface UserAchievements {
  user: UserInfo;
  achievements: Achievement[];
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