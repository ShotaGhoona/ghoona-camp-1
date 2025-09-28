/**
 * Query keys for title-related queries
 */
export const titleQueryKeys = {
  all: ['title'] as const,
  
  // Title queries
  titles: () => [...titleQueryKeys.all, 'titles'] as const,
  titlesList: (filters?: Record<string, any>) => 
    [...titleQueryKeys.titles(), 'list', filters] as const,
  titlesDetail: (id: string) => 
    [...titleQueryKeys.titles(), 'detail', id] as const,
  
  // Achievements queries
  achievements: () => [...titleQueryKeys.all, 'achievements'] as const,
  achievementsList: (filters?: Record<string, any>) => 
    [...titleQueryKeys.achievements(), 'list', filters] as const,
  achievementsDetail: (id: string) => 
    [...titleQueryKeys.achievements(), 'detail', id] as const,
  achievementsUser: (userId: string) => 
    [...titleQueryKeys.achievements(), 'user', userId] as const,
} as const;