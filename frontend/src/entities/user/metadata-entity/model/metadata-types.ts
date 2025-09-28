/**
 * User Metadata Types
 * 
 * ユーザーメタデータ関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * metadata-entity/index.ts からエクスポートされた型を使用してください。
 */

// === API Response Types ===

// ユーザーメタデータ（metadata-entity用）- 内部実装
interface UserMetadataResponse {
  id: string;
  user_id: string;
  display_name: string;
  profile_image_url: string;
  tagline: string;
  bio: string;
  vision: string;
  vision_public: boolean;
  timezone: string;
  skills: string[];
  interests: string[];
  created_at: string;
  updated_at: string;
}

// === GET /users/{userId}/metadata レスポンス（統一形式） ===
export interface UserMetadataDetailResponse {
  data: UserMetadataResponse;
  message: string;
  timestamp: string;
}

// === PUT /users/{userId}/metadata レスポンス（統一形式） ===
export interface UserMetadataUpdateResponse {
  data: {
    id: string;
    display_name: string;
    tagline: string;
    vision_public: boolean;
    skills: string[];
    interests: string[];
    updated_at: string;
  };
  message: string;
  timestamp: string;
}

// === Client Types（camelCase - フロントエンド用） ===
export interface UserMetadata {
  id: string;
  userId: string;
  displayName: string;
  profileImageUrl: string;
  tagline: string;
  bio: string;
  vision: string;
  visionPublic: boolean;
  timezone: string;
  skills: string[];
  interests: string[];
  createdAt: Date;
  updatedAt: Date;
}

// === DTO Types（API送信用） ===
export interface UpdateUserMetadataDto {
  display_name?: string;
  profile_image_url?: string;
  tagline?: string;
  bio?: string;
  vision?: string;
  vision_public?: boolean;
  timezone?: string;
  skills?: string[];
  interests?: string[];
}