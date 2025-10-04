import { useQuery } from '@tanstack/react-query';
import { getUserSocialLinks, socialLinksKeys } from '@/entities/user/social-links-entity';

export const useSocialLinksGet = (userId: string | undefined) => {
  return useQuery({
    queryKey: socialLinksKeys.list(userId!),
    queryFn: () => getUserSocialLinks(userId!),
    enabled: !!userId,
    staleTime: 1000 * 60 * 3, // 3分
  });
};