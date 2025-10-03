/**
 * User Rivals API Functions
 * 
 * ユーザーライバル関連のAPI呼び出しを提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * rivals-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { apiClient } from '@/shared/api';
import type { 
  UserRivalsListResponse,
  RivalCreateResponse,
  CreateRivalDto
} from '../model/rivals-types';

/** ユーザーのライバル一覧取得 */
/** GET /users/{userId}/rivals */
export const getUserRivals = async (userId: string): Promise<UserRivalsListResponse['data']> => {
  const { data } = await apiClient.get<UserRivalsListResponse>(`/users/${userId}/rivals`);
  return data.data;
};

/** 新しいライバルを追加 */
/** POST /users/{userId}/rivals */
export const createRival = async ({
  userId,
  data: rivalData,
}: {
  userId: string;
  data: CreateRivalDto;
}): Promise<RivalCreateResponse['data']> => {
  const { data } = await apiClient.post<RivalCreateResponse>(`/users/${userId}/rivals`, rivalData);
  return data.data;
};

/** ライバル関係を解除 */
/** DELETE /users/{userId}/rivals/{rivalId} */
export const deleteRival = async ({
  userId,
  rivalId,
}: {
  userId: string;
  rivalId: string;
}): Promise<void> => {
  await apiClient.delete(`/users/${userId}/rivals/${rivalId}`);
};