/**
 * User Metadata Types
 * 
 * ユーザーメタデータ関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * metadata-entity/index.ts からエクスポートされた型を使用してください。
 */

// === Core UserMetadata Type ===
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
  createdAt: string;
  updatedAt: string;
}

// === API Response Types ===

// === GET /users/{userId}/metadata レスポンス ===
export interface UserMetadataDetailResponse {
  data: UserMetadata;
  message: string;
  timestamp: string;
}

// === PUT /users/{userId}/metadata レスポンス ===
export interface UserMetadataUpdateResponse {
  data: {
    id: string;
    displayName: string;
    tagline: string;
    visionPublic: boolean;
    skills: string[];
    interests: string[];
    updatedAt: string;
  };
  message: string;
  timestamp: string;
}

// === DTO Types（API送信用） ===
export interface UpdateUserMetadataDto {
  displayName?: string;
  profileImageUrl?: string;
  tagline?: string;
  bio?: string;
  vision?: string;
  visionPublic?: boolean;
  timezone?: string;
  skills?: string[];
  interests?: string[];
}