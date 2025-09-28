/**
 * User Social Links Entity Query Keys
 * 
 * React Queryキー管理を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * social-links-entity/index.ts からエクスポートされたキーを使用してください。
 */

export const socialLinksKeys = {
  all: ['socialLinks'] as const,
  lists: () => [...socialLinksKeys.all, 'list'] as const,
  list: (userId: string) => [...socialLinksKeys.lists(), userId] as const,
  create: (userId: string) => [...socialLinksKeys.all, 'create', userId] as const,
  update: (userId: string, linkId: string) => [...socialLinksKeys.all, 'update', userId, linkId] as const,
  delete: (userId: string, linkId: string) => [...socialLinksKeys.all, 'delete', userId, linkId] as const,
} as const;