/**
 * Achievements API Functions
 * 
 * ユーザー称号実績関連のAPI呼び出しを提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * achievements-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { apiClient } from '@/shared/api';
import type { 
  UserAchievementsResponse,
  SetCurrentTitleResponse,
  SetCurrentTitleDto,
  AchievementsQueryParams
} from '../model/achievements-types';

/** ユーザーの称号取得履歴・現在設定中の称号を取得 */
/** GET /api/v1/users/{userId}/achievements */
export const getUserAchievements = async (
  userId: string, 
  params?: AchievementsQueryParams
): Promise<UserAchievementsResponse['data']> => {
  const { data } = await apiClient.get<UserAchievementsResponse>(`/users/${userId}/achievements`, { params });
  return data.data;
};

/** 現在表示する称号を変更 */
/** PUT /api/v1/users/{userId}/achievements/{titleId} */
export const setCurrentTitle = async ({
  userId,
  titleId,
}: {
  userId: string;
  titleId: string;
}): Promise<SetCurrentTitleResponse['data']> => {
  const requestData: SetCurrentTitleDto = { is_current: true };
  const { data } = await apiClient.put<SetCurrentTitleResponse>(`/users/${userId}/achievements/${titleId}`, requestData);
  return data.data;
};