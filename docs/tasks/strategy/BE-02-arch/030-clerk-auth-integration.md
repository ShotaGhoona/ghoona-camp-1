# BE-02-arch-03 戦略書 - Clerk認証システム統合

**タスクID**: BE-02-arch-03  
**タスク名**: Clerk認証システム統合  
**内容**: Clerk JWT検証、認証ミドルウェア、ユーザー情報連携の実装  
**工数見積**: 1時間（元見積6時間から大幅短縮）  
**優先度**: 🔴 高  

## 📊 スコープ確認（チェックリスト準拠）

### ✅ 対象範囲（実装必要）
- Clerk Go SDK導入
- auth_service.goの最小限修正
- 動作確認

### ❌ 対象外（既存コードに騙されない）
- 複雑なJWT手動実装 → 不要（SDKが処理）
- カスタムWrapper → 不要（公式SDKで十分）
- 独自検証ロジック → 不要（セキュリティリスク）
- 複数Phase実装 → 不要（オーバーエンジニアリング）

## 🔍 既存実装状況の分析

### 📁 既存ファイルの評価
1. **backend/internal/infrastructure/clerk/auth_service.go** (108行)
   - ✅ 基本構造は使える
   - ❌ 手動JWT実装部分は削除必要（61-82行）
   - ✅ ミドルウェア連携部分は完璧

2. **backend/internal/interface/middleware/auth.go** (137行)
   - ✅ 完全実装済み（変更不要）
   - ✅ Bearer token処理完璧
   - ✅ エラーハンドリング完璧

3. **backend/internal/infrastructure/config/app_config.go** (90行)
   - ✅ 設定完璧（変更不要）

### 🎯 実際に必要な作業
**ほぼ完成済み。auth_service.goの数行修正のみ。**

## 📁 影響を及ぼすファイル・ディレクトリ（修正・新規作成対象）

```
backend/
├── go.mod                                            # ✏️ Clerk SDK v2依存関係追加
├── go.sum                                            # ✏️ Clerk SDK v2チェックサム追加
└── internal/
    └── infrastructure/
        └── clerk/
            └── auth_service.go                       # ✏️ VerifyToken関数のSDK使用への修正（30行程度）
```

**修正対象：たった1ファイル（auth_service.go）のみ**

## 📋 シンプル実装計画

### 実装内容（1時間）

#### Step 1: SDK導入 (5分)
```bash
cd backend
go get github.com/clerk/clerk-sdk-go/v2/clerk
```

#### Step 2: auth_service.go修正 (30分)
```go
// 対象: backend/internal/infrastructure/clerk/auth_service.go
// 修正箇所: VerifyToken関数のみ（34-82行を以下に置換）

import "github.com/clerk/clerk-sdk-go/v2/clerk"

func (s *AuthService) VerifyToken(ctx context.Context, token string) (*ClerkUser, error) {
    if s.secretKey == "" {
        return nil, errors.New("Clerk secret key が設定されていません")
    }

    // 開発環境用のモック処理（既存維持）
    if !s.enabled {
        if token == "mock-clerk-token" {
            return &ClerkUser{
                ID:       "user_mock123",
                Email:    "test@example.com",
                Username: "testuser",
            }, nil
        }
        return nil, errors.New("無効なモックトークンです")
    }

    // Clerk SDK使用（これだけ）
    client := clerk.NewClient(s.secretKey)
    user, err := client.Users().VerifyToken(ctx, token)
    if err != nil {
        return nil, err
    }

    return &ClerkUser{
        ID:       user.ID,
        Email:    user.EmailAddresses[0].EmailAddress,
        Username: user.Username,
    }, nil
}
```

#### Step 3: 動作確認 (15分)
```bash
cd backend
go build ./cmd/api
# 簡単な動作確認
```

#### Step 4: ドキュメント更新 (10分)
- 変更内容の記録
- 設定方法の確認

## 🔧 技術的考慮事項

### 1. Clerk Go SDK
- SDK: `github.com/clerk/clerk-sdk-go/v2/clerk`
- 必要設定: CLERK_SECRET_KEYのみ
- 検証: SDK内蔵機能使用（手動実装不要）

### 2. セキュリティ
- SDK内蔵の署名検証（最も安全）
- 公式メンテナンスによるセキュリティパッチ
- 複雑な独自実装を避けることでリスク軽減

### 3. 保守性
- 公式ドキュメント準拠
- シンプルな実装で理解しやすい
- バグ発生リスク最小化

## 📝 既存コードに騙されないポイント

### ❌ 削除すべき複雑な実装
- 手動JWT検証（auth_service.go 61-67行）
- カスタムクレーム処理（68-81行）
- 独自エラーハンドリング

### ✅ 保持すべきシンプルな実装
- 基本構造（AuthService struct）
- モック機能（開発環境用）
- 既存ミドルウェア（完璧に動作）

## ✅ 完了条件

1. ✅ Clerk SDKが導入済み
2. ✅ auth_service.goのVerifyToken関数がSDK使用
3. ✅ 既存ミドルウェアが正常動作
4. ✅ モック機能が開発環境で動作
5. ✅ ビルドエラーなし

**これだけ。複雑な実装は一切不要。**

## 🚀 次のタスクへの連携

本タスク完了後、**BE-03-user-01**で認証されたユーザー情報を活用したユーザー管理機能を実装予定。

---

**重要**: このタスクは1時間で完了できます。既存の複雑なコードに騙されず、Clerk公式SDKのシンプルな使用方法を採用することで、セキュアで保守性の高い認証システムを実現します。