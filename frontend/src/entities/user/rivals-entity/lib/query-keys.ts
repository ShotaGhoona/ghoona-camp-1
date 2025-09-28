/**
 * User Rivals Entity Query Keys
 * 
 * React Queryキー管理を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * rivals-entity/index.ts からエクスポートされたキーを使用してください。
 */

export const rivalsKeys = {
  all: ['rivals'] as const,
  lists: () => [...rivalsKeys.all, 'list'] as const,
  list: (userId: string) => [...rivalsKeys.lists(), userId] as const,
  create: (userId: string) => [...rivalsKeys.all, 'create', userId] as const,
  delete: (userId: string, rivalId: string) => [...rivalsKeys.all, 'delete', userId, rivalId] as const,
} as const;