import { useMutation, useQueryClient } from '@tanstack/react-query';
import { createUser, userKeys, type CreateUserDto } from '@/entities/user/user-entity';

export const useUserCreate = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateUserDto) => createUser(data),
    onSuccess: () => {
      // ユーザー一覧を無効化（新しいユーザーが追加されたため）
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
      // セッション情報も無効化（作成したユーザーが現在のユーザーの場合）
      queryClient.invalidateQueries({ queryKey: userKeys.session() });
    },
    onError: (error) => {
      console.error('User creation failed:', error);
    },
  });
};