import { z } from 'zod';
import { createSocialLinkFormSchema, type SocialPlatform } from '@/entities/user/social-links-entity';

// Entity層のvalidationを基盤として使用
export const socialLinkCreateFormSchema = createSocialLinkFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
  // 例：重複URL防止、プラットフォーム固有のURL形式チェックなど
});

export type SocialLinkCreateFormData = z.infer<typeof socialLinkCreateFormSchema>;

// UI特有の型定義
export interface SocialLinkCreateFormProps {
  onSubmit: (data: SocialLinkCreateFormData) => void;
  isLoading?: boolean;
  availablePlatforms?: SocialPlatform[];
}

// プラットフォーム選択用の型
export interface PlatformOption {
  value: SocialPlatform;
  label: string;
  icon?: string;
  placeholder?: string;
}