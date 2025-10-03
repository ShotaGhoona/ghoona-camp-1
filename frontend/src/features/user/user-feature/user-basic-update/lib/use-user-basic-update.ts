import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateUser, userKeys, type UpdateUserDto } from '@/entities/user/user-entity';

export const useUserBasicUpdate = (userId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateUserDto) => updateUser({ id: userId, data }),
    onSuccess: () => {
      // ユーザー詳細とセッション情報を無効化
      queryClient.invalidateQueries({ queryKey: userKeys.detail(userId) });
      queryClient.invalidateQueries({ queryKey: userKeys.session() });
      // ユーザー一覧も更新
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
    },
    onError: (error) => {
      console.error('User basic update failed:', error);
    },
  });
};