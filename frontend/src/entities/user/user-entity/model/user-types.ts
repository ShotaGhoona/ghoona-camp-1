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
  username: string;
  avatarUrl: string;
  discordId?: string;
  isActive: boolean;
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
    users: Array<{
      id: string;
      displayName: string;
      username: string;
      avatarUrl: string;
      tagline: string;
      skills: string[];
      interests: string[];
      currentTitle: {
        level: number;
        nameJp: string;
        nameEn: string;
        colorTheme: string;
      };
      attendanceStats: {
        totalAttendanceDays: number;
        currentStreakDays: number;
      };
      createdAt: string;
    }>;
    pagination: {
      currentPage: number;
      totalPages: number;
      totalCount: number;
      limit: number;
      hasNext: boolean;
      hasPrev: boolean;
    };
  };
  message: string;
  timestamp: string;
}

// === GET /users/{userId} レスポンス ===
export interface UserDetailResponse {
  data: {
    id: string;
    displayName: string;
    username: string;
    avatarUrl: string;
    profileImageUrl: string;
    tagline: string;
    bio: string;
    vision: string;
    timezone: string;
    skills: string[];
    interests: string[];
    socialLinks: Array<{
      id: string;
      platform: string;
      url: string;
      title: string;
      isPublic: boolean;
    }>;
    currentTitle: {
      id: string;
      level: number;
      nameJp: string;
      nameEn: string;
      description: string;
      colorTheme: string;
      achievedAt: string;
    };
    attendanceStats: {
      totalAttendanceDays: number;
      currentStreakDays: number;
      maxStreakDays: number;
      firstAttendanceDate: string;
      lastAttendanceDate: string;
      totalDurationMinutes: number;
    };
    achievements: Array<{
      titleId: string;
      level: number;
      nameJp: string;
      achievedAt: string;
    }>;
    isRival: boolean;
    createdAt: string;
  };
  message: string;
  timestamp: string;
}

// === PUT /users/{userId} レスポンス ===
export interface UserUpdateResponse {
  data: {
    id: string;
    username: string;
    avatarUrl: string;
    updatedAt: string;
  };
  message: string;
  timestamp: string;
}

// === DTO Types（API送信用） ===
export interface UpdateUserDto {
  username?: string;
  avatarUrl?: string;
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
