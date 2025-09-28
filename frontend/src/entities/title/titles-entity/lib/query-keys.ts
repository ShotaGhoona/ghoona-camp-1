/**
 * Titles Entity Query Keys
 * 
 * React Queryキー管理を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * titles-entity/index.ts からエクスポートされたキーを使用してください。
 */

export const titlesKeys = {
  all: ['titles'] as const,
  lists: () => [...titlesKeys.all, 'list'] as const,
  list: (filters: Record<string, unknown>) => [...titlesKeys.lists(), { filters }] as const,
  details: () => [...titlesKeys.all, 'detail'] as const,
  detail: (id: string) => [...titlesKeys.details(), id] as const,
} as const;