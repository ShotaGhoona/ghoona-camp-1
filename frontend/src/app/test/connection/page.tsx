'use client';

import { Suspense, useState } from 'react';
import {
  useGetSession,
  useGetProfile,
  useUpdateProfile,
  useGetMetadata,
  useUpdateMetadata,
  useGetUsersList,
  useGetUserDetail,
  useGetRivals,
  useCreateRival,
  useDeleteRival,
  useGetSocialLinks,
  useCreateSocialLinks,
  useUpdateSocialLinks,
  useDeleteSocialLinks
} from '@/features/user';

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

// セッション確認コンポーネント
function SessionTest() {
  try {
    const sessionData = useGetSession();
    return <DataDisplay title="セッション情報" data={sessionData} status="✅ 成功" />;
  } catch (error: any) {
    return <DataDisplay title="セッション情報" data={{ error: error?.message || '不明なエラー' }} status="❌ エラー" />;
  }
}

// プロフィール確認コンポーネント
function ProfileTest({ userId }: { userId: string }) {
  const [updateStatus, setUpdateStatus] = useState<string>('');
  
  try {
    const profileData = useGetProfile(userId);
    const { updateProfile } = useUpdateProfile(userId);
    
    const handleUpdate = async () => {
      try {
        setUpdateStatus('更新中...');
        await updateProfile({
          display_name: '更新されたユーザー名',
          bio: '更新されたプロフィール'
        });
        setUpdateStatus('✅ 更新成功');
      } catch (error: any) {
        setUpdateStatus(`❌ 更新エラー: ${error?.message || '不明なエラー'}`);
      }
    };
    
    return (
      <div>
        <DataDisplay title="プロフィール情報" data={profileData} status="✅ 成功" />
        <button 
          onClick={handleUpdate}
          className="bg-blue-500 text-white px-4 py-2 rounded mb-2"
        >
          プロフィール更新テスト
        </button>
        {updateStatus && <p className="text-sm">{updateStatus}</p>}
      </div>
    );
  } catch (error: any) {
    return <DataDisplay title="プロフィール情報" data={{ error: error?.message || '不明なエラー' }} status="❌ エラー" />;
  }
}

// メタデータ確認コンポーネント
function MetadataTest({ userId }: { userId: string }) {
  const [updateStatus, setUpdateStatus] = useState<string>('');
  
  try {
    const metadataData = useGetMetadata(userId);
    const { updateMetadata } = useUpdateMetadata(userId);
    
    const handleUpdate = async () => {
      try {
        setUpdateStatus('更新中...');
        await updateMetadata({
          total_score: 1000,
          rank: 1,
          level: 10
        });
        setUpdateStatus('✅ 更新成功');
      } catch (error: any) {
        setUpdateStatus(`❌ 更新エラー: ${error?.message || '不明なエラー'}`);
      }
    };
    
    return (
      <div>
        <DataDisplay title="メタデータ" data={metadataData} status="✅ 成功" />
        <button 
          onClick={handleUpdate}
          className="bg-green-500 text-white px-4 py-2 rounded mb-2"
        >
          メタデータ更新テスト
        </button>
        {updateStatus && <p className="text-sm">{updateStatus}</p>}
      </div>
    );
  } catch (error: any) {
    return <DataDisplay title="メタデータ" data={{ error: error?.message || '不明なエラー' }} status="❌ エラー" />;
  }
}

// ユーザー一覧確認コンポーネント
function UsersTest() {
  try {
    const usersData = useGetUsersList({ limit: 5, offset: 0 });
    return <DataDisplay title="ユーザー一覧" data={usersData} status="✅ 成功" />;
  } catch (error: any) {
    return <DataDisplay title="ユーザー一覧" data={{ error: error?.message || '不明なエラー' }} status="❌ エラー" />;
  }
}

// ユーザー詳細確認コンポーネント
function UserDetailTest({ targetUserId }: { targetUserId: string }) {
  try {
    const userDetailData = useGetUserDetail(targetUserId);
    return <DataDisplay title="ユーザー詳細" data={userDetailData} status="✅ 成功" />;
  } catch (error: any) {
    return <DataDisplay title="ユーザー詳細" data={{ error: error?.message || '不明なエラー' }} status="❌ エラー" />;
  }
}

