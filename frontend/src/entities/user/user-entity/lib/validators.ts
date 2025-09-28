import { z } from 'zod';
import type { User, UpdateUserDto } from '../model/types';

// === Zodスキーマ ===

/** ユーザー名のバリデーション */
export const usernameSchema = z
  .string()
  .min(3, 'ユーザー名は3文字以上で入力してください')
  .max(50, 'ユーザー名は50文字以内で入力してください')
  .regex(/^[a-zA-Z0-9_-]+$/, '英数字、アンダースコア、ハイフンのみ使用可能です');

/** ユーザー基本情報更新フォームのバリデーション */
export const updateUserFormSchema = z.object({
  username: usernameSchema.optional(),
  avatar_url: z.string().url('有効なURLを入力してください').optional(),
});

/** ユーザーエンティティのスキーマ */
export const userSchema = z.object({
  id: z.string().uuid(),
  clerkId: z.string(),
  email: z.string().email(),
  username: z.string(),
  avatarUrl: z.string().url(),
  discordId: z.string().optional(),
  isActive: z.boolean(),
  createdAt: z.date(),
  updatedAt: z.date(),
});

// === Type Guards ===
export const isUser = (value: unknown): value is User => {
  return userSchema.safeParse(value).success;
};

export const isValidUpdateUserDto = (value: unknown): value is UpdateUserDto => {
  return updateUserFormSchema.safeParse(value).success;
};

// === Type exports ===
export type UpdateUserFormData = z.infer<typeof updateUserFormSchema>;