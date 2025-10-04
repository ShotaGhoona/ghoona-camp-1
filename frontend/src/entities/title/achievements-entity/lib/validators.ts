/**
 * Achievements Validators
 * 
 * ユーザー称号実績関連のバリデーション機能を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * achievements-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { z } from 'zod';
import type { 
  Achievement, 
  UserAchievements, 
  SetCurrentTitleDto,
  AchievementsQueryParams
} from '../model/achievements-types';

// === Zodスキーマ ===

/** 称号IDのバリデーション */
export const titleIdSchema = z
  .string()
  .uuid('有効な称号IDを入力してください');

/** ユーザー情報のスキーマ */
export const userInfoSchema = z.object({
  id: z.string().uuid(),
  displayName: z.string().min(1).max(100),
  username: z.string().min(1).max(50),
  avatarUrl: z.string().url(),
});

/** 実績エンティティのスキーマ（API仕様準拠） */
export const achievementSchema = z.object({
  id: z.string().uuid(),
  title: z.object({
    id: z.string().uuid(),
    level: z.number().int().min(1).max(8),
    nameJp: z.string().min(1).max(50),
    nameEn: z.string().min(1).max(50),
    description: z.string().min(1).max(2000),
    requiredDays: z.number().int().min(1).max(1000),
    colorTheme: z.string().regex(/^#[0-9A-Fa-f]{6}$/),
    imageUrl: z.string().url(),
    isActive: z.boolean(),
    createdAt: z.date(),
    updatedAt: z.date(),
  }),
  achievedAt: z.date(),
  isCurrent: z.boolean(),
  createdAt: z.date(),
  updatedAt: z.date(),
});

/** ユーザー実績エンティティのスキーマ（API仕様準拠） */
export const userAchievementsSchema = z.object({
  user: userInfoSchema,
  achievements: z.array(achievementSchema),
});

/** 称号設定DTOのスキーマ（API仕様準拠） */
export const setCurrentTitleDtoSchema = z.object({
  is_current: z.boolean(),
});

/** 実績クエリパラメータのスキーマ（API仕様準拠） */
export const achievementsQueryParamsSchema = z.object({
  include_progress: z.boolean().optional(),
  sort_by: z.enum(['achieved_at', 'level']).optional(),
  order: z.enum(['asc', 'desc']).optional(),
});

// === ビジネスルールバリデーション ===

/** 称号変更資格チェック（獲得済みの称号のみ設定可能） */
export const validateTitleChangeEligibility = (achievements: Achievement[], targetTitleId: string): boolean => {
  return achievements.some(achievement => achievement.title.id === targetTitleId);
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

/** 最高レベル称号のチェック */
export const isMaxLevelTitle = (level: number): boolean => {
  return level === 8;
};

/** 称号レベルの次のレベル取得 */
export const getNextTitleLevel = (currentLevel: number): number | null => {
  return currentLevel < 8 ? currentLevel + 1 : null;
};

/** 称号実績のソート（API仕様準拠） */
export const sortAchievements = (
  achievements: Achievement[], 
  sortBy: 'level' | 'achieved_at' = 'level',
  order: 'asc' | 'desc' = 'asc'
): Achievement[] => {
  return [...achievements].sort((a, b) => {
    let comparison = 0;
    
    if (sortBy === 'level') {
      comparison = a.title.level - b.title.level;
    } else {
      comparison = a.achievedAt.getTime() - b.achievedAt.getTime();
    }
    
    return order === 'desc' ? -comparison : comparison;
  });
};

// === Type Guards ===
export const isAchievement = (value: unknown): value is Achievement => {
  return achievementSchema.safeParse(value).success;
};

export const isUserAchievements = (value: unknown): value is UserAchievements => {
  return userAchievementsSchema.safeParse(value).success;
};

export const isValidSetCurrentTitleDto = (value: unknown): value is SetCurrentTitleDto => {
  return setCurrentTitleDtoSchema.safeParse(value).success;
};

export const isValidAchievementsQueryParams = (value: unknown): value is AchievementsQueryParams => {
  return achievementsQueryParamsSchema.safeParse(value).success;
};

// === Type exports ===
export type SetCurrentTitleFormData = z.infer<typeof setCurrentTitleDtoSchema>;
export type AchievementsQueryFormData = z.infer<typeof achievementsQueryParamsSchema>;