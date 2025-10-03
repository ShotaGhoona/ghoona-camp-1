/**
 * Rivals Get Hook
 * 
 * ライバル一覧取得を提供します。
 * React 19のuseフックパターンを使用してデータを取得します。
 */

import { use, useMemo } from 'react';
import { getUserRivals, mapUserRivalsListResponseToRivalsList } from '@/entities/user';

/**
 * ライバル一覧取得フック
 * 
 * @param userId - 取得対象のユーザーID
 * @returns ライバルデータ
 * 
 * @example
 * ```tsx
 * const { rivals, count, maxRivals } = useGetRivals(userId);
 * 
 * return (
 *   <div>
 *     <h2>ライバル ({count}/{maxRivals})</h2>
 *     {rivals.map(rival => (
 *       <RivalCard 
 *         key={rival.id} 
 *         rival={rival} 
 *         onRemove={() => handleRemove(rival.id)}
 *       />
 *     ))}
 *   </div>
 * );
 * ```
 */
export const useGetRivals = (userId: string) => {
  const rivalsPromise = useMemo(
    () => getUserRivals(userId).then(mapUserRivalsListResponseToRivalsList),
    [userId]
  );
  
  const rivalsData = use(rivalsPromise);
  
  return {
    rivals: rivalsData.rivals,
    count: rivalsData.count,
    maxRivals: rivalsData.maxRivals,
  };
};