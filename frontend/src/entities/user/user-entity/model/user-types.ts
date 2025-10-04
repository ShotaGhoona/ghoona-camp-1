/**
 * User Types
 * 
 * ユーザー関連の型定義を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * user-entity/index.ts からエクスポートされた型を使用してください。
 */

// === Core User Type ===
export interface User {
  id: string;
  clerkId: string;
  email: string;
  username: string | null;
  avatarUrl: string | null;
  discordId: string | null;
  status: string;
  createdAt: string;
  updatedAt: string;
}

// === API Response Types ===

// === GET /auth/me レスポンス ===
export interface AuthMeResponse {
  data: User & {
    metadata: unknown; // metadata-entityで定義
    attendanceStats: unknown; // attendance-entityで定義
  };
  message: string;
  timestamp: string;
}

// === GET /users レスポンス ===
export interface UsersListResponse {
  data: {
    users: User[];
    total: number;
  };
  message: string;
  timestamp: string;
}

// === GET /users/{userId} レスポンス ===
export interface UserDetailResponse {
  data: User & {
    metadata: unknown; // metadata-entityで定義
  };
  message: string;
  timestamp: string;
}

// === POST /users レスポンス ===
export interface UserCreateResponse {
  data: User;
  message: string;
  timestamp: string;
}

// === PUT /users/{userId} レスポンス ===
export interface UserUpdateResponse {
  data: User;
  message: string;
  timestamp: string;
}

// === DTO Types（API送信用） ===
export interface CreateUserDto {
  clerkId: string;
  email: string;
  username?: string;
  avatarUrl?: string;
  discordId?: string;
}

export interface UpdateUserDto {
  username?: string;
  avatarUrl?: string;
  discordId?: string;
}

// === Query Parameters Types ===
export interface UsersQueryParams extends Record<string, string | number | boolean | undefined> {
  page?: number;
  limit?: number;
  search?: string;
  skills?: string;
  interests?: string;
  sortBy?: 'name' | 'attendanceDays' | 'streakDays' | 'createdAt';
  order?: 'asc' | 'desc';
}
