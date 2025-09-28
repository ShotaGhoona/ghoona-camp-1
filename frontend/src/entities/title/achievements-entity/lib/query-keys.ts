/**
 * Achievements Entity Query Keys
 * 
 * React Queryキー管理を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * achievements-entity/index.ts からエクスポートされたキーを使用してください。
 */

export const achievementsKeys = {
  all: ['achievements'] as const,
  lists: () => [...achievementsKeys.all, 'list'] as const,
  list: (userId: string, filters?: Record<string, unknown>) => [...achievementsKeys.lists(), userId, { filters }] as const,
  setCurrentTitle: (userId: string, titleId: string) => [...achievementsKeys.all, 'setCurrentTitle', userId, titleId] as const,
} as const;