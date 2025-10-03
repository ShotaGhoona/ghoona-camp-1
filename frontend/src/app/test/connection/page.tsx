'use client';

import { Suspense, useState } from 'react';
import {
  // Auth Feature
  useAuthMeGet,
  
  // User Feature
  useUsersListGet,
  useUserDetailGet,
  useUserBasicUpdate,
  
  // Metadata Feature
  useMetadataGet,
  useMetadataUpdate,
  
  // Social Links Feature
  useSocialLinksGet,
  useSocialLinksCreate,
  useSocialLinksUpdate,
  useSocialLinksDelete,
  
  // Rivals Feature
  useRivalsGet,
  useRivalsCreate,
  useRivalsDelete
} from '@/features/user';

// Entity層から型定義をインポート
import type { UsersQueryParams } from '@/entities/user/user-entity';

// データ表示コンポーネント
function DataDisplay({ title, data, status }: { title: string; data: any; status?: string }) {
  return (
    <div className="border rounded p-4 mb-4">
      <h3 className="font-bold mb-2">{title}</h3>
      {status && <p className="text-sm text-gray-600 mb-2">ステータス: {status}</p>}
      <pre className="bg-gray-100 p-2 rounded text-sm overflow-auto max-h-40">
        {JSON.stringify(data, null, 2)}
      </pre>
    </div>
  );
}

// 認証テストコンポーネント
function AuthMeTest() {
  const { data, isLoading, error } = useAuthMeGet();
  
  if (isLoading) return <DataDisplay title="認証情報（GET /auth/me）" data="読み込み中..." status="🔄 Loading" />;
  if (error) return <DataDisplay title="認証情報（GET /auth/me）" data={{ error: error.message }} status="❌ Error" />;
  
  return <DataDisplay title="認証情報（GET /auth/me）" data={data} status="✅ Success" />;
}

// ユーザー一覧テストコンポーネント
function UsersListTest() {
  const [searchParams, setSearchParams] = useState<UsersQueryParams>({
    page: 1,
    limit: 5,
    search: '',
    sortBy: 'createdAt',
    order: 'desc'
  });
  
  const { data, isLoading, error } = useUsersListGet(searchParams);
  
  const handleSearch = () => {
    setSearchParams(prev => ({ ...prev, search: 'john' }));
  };
  
  const handleFilter = () => {
    setSearchParams(prev => ({ 
      ...prev, 
      skills: 'JavaScript,React', // カンマ区切り文字列として送信
      sortBy: 'attendanceDays'
    }));
  };
  
  if (isLoading) return <DataDisplay title="ユーザー一覧（GET /users）" data="読み込み中..." status="🔄 Loading" />;
  if (error) return <DataDisplay title="ユーザー一覧（GET /users）" data={{ error: error.message }} status="❌ Error" />;
  
  return (
    <div>
      <DataDisplay title="ユーザー一覧（GET /users）" data={data} status="✅ Success" />
      <div className="space-x-2 mb-2">
        <button 
          onClick={handleSearch}
          className="bg-blue-500 text-white px-4 py-2 rounded text-sm"
        >
          検索テスト (john)
        </button>
        <button 
          onClick={handleFilter}
          className="bg-green-500 text-white px-4 py-2 rounded text-sm"
        >
          フィルタテスト (JS/React)
        </button>
      </div>
    </div>
  );
}

// ユーザー詳細テストコンポーネント
function UserDetailTest({ userId }: { userId: string }) {
  const { data, isLoading, error } = useUserDetailGet(userId);
  
  if (!userId) return <DataDisplay title="ユーザー詳細（GET /users/{userId}）" data="ユーザーIDを入力してください" status="⏸️ Waiting" />;
  if (isLoading) return <DataDisplay title="ユーザー詳細（GET /users/{userId}）" data="読み込み中..." status="🔄 Loading" />;
  if (error) return <DataDisplay title="ユーザー詳細（GET /users/{userId}）" data={{ error: error.message }} status="❌ Error" />;
  
  return <DataDisplay title="ユーザー詳細（GET /users/{userId}）" data={data} status="✅ Success" />;
}

