/**
 * Profile Update Hook
 * 
 * ユーザー基本情報更新を提供します。
 * React Queryのmutationを使用してデータを更新し、適切なキャッシュ無効化を行います。
 */

import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateUser, userKeys, type UpdateUserDto } from '@/entities/user';

/**
 * プロフィール更新フック
 * 
 * @param userId - 更新対象のユーザーID
 * @returns mutation関数とステータス
 * 
 * @example
 * ```tsx
 * const updateProfile = useUpdateProfile(userId);
 * 
 * const handleUpdate = (data: UpdateUserDto) => {
 *   updateProfile.mutate(data, {
 *     onSuccess: () => {
 *       toast({ title: "プロフィールを更新しました" });
 *     },
 *     onError: (error) => {
 *       toast({ title: "更新に失敗しました", variant: "destructive" });
 *     },
 *   });
 * };
 * ```
 */
export const useUpdateProfile = (userId: string) => {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (data: UpdateUserDto) => 
      updateUser({ id: userId, data }),
    onSuccess: () => {
      // ユーザー詳細情報のキャッシュを無効化
      queryClient.invalidateQueries({
        queryKey: userKeys.detail(userId),
      });
      
      // セッション情報のキャッシュを無効化（自分のプロフィール更新の場合）
      queryClient.invalidateQueries({
        queryKey: userKeys.session(),
      });
      
      // ユーザー一覧のキャッシュも無効化（一覧に表示される情報が変更される可能性）
      queryClient.invalidateQueries({
        queryKey: userKeys.list(),
      });
    },
  });
};