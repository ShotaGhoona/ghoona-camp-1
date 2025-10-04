/**
 * User Social Links Types
 * 
 * ユーザーソーシャルリンク関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * social-links-entity/index.ts からエクスポートされた型を使用してください。
 */

// === Core SocialLink Type ===
export interface SocialLink {
  id: string;
  userId: string;
  platform: string;
  url: string;
  title: string | null;
  isPublic: boolean;
  createdAt: string;
  updatedAt: string;
}

// === API Response Types ===

// === GET /users/{userId}/social-links レスポンス ===
export interface UserSocialLinksListResponse {
  data: {
    socialLinks: SocialLink[];
    total: number;
  };
  message: string;
  timestamp: string;
}

// === POST /users/{userId}/social-links レスポンス ===
export interface SocialLinkCreateResponse {
  data: SocialLink;
  message: string;
  timestamp: string;
}

// === PUT /users/{userId}/social-links/{linkId} レスポンス ===
export interface SocialLinkUpdateResponse {
  data: SocialLink;
  message: string;
  timestamp: string;
}

// === DTO Types（API送信用） ===
export interface CreateSocialLinkDto {
  platform: string;
  url: string;
  title?: string;
  isPublic?: boolean;
}

export interface UpdateSocialLinkDto {
  url?: string;
  title?: string;
  isPublic?: boolean;
}

// === Platform Types ===
export type SocialPlatform = 
  | 'twitter'
  | 'instagram' 
  | 'github'
  | 'linkedin'
  | 'website'
  | 'blog'
  | 'youtube'
  | 'facebook'
  | 'discord'
  | 'twitch';

export const SOCIAL_PLATFORMS: SocialPlatform[] = [
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
] as const;