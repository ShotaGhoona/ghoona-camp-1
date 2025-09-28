/**
 * Titles Validators
 * 
 * 称号関連のバリデーション機能を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * titles-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { z } from 'zod';
import type { Title, TitleLevel, TitlesQueryParams } from '../model/titles-types';
import { TITLE_LEVELS, TITLE_REQUIRED_DAYS } from '../model/titles-types';

// === Zodスキーマ ===

/** 称号レベルのバリデーション */
export const titleLevelSchema = z
  .number()
  .int('レベルは整数である必要があります')
  .min(TITLE_LEVELS.MIN, `レベルは${TITLE_LEVELS.MIN}以上である必要があります`)
  .max(TITLE_LEVELS.MAX, `レベルは${TITLE_LEVELS.MAX}以下である必要があります`);

/** 称号名（日本語）のバリデーション */
export const titleNameJpSchema = z
  .string()
  .min(1, '称号名（日本語）は必須です')
  .max(50, '称号名（日本語）は50文字以内で入力してください');

/** 称号名（英語）のバリデーション */
export const titleNameEnSchema = z
  .string()
  .min(1, '称号名（英語）は必須です')
  .max(50, '称号名（英語）は50文字以内で入力してください');

/** 称号説明のバリデーション */
export const titleDescriptionSchema = z
  .string()
  .min(1, '称号説明は必須です')
  .max(2000, '称号説明は2000文字以内で入力してください');

/** 必要参加日数のバリデーション */
export const requiredDaysSchema = z
  .number()
  .int('必要参加日数は整数である必要があります')
  .min(1, '必要参加日数は1以上である必要があります')
  .max(1000, '必要参加日数は1000以下である必要があります');

/** カラーテーマのバリデーション */
export const colorThemeSchema = z
  .string()
  .regex(/^#[0-9A-Fa-f]{6}$/, '有効なカラーコードを入力してください（例: #FF0000）');

/** 称号クエリパラメータのバリデーション（API仕様準拠） */
export const titlesQueryParamsSchema = z.object({
  include_inactive: z.boolean().optional(),
  user_id: z.string().uuid().optional(),
});

/** 称号エンティティのスキーマ（API仕様準拠） */
export const titleSchema = z.object({
  id: z.string().uuid(),
  level: titleLevelSchema,
  nameJp: titleNameJpSchema,
  nameEn: titleNameEnSchema,
  description: titleDescriptionSchema,
  requiredDays: requiredDaysSchema,
  imageUrl: z.string().url(),
  colorTheme: colorThemeSchema,
  isActive: z.boolean(),
  createdAt: z.date(),
  updatedAt: z.date(),
});

// === ビジネスルールバリデーション ===

/** 称号レベルの連続性チェック */
export const validateTitleLevelSequence = (titles: Title[]): boolean => {
  const sortedTitles = titles.sort((a, b) => a.level - b.level);
  for (let i = 0; i < sortedTitles.length; i++) {
    if (sortedTitles[i].level !== i + 1) {
      return false;
    }
  }
  return true;
};

/** 必要日数の順序チェック（API仕様準拠） */
export const validateRequiredDaysSequence = (titles: Title[]): boolean => {
  const sortedTitles = titles.sort((a, b) => a.level - b.level);
  for (let i = 0; i < sortedTitles.length; i++) {
    const expectedDays = TITLE_REQUIRED_DAYS[sortedTitles[i].level as TitleLevel];
    if (sortedTitles[i].requiredDays !== expectedDays) {
      return false;
    }
  }
  return true;
};

/** 称号取得資格チェック */
export const validateTitleEligibility = (userAttendanceDays: number, requiredDays: number): boolean => {
  return userAttendanceDays >= requiredDays;
};

/** 段階的獲得ルールのチェック（下位称号を先に獲得する必要がある） */
export const validateProgressiveAchievement = (
  userAchievements: Array<{ level: number }>,
  targetLevel: number
): boolean => {
  // レベル1は常に取得可能
  if (targetLevel === 1) return true;
  
  // 1つ下のレベルが獲得済みかチェック
  return userAchievements.some(achievement => achievement.level === targetLevel - 1);
};

/** 進捗率の計算 */
export const calculateProgressRate = (currentDays: number, requiredDays: number): number => {
  if (requiredDays <= 0) return 0;
  return Math.min((currentDays / requiredDays) * 100, 100);
};

/** 残り日数の計算 */
export const calculateRemainingDays = (currentDays: number, requiredDays: number): number => {
  return Math.max(requiredDays - currentDays, 0);
};

// === Type Guards ===
export const isTitle = (value: unknown): value is Title => {
  return titleSchema.safeParse(value).success;
};

export const isTitleLevel = (value: unknown): value is TitleLevel => {
  return titleLevelSchema.safeParse(value).success;
};

export const isValidTitlesQueryParams = (value: unknown): value is TitlesQueryParams => {
  return titlesQueryParamsSchema.safeParse(value).success;
};

// === Type exports ===
export type TitlesQueryFormData = z.infer<typeof titlesQueryParamsSchema>;