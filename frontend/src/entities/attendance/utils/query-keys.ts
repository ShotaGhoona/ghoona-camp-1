/**
 * Query keys for attendance-related queries
 */
export const attendanceQueryKeys = {
  all: ['attendance'] as const,
  
  // Logs queries
  logs: () => [...attendanceQueryKeys.all, 'logs'] as const,
  logsList: (filters?: Record<string, any>) => 
    [...attendanceQueryKeys.logs(), 'list', filters] as const,
  logsDetail: (id: string) => 
    [...attendanceQueryKeys.logs(), 'detail', id] as const,
  
  // Ranking queries
  ranking: () => [...attendanceQueryKeys.all, 'ranking'] as const,
  rankingList: (filters?: Record<string, any>) => 
    [...attendanceQueryKeys.ranking(), 'list', filters] as const,
  
  // Statistics queries
  statistics: () => [...attendanceQueryKeys.all, 'statistics'] as const,
  statisticsOverview: (period?: string) => 
    [...attendanceQueryKeys.statistics(), 'overview', period] as const,
  
  // Summaries queries
  summaries: () => [...attendanceQueryKeys.all, 'summaries'] as const,
  summariesList: (period?: string) => 
    [...attendanceQueryKeys.summaries(), 'list', period] as const,
} as const;