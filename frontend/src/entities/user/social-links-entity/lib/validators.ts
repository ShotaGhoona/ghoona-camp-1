/**
 * User Social Links Validators
 * 
 * ユーザーソーシャルリンク関連のバリデーション機能を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * social-links-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { z } from 'zod';
import type { SocialLink, CreateSocialLinkDto, UpdateSocialLinkDto, SocialPlatform } from '../model/social-links-types';
import { SOCIAL_PLATFORMS } from '../model/social-links-types';

// === Zodスキーマ ===

/** プラットフォームのバリデーション */
export const platformSchema = z.enum(SOCIAL_PLATFORMS as [SocialPlatform, ...SocialPlatform[]], {
  message: '有効なプラットフォームを選択してください',
});

/** URLのバリデーション */
export const urlSchema = z
  .string()
  .url('有効なURLを入力してください')
  .min(1, 'URLは必須です');

/** タイトルのバリデーション */
export const titleSchema = z
  .string()
  .max(100, 'タイトルは100文字以内で入力してください');

/** ソーシャルリンク作成フォームのバリデーション */
export const createSocialLinkFormSchema = z.object({
  platform: platformSchema,
  url: urlSchema,
  title: titleSchema.optional(),
  isPublic: z.boolean().default(true),
});

/** ソーシャルリンク更新フォームのバリデーション */
export const updateSocialLinkFormSchema = z.object({
  url: urlSchema.optional(),
  title: titleSchema.optional(),
  isPublic: z.boolean().optional(),
});

/** ソーシャルリンクエンティティのスキーマ */
export const socialLinkSchema = z.object({
  id: z.string().uuid(),
  userId: z.string().uuid(),
  platform: platformSchema,
  url: z.string().url(),
  title: z.string().nullable(),
  isPublic: z.boolean(),
  createdAt: z.string(),
  updatedAt: z.string(),
});

// === Type Guards ===
export const isSocialLink = (value: unknown): value is SocialLink => {
  return socialLinkSchema.safeParse(value).success;
};

export const isValidCreateSocialLinkDto = (value: unknown): value is CreateSocialLinkDto => {
  return createSocialLinkFormSchema.safeParse(value).success;
};

export const isValidUpdateSocialLinkDto = (value: unknown): value is UpdateSocialLinkDto => {
  return updateSocialLinkFormSchema.safeParse(value).success;
};

export const isSocialPlatform = (value: unknown): value is SocialPlatform => {
  return platformSchema.safeParse(value).success;
};

// === Type exports ===
export type CreateSocialLinkFormData = z.infer<typeof createSocialLinkFormSchema>;
export type UpdateSocialLinkFormData = z.infer<typeof updateSocialLinkFormSchema>;