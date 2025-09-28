# 🔐 Ghoona Camp 認証システム完全ガイド

**対象**: プログラミング初学者  
**目的**: 認証システムの基本からGhoona Campでの実装まで段階的に理解する

---

## 📚 目次

1. [日常例から理解する認証](#1-日常例から理解する認証)
2. [一般的な認証システム](#2-一般的な認証システム)
3. [Clerkとは何か](#3-clerkとは何か)
4. [Ghoona Campでの認証の流れ](#4-ghoona-campでの認証の流れ)
5. [実装の詳細](#5-実装の詳細)
6. [開発者が知るべきポイント](#6-開発者が知るべきポイント)

---

## 1. 日常例から理解する認証

### 🏠 家の鍵で理解する認証の基本

認証システムは「家の鍵」のようなものです。

```mermaid
graph LR
    A[あなた] --> B[鍵を使う]
    B --> C{正しい鍵？}
    C -->|はい| D[家に入れる]
    C -->|いいえ| E[入れない]
    
    F[家の所有者] --> G[鍵を発行]
    G --> H[あなたに渡す]
```

**日常の例**:
- **鍵** = ログイン情報（メール/パスワード）
- **家** = アプリケーション
- **鍵の確認** = 認証
- **家に入る** = アプリが使える

### 🏢 会社の入館証で理解する

```mermaid
graph TD
    A[社員] --> B[入館証をかざす]
    B --> C[受付システム]
    C --> D{有効な入館証？}
    D -->|はい| E[入館許可]
    D -->|いいえ| F[入館拒否]
    
    G[人事部] --> H[入館証を発行]
    H --> I[社員に配布]
```

**重要なポイント**:
- **入館証** = セッショントークン
- **受付システム** = 認証サーバー
- **人事部** = 認証プロバイダー（Clerk）

---

## 2. 一般的な認証システム

### 🔑 認証の3つの要素

```mermaid
graph TD
    A[認証システム] --> B[認証 Authentication]
    A --> C[認可 Authorization]
    A --> D[セッション管理]
    
    B --> B1[あなたは誰？]
    C --> C1[何ができる？]
    D --> D1[いつまで有効？]
    
    B1 --> B2[メール + パスワード]
    C1 --> C2[管理者 or 一般ユーザー]
    D1 --> D2[24時間有効]
```

### 💳 トークンベース認証の仕組み

**クレジットカードの例**で理解しましょう。

```mermaid
sequenceDiagram
    participant U as ユーザー
    participant A as アプリ
    participant B as 銀行（認証サーバー）
    participant S as お店（リソース）
    
    U->>A: ログイン要求
    A->>B: 認証情報確認
    B->>B: 身元確認
    B->>A: クレジットカード発行
    A->>U: カード受け取り
    
    Note over U,A: 以降、カードで買い物可能
    
    U->>S: カードで支払い
    S->>B: カード有効性確認
    B->>S: 有効です
    S->>U: 商品提供
```

**実際の認証システム**:
- **クレジットカード** = JWTトークン
- **銀行** = Clerk認証サーバー
- **お店** = Ghoona Camp API
- **商品** = ユーザーデータ

---

## 3. Clerkとは何か

### 🏦 Clerkを銀行で例える

```mermaid
graph TD
    A[あなた] --> B[銀行口座開設]
    B --> C[Clerk銀行]
    C --> D[身元確認]
    D --> E[クレジットカード発行]
    
    F[様々なお店] --> G[Clerkカード対応]
    G --> H[安全な決済]
    
    C --> I[セキュリティ管理]
    C --> J[不正利用防止]
    C --> K[24時間監視]
    
    style C fill:#e1f5fe
    style I fill:#f3e5f5
    style J fill:#f3e5f5
    style K fill:#f3e5f5
```

### 🎯 Clerkの利点

**自分で認証システムを作る場合**:
```mermaid
graph TD
    A[開発者] --> B[パスワード保存]
    A --> C[暗号化実装]
    A --> D[セッション管理]
    A --> E[セキュリティ対策]
    A --> F[パスワードリセット]
    A --> G[ソーシャルログイン]
    
    B --> H[🚨 セキュリティリスク]
    C --> H
    D --> H
    E --> H
    
    style H fill:#ffebee
    style A fill:#fff3e0
```

**Clerkを使う場合**:
```mermaid
graph TD
    A[開発者] --> B[Clerk SDK導入]
    B --> C[5分で認証完成]
    
    D[Clerk] --> E[プロのセキュリティ]
    D --> F[自動アップデート]
    D --> G[99.9%稼働率]
    D --> H[企業レベル保護]
    
    style C fill:#e8f5e8
    style E fill:#e3f2fd
    style F fill:#e3f2fd
    style G fill:#e3f2fd
    style H fill:#e3f2fd
```

---

## 4. Ghoona Campでの認証の流れ

### 🌅 朝活アプリでの具体的な流れ

```mermaid
sequenceDiagram
    participant U as 田中さん
    participant F as フロントエンド<br/>（ブラウザ）
    participant B as バックエンド<br/>（Goサーバー）
    participant C as Clerk
    participant D as データベース
    
    Note over U,D: 1. 初回ログイン
    U->>F: メール入力
    F->>C: ログイン要求
    C->>C: メール認証
    C->>F: セッショントークン
    F->>U: ログイン完了
    
    Note over U,D: 2. 朝活チェックイン
    U->>F: 「今日も参加！」ボタン
    F->>B: トークン付きでAPI呼び出し
    B->>C: トークン有効性確認
    C->>B: 有効（田中さんのID返却）
    B->>D: 田中さんの出席記録保存
    D->>B: 保存完了
    B->>F: チェックイン成功
    F->>U: 「お疲れ様！」表示
```

### 🔄 セッション継続の仕組み

```mermaid
graph TD
    A[田中さんログイン] --> B[Clerkトークン取得]
    B --> C[ブラウザに保存]
    
    D[翌日アクセス] --> E{トークン有効？}
    E -->|はい| F[そのまま使用可能]
    E -->|いいえ| G[再ログイン要求]
    
    F --> H[朝活データ表示]
    G --> I[メール認証]
    I --> J[新しいトークン]
    J --> H
    
    style C fill:#e8f5e8
    style F fill:#e8f5e8
    style H fill:#e1f5fe
```

---

## 5. 実装の詳細

### 🏗️ Ghoona Campのアーキテクチャ

```mermaid
graph TB
    subgraph "フロントエンド（Next.js）"
        A[ログインページ]
        B[朝活ダッシュボード]
        C[プロフィール]
    end
    
    subgraph "バックエンド（Go）"
        D[認証ミドルウェア]
        E[APIハンドラー]
        F[Clerk認証サービス]
    end
    
    subgraph "外部サービス"
        G[Clerk認証サーバー]
        H[PostgreSQL DB]
    end
    
    A --> D
    B --> D
    C --> D
    
    D --> F
    F --> G
    E --> H
    
    style D fill:#ffebee
    style F fill:#e3f2fd
    style G fill:#e8f5e8
```

### 🔍 実際のコードの流れ

**1. ユーザーがAPIを呼び出す**
```go
// フロントエンドから
fetch('/api/v1/attendance', {
    headers: {
        'Authorization': 'Bearer clerk_session_token_here'
    }
})
```

**2. ミドルウェアが認証をチェック**
```go
// backend/internal/interface/middleware/auth.go
func AuthMiddleware(authService *clerk.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Authorizationヘッダー取得
        authHeader := c.GetHeader("Authorization")
        
        // 2. Bearer トークン抽出
        token := strings.TrimPrefix(authHeader, "Bearer ")
        
        // 3. Clerkで検証
        user, err := authService.VerifyToken(c.Request.Context(), token)
        
        // 4. ユーザー情報をコンテキストに保存
        c.Set("user", user)
        c.Set("user_id", user.ID)
    }
}
```

**3. Clerkサービスがトークンを検証**
```go
// backend/internal/infrastructure/clerk/auth_service.go
func (s *AuthService) VerifyToken(ctx context.Context, token string) (*ClerkUser, error) {
    // Clerk SDK使用
    clerk.SetKey(s.secretKey)
    
    claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
        Token: token,
    })
    
    return &ClerkUser{
        ID: claims.Subject,
        // その他のユーザー情報
    }, nil
}
```

### 🛡️ セキュリティの仕組み

```mermaid
graph TD
    A[ユーザーのリクエスト] --> B[HTTPSで暗号化]
    B --> C[認証ミドルウェア]
    C --> D{Authorizationヘッダーあり？}
    D -->|なし| E[401 Unauthorized]
    D -->|あり| F{Bearer形式？}
    F -->|違う| E
    F -->|正しい| G[トークン抽出]
    G --> H[Clerk SDK検証]
    H --> I{有効なトークン？}
    I -->|無効| E
    I -->|有効| J[ユーザー情報取得]
    J --> K[APIハンドラー実行]
    
    style E fill:#ffebee
    style K fill:#e8f5e8
    style H fill:#e3f2fd
```

---

## 6. 開発者が知るべきポイント

### 💡 開発環境での動作

```mermaid
graph LR
    subgraph "開発モード"
        A[CLERK_SECRET_KEY未設定]
        B[モック認証有効]
        C["mock-clerk-token"で認証]
        D[テストユーザー自動生成]
    end
    
    subgraph "本番モード"
        E[CLERK_SECRET_KEY設定済み]
        F[実際のClerk認証]
        G[実ユーザーのトークン]
        H[セキュアな検証]
    end
    
    A --> B --> C --> D
    E --> F --> G --> H
    
    style A fill:#fff3e0
    style E fill:#e8f5e8
```

### 🔧 よくあるエラーと対処法

**1. 401 Unauthorized エラー**
```mermaid
graph TD
    A[401エラー] --> B{原因は？}
    B --> C[トークンなし]
    B --> D[トークン無効]
    B --> E[Bearer形式違い]
    
    C --> F[フロントエンドのログイン確認]
    D --> G[トークン期限切れ → 再ログイン]
    E --> H[Authorization: Bearer XXX 形式確認]
    
    style A fill:#ffebee
    style F fill:#e1f5fe
    style G fill:#e1f5fe
    style H fill:#e1f5fe
```

**2. 開発環境での認証テスト**
```bash
# curlでのテスト例
curl -H "Authorization: Bearer mock-clerk-token" \
     http://localhost:8080/api/v1/ping

# 期待される結果
{"message": "pong", "user_id": "user_mock123"}
```

### 📝 次のステップ

```mermaid
graph TD
    A[現在: 認証基盤完成] --> B[次: ユーザー管理機能]
    B --> C[ユーザープロフィール]
    B --> D[出席記録機能]
    B --> E[ランキング機能]
    
    C --> F[BE-03-user-01]
    D --> G[BE-04-attend-01]
    E --> H[BE-04-attend-05]
    
    style A fill:#e8f5e8
    style F fill:#e1f5fe
    style G fill:#e1f5fe
    style H fill:#e1f5fe
```

### 🎯 実用的なTips

**1. 認証情報の取得方法**
```go
// コントローラーでユーザー情報取得
func (h *Handler) GetMyProfile(c *gin.Context) {
    // ミドルウェアで設定されたユーザー情報を取得
    userID, exists := middleware.GetUserID(c)
    if !exists {
        c.JSON(401, gin.H{"error": "認証が必要です"})
        return
    }
    
    user, exists := middleware.GetUser(c)
    if !exists {
        c.JSON(401, gin.H{"error": "ユーザー情報が見つかりません"})
        return
    }
    
    // userIDやuserを使ってビジネスロジック実行
}
```

**2. フロントエンドでの認証状態管理**
```javascript
// Next.jsでの例
import { useAuth } from '@clerk/nextjs'

function Dashboard() {
    const { isSignedIn, user, getToken } = useAuth()
    
    if (!isSignedIn) {
        return <div>ログインしてください</div>
    }
    
    // APIコール時にトークンを含める
    const callAPI = async () => {
        const token = await getToken()
        const response = await fetch('/api/attendance', {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
    }
}
```

---

## 📚 まとめ

**認証システムは家の鍵のようなもの**。Clerkという信頼できる鍵屋さんに鍵の管理を任せることで、安全で簡単な認証システムを実現できました。

**Ghoona Campでの認証フロー**:
1. ユーザーがフロントエンドでログイン
2. Clerkがセッショントークンを発行
3. API呼び出し時にトークンを添付
4. バックエンドがClerkでトークンを検証
5. 認証済みユーザーとしてAPIを実行

この基盤の上に、ユーザー管理や出席記録などの機能を安全に構築していきます。