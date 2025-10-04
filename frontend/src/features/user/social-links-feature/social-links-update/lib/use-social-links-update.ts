import { useMutation, useQueryClient } from '@tanstack/react-query';
import { updateSocialLink, socialLinksKeys, type UpdateSocialLinkDto } from '@/entities/user/social-links-entity';
import { userKeys } from '@/entities/user/user-entity';

export const useSocialLinksUpdate = (userId: string, linkId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateSocialLinkDto) => 
      updateSocialLink({ userId, linkId, data }),
    onSuccess: () => {
      // ソーシャルリンク一覧とユーザー詳細を無効化
      queryClient.invalidateQueries({ queryKey: socialLinksKeys.list(userId) });
      queryClient.invalidateQueries({ queryKey: userKeys.detail(userId) });
      // セッション情報も更新（自分のソーシャルリンク更新の場合）
      queryClient.invalidateQueries({ queryKey: userKeys.session() });
    },
    onError: (error) => {
      console.error('Social link update failed:', error);
    },
  });
};