/**
 * Query keys for goal-related queries
 */
export const goalQueryKeys = {
  all: ['goal'] as const,
  
  // Goal queries
  goals: () => [...goalQueryKeys.all, 'goals'] as const,
  goalsList: (filters?: Record<string, unknown>) => 
    [...goalQueryKeys.goals(), 'list', filters] as const,
  goalsDetail: (id: string) => 
    [...goalQueryKeys.goals(), 'detail', id] as const,
  
  // Goal progress queries
  progress: () => [...goalQueryKeys.all, 'progress'] as const,
  progressByGoal: (goalId: string) => 
    [...goalQueryKeys.progress(), 'goal', goalId] as const,
} as const;