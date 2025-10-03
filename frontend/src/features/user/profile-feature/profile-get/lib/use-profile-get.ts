/**
 * Profile Get Hook
 * 
 * ユーザー基本情報取得を提供します。
 * 自分・他ユーザー両方のプロフィール取得に対応。
 * React 19のuseフックパターンを使用してデータを取得します。
 */

import { use, useMemo } from 'react';
import { getUserDetail, mapUserDetailResponseToUser } from '@/entities/user';

/**
 * プロフィール取得フック
 * 
 * @param userId - 取得対象のユーザーID
 * @returns プロフィールデータ
 * 
 * @example
 * ```tsx
 * const { profile } = useGetProfile(userId);
 * 
 * return (
 *   <div>
 *     <img src={profile.avatarUrl} alt={profile.displayName} />
 *     <h1>{profile.displayName}</h1>
 *     <p>@{profile.username}</p>
 *   </div>
 * );
 * ```
 */
export const useGetProfile = (userId: string) => {
  const profilePromise = useMemo(
    () => getUserDetail(userId).then(mapUserDetailResponseToUser),
    [userId]
  );
  
  const profileData = use(profilePromise);
  
  return {
    profile: profileData,
  };
};