/**
 * Social Links Create Hook
 * 
 * SNSリンク作成を提供します。
 * React 19のuseフックパターンを使用してサーバー操作を実行します。
 */

import { useCallback } from 'react';
import { createSocialLink, type CreateSocialLinkDto } from '@/entities/user';

/**
 * SNSリンク作成フック
 * 
 * @param userId - SNSリンクを作成するユーザーID
 * @returns SNSリンク作成関数
 * 
 * @example
 * ```tsx
 * const { createSocialLinks } = useCreateSocialLinks(userId);
 * 
 * const handleAddSocialLink = async (platform: string, url: string) => {
 *   try {
 *     await createSocialLinks({
 *       platform,
 *       url,
 *       display_order: 1
 *     });
 *     // 成功時の処理
 *   } catch (error) {
 *     // エラー時の処理
 *   }
 * };
 * ```
 */
export const useCreateSocialLinks = (userId: string) => {
  const createSocialLinkAction = useCallback(
    (data: CreateSocialLinkDto) => createSocialLink({ userId, data }),
    [userId]
  );
  
  return {
    createSocialLinks: createSocialLinkAction,
  };
};