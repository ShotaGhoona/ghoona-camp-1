import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateUserMetadata, metadataKeys, type UpdateUserMetadataDto } from '@/entities/user/metadata-entity';
import { userKeys } from '@/entities/user/user-entity';

export const useMetadataUpdate = (userId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateUserMetadataDto) => 
      updateUserMetadata({ userId, data }),
    onSuccess: () => {
      // メタデータとユーザー詳細を無効化
      queryClient.invalidateQueries({ queryKey: metadataKeys.detail(userId) });
      queryClient.invalidateQueries({ queryKey: userKeys.detail(userId) });
      // ユーザー一覧も更新（プロフィール情報が表示される場合）
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
    },
    onError: (error) => {
      console.error('Metadata update failed:', error);
    },
  });
};