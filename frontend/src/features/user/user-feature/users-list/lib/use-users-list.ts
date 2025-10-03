/**
 * Users List Hook
 * 
 * ユーザー一覧取得・検索・フィルター機能を提供します。
 * React 19のuseフックパターンを使用してデータを取得します。
 */

import { use, useMemo } from 'react';
import { getUsers, mapUsersListResponseToUsers, type UsersQueryParams } from '@/entities/user';

/**
 * ユーザー一覧取得フック
 * 
 * @param params - クエリパラメータ（検索・フィルター・ページング）
 * @returns ユーザー一覧データ
 * 
 * @example
 * ```tsx
 * const params = {
 *   search: 'プログラミング',
 *   skills: 'JavaScript,React',
 *   page: 1,
 *   limit: 20
 * };
 * const { users, pagination } = useGetUsersList(params);
 * 
 * return (
 *   <div>
 *     {users.map(user => (
 *       <UserCard key={user.id} user={user} />
 *     ))}
 *     <Pagination {...pagination} />
 *   </div>
 * );
 * ```
 */
export const useGetUsersList = (params?: UsersQueryParams) => {
  const usersPromise = useMemo(
    () => getUsers(params),
    [params]
  );
  
  const usersData = use(usersPromise);
  
  return {
    users: mapUsersListResponseToUsers(usersData),
    pagination: usersData.pagination,
  };
};