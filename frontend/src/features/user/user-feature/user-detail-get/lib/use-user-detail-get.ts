import { useQuery } from '@tanstack/react-query';
import { getUserDetail, userKeys } from '@/entities/user/user-entity';

export const useUserDetailGet = (userId: string | undefined) => {
  return useQuery({
    queryKey: userKeys.detail(userId!),
    queryFn: () => getUserDetail(userId!),
    enabled: !!userId,
    staleTime: 1000 * 60 * 3, // 3分
  });
};