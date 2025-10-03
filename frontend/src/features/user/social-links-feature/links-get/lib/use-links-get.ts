/**
 * Social Links Get Hook
 * 
 * SNSリンク一覧取得を提供します。
 * React 19のuseフックパターンを使用してデータを取得します。
 */

import { use, useMemo } from 'react';
import { getUserSocialLinks, mapUserSocialLinksListResponseToSocialLinks } from '@/entities/user';

/**
 * SNSリンク一覧取得フック
 * 
 * @param userId - 取得対象のユーザーID
 * @returns SNSリンクデータ
 * 
 * @example
 * ```tsx
 * const { socialLinks } = useGetSocialLinks(userId);
 * 
 * return (
 *   <div>
 *     {socialLinks.map(link => (
 *       <SocialLinkItem 
 *         key={link.id} 
 *         link={link}
 *         onEdit={() => handleEdit(link)}
 *         onDelete={() => handleDelete(link.id)}
 *       />
 *     ))}
 *   </div>
 * );
 * ```
 */
export const useGetSocialLinks = (userId: string) => {
  const socialLinksPromise = useMemo(
    () => getUserSocialLinks(userId).then(mapUserSocialLinksListResponseToSocialLinks),
    [userId]
  );
  
  const socialLinksData = use(socialLinksPromise);
  
  return {
    socialLinks: socialLinksData,
  };
};