/**
 * Session Get Hook
 * 
 * 認証セッション管理・現在ユーザー情報取得を提供します。
 * React 19のuseフックパターンを使用してデータを取得します。
 */

import { use, useMemo } from 'react';
import { getCurrentUser, mapAuthMeResponseToUser, type User } from '@/entities/user';

interface SessionData {
  user: User;
  isAuthenticated: boolean;
}

/**
 * 認証セッション取得フック
 * 
 * @returns セッションデータ（ユーザー情報、認証状態）
 * 
 * @example
 * ```tsx
 * const { user, isAuthenticated } = useGetSession();
 * 
 * if (!isAuthenticated) {
 *   return <LoginRequired />;
 * }
 * 
 * return <div>Welcome, {user.username}!</div>;
 * ```
 */
export const useGetSession = (): SessionData => {
  const sessionPromise = useMemo(
    () => getCurrentUser().then(mapAuthMeResponseToUser),
    []
  );
  
  const userData = use(sessionPromise);
  
  return {
    user: userData,
    isAuthenticated: !!userData,
  };
};