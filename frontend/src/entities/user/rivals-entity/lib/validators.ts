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
  clerkId: z.string().min(1),
  email: z.string().email(),
  username: z.string().nullable(),
  avatarUrl: z.string().url().nullable(),
  discordId: z.string().nullable(),
  status: z.enum(['active', 'inactive', 'suspended']).default('active'),
  createdAt: z.string(),
  updatedAt: z.string(),
});

/** ライバルエンティティのスキーマ */
export const rivalSchema = z.object({
  id: z.string().uuid(),
  userId: z.string().uuid(),
  rivalUser: rivalUserSchema,
  createdAt: z.string(),
});

/** ライバルリストのスキーマ */
export const rivalsListSchema = z.object({
  rivals: z.array(rivalSchema),
  total: z.number().min(0),
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