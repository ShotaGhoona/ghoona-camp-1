/**
 * User Rivals Validators
 * 
 * ユーザーライバル関連のバリデーション機能を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * rivals-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { z } from 'zod';
import type { Rival, RivalsList, CreateRivalDto } from '../model/rivals-types';
import { MAX_RIVALS_COUNT } from '../model/rivals-types';

// === Zodスキーマ ===

/** ライバルユーザーIDのバリデーション */
export const rivalUserIdSchema = z
  .string()
  .uuid('有効なユーザーIDを入力してください');

/** ライバル作成フォームのバリデーション */
export const createRivalFormSchema = z.object({
  rivalUserId: rivalUserIdSchema,
});

/** ライバルユーザーのスキーマ */
export const rivalUserSchema = z.object({
  id: z.string().uuid(),
  displayName: z.string(),
  username: z.string(),
  avatarUrl: z.string().url(),
  currentTitle: z.object({
    level: z.number(),
    nameJp: z.string(),
    colorTheme: z.string(),
  }),
  attendanceStats: z.object({
    totalAttendanceDays: z.number(),
    currentStreakDays: z.number(),
  }),
});

/** ライバルエンティティのスキーマ */
export const rivalSchema = z.object({
  id: z.string().uuid(),
  rivalUser: rivalUserSchema,
  createdAt: z.date(),
});

/** ライバルリストのスキーマ */
export const rivalsListSchema = z.object({
  rivals: z.array(rivalSchema),
  count: z.number().min(0).max(MAX_RIVALS_COUNT),
  maxRivals: z.number().default(MAX_RIVALS_COUNT),
});

// === ビジネスルールバリデーション ===

/** ライバル数の上限チェック */
export const validateRivalCount = (currentCount: number): boolean => {
  return currentCount < MAX_RIVALS_COUNT;
};

/** 自分自身をライバルに設定できないチェック */
export const validateNotSelfRival = (userId: string, rivalUserId: string): boolean => {
  return userId !== rivalUserId;
};

/** 重複ライバルのチェック */
export const validateDuplicateRival = (existingRivals: Rival[], newRivalUserId: string): boolean => {
  return !existingRivals.some(rival => rival.rivalUser.id === newRivalUserId);
};

// === Type Guards ===
export const isRival = (value: unknown): value is Rival => {
  return rivalSchema.safeParse(value).success;
};

export const isRivalsList = (value: unknown): value is RivalsList => {
  return rivalsListSchema.safeParse(value).success;
};

export const isValidCreateRivalDto = (value: unknown): value is CreateRivalDto => {
  return createRivalFormSchema.safeParse(value).success;
};

// === Type exports ===
export type CreateRivalFormData = z.infer<typeof createRivalFormSchema>;