// === API Response Types ===

// 基本ユーザー情報（user-entity用）
export interface UserResponse {
  id: string;
  clerk_id: string;
  email: string;
  username: string;
  avatar_url: string;
  discord_id?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

// === GET /auth/me レスポンス（統一形式） ===
export interface AuthMeResponse {
  data: UserResponse & {
    metadata: unknown; // metadata-entityで定義
    attendance_stats: unknown; // attendance-entityで定義
  };
  message: string;
  timestamp: string;
}

// === GET /users レスポンス（統一形式） ===
export interface UsersListResponse {
  data: {
    users: Array<{
      id: string;
      display_name: string;
      username: string;
      avatar_url: string;
      tagline: string;
      skills: string[];
      interests: string[];
      current_title: {
        level: number;
        name_jp: string;
        name_en: string;
        color_theme: string;
      };
      attendance_stats: {
        total_attendance_days: number;
        current_streak_days: number;
      };
      created_at: string;
    }>;
    pagination: {
      current_page: number;
      total_pages: number;
      total_count: number;
      limit: number;
      has_next: boolean;
      has_prev: boolean;
    };
  };
  message: string;
  timestamp: string;
}

// === GET /users/{userId} レスポンス（dataのみ） ===
export interface UserDetailResponse {
  data: {
    id: string;
    display_name: string;
    username: string;
    avatar_url: string;
    profile_image_url: string;
    tagline: string;
    bio: string;
    vision: string;
    timezone: string;
    skills: string[];
    interests: string[];
    social_links: Array<{
      id: string;
      platform: string;
      url: string;
      title: string;
      is_public: boolean;
    }>;
    current_title: {
      id: string;
      level: number;
      name_jp: string;
      name_en: string;
      description: string;
      color_theme: string;
      achieved_at: string;
    };
    attendance_stats: {
      total_attendance_days: number;
      current_streak_days: number;
      max_streak_days: number;
      first_attendance_date: string;
      last_attendance_date: string;
      total_duration_minutes: number;
    };
    achievements: Array<{
      title_id: string;
      level: number;
      name_jp: string;
      achieved_at: string;
    }>;
    is_rival: boolean;
    created_at: string;
  };
}

// === PUT /users/{userId} レスポンス（統一形式） ===
export interface UserUpdateResponse {
  data: {
    id: string;
    username: string;
    avatar_url: string;
    updated_at: string;
  };
  message: string;
  timestamp: string;
}

// === Client Types（camelCase - フロントエンド用） ===
export interface User {
  id: string;
  clerkId: string;
  email: string;
  username: string;
  avatarUrl: string;
  discordId?: string;
  isActive: boolean;
  createdAt: Date;
  updatedAt: Date;
}

// === DTO Types（API送信用） ===
export interface UpdateUserDto {
  username?: string;
  avatar_url?: string;
}

// === Query Parameters Types ===
export interface UsersQueryParams extends Record<string, string | number | boolean | undefined> {
  page?: number;
  limit?: number;
  search?: string;
  skills?: string;
  interests?: string;
  sort_by?: 'name' | 'attendance_days' | 'streak_days' | 'created_at';
  order?: 'asc' | 'desc';
}
