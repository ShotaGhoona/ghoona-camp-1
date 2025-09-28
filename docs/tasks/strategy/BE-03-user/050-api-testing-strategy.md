# BE-03-user-05 API動作確認戦略書

## 概要
エラーハンドリングバグ修正後のユーザー管理API全体の動作確認を行い、BE-03-user-05を完了させる。

## 目的
1. **修正したエラーハンドリングの動作確認**
2. **API仕様書通りの動作検証**
3. **エッジケース・エラーケースの網羅確認**
4. **本番運用前の品質保証**

## テスト戦略

### 🛠️ **テストツール選定**

**メインツール: Postman** ✅
- **理由**: 
  - GUI で直感的
  - コレクション機能でテストケース管理
  - 環境変数でベースURL・認証トークン管理
  - レスポンス検証スクリプト実行可能
  - チーム共有可能

**補完ツール**:
- **curl**: 単発テスト用
- **httpie**: 見やすい出力確認用

### 📋 **テスト対象API一覧**

| カテゴリ | メソッド | エンドポイント | 優先度 |
|---------|---------|----------------|--------|
| **認証** | GET | `/api/v1/auth/me` | 🔴高 |
| **User基本** | GET | `/api/v1/users` | 🔴高 |
| **User基本** | POST | `/api/v1/users` | 🔴高 |
| **User基本** | GET | `/api/v1/users/{userId}` | 🔴高 |
| **User基本** | PUT | `/api/v1/users/{userId}` | 🔴高 |
| **Metadata** | GET | `/api/v1/users/{userId}/metadata` | 🟡中 |
| **Metadata** | POST | `/api/v1/users/{userId}/metadata` | 🟡中 |
| **Metadata** | PUT | `/api/v1/users/{userId}/metadata` | 🟡中 |
| **Social** | GET | `/api/v1/users/{userId}/social-links` | 🟢低 |
| **Social** | POST | `/api/v1/users/{userId}/social-links` | 🟢低 |
| **Social** | PUT | `/api/v1/users/{userId}/social-links/{linkId}` | 🟢低 |
| **Social** | DELETE | `/api/v1/users/{userId}/social-links/{linkId}` | 🟢低 |
| **Rival** | GET | `/api/v1/users/{userId}/rivals` | 🟢低 |
| **Rival** | POST | `/api/v1/users/{userId}/rivals` | 🟢低 |
| **Rival** | DELETE | `/api/v1/users/{userId}/rivals/{rivalId}` | 🟢低 |

### 🎯 **テストケース設計**

#### **1. 正常系テスト (Happy Path)**

```
Postmanコレクション: "User Management - Happy Path"

1. CreateUser → GetUser → UpdateUser → GetUsers の流れ
2. CreateMetadata → GetMetadata → UpdateMetadata の流れ  
3. CreateSocialLink → GetSocialLinks → UpdateSocialLink → DeleteSocialLink の流れ
4. AddRival → GetRivals → RemoveRival の流れ
```

#### **2. エラー系テスト (Error Cases)**

```
Postmanコレクション: "User Management - Error Cases"

🔴 最重要: 修正したエラーハンドリング確認
- ClerkID重複エラー (ErrDuplicateClerkID)
- メタデータ重複エラー (ErrUserMetadataAlreadyExists)
- ユーザー未存在エラー (ErrUserNotFound)

🟡 重要: バリデーションエラー  
- 不正なUUID形式
- 必須フィールド欠如
- 文字数制限違反

🟢 補完: その他エラー
- 認証エラー
- 権限エラー
- 不正なJSON形式
```

#### **3. エッジケーステスト**

```
- 空文字・null値の処理
- 特殊文字・絵文字入力
- 極端に長い文字列
- 存在しないリソースへのアクセス
- 権限境界テスト（他人のデータアクセス）
```

## 📝 **Postman設定手順**

### **Step 1: 環境設定**
```json
{
  "name": "Ghoona Camp Local",
  "values": [
    {
      "key": "BASE_URL",
      "value": "http://localhost:8080/api/v1"
    },
    {
      "key": "CLERK_TOKEN", 
      "value": "{{取得したJWTトークン}}"
    },
    {
      "key": "TEST_USER_ID",
      "value": "{{動的に設定}}"
    }
  ]
}
```