// ライバル機能確認コンポーネント
function RivalsTest({ userId }: { userId: string }) {
  const [actionStatus, setActionStatus] = useState<string>('');
  
  try {
    const rivalsData = useGetRivals(userId);
    const { createRival } = useCreateRival(userId);
    const { deleteRival } = useDeleteRival(userId);
    
    const handleCreateRival = async () => {
      try {
        setActionStatus('ライバル追加中...');
        await createRival({ rival_user_id: 'test-rival-id' });
        setActionStatus('✅ 追加成功');
      } catch (error: any) {
        setActionStatus(`❌ 追加エラー: ${error?.message || '不明なエラー'}`);
      }
    };
    
    const handleDeleteRival = async (rivalId: string) => {
      try {
        setActionStatus('ライバル削除中...');
        await deleteRival(rivalId);
        setActionStatus('✅ 削除成功');
      } catch (error: any) {
        setActionStatus(`❌ 削除エラー: ${error?.message || '不明なエラー'}`);
      }
    };
    
    return (
      <div>
        <DataDisplay title="ライバル情報" data={rivalsData} status="✅ 成功" />
        <div className="space-x-2 mb-2">
          <button 
            onClick={handleCreateRival}
            className="bg-purple-500 text-white px-4 py-2 rounded"
          >
            ライバル追加テスト
          </button>
          {rivalsData.rivals?.length > 0 && (
            <button 
              onClick={() => handleDeleteRival(rivalsData.rivals[0].id)}
              className="bg-red-500 text-white px-4 py-2 rounded"
            >
              最初のライバル削除テスト
            </button>
          )}
        </div>
        {actionStatus && <p className="text-sm">{actionStatus}</p>}
      </div>
    );
  } catch (error: any) {
    return <DataDisplay title="ライバル情報" data={{ error: error?.message || '不明なエラー' }} status="❌ エラー" />;
  }
}

// SNSリンク機能確認コンポーネント
function SocialLinksTest({ userId }: { userId: string }) {
  const [actionStatus, setActionStatus] = useState<string>('');
  
  try {
    const socialLinksData = useGetSocialLinks(userId);
    const { createSocialLinks } = useCreateSocialLinks(userId);
    const { updateSocialLinks } = useUpdateSocialLinks(userId);
    const { deleteSocialLinks } = useDeleteSocialLinks(userId);
    
    const handleCreateLink = async () => {
      try {
        setActionStatus('SNSリンク作成中...');
        await createSocialLinks({
          platform: 'twitter',
          url: 'https://twitter.com/test',
          display_order: 1
        });
        setActionStatus('✅ 作成成功');
      } catch (error: any) {
        setActionStatus(`❌ 作成エラー: ${error?.message || '不明なエラー'}`);
      }
    };
    
    const handleUpdateLink = async (linkId: string) => {
      try {
        setActionStatus('SNSリンク更新中...');
        await updateSocialLinks(linkId, {
          url: 'https://twitter.com/updated',
          display_order: 2
        });
        setActionStatus('✅ 更新成功');
      } catch (error: any) {
        setActionStatus(`❌ 更新エラー: ${error?.message || '不明なエラー'}`);
      }
    };
    
    const handleDeleteLink = async (linkId: string) => {
      try {
        setActionStatus('SNSリンク削除中...');
        await deleteSocialLinks(linkId);
        setActionStatus('✅ 削除成功');
      } catch (error: any) {
        setActionStatus(`❌ 削除エラー: ${error?.message || '不明なエラー'}`);
      }
    };
    
    return (
      <div>
        <DataDisplay title="SNSリンク情報" data={socialLinksData} status="✅ 成功" />
        <div className="space-x-2 mb-2">
          <button 
            onClick={handleCreateLink}
            className="bg-blue-500 text-white px-4 py-2 rounded"
          >
            SNSリンク作成テスト
          </button>
          {socialLinksData.socialLinks?.length > 0 && (
            <>
              <button 
                onClick={() => handleUpdateLink(socialLinksData.socialLinks[0].id)}
                className="bg-yellow-500 text-white px-4 py-2 rounded"
              >
                最初のリンク更新テスト
              </button>
              <button 
                onClick={() => handleDeleteLink(socialLinksData.socialLinks[0].id)}
                className="bg-red-500 text-white px-4 py-2 rounded"
              >
                最初のリンク削除テスト
              </button>
            </>
          )}
        </div>
        {actionStatus && <p className="text-sm">{actionStatus}</p>}
      </div>
    );
  } catch (error: any) {
    return <DataDisplay title="SNSリンク情報" data={{ error: error?.message || '不明なエラー' }} status="❌ エラー" />;
  }
}

