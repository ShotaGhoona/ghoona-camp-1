import { apiClient } from '@/shared/api';
import type { 
  AuthMeResponse,
  UsersListResponse,
  UserDetailResponse,
  UserUpdateResponse,
  UpdateUserDto,
  UsersQueryParams
} from '../model/types';

/** 現在のユーザー情報を取得（認証状態確認） */
export const getCurrentUser = async (): Promise<AuthMeResponse['data']> => {
  const { data } = await apiClient.get<AuthMeResponse>('/auth/me');
  return data.data;
};

/** ユーザー一覧を取得 */
export const getUsers = async (params?: UsersQueryParams): Promise<UsersListResponse['data']> => {
  // undefinedを除去してAPIクライアントに渡す
  const cleanParams = params ? Object.fromEntries(
    Object.entries(params).filter(([_, value]) => value !== undefined)
  ) as Record<string, string | number | boolean> : undefined;
  
  const { data } = await apiClient.get<UsersListResponse>('/users', { params: cleanParams });
  return data.data;
};

/** 特定ユーザーの詳細情報を取得 */
export const getUserDetail = async (id: string): Promise<UserDetailResponse['data']> => {
  const { data } = await apiClient.get<UserDetailResponse>(`/users/${id}`);
  return data.data;
};

/** ユーザーの基本情報を更新 */
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