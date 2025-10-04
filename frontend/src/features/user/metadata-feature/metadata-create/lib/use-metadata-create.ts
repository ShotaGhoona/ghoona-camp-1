import { useMutation, useQueryClient } from '@tanstack/react-query';
import { createUserMetadata, metadataKeys, type CreateUserMetadataDto } from '@/entities/user/metadata-entity';
import { userKeys } from '@/entities/user/user-entity';

export const useMetadataCreate = (userId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateUserMetadataDto) => 
      createUserMetadata({ userId, data }),
    onSuccess: () => {
      // メタデータとユーザー詳細を無効化
      queryClient.invalidateQueries({ queryKey: metadataKeys.detail(userId) });
      queryClient.invalidateQueries({ queryKey: userKeys.detail(userId) });
      // ユーザー一覧も更新（プロフィール情報が表示される場合）
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
      // セッション情報も更新（自分のメタデータ作成の場合）
      queryClient.invalidateQueries({ queryKey: userKeys.session() });
    },
    onError: (error) => {
      console.error('Metadata creation failed:', error);
    },
  });
};