// ユーザー基本更新テストコンポーネント
function UserBasicUpdateTest({ userId }: { userId: string }) {
  const { data: userData } = useUserDetailGet(userId);
  const updateMutation = useUserBasicUpdate(userId);
  
  const handleUpdate = () => {
    updateMutation.mutate({
      username: `updated_${Date.now()}`,
      avatarUrl: 'https://example.com/new-avatar.jpg'
    });
  };
  
  if (!userId) return <DataDisplay title="ユーザー基本更新（PUT /users/{userId}）" data="ユーザーIDを入力してください" status="⏸️ Waiting" />;
  
  return (
    <div>
      <DataDisplay 
        title="ユーザー基本更新（PUT /users/{userId}）" 
        data={{
          currentUser: userData,
          mutationState: {
            isPending: updateMutation.isPending,
            error: updateMutation.error?.message,
            isSuccess: updateMutation.isSuccess
          }
        }} 
        status={updateMutation.isPending ? "🔄 Updating" : updateMutation.isSuccess ? "✅ Updated" : "⏸️ Ready"} 
      />
      <button 
        onClick={handleUpdate}
        disabled={updateMutation.isPending}
        className="bg-blue-500 text-white px-4 py-2 rounded text-sm disabled:opacity-50"
      >
        {updateMutation.isPending ? '更新中...' : 'ユーザー情報更新テスト'}
      </button>
    </div>
  );
}

// メタデータテストコンポーネント
function MetadataTest({ userId }: { userId: string }) {
  const { data, isLoading, error } = useMetadataGet(userId);
  const updateMutation = useMetadataUpdate(userId);
  
  const handleUpdate = () => {
    updateMutation.mutate({
      displayName: `Updated User ${Date.now()}`,
      tagline: 'Updated tagline from test',
      bio: 'This is an updated bio from the connection test',
      skills: ['JavaScript', 'React', 'TypeScript', 'Node.js'],
      interests: ['Web Development', 'AI/ML', 'Open Source'],
      visionPublic: true
    });
  };
  
  if (!userId) return <DataDisplay title="メタデータ（GET/PUT /users/{userId}/metadata）" data="ユーザーIDを入力してください" status="⏸️ Waiting" />;
  if (isLoading) return <DataDisplay title="メタデータ（GET/PUT /users/{userId}/metadata）" data="読み込み中..." status="🔄 Loading" />;
  if (error) return <DataDisplay title="メタデータ（GET/PUT /users/{userId}/metadata）" data={{ error: error.message }} status="❌ Error" />;
  
  return (
    <div>
      <DataDisplay 
        title="メタデータ（GET/PUT /users/{userId}/metadata）" 
        data={{
          metadata: data,
          mutationState: {
            isPending: updateMutation.isPending,
            error: updateMutation.error?.message,
            isSuccess: updateMutation.isSuccess
          }
        }} 
        status="✅ Success" 
      />
      <button 
        onClick={handleUpdate}
        disabled={updateMutation.isPending}
        className="bg-green-500 text-white px-4 py-2 rounded text-sm disabled:opacity-50"
      >
        {updateMutation.isPending ? '更新中...' : 'メタデータ更新テスト'}
      </button>
    </div>
  );
}

