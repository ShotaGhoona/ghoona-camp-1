'use client';

import { Suspense, useState } from 'react';
import {
  // Auth Feature
  useAuthMeGet,
  
  // User Feature
  useUserCreate,
  useUsersListGet,
  useUserDetailGet,
  useUserBasicUpdate,
  
  // Metadata Feature
  useMetadataGet,
  useMetadataCreate,
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

// 実際のテストデータ
const TEST_USER_ID = '5298ae80-faeb-4f65-8894-695b652f8115';
const TEST_CLERK_ID = 'clerk_dev_user_001';
const TEST_EMAIL = 'developer@ghoona.camp';

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

// ユーザー作成テストコンポーネント
function UserCreateTest() {
  const createMutation = useUserCreate();
  
  const handleCreate = () => {
    createMutation.mutate({
      clerkId: `test_clerk_${Date.now()}`,
      email: `test.user.${Date.now()}@example.com`,
      username: `TestUser${Date.now()}`,
      avatarUrl: 'https://example.com/avatar.jpg',
      discordId: `discord_${Date.now()}`
    });
  };
  
  return (
    <div>
      <DataDisplay 
        title="ユーザー作成（POST /users）" 
        data={{
          mutationState: {
            isPending: createMutation.isPending,
            error: createMutation.error?.message,
            isSuccess: createMutation.isSuccess,
            data: createMutation.data
          }
        }} 
        status={createMutation.isPending ? "🔄 Creating" : createMutation.isSuccess ? "✅ Created" : "⏸️ Ready"} 
      />
      <button 
        onClick={handleCreate}
        disabled={createMutation.isPending}
        className="bg-green-500 text-white px-4 py-2 rounded text-sm disabled:opacity-50"
      >
        {createMutation.isPending ? '作成中...' : '新規ユーザー作成テスト'}
      </button>
    </div>
  );
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
    setSearchParams(prev => ({ ...prev, search: 'developer' }));
  };
  
  const handleFilter = () => {
    setSearchParams(prev => ({ 
      ...prev, 
      skills: 'Go,React', 
      sortBy: 'name'
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
          検索テスト (developer)
        </button>
        <button 
          onClick={handleFilter}
          className="bg-green-500 text-white px-4 py-2 rounded text-sm"
        >
          フィルタテスト (Go/React)
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
      username: `UpdatedUser_${Date.now()}`,
      avatarUrl: 'https://example.com/updated-avatar.jpg',
      discordId: `discord_updated_${Date.now()}`
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
            isSuccess: updateMutation.isSuccess,
            data: updateMutation.data
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

// メタデータ作成テストコンポーネント
function MetadataCreateTest({ userId }: { userId: string }) {
  const createMutation = useMetadataCreate(userId);
  
  const handleCreate = () => {
    createMutation.mutate({
      displayName: `New User ${Date.now()}`,
      profileImageUrl: 'https://example.com/profile.jpg',
      tagline: '新規プロフィール作成テスト',
      bio: 'これは新規作成されたメタデータのテストです。',
      vision: 'テクノロジーで世界をより良くする',
      visionPublic: true,
      timezone: 'Asia/Tokyo',
      skills: ['JavaScript', 'TypeScript', 'React', 'Node.js'],
      interests: ['プログラミング', 'AI', 'Web開発', 'オープンソース']
    });
  };
  
  if (!userId) return <DataDisplay title="メタデータ作成（POST /users/{userId}/metadata）" data="ユーザーIDを入力してください" status="⏸️ Waiting" />;
  
  return (
    <div>
      <DataDisplay 
        title="メタデータ作成（POST /users/{userId}/metadata）" 
        data={{
          mutationState: {
            isPending: createMutation.isPending,
            error: createMutation.error?.message,
            isSuccess: createMutation.isSuccess,
            data: createMutation.data
          }
        }} 
        status={createMutation.isPending ? "🔄 Creating" : createMutation.isSuccess ? "✅ Created" : "⏸️ Ready"} 
      />
      <button 
        onClick={handleCreate}
        disabled={createMutation.isPending}
        className="bg-green-500 text-white px-4 py-2 rounded text-sm disabled:opacity-50"
      >
        {createMutation.isPending ? '作成中...' : 'メタデータ作成テスト'}
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
      skills: ['Go', 'React', 'TypeScript', 'Docker', 'AWS'],
      interests: ['朝活', 'プログラミング', 'コミュニティ', '健康'],
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
            isSuccess: updateMutation.isSuccess,
            data: updateMutation.data
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
  const updateMutation = useSocialLinksUpdate(userId, data?.socialLinks?.[0]?.id || '');
  const deleteMutation = useSocialLinksDelete(userId);
  
  const handleCreate = () => {
    createMutation.mutate({
      platform: 'twitter',
      url: `https://twitter.com/test-user-${Date.now()}`,
      title: 'Test Twitter Profile',
      isPublic: true
    });
  };
  
  const handleUpdate = () => {
    if (data?.socialLinks && data.socialLinks.length > 0) {
      updateMutation.mutate({
        url: `https://twitter.com/updated-user-${Date.now()}`,
        title: 'Updated Twitter Profile',
        isPublic: false
      });
    }
  };
  
  const handleDelete = () => {
    if (data?.socialLinks && data.socialLinks.length > 0) {
      deleteMutation.mutate(data.socialLinks[0].id);
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
          disabled={updateMutation.isPending || !data?.socialLinks || data.socialLinks.length === 0}
          className="bg-yellow-500 text-white px-3 py-1 rounded text-sm disabled:opacity-50"
        >
          {updateMutation.isPending ? '更新中...' : 'UPDATE'}
        </button>
        <button 
          onClick={handleDelete}
          disabled={deleteMutation.isPending || !data?.socialLinks || data.socialLinks.length === 0}
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
    // 他の実際のユーザーIDを使用（テストデータから）
    createMutation.mutate({
      rivalUserId: 'b0e2d712-002a-4073-85f2-4daefcddf249' // Test Expert
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
          disabled={createMutation.isPending || (data?.total || 0) >= (data?.maxRivals || 3)}
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
        ライバル数: {data?.total || 0} / {data?.maxRivals || 3}
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
  const [testUserId, setTestUserId] = useState(TEST_USER_ID);

  return (
    <div className="container mx-auto p-6 max-w-7xl">
      <h1 className="text-3xl font-bold mb-6">User Features 接続テスト (Updated)</h1>
      
      <div className="mb-6 p-4 bg-blue-50 rounded border">
        <h3 className="font-semibold mb-2">🔧 テスト設定</h3>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-sm">
          <div>
            <strong>User ID:</strong><br />
            <code className="bg-gray-100 px-2 py-1 rounded">{TEST_USER_ID}</code>
          </div>
          <div>
            <strong>Clerk ID:</strong><br />
            <code className="bg-gray-100 px-2 py-1 rounded">{TEST_CLERK_ID}</code>
          </div>
          <div>
            <strong>Email:</strong><br />
            <code className="bg-gray-100 px-2 py-1 rounded">{TEST_EMAIL}</code>
          </div>
        </div>
        
        <div className="mt-4">
          <label className="block text-sm font-medium mb-2">カスタムテスト用ユーザーID:</label>
          <input 
            type="text"
            value={testUserId}
            onChange={(e) => setTestUserId(e.target.value)}
            className="border rounded px-3 py-2 w-80"
            placeholder="テスト用のユーザーIDを入力"
          />
          <button 
            onClick={() => setTestUserId(TEST_USER_ID)}
            className="ml-2 bg-gray-500 text-white px-3 py-2 rounded text-sm"
          >
            デフォルトに戻す
          </button>
        </div>
      </div>

      <div className="space-y-8">
        {/* 認証機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-blue-700 border-b-2 border-blue-200 pb-2">🔐 Auth Feature</h2>
          
          <div className="overflow-x-auto">
            <div className="flex gap-4 min-w-fit">
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <AuthMeTest />
                </ErrorBoundary>
              </div>
            </div>
          </div>
        </div>

        {/* ユーザー機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-green-700 border-b-2 border-green-200 pb-2">👥 User Feature</h2>
          
          <div className="overflow-x-auto">
            <div className="flex gap-4 min-w-fit">
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <UserCreateTest />
                </ErrorBoundary>
              </div>
              
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <UsersListTest />
                </ErrorBoundary>
              </div>
              
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <UserDetailTest userId={testUserId} />
                </ErrorBoundary>
              </div>
              
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <UserBasicUpdateTest userId={testUserId} />
                </ErrorBoundary>
              </div>
            </div>
          </div>
        </div>

        {/* メタデータ機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-purple-700 border-b-2 border-purple-200 pb-2">📝 Metadata Feature</h2>
          
          <div className="overflow-x-auto">
            <div className="flex gap-4 min-w-fit">
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <MetadataCreateTest userId={testUserId} />
                </ErrorBoundary>
              </div>
              
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <MetadataTest userId={testUserId} />
                </ErrorBoundary>
              </div>
            </div>
          </div>
        </div>

        {/* ソーシャルリンク機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-orange-700 border-b-2 border-orange-200 pb-2">🔗 Social Links Feature</h2>
          
          <div className="overflow-x-auto">
            <div className="flex gap-4 min-w-fit">
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <SocialLinksTest userId={testUserId} />
                </ErrorBoundary>
              </div>
            </div>
          </div>
        </div>

        {/* ライバル機能 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-red-700 border-b-2 border-red-200 pb-2">⚔️ Rivals Feature</h2>
          
          <div className="overflow-x-auto">
            <div className="flex gap-4 min-w-fit">
              <div className="flex-shrink-0 w-96">
                <ErrorBoundary>
                  <RivalsTest userId={testUserId} />
                </ErrorBoundary>
              </div>
            </div>
          </div>
        </div>

        {/* テスト手順説明 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-gray-700 border-b-2 border-gray-200 pb-2">📋 テスト手順 (Updated)</h2>
          
          <div className="overflow-x-auto">
            <div className="flex gap-4 min-w-fit">
              <div className="flex-shrink-0 w-96">
                <div className="bg-gray-50 p-4 rounded border h-full">
                  <h3 className="font-semibold mb-3">🆕 新機能テスト</h3>
                  <div className="space-y-2 text-sm">
                    <div><strong>ユーザー作成:</strong> POST /users で新規ユーザー追加</div>
                    <div><strong>メタデータ作成:</strong> POST /users/{'{userId}'}/metadata で初回プロフィール作成</div>
                    <div><strong>camelCase対応:</strong> 全APIレスポンスがcamelCase形式</div>
                  </div>
                </div>
              </div>
              
              <div className="flex-shrink-0 w-96">
                <div className="bg-gray-50 p-4 rounded border h-full">
                  <h3 className="font-semibold mb-3">🎯 主要テスト項目</h3>
                  <div className="space-y-2 text-sm">
                    <div><strong>認証:</strong> セッション情報の取得</div>
                    <div><strong>ユーザー:</strong> 作成・一覧・詳細・基本情報更新</div>
                    <div><strong>メタデータ:</strong> 作成・取得・更新</div>
                    <div><strong>ソーシャルリンク:</strong> CRUD操作（作成・読取・更新・削除）</div>
                    <div><strong>ライバル:</strong> 一覧・追加・削除</div>
                  </div>
                </div>
              </div>
              
              <div className="flex-shrink-0 w-96">
                <div className="bg-gray-50 p-4 rounded border h-full">
                  <h3 className="font-semibold mb-3">✅ 確認ポイント</h3>
                  <ul className="list-disc pl-5 space-y-1 text-sm">
                    <li>各APIの正常なレスポンス（✅ Success）</li>
                    <li>新機能（CREATE操作）の動作確認</li>
                    <li>camelCase形式のレスポンス確認</li>
                    <li>ローディング状態（🔄 Loading）</li>
                    <li>エラーハンドリング（❌ Error）</li>
                    <li>リアルタイム状態更新とキャッシュ無効化</li>
                  </ul>
                </div>
              </div>
              
              <div className="flex-shrink-0 w-96">
                <div className="bg-gray-50 p-4 rounded border h-full">
                  <h3 className="font-semibold mb-3">🧪 実テストデータ</h3>
                  <div className="text-xs space-y-1">
                    <div><strong>Base User:</strong> {TEST_EMAIL} (developer@ghoona.camp)</div>
                    <div><strong>Test Rival:</strong> Test Expert (tester@ghoona.camp)</div>
                    <div><strong>Platforms:</strong> Twitter, GitHub等の実際のプラットフォーム</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}