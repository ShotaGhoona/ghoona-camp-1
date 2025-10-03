/**
 * User Detail Hook
 * 
 * 他ユーザーの詳細情報取得を提供します。
 * React 19のuseフックパターンを使用してデータを取得します。
 */

import { use, useMemo } from 'react';
import { getUserDetail, mapUserDetailResponseToUser } from '@/entities/user';

/**
 * ユーザー詳細取得フック
 * 
 * @param userId - 取得対象のユーザーID
 * @returns ユーザー詳細データ
 * 
 * @example
 * ```tsx
 * const { userDetail } = useGetUserDetail(userId);
 * 
 * return (
 *   <div>
 *     <img src={userDetail.avatarUrl} alt={userDetail.displayName} />
 *     <h1>{userDetail.displayName}</h1>
 *     <p>@{userDetail.username}</p>
 *     <p>{userDetail.bio}</p>
 *     <div>
 *       {userDetail.socialLinks?.map(link => (
 *         <a key={link.id} href={link.url}>{link.platform}</a>
 *       ))}
 *     </div>
 *   </div>
 * );
 * ```
 */
export const useGetUserDetail = (userId: string) => {
  const userDetailPromise = useMemo(
    () => getUserDetail(userId),
    [userId]
  );
  
  const userDetailData = use(userDetailPromise);
  
  return {
    userDetail: userDetailData, // 詳細情報なので変換せずそのまま
  };
};