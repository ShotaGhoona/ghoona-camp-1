/**
 * Metadata Update Hook
 * 
 * ユーザーメタデータ更新を提供します。
 * React Queryのmutationを使用してデータを更新し、適切なキャッシュ無効化を行います。
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateUserMetadata, metadataKeys, userKeys, type UpdateUserMetadataDto } from '@/entities/user';

/**
 * メタデータ更新フック
 * 
 * @param userId - 更新対象のユーザーID
 * @returns mutation関数とステータス
 * 
 * @example
 * ```tsx
 * const updateMetadata = useUpdateMetadata(userId);
 * 
 * const handleUpdate = (data: UpdateUserMetadataDto) => {
 *   updateMetadata.mutate(data, {
 *     onSuccess: () => {
 *       toast({ title: "メタデータを更新しました" });
 *     },
 *     onError: (error) => {
 *       toast({ title: "更新に失敗しました", variant: "destructive" });
 *     },
 *   });
 * };
 * ```
 */
export const useUpdateMetadata = (userId: string) => {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (data: UpdateUserMetadataDto) => 
      updateUserMetadata({ userId, data }),
    onSuccess: () => {
      // メタデータのキャッシュを無効化
      queryClient.invalidateQueries({
        queryKey: metadataKeys.detail(userId),
      });
      
      // ユーザー詳細情報のキャッシュも無効化（メタデータが含まれる場合）
      queryClient.invalidateQueries({
        queryKey: userKeys.detail(userId),
      });
      
      // セッション情報のキャッシュも無効化（自分のメタデータ更新の場合）
      queryClient.invalidateQueries({
        queryKey: userKeys.session(),
      });
    },
  });
};