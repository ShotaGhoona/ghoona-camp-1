# Controller層実装戦略

## 概要
Presentation層のController実装方針。オーバーエンジニアリングを避け、機械的に実装可能な部分を明確にする。

## 実装方針

### 1. 基本構造
```go
type UserController struct {
    // TODO: UseCaseの依存注入（UseCase実装後）
}

func (uc *UserController) GetUsers(c *gin.Context) {
    // 1. パラメータ抽出（機械的）
    // 2. バリデーション（機械的）
    // 3. UseCase呼び出し（TODO）
    // 4. レスポンス生成（機械的）
}
```

### 2. 機械的に実装可能な部分

#### パラメータ抽出
- URLパラメータ: `c.Param("userId")`
- クエリパラメータ: `c.Query("search")`
- リクエストボディ: `c.ShouldBindJSON(&req)`

#### レスポンス形式統一
```go
// 成功レスポンス
c.JSON(200, gin.H{
    "data": result,
    "message": "success",
    "timestamp": time.Now().UTC().Format(time.RFC3339),
})

// エラーレスポンス
c.JSON(400, gin.H{
    "error": gin.H{
        "code": "VALIDATION_ERROR",
        "message": "Invalid request",
    },
    "timestamp": time.Now().UTC().Format(time.RFC3339),
})
```

#### HTTPステータスコード
- GET: 200 (成功), 404 (NotFound)
- POST: 201 (作成), 400 (バリデーションエラー)
- PUT: 200 (更新), 404 (NotFound)
- DELETE: 204 (削除), 404 (NotFound)

### 3. TODO部分

#### UseCase呼び出し
```go
// TODO: UseCase実装後に追加
// result, err := uc.userUseCase.GetUsers(ctx, req)
```

#### 詳細バリデーション
```go
// TODO: 複雑なビジネスルール検証
// - 日付範囲チェック
- 文字数制限
- 権限チェック（基本認証はmiddlewareで実装済み）
```

## 実装順序

### Phase 1: 構造体定義
- 各Controller構造体を空で作成
- import エラーを解決

### Phase 2: 基本ハンドラー
- パラメータ抽出
- レスポンス生成
- エラーハンドリング

### Phase 3: UseCase連携
- UseCase実装完了後
- 実際のビジネスロジック呼び出し

## ファイル構成

```
internal/presentation/controller/
├── user_controller.go
├── goal_controller.go
├── event_controller.go
├── title_controller.go
├── attendance_controller.go
├── notification_controller.go
└── types.go (共通型定義)
```

## 共通型定義

### リクエスト型
- クエリパラメータ用構造体
- ページネーション用構造体

### レスポンス型
- 統一レスポンス形式
- エラーレスポンス形式

## 注意点

### オーバーエンジニアリング回避
- 複雑な抽象化は避ける
- 共通処理は最小限に留める
- 早期最適化を避ける

### 一貫性保持
- 全Controllerで同じパターンを使用
- レスポンス形式を統一
- エラーハンドリングを統一

### 拡張性考慮
- UseCase インターフェース対応
- テスト可能な構造
- 依存注入対応