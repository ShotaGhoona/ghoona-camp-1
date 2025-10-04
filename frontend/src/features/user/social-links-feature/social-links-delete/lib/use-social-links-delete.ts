import { useMutation, useQueryClient } from '@tanstack/react-query';
import { deleteSocialLink, socialLinksKeys } from '@/entities/user/social-links-entity';
import { userKeys } from '@/entities/user/user-entity';

export const useSocialLinksDelete = (userId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (linkId: string) => 
      deleteSocialLink({ userId, linkId }),
    onSuccess: () => {
      // ソーシャルリンク一覧とユーザー詳細を無効化
      queryClient.invalidateQueries({ queryKey: socialLinksKeys.list(userId) });
      queryClient.invalidateQueries({ queryKey: userKeys.detail(userId) });
      // セッション情報も更新（自分のソーシャルリンク削除の場合）
      queryClient.invalidateQueries({ queryKey: userKeys.session() });
    },
    onError: (error) => {
      console.error('Social link deletion failed:', error);
    },
  });
};