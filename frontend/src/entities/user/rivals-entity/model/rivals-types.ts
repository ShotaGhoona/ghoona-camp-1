/**
 * User Rivals Types
 * 
 * ユーザーライバル関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * rivals-entity/index.ts からエクスポートされた型を使用してください。
 */

// === API Response Types ===

// ライバル関係（rivals-entity用）- 内部実装
interface RivalResponse {
  id: string;
  rival_user: {
    id: string;
    display_name: string;
    username: string;
    avatar_url: string;
    current_title: {
      level: number;
      name_jp: string;
      color_theme: string;
    };
    attendance_stats: {
      total_attendance_days: number;
      current_streak_days: number;
    };
  };
  created_at: string;
}

// === GET /users/{userId}/rivals レスポンス（統一形式） ===
export interface UserRivalsListResponse {
  data: {
    rivals: RivalResponse[];
    count: number;
    max_rivals: number;
  };
  message: string;
  timestamp: string;
}

// === POST /users/{userId}/rivals レスポンス（統一形式） ===
export interface RivalCreateResponse {
  data: {
    id: string;
    rival_user: {
      id: string;
      display_name: string;
      username: string;
      avatar_url: string;
      current_title: {
        level: number;
        name_jp: string;
        color_theme: string;
      };
    };
    created_at: string;
  };
  message: string;
  timestamp: string;
}

// === Client Types（camelCase - フロントエンド用） ===
export interface RivalUser {
  id: string;
  displayName: string;
  username: string;
  avatarUrl: string;
  currentTitle: {
    level: number;
    nameJp: string;
    colorTheme: string;
  };
  attendanceStats: {
    totalAttendanceDays: number;
    currentStreakDays: number;
  };
}

export interface Rival {
  id: string;
  rivalUser: RivalUser;
  createdAt: Date;
}

export interface RivalsList {
  rivals: Rival[];
  count: number;
  maxRivals: number;
}

// === DTO Types（API送信用） ===
export interface CreateRivalDto {
  rival_user_id: string;
}

// === Validation Constants ===
export const MAX_RIVALS_COUNT = 3;