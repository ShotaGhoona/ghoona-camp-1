import { useMutation, useQueryClient } from '@tanstack/react-query';
import { createSocialLink, socialLinksKeys, type CreateSocialLinkDto } from '@/entities/user/social-links-entity';
import { userKeys } from '@/entities/user/user-entity';

export const useSocialLinksCreate = (userId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateSocialLinkDto) => 
      createSocialLink({ userId, data }),
    onSuccess: () => {
      // ソーシャルリンク一覧とユーザー詳細を無効化
      queryClient.invalidateQueries({ queryKey: socialLinksKeys.list(userId) });
      queryClient.invalidateQueries({ queryKey: userKeys.detail(userId) });
    },
    onError: (error) => {
      console.error('Social link creation failed:', error);
    },
  });
};