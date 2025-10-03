import { useMutation, useQueryClient } from '@tanstack/react-query';
import { deleteRival, rivalsKeys } from '@/entities/user/rivals-entity';
import { userKeys } from '@/entities/user/user-entity';

export const useRivalsDelete = (userId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (rivalId: string) => 
      deleteRival({ userId, rivalId }),
    onSuccess: () => {
      // ライバル一覧とユーザー詳細を無効化
      queryClient.invalidateQueries({ queryKey: rivalsKeys.list(userId) });
      queryClient.invalidateQueries({ queryKey: userKeys.detail(userId) });
      // ユーザー一覧も更新（ライバル情報が表示される場合）
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
    },
    onError: (error) => {
      console.error('Rival deletion failed:', error);
    },
  });
};