import { useQuery } from '@tanstack/react-query';
import { getCurrentUser, userKeys } from '@/entities/user/user-entity';

export const useAuthMeGet = () => {
  return useQuery({
    queryKey: userKeys.session(),
    queryFn: getCurrentUser,
    staleTime: 1000 * 60 * 5, // 5分
    retry: 2,
  });
};