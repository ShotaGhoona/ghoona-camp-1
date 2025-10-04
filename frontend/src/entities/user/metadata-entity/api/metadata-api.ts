/**
 * User Metadata API Functions
 * 
 * ユーザーメタデータ関連のAPI呼び出しを提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * metadata-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { apiClient } from '@/shared/api';
import type { 
  UserMetadataDetailResponse,
  UserMetadataCreateResponse,
  UserMetadataUpdateResponse,
  CreateUserMetadataDto,
  UpdateUserMetadataDto
} from '../model/metadata-types';

/** ユーザーメタデータ取得 */
/** GET /users/{userId}/metadata */
export const getUserMetadata = async (userId: string): Promise<UserMetadataDetailResponse['data']> => {
  const { data } = await apiClient.get<UserMetadataDetailResponse>(`/users/${userId}/metadata`);
  return data.data;
};

/** ユーザーメタデータ作成 */
/** POST /users/{userId}/metadata */
export const createUserMetadata = async ({
  userId,
  data: metadataData,
}: {
  userId: string;
  data: CreateUserMetadataDto;
}): Promise<UserMetadataCreateResponse['data']> => {
  const { data } = await apiClient.post<UserMetadataCreateResponse>(`/users/${userId}/metadata`, metadataData);
  return data.data;
};

/** ユーザーメタデータ更新 */
/** PUT /users/{userId}/metadata */
export const updateUserMetadata = async ({
  userId,
  data: updateData,
}: {
  userId: string;
  data: UpdateUserMetadataDto;
}): Promise<UserMetadataUpdateResponse['data']> => {
  const { data } = await apiClient.put<UserMetadataUpdateResponse>(`/users/${userId}/metadata`, updateData);
  return data.data;
};