// ソーシャルリンクテストコンポーネント
function SocialLinksTest({ userId }: { userId: string }) {
  const { data, isLoading, error } = useSocialLinksGet(userId);
  const createMutation = useSocialLinksCreate(userId);
  const updateMutation = useSocialLinksUpdate(userId, data?.[0]?.id || '');
  const deleteMutation = useSocialLinksDelete(userId);
  
  const handleCreate = () => {
    createMutation.mutate({
      platform: 'github',
      url: `https://github.com/test-user-${Date.now()}`,
      title: 'Test GitHub Profile',
      isPublic: true
    });
  };
  
  const handleUpdate = () => {
    if (data && data.length > 0) {
      updateMutation.mutate({
        url: `https://github.com/updated-user-${Date.now()}`,
        title: 'Updated GitHub Profile',
        isPublic: false
      });
    }
  };
  
  const handleDelete = () => {
    if (data && data.length > 0) {
      deleteMutation.mutate(data[0].id);
    }
  };
  
  if (!userId) return <DataDisplay title="ソーシャルリンク（/users/{userId}/social-links）" data="ユーザーIDを入力してください" status="⏸️ Waiting" />;
  if (isLoading) return <DataDisplay title="ソーシャルリンク（/users/{userId}/social-links）" data="読み込み中..." status="🔄 Loading" />;
  if (error) return <DataDisplay title="ソーシャルリンク（/users/{userId}/social-links）" data={{ error: error.message }} status="❌ Error" />;
  
  return (
    <div>
      <DataDisplay 
        title="ソーシャルリンク（/users/{userId}/social-links）" 
        data={{
          socialLinks: data,
          mutations: {
            create: { isPending: createMutation.isPending, isSuccess: createMutation.isSuccess },
            update: { isPending: updateMutation.isPending, isSuccess: updateMutation.isSuccess },
            delete: { isPending: deleteMutation.isPending, isSuccess: deleteMutation.isSuccess }
          }
        }} 
        status="✅ Success" 
      />
      <div className="space-x-2 mb-2">
        <button 
          onClick={handleCreate}
          disabled={createMutation.isPending}
          className="bg-blue-500 text-white px-3 py-1 rounded text-sm disabled:opacity-50"
        >
          {createMutation.isPending ? '作成中...' : 'CREATE'}
        </button>
        <button 
          onClick={handleUpdate}
          disabled={updateMutation.isPending || !data || data.length === 0}
          className="bg-yellow-500 text-white px-3 py-1 rounded text-sm disabled:opacity-50"
        >
          {updateMutation.isPending ? '更新中...' : 'UPDATE'}
        </button>
        <button 
          onClick={handleDelete}
          disabled={deleteMutation.isPending || !data || data.length === 0}
          className="bg-red-500 text-white px-3 py-1 rounded text-sm disabled:opacity-50"
        >
          {deleteMutation.isPending ? '削除中...' : 'DELETE'}
        </button>
      </div>
    </div>
  );
}

// ライバルテストコンポーネント
function RivalsTest({ userId }: { userId: string }) {
  const { data, isLoading, error } = useRivalsGet(userId);
  const createMutation = useRivalsCreate(userId);
  const deleteMutation = useRivalsDelete(userId);
  
  const handleCreate = () => {
    // テスト用の仮のライバルユーザーID
    createMutation.mutate({
      rivalUserId: 'test-rival-user-id'
    });
  };
  
  const handleDelete = () => {
    if (data?.rivals && data.rivals.length > 0) {
      deleteMutation.mutate(data.rivals[0].id);
    }
  };
  
  if (!userId) return <DataDisplay title="ライバル（/users/{userId}/rivals）" data="ユーザーIDを入力してください" status="⏸️ Waiting" />;
  if (isLoading) return <DataDisplay title="ライバル（/users/{userId}/rivals）" data="読み込み中..." status="🔄 Loading" />;
  if (error) return <DataDisplay title="ライバル（/users/{userId}/rivals）" data={{ error: error.message }} status="❌ Error" />;
  
  return (
    <div>
      <DataDisplay 
        title="ライバル（/users/{userId}/rivals）" 
        data={{
          rivals: data,
          mutations: {
            create: { isPending: createMutation.isPending, isSuccess: createMutation.isSuccess },
            delete: { isPending: deleteMutation.isPending, isSuccess: deleteMutation.isSuccess }
          }
        }} 
        status="✅ Success" 
      />
      <div className="space-x-2 mb-2">
        <button 
          onClick={handleCreate}
          disabled={createMutation.isPending || (data?.count || 0) >= 3}
          className="bg-purple-500 text-white px-3 py-1 rounded text-sm disabled:opacity-50"
        >
          {createMutation.isPending ? '追加中...' : 'CREATE RIVAL'}
        </button>
        <button 
          onClick={handleDelete}
          disabled={deleteMutation.isPending || !data?.rivals || data.rivals.length === 0}
          className="bg-red-500 text-white px-3 py-1 rounded text-sm disabled:opacity-50"
        >
          {deleteMutation.isPending ? '削除中...' : 'DELETE RIVAL'}
        </button>
      </div>
      <p className="text-xs text-gray-500">
        ライバル数: {data?.count || 0} / {data?.maxRivals || 3}
      </p>
    </div>
  );
}

