import { useQuery } from '@tanstack/react-query';
import { getUserMetadata, metadataKeys } from '@/entities/user/metadata-entity';

export const useMetadataGet = (userId: string | undefined) => {
  return useQuery({
    queryKey: metadataKeys.detail(userId!),
    queryFn: () => getUserMetadata(userId!),
    enabled: !!userId,
    staleTime: 1000 * 60 * 5, // 5分
  });
};