/**
 * User Metadata Entity Query Keys
 * 
 * React Queryキー管理を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * metadata-entity/index.ts からエクスポートされたキーを使用してください。
 */

export const metadataKeys = {
  all: ['metadata'] as const,
  details: () => [...metadataKeys.all, 'detail'] as const,
  detail: (userId: string) => [...metadataKeys.details(), userId] as const,
  create: (userId: string) => [...metadataKeys.all, 'create', userId] as const,
  update: (userId: string) => [...metadataKeys.all, 'update', userId] as const,
} as const;