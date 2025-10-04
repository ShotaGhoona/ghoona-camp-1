import { z } from 'zod';
import { updateSocialLinkFormSchema, type SocialLink, type SocialPlatform } from '@/entities/user/social-links-entity';
import { platformOptions, type PlatformOption, type UrlValidationResult } from '../../social-links-create/model/social-link-form-types';

// Entity層のvalidationを基盤として使用
export const socialLinkUpdateFormSchema = updateSocialLinkFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
  // 例：変更差分チェック、URL重複防止など
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
  editingLinkId: string | null;
  hasChanges: boolean;
}

// インライン編集用の型
export interface InlineEditProps {
  value: string;
  onSave: (value: string) => void;
  onCancel: () => void;
  isLoading?: boolean;
  placeholder?: string;
  multiline?: boolean;
  maxLength?: number;
}

// ソーシャルリンクカード用の型
export interface SocialLinkCardProps {
  socialLink: SocialLink;
  isEditing: boolean;
  onEdit: () => void;
  onSave: (data: SocialLinkUpdateFormData) => void;
  onCancel: () => void;
  onDelete: () => void;
  isLoading?: boolean;
}

// プラットフォーム情報表示用
export interface PlatformDisplayInfo {
  platform: SocialPlatform;
  option: PlatformOption;
  isEditable: boolean; // プラットフォーム自体は通常編集不可
}

// URL編集用の型
export interface UrlEditProps {
  currentUrl: string;
  platform: SocialPlatform;
  onUrlChange: (url: string) => void;
  onValidation: (result: UrlValidationResult) => void;
  isLoading?: boolean;
}

// タイトル編集用の型
export interface TitleEditProps {
  currentTitle?: string;
  platform: SocialPlatform;
  url: string;
  onTitleChange: (title: string) => void;
  autoSuggest?: boolean;
}

// 公開設定編集用の型
export interface VisibilityEditProps {
  isPublic: boolean;
  onVisibilityChange: (isPublic: boolean) => void;
  description?: string;
}

// 変更追跡用の型
export interface SocialLinkChanges {
  hasChanges: boolean;
  changedFields: Set<keyof SocialLinkUpdateFormData>;
  originalData: SocialLinkUpdateFormData;
}

// 一括編集用の型
export interface BulkEditProps {
  socialLinks: SocialLink[];
  selectedLinkIds: Set<string>;
  onSelectionChange: (linkIds: Set<string>) => void;
  onBulkVisibilityChange: (isPublic: boolean) => void;
  onBulkDelete: () => void;
}

// ソート・フィルター用の型
export interface SocialLinksViewOptions {
  sortBy: 'platform' | 'title' | 'createdAt' | 'isPublic';
  sortOrder: 'asc' | 'desc';
  filterPlatform?: SocialPlatform;
  showPublicOnly?: boolean;
}

// エクスポート用の型
export interface SocialLinksExportData {
  format: 'json' | 'csv' | 'markdown';
  includePrivate: boolean;
  groupByPlatform: boolean;
}

// 定数の再エクスポート
export { platformOptions };
export type { PlatformOption, UrlValidationResult };