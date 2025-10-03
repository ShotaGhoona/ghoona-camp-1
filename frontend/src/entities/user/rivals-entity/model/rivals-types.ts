/**
 * User Rivals Types
 * 
 * ユーザーライバル関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * rivals-entity/index.ts からエクスポートされた型を使用してください。
 */

// === Core Types ===
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
  createdAt: string;
}

export interface RivalsList {
  rivals: Rival[];
  count: number;
  maxRivals: number;
}

// === API Response Types ===

// === GET /users/{userId}/rivals レスポンス ===
export interface UserRivalsListResponse {
  data: RivalsList;
  message: string;
  timestamp: string;
}

// === POST /users/{userId}/rivals レスポンス ===
export interface RivalCreateResponse {
  data: Rival;
  message: string;
  timestamp: string;
}

// === DTO Types（API送信用） ===
export interface CreateRivalDto {
  rivalUserId: string;
}

// === Validation Constants ===
export const MAX_RIVALS_COUNT = 3;