/**
 * User API Functions
 * 
 * ユーザー関連のAPI呼び出しを提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * user-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { apiClient } from '@/shared/api';
import type { 
  AuthMeResponse,
  UsersListResponse,
  UserDetailResponse,
  UserUpdateResponse,
  UpdateUserDto,
  UsersQueryParams
} from '../model/user-types';

/** 現在のユーザー情報を取得（認証状態確認） */
/** GET /auth/me */
export const getCurrentUser = async (): Promise<AuthMeResponse['data']> => {
  const { data } = await apiClient.get<AuthMeResponse>('/auth/me');
  return data.data;
};

/** ユーザー一覧を取得 */
/** GET /users */ 
export const getUsers = async (params?: UsersQueryParams): Promise<UsersListResponse['data']> => {
  const { data } = await apiClient.get<UsersListResponse>('/users', { params });
  return data.data;
};

/** 特定ユーザーの詳細情報を取得 */
/** GET /users/{userId} */
export const getUserDetail = async (id: string): Promise<UserDetailResponse['data']> => {
  const { data } = await apiClient.get<UserDetailResponse>(`/users/${id}`);
  return data.data;
};

/** ユーザーの基本情報を更新 */
/** PUT /users/{userId} */
export const updateUser = async ({
  id,
  data: updateData,
}: {
  id: string;
  data: UpdateUserDto;
}): Promise<UserUpdateResponse['data']> => {
  const { data } = await apiClient.put<UserUpdateResponse>(`/users/${id}`, updateData);
  return data.data;
};