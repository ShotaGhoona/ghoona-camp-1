/**
 * Metadata Update Form Types
 * 
 * メタデータ更新フォーム専用の型定義を提供します。
 * フォームライブラリ（react-hook-form）との連携用です。
 */

import { z } from 'zod';
import { updateUserMetadataFormSchema, type UpdateUserMetadataDto } from '@/entities/user';

// === Form Data Types ===

/**
 * メタデータ更新フォームのデータ型
 * react-hook-formで使用
 */
export type UpdateMetadataFormData = z.infer<typeof updateUserMetadataFormSchema>;

/**
 * APIへ送信するDTO型の再エクスポート
 */
export type { UpdateUserMetadataDto };

// === Form Schema Re-export ===

/**
 * バリデーションスキーマの再エクスポート
 */
export { updateUserMetadataFormSchema };

// === Form Configuration ===

/**
 * フォームのデフォルト値
 */
export const defaultMetadataFormValues: Partial<UpdateMetadataFormData> = {
  display_name: '',
  tagline: '',
  bio: '',
  vision: '',
  vision_public: false,
  timezone: 'Asia/Tokyo',
  skills: [],
  interests: [],
};

/**
 * フォームフィールドの制限値
 */
export const METADATA_FORM_LIMITS = {
  DISPLAY_NAME_MAX: 100,
  TAGLINE_MAX: 150,
  BIO_MAX: 1000,
  VISION_MAX: 2000,
  SKILLS_MAX: 20,
  INTERESTS_MAX: 20,
  SKILL_ITEM_MAX: 50,
  INTEREST_ITEM_MAX: 50,
} as const;