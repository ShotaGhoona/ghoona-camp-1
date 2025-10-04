/**
 * User Metadata Validators
 * 
 * ユーザーメタデータ関連のバリデーション機能を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * metadata-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { z } from 'zod';
import type { UserMetadata, CreateUserMetadataDto, UpdateUserMetadataDto } from '../model/metadata-types';

// === Zodスキーマ ===

/** 表示名のバリデーション */
export const displayNameSchema = z
  .string()
  .min(1, '表示名は必須です')
  .max(100, '表示名は100文字以内で入力してください');

/** 一言プロフィールのバリデーション */
export const taglineSchema = z
  .string()
  .min(1, '一言プロフィールを入力してください')
  .max(150, '一言プロフィールは150文字以内で入力してください');

/** 自己紹介のバリデーション */
export const bioSchema = z
  .string()
  .min(1, '自己紹介を入力してください')
  .max(1000, '自己紹介は1000文字以内で入力してください');

/** ビジョンのバリデーション */
export const visionSchema = z
  .string()
  .min(1, 'ビジョンを入力してください')
  .max(2000, 'ビジョンは2000文字以内で入力してください');

/** スキル配列のバリデーション */
export const skillsSchema = z
  .array(z.string().min(1, 'スキル名を入力してください').max(50, 'スキルは50文字以内で入力してください'))
  .max(20, 'スキルは最大20個まで設定できます');

/** 興味・関心配列のバリデーション */
export const interestsSchema = z
  .array(z.string().min(1, '興味・関心を入力してください').max(50, '興味・関心は50文字以内で入力してください'))
  .max(20, '興味・関心は最大20個まで設定できます');

/** タイムゾーンのバリデーション */
export const timezoneSchema = z
  .string()
  .regex(/^[A-Za-z_]+\/[A-Za-z_0-9]+$/, '有効なタイムゾーンを入力してください')
  .default('Asia/Tokyo');

/** ユーザーメタデータ作成フォームのバリデーション */
export const createUserMetadataFormSchema = z.object({
  displayName: displayNameSchema.optional(),
  profileImageUrl: z.string().url('有効なURLを入力してください').optional(),
  tagline: taglineSchema.optional(),
  bio: bioSchema.optional(),
  vision: visionSchema.optional(),
  visionPublic: z.boolean().optional(),
  timezone: timezoneSchema.optional(),
  skills: skillsSchema.optional(),
  interests: interestsSchema.optional(),
});

/** ユーザーメタデータ更新フォームのバリデーション */
export const updateUserMetadataFormSchema = z.object({
  displayName: displayNameSchema.optional(),
  profileImageUrl: z.string().url('有効なURLを入力してください').optional(),
  tagline: taglineSchema.optional(),
  bio: bioSchema.optional(),
  vision: visionSchema.optional(),
  visionPublic: z.boolean().optional(),
  timezone: timezoneSchema.optional(),
  skills: skillsSchema.optional(),
  interests: interestsSchema.optional(),
});

/** ユーザーメタデータエンティティのスキーマ */
export const userMetadataSchema = z.object({
  id: z.string().uuid(),
  userId: z.string().uuid(),
  displayName: z.string().nullable(),
  profileImageUrl: z.string().url().nullable(),
  tagline: z.string().nullable(),
  bio: z.string().nullable(),
  vision: z.string().nullable(),
  visionPublic: z.boolean(),
  timezone: z.string(),
  skills: z.array(z.string()),
  interests: z.array(z.string()),
  createdAt: z.string(),
  updatedAt: z.string(),
});

// === Type Guards ===
export const isUserMetadata = (value: unknown): value is UserMetadata => {
  return userMetadataSchema.safeParse(value).success;
};

export const isValidCreateUserMetadataDto = (value: unknown): value is CreateUserMetadataDto => {
  return createUserMetadataFormSchema.safeParse(value).success;
};

export const isValidUpdateUserMetadataDto = (value: unknown): value is UpdateUserMetadataDto => {
  return updateUserMetadataFormSchema.safeParse(value).success;
};

// === Type exports ===
export type CreateUserMetadataFormData = z.infer<typeof createUserMetadataFormSchema>;
export type UpdateUserMetadataFormData = z.infer<typeof updateUserMetadataFormSchema>;