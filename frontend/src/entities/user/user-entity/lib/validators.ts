/**
 * User Validators
 * 
 * ユーザー関連のバリデーション機能を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * user-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { z } from 'zod';
import type { User, CreateUserDto, UpdateUserDto } from '../model/user-types';

// === Zodスキーマ ===

/** ユーザー名のバリデーション */
export const usernameSchema = z
  .string()
  .min(1, 'ユーザー名は必須です')
  .max(50, 'ユーザー名は50文字以内で入力してください');

/** ユーザー作成フォームのバリデーション */
export const createUserFormSchema = z.object({
  clerkId: z.string().min(1, 'Clerk IDは必須です'),
  email: z.string().email('有効なメールアドレスを入力してください'),
  username: usernameSchema.optional(),
  avatarUrl: z.string().url('有効なURLを入力してください').optional(),
  discordId: z.string().optional(),
});

/** ユーザー基本情報更新フォームのバリデーション */
export const updateUserFormSchema = z.object({
  username: usernameSchema.optional(),
  avatarUrl: z.string().url('有効なURLを入力してください').optional(),
  discordId: z.string().optional(),
});

/** ユーザーエンティティのスキーマ */
export const userSchema = z.object({
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

// === Type Guards ===
export const isUser = (value: unknown): value is User => {
  return userSchema.safeParse(value).success;
};

export const isValidCreateUserDto = (value: unknown): value is CreateUserDto => {
  return createUserFormSchema.safeParse(value).success;
};

export const isValidUpdateUserDto = (value: unknown): value is UpdateUserDto => {
  return updateUserFormSchema.safeParse(value).success;
};

// === Type exports ===
export type CreateUserFormData = z.infer<typeof createUserFormSchema>;
export type UpdateUserFormData = z.infer<typeof updateUserFormSchema>;