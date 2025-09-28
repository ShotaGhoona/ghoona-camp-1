/**
 * Query keys for user-related queries
 */
export const userQueryKeys = {
  all: ['user'] as const,
  
  // User queries
  users: () => [...userQueryKeys.all, 'users'] as const,
  usersList: (filters?: Record<string, any>) => 
    [...userQueryKeys.users(), 'list', filters] as const,
  usersDetail: (id: string) => 
    [...userQueryKeys.users(), 'detail', id] as const,
  usersProfile: (id: string) => 
    [...userQueryKeys.users(), 'profile', id] as const,
  
  // Metadata queries
  metadata: () => [...userQueryKeys.all, 'metadata'] as const,
  metadataByUser: (userId: string) => 
    [...userQueryKeys.metadata(), 'user', userId] as const,
  
  // Rivals queries
  rivals: () => [...userQueryKeys.all, 'rivals'] as const,
  rivalsList: (userId: string) => 
    [...userQueryKeys.rivals(), 'list', userId] as const,
  rivalsDetail: (userId: string, rivalId: string) => 
    [...userQueryKeys.rivals(), 'detail', userId, rivalId] as const,
  
  // Social Links queries
  socialLinks: () => [...userQueryKeys.all, 'socialLinks'] as const,
  socialLinksByUser: (userId: string) => 
    [...userQueryKeys.socialLinks(), 'user', userId] as const,
} as const;