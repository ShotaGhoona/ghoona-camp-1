import { z } from 'zod';
import { updateSocialLinkFormSchema, type SocialLink } from '@/entities/user/social-links-entity';

// Entity層のvalidationを基盤として使用
export const socialLinkUpdateFormSchema = updateSocialLinkFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
});

export type SocialLinkUpdateFormData = z.infer<typeof socialLinkUpdateFormSchema>;

// UI特有の型定義
export interface SocialLinkUpdateFormProps {
  socialLink: SocialLink;
  onSubmit: (data: SocialLinkUpdateFormData) => void;
  isLoading?: boolean;
}

// 編集モード用の型
export interface SocialLinkEditState {
  isEditing: boolean;
  linkId: string | null;
}