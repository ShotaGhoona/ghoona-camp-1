/**
 * Metadata Get Hook
 * 
 * ユーザーメタデータ取得を提供します。
 * React 19のuseフックパターンを使用してデータを取得します。
 */

import { use, useMemo } from 'react';
import { getUserMetadata, mapUserMetadataDetailResponseToUserMetadata } from '@/entities/user';

/**
 * メタデータ取得フック
 * 
 * @param userId - 取得対象のユーザーID
 * @returns メタデータ
 * 
 * @example
 * ```tsx
 * const { metadata } = useGetMetadata(userId);
 * 
 * return (
 *   <div>
 *     <h1>{metadata.displayName}</h1>
 *     <p>{metadata.tagline}</p>
 *     <p>{metadata.bio}</p>
 *     <div>スキル: {metadata.skills.join(', ')}</div>
 *   </div>
 * );
 * ```
 */
export const useGetMetadata = (userId: string) => {
  const metadataPromise = useMemo(
    () => getUserMetadata(userId).then(mapUserMetadataDetailResponseToUserMetadata),
    [userId]
  );
  
  const metadataData = use(metadataPromise);
  
  return {
    metadata: metadataData,
  };
};