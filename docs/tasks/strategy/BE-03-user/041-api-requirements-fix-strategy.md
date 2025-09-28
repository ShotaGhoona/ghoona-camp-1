# BE-03-user-041 API要件適合修正 - 戦略書【修正版】

## 📋 修正概要

**タスクID**: BE-03-user-041  
**作成日**: 2025-09-28（修正版）  
**目的**: プレースホルダー実装と欠落機能の完成  
**推定工数**: 30分  
**優先度**: 🟡 Medium（アーキテクチャは完成済み、細部調整のみ）

## 🎯 修正目標

既存の成熟したアーキテクチャにおいて、プレースホルダー実装の完成と要件書との最終的な適合を実現する。

## ✅ 既存実装状況（再確認済み）

### 完全実装済み
- **UseCase層**: 全メソッド実装済み（user_service.go）
- **DTO層**: 全リクエスト/レスポンス実装済み  
- **Repository層**: インターフェース＋GORM実装完了（275行）
- **Domain Services**: バリデーション＋ビジネスロジック完了
- **Controller層**: 全エンドポイント実装済み（571行）
- **Routing層**: 完全なREST API設定済み

### 実際の不足部分（プレースホルダー実装）

#### 1. GetUsersメソッドの実装不完全
- **現状**: プレースホルダー実装（空のレスポンス返却）
- **必要**: 実際のユーザー一覧取得＋検索＋ページネーション

#### 2. ユーザー登録エンドポイント欠落
- **現状**: CreateUser UseCaseは存在するが公開エンドポイントなし
- **必要**: `POST /users` 登録エンドポイント追加

#### 3. 認証統合の未完成
- **現状**: ClerkID→UUID変換がTODO状態
- **必要**: `requireSelfAccess`の完全実装

## 🏗 影響範囲ツリー【修正版】

```
backend/
├── internal/
│   ├── application/usecase/
│   │   └── user_service.go                    # ✏️ GetUsersメソッド実装完成
│   ├── infrastructure/gorm/repository/
│   │   └── user_repository.go                 # ✏️ GetAllメソッド実装（既存）
│   └── interface/
│       ├── controller/
│       │   └── user_controller.go             # ✏️ CreateUser公開エンドポイント追加
│       │                                      # ✏️ requireSelfAccess修正
│       └── router/
│           └── user_routes.go                 # ✏️ POST /usersルート追加
├── docs/tasks/strategy/BE-03-user/
│   └── 041-api-requirements-fix-strategy.md   # ✏️ この戦略書（修正版）
└── docs/tasks/report/BE-03-user/
    └── 041-api-requirements-fix-report.md     # 🆕 修正完了レポート（予定）
```

**注意**: DTOやRepositoryインターフェースは既に完全実装済みのため変更不要

## 📚 修正実装計画【簡素化版】

### Task 1: GetUsersメソッド実装完成 【HIGH】
**推定時間**: 15分

**現状確認**:
```go
// 現在のプレースホルダー実装
func (u *userUseCase) GetUsers(ctx context.Context) (*user.UserListResponse, error) {
    // PLACEHOLDER IMPLEMENTATION
    return &user.UserListResponse{
        Users: []user.UserResponse{},
        Total: 0,
    }, nil
}
```

**実装内容**:
- 既存のRepository `GetAll`メソッドを活用
- 既存のDTO `UserListResponse`を活用
- ページネーション・検索機能は既存DTOで対応済み

### Task 2: ユーザー登録エンドポイント追加 【MEDIUM】
**推定時間**: 10分

**実装内容**:
```go
// user_routes.go に追加
usersGroup.POST("", r.container.UserController.CreateUser)

// アクセス制御: 🔐 認証済みユーザー（要件書準拠）
```

### Task 3: ClerkID→UUID変換機能完成 【MEDIUM】
**推定時間**: 5分

**現状**:
```go
// requireSelfAccess内のTODO
// TODO: ClerkIDから内部UUIDへの変換を実装する必要がある
```

**実装内容**:
- 既存のUserRepository `GetByClerkID`メソッドを活用
- 認証ユーザーの内部UUID取得ロジック実装

## 🔧 実装詳細【既存コード活用】

### Task 1実装: GetUsersメソッド
```go
// user_service.go 修正
func (u *userUseCase) GetUsers(ctx context.Context) (*user.UserListResponse, error) {
    // 既存のGetAllメソッドを活用
    users, err := u.userRepo.GetAll(ctx, repository.UserFilters{})
    if err != nil {
        return nil, err
    }
    
    // 既存DTOに変換
    userResponses := make([]user.UserResponse, len(users))
    for i, usr := range users {
        userResponses[i] = user.UserResponse{
            ID:       usr.ID.String(),
            Username: usr.Username,
            Email:    usr.Email,
            // ... 既存の変換ロジック
        }
    }
    
    return &user.UserListResponse{
        Users: userResponses,
        Total: len(users),
    }, nil
}
```

### Task 2実装: ユーザー登録ルート
```go
// user_routes.go に1行追加
usersGroup.POST("", r.container.UserController.CreateUser)
```

### Task 3実装: ClerkID変換
```go
// user_controller.go の requireSelfAccess修正
func requireSelfAccess(ctx *gin.Context, targetUserID uuid.UUID) bool {
    clerkID, err := getCurrentClerkID(ctx)
    if err != nil {
        // エラーレスポンス処理（既存）
        return false
    }
    
    // 既存のGetByClerkIDメソッドを活用
    currentUser, err := userRepo.GetByClerkID(ctx.Request.Context(), clerkID)
    if err != nil || currentUser.ID != targetUserID {
        // 権限エラー
        return false
    }
    
    return true
}
```

## ✅ 実装完了確認

### 確認項目
- [ ] GetUsersメソッドがプレースホルダーでなく実際のデータを返す
- [ ] `POST /users` エンドポイントが追加されている
- [ ] ClerkID→UUID変換が機能する
- [ ] 要件書の全エンドポイントが正常動作する

## ⚠️ 最小限のリスク

### 1. 既存システムとの互換性
- **影響**: 既存のエンドポイント動作は全て維持
- **対策**: 新規追加と内部実装のみ修正

### 2. 認証統合
- **影響**: ClerkID変換ロジック追加でわずかな処理追加
- **対策**: 既存のRepository活用で最小限の変更

## 🔄 次のステップ【簡素化】

1. **Task 1**: GetUsersメソッド実装完成（15分）
2. **Task 2**: POST /usersルート追加（5分）  
3. **Task 3**: ClerkID変換実装（10分）
4. **レポート作成**: 041完了レポート作成

**総工数**: 約30分（元の2時間から大幅短縮）

---

**戦略書修正完了**: BE-03-user-041 API要件適合修正【修正版】  
**実際の作業**: 既存アーキテクチャでの最小限修正のみ