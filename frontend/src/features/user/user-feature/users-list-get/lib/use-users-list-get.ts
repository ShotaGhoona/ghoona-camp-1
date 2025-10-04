import { useQuery } from '@tanstack/react-query';
import { getUsers, userKeys, type UsersQueryParams } from '@/entities/user/user-entity';

export const useUsersListGet = (params?: UsersQueryParams) => {
  return useQuery({
    queryKey: userKeys.list(params),
    queryFn: () => getUsers(params),
    enabled: true,
    staleTime: 1000 * 60 * 2, // 2分
  });
};