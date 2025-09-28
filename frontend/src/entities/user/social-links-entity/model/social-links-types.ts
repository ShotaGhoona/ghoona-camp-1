/**
 * User Social Links Types
 * 
 * ユーザーソーシャルリンク関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * social-links-entity/index.ts からエクスポートされた型を使用してください。
 */

// === API Response Types ===

// ソーシャルリンク（social-links-entity用）- 内部実装
interface SocialLinkResponse {
  id: string;
  platform: string;
  url: string;
  title: string;
  is_public: boolean;
  created_at: string;
  updated_at: string;
}

// === GET /users/{userId}/social-links レスポンス（統一形式） ===
export interface UserSocialLinksListResponse {
  data: {
    social_links: SocialLinkResponse[];
  };
  message: string;
  timestamp: string;
}

// === POST /users/{userId}/social-links レスポンス（統一形式） ===
export interface SocialLinkCreateResponse {
  data: {
    id: string;
    platform: string;
    url: string;
    title: string;
    is_public: boolean;
    created_at: string;
  };
  message: string;
  timestamp: string;
}

// === PUT /users/{userId}/social-links/{linkId} レスポンス（統一形式） ===
export interface SocialLinkUpdateResponse {
  data: {
    id: string;
    platform: string;
    url: string;
    title: string;
    is_public: boolean;
    updated_at: string;
  };
  message: string;
  timestamp: string;
}

// === Client Types（camelCase - フロントエンド用） ===
export interface SocialLink {
  id: string;
  platform: string;
  url: string;
  title: string;
  isPublic: boolean;
  createdAt: Date;
  updatedAt: Date;
}

// === DTO Types（API送信用） ===
export interface CreateSocialLinkDto {
  platform: string;
  url: string;
  title: string;
  is_public?: boolean;
}

export interface UpdateSocialLinkDto {
  url?: string;
  title?: string;
  is_public?: boolean;
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