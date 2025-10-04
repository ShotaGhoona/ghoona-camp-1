import { z } from 'zod';
import { createSocialLinkFormSchema, type SocialPlatform } from '@/entities/user/social-links-entity';

// Entity層のvalidationを基盤として使用
export const socialLinkCreateFormSchema = createSocialLinkFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
  // 例：重複URL防止、プラットフォーム固有のURL形式チェック、リンク検証など
});

export type SocialLinkCreateFormData = z.infer<typeof socialLinkCreateFormSchema>;

// UI特有の型定義
export interface SocialLinkCreateFormProps {
  onSubmit: (data: SocialLinkCreateFormData) => void;
  isLoading?: boolean;
  availablePlatforms?: SocialPlatform[];
  existingPlatforms?: SocialPlatform[];
}

// プラットフォーム選択用の型
export interface PlatformOption {
  value: SocialPlatform;
  label: string;
  icon?: string;
  placeholder?: string;
  urlPattern?: RegExp;
  description?: string;
}

// プラットフォーム設定（create用に特化）
export const platformOptions: PlatformOption[] = [
  {
    value: 'twitter',
    label: 'Twitter/X',
    icon: '𝕏',
    placeholder: 'https://twitter.com/username',
    urlPattern: /^https?:\/\/(twitter\.com|x\.com)\/.+/,
    description: 'Twitter/Xプロフィール',
  },
  {
    value: 'github',
    label: 'GitHub',
    icon: '🐙',
    placeholder: 'https://github.com/username',
    urlPattern: /^https?:\/\/github\.com\/.+/,
    description: 'GitHubプロフィール',
  },
  {
    value: 'linkedin',
    label: 'LinkedIn',
    icon: '💼',
    placeholder: 'https://linkedin.com/in/username',
    urlPattern: /^https?:\/\/linkedin\.com\/in\/.+/,
    description: 'LinkedInプロフィール',
  },
  {
    value: 'instagram',
    label: 'Instagram',
    icon: '📸',
    placeholder: 'https://instagram.com/username',
    urlPattern: /^https?:\/\/instagram\.com\/.+/,
    description: 'Instagramプロフィール',
  },
  {
    value: 'youtube',
    label: 'YouTube',
    icon: '📺',
    placeholder: 'https://youtube.com/@username',
    urlPattern: /^https?:\/\/(youtube\.com|youtu\.be)\/.+/,
    description: 'YouTubeチャンネル',
  },
  {
    value: 'website',
    label: 'Website',
    icon: '🌐',
    placeholder: 'https://your-website.com',
    urlPattern: /^https?:\/\/.+/,
    description: '個人ウェブサイト',
  },
  {
    value: 'blog',
    label: 'Blog',
    icon: '📝',
    placeholder: 'https://your-blog.com',
    urlPattern: /^https?:\/\/.+/,
    description: 'ブログ',
  },
  {
    value: 'discord',
    label: 'Discord',
    icon: '💬',
    placeholder: 'https://discord.gg/server',
    urlPattern: /^https?:\/\/discord\.(gg|com)\/.+/,
    description: 'Discordサーバー',
  },
  {
    value: 'twitch',
    label: 'Twitch',
    icon: '🎮',
    placeholder: 'https://twitch.tv/username',
    urlPattern: /^https?:\/\/twitch\.tv\/.+/,
    description: 'Twitchチャンネル',
  },
  {
    value: 'facebook',
    label: 'Facebook',
    icon: '👥',
    placeholder: 'https://facebook.com/username',
    urlPattern: /^https?:\/\/facebook\.com\/.+/,
    description: 'Facebookページ',
  },
];

// プラットフォーム別URL検証
export interface UrlValidationResult {
  isValid: boolean;
  suggestedUrl?: string;
  errorMessage?: string;
}

// プラットフォーム選択ステップ管理
export interface PlatformSelectionStep {
  selectedPlatform: SocialPlatform | null;
  step: 'select-platform' | 'enter-details' | 'confirm';
}

// プレビュー表示用の型
export interface SocialLinkPreviewProps {
  platform: SocialPlatform;
  url: string;
  title?: string;
  isPublic: boolean;
  onEdit: () => void;
}

// バリデーション関数の型
export interface SocialLinkValidation {
  validateUrl: (platform: SocialPlatform, url: string) => UrlValidationResult;
  checkDuplicatePlatform: (platform: SocialPlatform, existingPlatforms: SocialPlatform[]) => boolean;
  suggestTitle: (platform: SocialPlatform, url: string) => string;
}

// Feature層用の定数定義
export const FEATURE_SOCIAL_PLATFORMS: SocialPlatform[] = [
  'twitter',
  'instagram',
  'github',
  'linkedin',
  'website',
  'blog',
  'youtube',
  'facebook',
  'discord',
  'twitch',
];

export type { SocialPlatform };