// エラーバウンダリコンポーネント
function ErrorBoundary({ children }: { children: React.ReactNode }) {
  return (
    <div className="border-2 border-gray-200 rounded p-4 mb-4">
      <Suspense fallback={
        <div className="text-blue-600">🔄 データ読み込み中...</div>
      }>
        {children}
      </Suspense>
    </div>
  );
}

export default function ConnectionTestPage() {
  const [testUserId, setTestUserId] = useState('user-123');

  return (
    <div className="container mx-auto p-6 max-w-7xl">
      <h1 className="text-3xl font-bold mb-6">User Features 接続テスト</h1>
      
      <div className="mb-6">
        <label className="block text-sm font-medium mb-2">テスト用ユーザーID:</label>
        <input 
          type="text"
          value={testUserId}
          onChange={(e) => setTestUserId(e.target.value)}
          className="border rounded px-3 py-2 w-80"
          placeholder="テスト用のユーザーIDを入力"
        />
        <p className="text-xs text-gray-500 mt-1">
          実際のユーザーIDを入力するか、テスト用の値のままで動作確認できます
        </p>
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-2 gap-6">
        {/* 認証機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-blue-700">🔐 Auth Feature</h2>
          
          <ErrorBoundary>
            <AuthMeTest />
          </ErrorBoundary>
        </div>

        {/* ユーザー機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-green-700">👥 User Feature</h2>
          
          <ErrorBoundary>
            <UsersListTest />
          </ErrorBoundary>
          
          <ErrorBoundary>
            <UserDetailTest userId={testUserId} />
          </ErrorBoundary>
          
          <ErrorBoundary>
            <UserBasicUpdateTest userId={testUserId} />
          </ErrorBoundary>
        </div>

        {/* メタデータ機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-purple-700">📝 Metadata Feature</h2>
          
          <ErrorBoundary>
            <MetadataTest userId={testUserId} />
          </ErrorBoundary>
        </div>

        {/* ソーシャルリンク機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-orange-700">🔗 Social Links Feature</h2>
          
          <ErrorBoundary>
            <SocialLinksTest userId={testUserId} />
          </ErrorBoundary>
        </div>

        {/* ライバル機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-red-700">⚔️ Rivals Feature</h2>
          
          <ErrorBoundary>
            <RivalsTest userId={testUserId} />
          </ErrorBoundary>
        </div>

        {/* テスト手順説明 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-gray-700">📋 テスト手順</h2>
          <div className="bg-gray-50 p-4 rounded border">
            <h3 className="font-semibold mb-3">🎯 テスト項目</h3>
            <div className="space-y-2 text-sm">
              <div><strong>認証:</strong> セッション情報の取得</div>
              <div><strong>ユーザー:</strong> 一覧・詳細・基本情報更新</div>
              <div><strong>メタデータ:</strong> プロフィール詳細の取得・更新</div>
              <div><strong>ソーシャルリンク:</strong> CRUD操作（作成・読取・更新・削除）</div>
              <div><strong>ライバル:</strong> 一覧・追加・削除</div>
            </div>
            
            <h3 className="font-semibold mt-4 mb-3">✅ 確認ポイント</h3>
            <ul className="list-disc pl-5 space-y-1 text-sm">
              <li>各APIの正常なレスポンス（✅ Success）</li>
              <li>ローディング状態（🔄 Loading）</li>
              <li>エラーハンドリング（❌ Error）</li>
              <li>リアルタイム状態更新</li>
              <li>ボタンクリックでの操作実行</li>
              <li>キャッシュ無効化とデータ再取得</li>
            </ul>

            <h3 className="font-semibold mt-4 mb-3">🔧 デバッグ</h3>
            <ul className="list-disc pl-5 space-y-1 text-sm">
              <li>ブラウザコンソールでエラー詳細確認</li>
              <li>ネットワークタブでAPI呼び出し確認</li>
              <li>React Query DevTools使用推奨</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}