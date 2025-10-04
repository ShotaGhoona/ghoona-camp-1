import { useQuery } from '@tanstack/react-query';
import { getUserRivals, rivalsKeys } from '@/entities/user/rivals-entity';

export const useRivalsGet = (userId: string | undefined) => {
  return useQuery({
    queryKey: rivalsKeys.list(userId!),
    queryFn: () => getUserRivals(userId!),
    enabled: !!userId,
    staleTime: 1000 * 60 * 2, // 2分（ライバル情報は比較的変更頻度が低い）
  });
};