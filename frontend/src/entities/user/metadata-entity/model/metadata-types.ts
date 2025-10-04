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
  displayName: string | null;
  profileImageUrl: string | null;
  tagline: string | null;
  bio: string | null;
  vision: string | null;
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

// === POST /users/{userId}/metadata レスポンス ===
export interface UserMetadataCreateResponse {
  data: UserMetadata;
  message: string;
  timestamp: string;
}

// === PUT /users/{userId}/metadata レスポンス ===
export interface UserMetadataUpdateResponse {
  data: UserMetadata;
  message: string;
  timestamp: string;
}

// === DTO Types（API送信用） ===
export interface CreateUserMetadataDto {
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