/**
 * Rivals Delete Hook
 * 
 * ライバル削除を提供します。
 * React 19のuseフックパターンを使用してサーバー操作を実行します。
 */

import { useCallback } from 'react';
import { deleteRival } from '@/entities/user';

/**
 * ライバル削除フック
 * 
 * @param userId - ライバルを削除するユーザーID
 * @returns ライバル削除関数
 * 
 * @example
 * ```tsx
 * const { deleteRival } = useDeleteRival(userId);
 * 
 * const handleRemoveRival = async (rivalId: string) => {
 *   try {
 *     await deleteRival(rivalId);
 *     // 成功時の処理
 *   } catch (error) {
 *     // エラー時の処理
 *   }
 * };
 * ```
 */
export const useDeleteRival = (userId: string) => {
  const deleteRivalAction = useCallback(
    (rivalId: string) => deleteRival({ userId, rivalId }),
    [userId]
  );
  
  return {
    deleteRival: deleteRivalAction,
  };
};