// エラーバウンダリコンポーネント
function ErrorBoundary({ children }: { children: React.ReactNode }) {
  return (
    <div className="border-2 border-red-200 rounded p-4 mb-4">
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
  const [targetUserId, setTargetUserId] = useState('target-user-456');

  return (
    <div className="container mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6">ユーザー機能 接続テスト</h1>
      
      <div className="mb-6 space-y-4">
        <div>
          <label className="block text-sm font-medium mb-1">テスト用ユーザーID:</label>
          <input 
            type="text"
            value={testUserId}
            onChange={(e) => setTestUserId(e.target.value)}
            className="border rounded px-3 py-2 w-64"
            placeholder="テスト用のユーザーIDを入力"
          />
        </div>
        <div>
          <label className="block text-sm font-medium mb-1">対象ユーザーID（詳細表示用）:</label>
          <input 
            type="text"
            value={targetUserId}
            onChange={(e) => setTargetUserId(e.target.value)}
            className="border rounded px-3 py-2 w-64"
            placeholder="詳細表示用のユーザーIDを入力"
          />
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* 認証・プロフィール確認 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold">認証・プロフィール</h2>
          
          <ErrorBoundary>
            <SessionTest />
          </ErrorBoundary>
          
          <ErrorBoundary>
            <ProfileTest userId={testUserId} />
          </ErrorBoundary>
          
          <ErrorBoundary>
            <MetadataTest userId={testUserId} />
          </ErrorBoundary>
        </div>

        {/* ユーザー管理確認 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold">ユーザー管理</h2>
          
          <ErrorBoundary>
            <UsersTest />
          </ErrorBoundary>
          
          <ErrorBoundary>
            <UserDetailTest targetUserId={targetUserId} />
          </ErrorBoundary>
        </div>

        {/* ソーシャル機能確認 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold">ソーシャル機能</h2>
          
          <ErrorBoundary>
            <RivalsTest userId={testUserId} />
          </ErrorBoundary>
          
          <ErrorBoundary>
            <SocialLinksTest userId={testUserId} />
          </ErrorBoundary>
        </div>

        {/* テスト手順説明 */}
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold">テスト手順</h2>
          <div className="bg-gray-50 p-4 rounded">
            <h3 className="font-semibold mb-2">テスト方法:</h3>
            <ul className="list-disc pl-5 space-y-1 text-sm">
              <li>上記の入力欄に有効なユーザーIDを入力してください</li>
              <li>データが正しく読み込まれるか確認してください（緑色ステータス）</li>
              <li>ボタンを使用して変更操作をテストしてください</li>
              <li>ローディング状態とエラーハンドリングを確認してください</li>
              <li>詳細なエラーメッセージはブラウザコンソールで確認してください</li>
              <li>API レスポンスはネットワークタブで確認してください</li>
            </ul>
            
            <h3 className="font-semibold mt-4 mb-2">期待される動作:</h3>
            <ul className="list-disc pl-5 space-y-1 text-sm">
              <li>✅ 成功: データが正しく読み込まれ表示される</li>
              <li>🔄 読み込み中: データ取得中に表示される</li>
              <li>❌ エラー: API呼び出し失敗時やデータが無効な時に表示される</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}