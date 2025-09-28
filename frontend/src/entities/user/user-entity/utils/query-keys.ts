/**
 * User Entity Query Keys
 * React Queryキー管理
 */

export const userKeys = {
  all: ['users'] as const,
  lists: () => [...userKeys.all, 'list'] as const,
  list: (filters: Record<string, unknown>) => [...userKeys.lists(), { filters }] as const,
  details: () => [...userKeys.all, 'detail'] as const,
  detail: (id: string) => [...userKeys.details(), id] as const,
  update: (id: string) => [...userKeys.all, 'update', id] as const,
} as const;

// 認証関連（Clerkとの区別）
export const authKeys = {
  all: ['auth'] as const,
  me: () => [...authKeys.all, 'me'] as const, // GET /auth/me用
} as const;