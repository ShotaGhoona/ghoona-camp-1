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
  UserMetadataUpdateResponse,
  UpdateUserMetadataDto
} from '../model/metadata-types';

/** ユーザーメタデータ取得 */
/** GET /api/v1/users/{userId}/metadata */
export const getUserMetadata = async (userId: string): Promise<UserMetadataDetailResponse['data']> => {
  const { data } = await apiClient.get<UserMetadataDetailResponse>(`/users/${userId}/metadata`);
  return data.data;
};

/** ユーザーメタデータ更新 */
/** PUT /api/v1/users/{userId}/metadata */
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