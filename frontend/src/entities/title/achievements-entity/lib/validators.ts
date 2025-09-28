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
  CurrentTitle, 
  NextTitleProgress, 
  UserAchievements, 
  SetCurrentTitleDto,
  AchievementsQueryParams,
  AchievementStatistics,
  UserInfo
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
    colorTheme: z.string().regex(/^#[0-9A-Fa-f]{6}$/),
    imageUrl: z.string().url(),
  }),
  achievedAt: z.date(),
  isCurrent: z.boolean(),
  achievementRank: z.number().int().min(1),
  totalAchieversAtTime: z.number().int().min(0),
});

/** 現在称号エンティティのスキーマ（API仕様準拠） */
export const currentTitleSchema = z.object({
  id: z.string().uuid(),
  level: z.number().int().min(1).max(8),
  nameJp: z.string().min(1).max(50),
  nameEn: z.string().min(1).max(50),
  colorTheme: z.string().regex(/^#[0-9A-Fa-f]{6}$/),
  imageUrl: z.string().url(),
  achievedAt: z.date(),
  daysHeld: z.number().int().min(0),
});

/** 次の称号進捗エンティティのスキーマ */
export const nextTitleProgressSchema = z.object({
  title: z.object({
    id: z.string().uuid(),
    level: z.number().int().min(1).max(8),
    nameJp: z.string().min(1).max(50),
    requiredDays: z.number().int().min(1),
  }),
  progress: z.object({
    currentDays: z.number().int().min(0),
    remainingDays: z.number().int().min(0),
    progressRate: z.number().min(0).max(100),
    estimatedAchievementDate: z.date(),
  }),
});

/** 実績統計情報のスキーマ */
export const achievementStatisticsSchema = z.object({
  totalAchieved: z.number().int().min(0),
  maxLevelAchieved: z.number().int().min(1).max(8),
  consistencyScore: z.number().min(0).max(100),
  totalAttendanceDays: z.number().int().min(0),
});

/** ユーザー実績エンティティのスキーマ（API仕様準拠） */
export const userAchievementsSchema = z.object({
  user: userInfoSchema,
  currentTitle: currentTitleSchema,
  achievements: z.array(achievementSchema),
  nextTitleProgress: nextTitleProgressSchema,
  statistics: achievementStatisticsSchema,
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


/** 一貫性スコアの計算（出席率ベース） */
export const calculateConsistencyScore = (totalAttendanceDays: number, totalPossibleDays: number): number => {
  if (totalPossibleDays <= 0) return 0;
  return Math.min((totalAttendanceDays / totalPossibleDays) * 100, 100);
};

// === Type Guards ===
export const isAchievement = (value: unknown): value is Achievement => {
  return achievementSchema.safeParse(value).success;
};

export const isCurrentTitle = (value: unknown): value is CurrentTitle => {
  return currentTitleSchema.safeParse(value).success;
};

export const isNextTitleProgress = (value: unknown): value is NextTitleProgress => {
  return nextTitleProgressSchema.safeParse(value).success;
};

export const isUserAchievements = (value: unknown): value is UserAchievements => {
  return userAchievementsSchema.safeParse(value).success;
};

export const isAchievementStatistics = (value: unknown): value is AchievementStatistics => {
  return achievementStatisticsSchema.safeParse(value).success;
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