### **Step 2: 認証設定**
```javascript
// Collection-level Auth Script
pm.request.headers.add({
    key: 'Authorization',
    value: 'Bearer ' + pm.environment.get('CLERK_TOKEN')
});
```

### **Step 3: テストスクリプトテンプレート**
```javascript
// 正常系テスト
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Response has required fields", function () {
    const responseJson = pm.response.json();
    pm.expect(responseJson).to.have.property('data');
    pm.expect(responseJson).to.have.property('message', 'success');
    pm.expect(responseJson).to.have.property('timestamp');
});

// エラー系テスト  
pm.test("Status code is 409 for duplicate", function () {
    pm.response.to.have.status(409);
});

pm.test("Error response format is correct", function () {
    const responseJson = pm.response.json();
    pm.expect(responseJson).to.have.property('error');
    pm.expect(responseJson.error).to.have.property('code');
    pm.expect(responseJson.error).to.have.property('message');
});

// 修正したエラーハンドリング確認
pm.test("ClerkID duplicate returns correct error", function () {
    const responseJson = pm.response.json();
    pm.expect(responseJson.error.code).to.eql('DUPLICATE_CLERK_ID');
});
```

## 🔍 **重点確認項目**

### **1. エラーハンドリング修正確認** 🔴最重要

```bash
# Test Case 1: ClerkID重複エラー
POST /api/v1/users
{
  "clerkId": "existing_clerk_id",
  "email": "new@example.com"
}
期待結果: 409 Conflict, code: "DUPLICATE_CLERK_ID"

# Test Case 2: メタデータ重複エラー  
POST /api/v1/users/{userId}/metadata
(既にメタデータ存在するユーザー)
期待結果: 409 Conflict, code: "USER_METADATA_ALREADY_EXISTS"

# Test Case 3: メタデータ取得エラー時のログ確認
GET /api/v1/users/{userId}
(メタデータが壊れているユーザー)
期待結果: 200 OK, ログに警告出力
```

### **2. APIレスポンス一貫性確認**

```javascript
// 成功レスポンス形式
{
  "data": { /* actual data */ },
  "message": "success",
  "timestamp": "2025-01-21T12:00:00Z"
}

// エラーレスポンス形式
{
  "error": {
    "code": "ERROR_CODE",
    "message": "エラーメッセージ"
  },
  "timestamp": "2025-01-21T12:00:00Z"
}
```

### **3. 認証・認可確認**

```bash
# 未認証アクセス
Authorization ヘッダーなしでリクエスト
期待結果: 401 Unauthorized

# 他人のデータアクセス
GET /api/v1/users/{otherUserId}/metadata
期待結果: 403 Forbidden
```

## 📊 **テスト実行計画**

### **Phase 1: 基本動作確認** (30分)
1. サーバー起動確認
2. 認証機能確認
3. User基本CRUD動作確認

### **Phase 2: エラーハンドリング確認** (45分)
1. 修正したエラーケース全確認
2. バリデーションエラー確認
3. ログ出力確認

### **Phase 3: 総合シナリオ確認** (30分)
1. ユーザー登録→プロフィール設定→更新の一連フロー
2. エラー発生→正常復帰フロー
3. 権限チェック

### **Phase 4: レポート作成** (15分)
1. テスト結果まとめ
2. 発見した問題の記録
3. 改善提案

## 📈 **成功基準**

### **必須項目 (Must Have)**
- ✅ 修正したエラーハンドリングが正しく動作
- ✅ 全ての正常系APIが仕様通り動作
- ✅ 認証・認可が適切に機能
- ✅ レスポンス形式が一貫している

### **望ましい項目 (Nice to Have)**
- ✅ エッジケースも適切に処理
- ✅ ログが適切に出力される
- ✅ パフォーマンスが許容範囲内

## 🚀 **次のアクション**

1. **Postmanコレクション作成** (15分)
2. **テスト実行** (2時間)
3. **問題修正** (必要に応じて)
4. **BE-03-user-05完了報告**

## 📝 **成果物**

1. **Postmanコレクション** (.json)
2. **テスト実行結果レポート** (.md)
3. **発見した問題リスト** (あれば)
4. **API動作確認完了証明**

---

この戦略に従ってテストを実行することで、BE-03-user-05の「統合・動作確認」が完了し、ユーザー管理機能の品質が保証されます。