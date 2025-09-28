/**
 * Titles Types
 * 
 * 称号関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * titles-entity/index.ts からエクスポートされた型を使用してください。
 */

// === API Response Types（API仕様準拠） ===

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
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

// === GET /titles レスポンス（API仕様準拠） ===
export interface TitlesListResponse {
  data: TitleResponse[];
  message: string;
  timestamp: string;
}

// === GET /titles/{titleId} レスポンス（API仕様準拠） ===
export interface TitleDetailResponse {
  data: TitleResponse;
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
  isActive: boolean;
  createdAt: Date;
  updatedAt: Date;
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