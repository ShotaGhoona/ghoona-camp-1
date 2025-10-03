/**
 * Social Links Update Hook
 * 
 * SNSリンク更新を提供します。
 * React 19のuseフックパターンを使用してサーバー操作を実行します。
 */

import { useCallback } from 'react';
import { updateSocialLink, type UpdateSocialLinkDto } from '@/entities/user';

/**
 * SNSリンク更新フック
 * 
 * @param userId - SNSリンクを更新するユーザーID
 * @returns SNSリンク更新関数
 * 
 * @example
 * ```tsx
 * const { updateSocialLinks } = useUpdateSocialLinks(userId);
 * 
 * const handleEditSocialLink = async (linkId: string, updates: { url?: string; display_order?: number }) => {
 *   try {
 *     await updateSocialLinks(linkId, updates);
 *     // 成功時の処理
 *   } catch (error) {
 *     // エラー時の処理
 *   }
 * };
 * ```
 */
export const useUpdateSocialLinks = (userId: string) => {
  const updateSocialLinks = useCallback(
    (linkId: string, data: UpdateSocialLinkDto) => 
      updateSocialLink({ userId, linkId, data }),
    [userId]
  );
  
  return {
    updateSocialLinks,
  };
};