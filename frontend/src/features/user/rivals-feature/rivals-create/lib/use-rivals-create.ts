/**
 * Rivals Create Hook
 * 
 * ライバル追加を提供します。
 * React 19のuseフックパターンを使用してサーバー操作を実行します。
 */

import { use, useMemo, useCallback } from 'react';
import { createRival, type CreateRivalDto } from '@/entities/user';

/**
 * ライバル追加フック
 * 
 * @param userId - ライバルを追加するユーザーID
 * @returns ライバル追加関数
 * 
 * @example
 * ```tsx
 * const { createRival } = useCreateRival(userId);
 * 
 * const handleAddRival = async (rivalUserId: string) => {
 *   try {
 *     await createRival({ rival_user_id: rivalUserId });
 *     // 成功時の処理
 *   } catch (error) {
 *     // エラー時の処理
 *   }
 * };
 * ```
 */
export const useCreateRival = (userId: string) => {
  const createRivalAction = useCallback(
    (data: CreateRivalDto) => createRival({ userId, data }),
    [userId]
  );
  
  return {
    createRival: createRivalAction,
  };
};