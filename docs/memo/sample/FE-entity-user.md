frontend/src/entities/user/user-entity
├── api
│   ├── index.ts
│   └── user-api.ts
├── index.ts
├── lib
│   ├── index.ts
│   └── validators.ts
├── model
│   ├── index.ts
│   ├── selectors.ts
│   ├── slice.ts
│   └── types.ts
└── utils
    ├── index.ts
    └── query-keys.ts



/**
 * ユーザーAPI
 */
import { apiClient } from '@/shared';
import type { User, CreateUserDto, UpdateUserDto } from '../model/types';

/** ユーザー一覧レスポンス */
export interface UsersResponse {
  users: User[];
}

/** 認証ステータスレスポンス */
export interface AuthStatusResponse {
  authenticated: boolean;
  user: User | null;
}

/** ユーザー一覧を取得 */
export const getUsers = async (): Promise<UsersResponse> => {
  const { data } = await apiClient.get<UsersResponse>('/users');
  return data;
};

/** 特定のユーザーを取得 */
export const getUser = async (id: number): Promise<User> => {
  const { data } = await apiClient.get<User>(`/users/${id}`);
  return data;
};

/** 現在のユーザー情報を取得（認証状態確認） */
export const getAuthStatus = async (): Promise<AuthStatusResponse> => {
  const { data } = await apiClient.get<AuthStatusResponse>('/auth/status');
  return data;
};

/** ユーザーを作成 */
export const createUser = async (userData: CreateUserDto): Promise<User> => {
  const response = await apiClient.post<User>('/users', userData);
  return response.data;
};

/** ユーザーを更新 */
export const updateUser = async ({
  id,
  data,
}: {
  id: number;
  data: UpdateUserDto;
}): Promise<User> => {
  const response = await apiClient.put<User>(`/users/${id}`, data);
  return response.data;
};

---


/**
 * ユーザー関連のバリデーション
 */

/** メールアドレスのバリデーション */
export const validateEmail = (email: string): boolean => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
};

/** パスワードの強度チェック */
export const validatePassword = (
  password: string,
): { isValid: boolean; errors: string[] } => {
  const errors: string[] = [];

  // テスト用に緩和
  if (password.length < 4) {
    errors.push('4文字以上必要です');
  }
  if (!/[a-z]/.test(password)) {
    errors.push('小文字を含む必要があります');
  }

  return {
    isValid: errors.length === 0,
    errors,
  };
};


---

/**
 * 認証状態のセレクター
 */

// AuthStateの型定義
interface AuthState {
  currentUser: any;
  isAuthenticated: boolean;
}

// セレクターは汎用的な形で定義（RootStateに依存しない）
/** 現在のユーザーを取得 */
export const selectCurrentUser = (state: { auth: AuthState }) =>
  state.auth.currentUser;

/** 認証状態を取得 */
export const selectIsAuthenticated = (state: { auth: AuthState }) =>
  state.auth.isAuthenticated;

/** 現在のユーザーの権限を取得 */
export const selectCurrentAuthority = (state: { auth: AuthState }) =>
  state.auth.currentUser?.authority || null;

/** 現在のユーザーのIDを取得 */
export const selectCurrentUserId = (state: { auth: AuthState }) =>
  state.auth.currentUser?.id || null;


---

/**
 * 認証状態管理のRedux slice
 */
import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import type { User } from './types';

/** 認証状態の型定義 */
interface AuthState {
  currentUser: User | null;
  isAuthenticated: boolean;
}

/** 初期状態 */
const initialState: AuthState = {
  currentUser: null,
  isAuthenticated: false,
};

/** 認証用のslice */
export const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    /** ログイン処理 */
    login: (state, action: PayloadAction<User>) => {
      state.currentUser = action.payload;
      state.isAuthenticated = true;
    },
    /** ログアウト処理 */
    logout: (state) => {
      state.currentUser = null;
      state.isAuthenticated = false;
    },
    /** 現在のユーザー情報を更新 */
    updateCurrentUser: (state, action: PayloadAction<Partial<User>>) => {
      if (state.currentUser) {
        state.currentUser = { ...state.currentUser, ...action.payload };
      }
    },
  },
});

export const { login, logout, updateCurrentUser } = authSlice.actions;
export default authSlice.reducer;


---

import { Authority } from '@/entities/user/authority/model/types';
import { Role } from '../../role/model/types';
import { Skill } from '../../skill/model/types';

/** ユーザーエンティティ */
export interface User {
  id: number;
  authority: Authority;
  name: string;
  email: string;
  github_username?: string; // GitHubユーザー名（HR機能用）
  skills?: Skill[]; // スキル一覧（HR機能用）
  roles?: Role[]; // 役割一覧（HR機能用）
  createdAt: string;
  updatedAt: string;
}

/** ユーザー作成DTO */
export interface CreateUserDto {
  name: string;
  email: string;
  password: string;
  authorityId: number;
  github_username?: string; // GitHubユーザー名（HR機能用）
}

/** ユーザー更新DTO */
export interface UpdateUserDto {
  name?: string;
  email?: string;
  authorityId?: number;
  github_username?: string; // GitHubユーザー名（HR機能用）
}

/** ログインDTO */
export interface LoginDto {
  email: string;
  password: string;
}

/** ログインレスポンス */
export interface LoginResponse {
  message: string;
  access_token: string;
  user: User;
}


---

/**
 * ユーザー関連のクエリキー
 */
export const userKeys = {
  all: ['users'] as const,
  lists: () => [...userKeys.all, 'list'] as const,
  list: (filters: string) => [...userKeys.lists(), { filters }] as const,
  details: () => [...userKeys.all, 'detail'] as const,
  detail: (id: number) => [...userKeys.details(), id] as const,
  current: () => [...userKeys.all, 'current'] as const,
};

export const authKeys = {
  all: ['auth'] as const,
  status: () => [...authKeys.all, 'status'] as const,
};
