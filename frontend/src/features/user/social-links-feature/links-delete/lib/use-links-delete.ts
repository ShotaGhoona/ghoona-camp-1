/**
 * Social Links Delete Hook
 * 
 * SNSリンク削除を提供します。
 * React 19のuseフックパターンを使用してサーバー操作を実行します。
 */

import { useCallback } from 'react';
import { deleteSocialLink } from '@/entities/user';

/**
 * SNSリンク削除フック
 * 
 * @param userId - SNSリンクを削除するユーザーID
 * @returns SNSリンク削除関数
 * 
 * @example
 * ```tsx
 * const { deleteSocialLinks } = useDeleteSocialLinks(userId);
 * 
 * const handleRemoveSocialLink = async (linkId: string) => {
 *   try {
 *     await deleteSocialLinks(linkId);
 *     // 成功時の処理
 *   } catch (error) {
 *     // エラー時の処理
 *   }
 * };
 * ```
 */
export const useDeleteSocialLinks = (userId: string) => {
  const deleteSocialLinks = useCallback(
    (linkId: string) => deleteUserSocialLinks({ userId, linkId }),
    [userId]
  );
  
  return {
    deleteSocialLinks,
  };
};