/**
 * User Entity Mappers
 * 
 * APIレスポンス型からフロントエンド用型への変換関数を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * user-entity/index.ts からエクスポートされた関数を使用してください。
 */

import type { User, AuthMeResponse, UserDetailResponse, UsersListResponse, UserUpdateResponse } from '../model/user-types';

/**
 * AuthMeResponse のデータをUser型に変換
 */
export const mapAuthMeResponseToUser = (
  authData: AuthMeResponse['data']
): User => {
  return {
    id: authData.id,
    clerkId: authData.clerk_id,
    email: authData.email,
    username: authData.username,
    avatarUrl: authData.avatar_url,
    discordId: authData.discord_id,
    isActive: authData.is_active,
    createdAt: new Date(authData.created_at),
    updatedAt: new Date(authData.updated_at),
  };
};

/**
 * UserDetailResponse のデータをUser型に変換
 */
export const mapUserDetailResponseToUser = (
  userData: UserDetailResponse['data']
): User => {
  return {
    id: userData.id,
    clerkId: '', // UserDetailResponseには含まれていない
    email: '', // UserDetailResponseには含まれていない
    username: userData.username,
    avatarUrl: userData.avatar_url,
    discordId: undefined, // UserDetailResponseには含まれていない
    isActive: true, // UserDetailResponseには含まれていない（公開されているため有効と仮定）
    createdAt: new Date(userData.created_at),
    updatedAt: new Date(), // UserDetailResponseには含まれていない
  };
};

/**
 * UsersListResponse のデータをUser配列に変換（簡易版）
 */
export const mapUsersListResponseToUsers = (
  usersData: UsersListResponse['data']
): Array<Partial<User>> => {
  return usersData.users.map((userData) => ({
    id: userData.id,
    username: userData.username,
    avatarUrl: userData.avatar_url,
    createdAt: new Date(userData.created_at),
    // Note: 一覧取得では限定された情報のみ
  }));
};

/**
 * UserUpdateResponse のデータをUser型（部分的）に変換
 */
export const mapUserUpdateResponseToUser = (
  updateData: UserUpdateResponse['data']
): Partial<User> => {
  return {
    id: updateData.id,
    username: updateData.username,
    avatarUrl: updateData.avatar_url,
    updatedAt: new Date(updateData.updated_at),
  };
};