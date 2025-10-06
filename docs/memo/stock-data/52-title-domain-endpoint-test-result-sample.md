# Title Domain Endpoint Test Result Sample

## Overview
称号管理機能のエンドポイントテスト結果サンプル集です。
実際のAPIレスポンスをPostmanでテストして記録します。

## Test Environment
- **Base URL**: `http://localhost:8080/api/v1`
- **Authorization**: `Bearer {clerk_token}`
- **Content-Type**: `application/json`

---

## 1. Title Management Endpoints

### 1.1 GET /titles - 称号一覧取得

#### Request
```http
GET /api/v1/titles
Authorization: Bearer {clerk_token}
```

#### Query Parameters Test Cases
```http
# Case 1: 基本取得（アクティブのみ）
GET /api/v1/titles

# Case 2: 非アクティブも含める
GET /api/v1/titles?include_inactive=true
```


#### Actual Response
```json
{
    "data": [
        {
            "id": "b43a2aa6-ac14-4355-97e1-c6721df6080c",
            "level": 1,
            "nameJp": "まどろみ見習い",
            "nameEn": "Sleeper",
            "description": "朝活の世界への第一歩。まだ眠りの世界から抜け出せないあなたに贈る称号です。",
            "requiredDays": 1,
            "imageUrl": null,
            "colorTheme": "#8B5A3C",
            "isActive": true,
            "createdAt": "2025-09-28T19:13:01.948215+09:00",
            "updatedAt": "2025-09-28T19:13:01.948215+09:00"
        },
        {
            "id": "1164fb37-9ac0-49ff-afa6-dc9051947590",
            "level": 2,
            "nameJp": "夜明け前の戦士",
            "nameEn": "Dawn Warrior",
            "description": "暗闇の中でも立ち上がる勇気を見せたあなた。夜明け前の静寂を味方にする戦士です。",
            "requiredDays": 7,
            "imageUrl": null,
            "colorTheme": "#4A5568",
            "isActive": true,
            "createdAt": "2025-09-28T19:13:01.948215+09:00",
            "updatedAt": "2025-09-28T19:13:01.948215+09:00"
        },
        {
            "id": "edcc23e1-f5a4-493d-a4b7-e352f96089f7",
            "level": 3,
            "nameJp": "朝日の使者",
            "nameEn": "Sunrise Messenger",
            "description": "朝日と共に動き出すあなたは、新しい一日の使者として認められました。",
            "requiredDays": 14,
            "imageUrl": null,
            "colorTheme": "#F6AD55",
            "isActive": true,
            "createdAt": "2025-09-28T19:13:01.948215+09:00",
            "updatedAt": "2025-09-28T19:13:01.948215+09:00"
        },
        {
            "id": "efbad72e-fb40-44f2-ba4d-36fa596e5542",
            "level": 4,
            "nameJp": "早起きの達人",
            "nameEn": "Early Bird Master",
            "description": "早起きのコツを掴み、朝の時間を有効活用できるようになった達人です。",
            "requiredDays": 30,
            "imageUrl": null,
            "colorTheme": "#48BB78",
            "isActive": true,
            "createdAt": "2025-09-28T19:13:01.948215+09:00",
            "updatedAt": "2025-09-28T19:13:01.948215+09:00"
        },
        {
            "id": "27f7b73d-67d0-4852-a5a8-6be4bb526858",
            "level": 5,
            "nameJp": "朝活の戦士",
            "nameEn": "Morning Warrior",
            "description": "継続的な朝活により、真の朝活戦士として認定されました。",
            "requiredDays": 60,
            "imageUrl": null,
            "colorTheme": "#4299E1",
            "isActive": true,
            "createdAt": "2025-09-28T19:13:01.948215+09:00",
            "updatedAt": "2025-09-28T19:13:01.948215+09:00"
        },
        {
            "id": "1e64d24d-0dfe-4678-811c-ff5a9a099e30",
            "level": 6,
            "nameJp": "黎明の守護者",
            "nameEn": "Dawn Guardian",
            "description": "朝活コミュニティを支える存在として、黎明を守護する者です。",
            "requiredDays": 120,
            "imageUrl": null,
            "colorTheme": "#9F7AEA",
            "isActive": true,
            "createdAt": "2025-09-28T19:13:01.948215+09:00",
            "updatedAt": "2025-09-28T19:13:01.948215+09:00"
        },
        {
            "id": "868c150c-2bae-4ed2-974b-0724000c9605",
            "level": 7,
            "nameJp": "朝の賢者",
            "nameEn": "Morning Sage",
            "description": "長年の朝活経験により培われた智慧を持つ、朝の賢者です。",
            "requiredDays": 200,
            "imageUrl": null,
            "colorTheme": "#F56565",
            "isActive": true,
            "createdAt": "2025-09-28T19:13:01.948215+09:00",
            "updatedAt": "2025-09-28T19:13:01.948215+09:00"
        },
        {
            "id": "f9a7b4d7-6995-480b-9550-6859c2555a22",
            "level": 8,
            "nameJp": "朝活の伝説",
            "nameEn": "Legend of Morning",
            "description": "朝活界の伝説的存在。あなたの継続力は多くの人に勇気を与えています。",
            "requiredDays": 365,
            "imageUrl": null,
            "colorTheme": "#FFD700",
            "isActive": true,
            "createdAt": "2025-09-28T19:13:01.948215+09:00",
            "updatedAt": "2025-09-28T19:13:01.948215+09:00"
        }
    ],
    "message": "success",
    "timestamp": "2025-10-05T03:56:50Z"
}
```

---

### 1.2 GET /titles/{titleId} - 称号詳細取得

#### Request
```http
GET /api/v1/titles/{titleId}
Authorization: Bearer {clerk_token}
```

#### Test Cases
```http
# Case 1: 存在する称号ID
GET /api/v1/titles/550e8400-e29b-41d4-a716-446655440001

# Case 2: 存在しない称号ID
GET /api/v1/titles/550e8400-e29b-41d4-a716-446655440999

# Case 3: 無効なUUID形式
GET /api/v1/titles/invalid-uuid
```

#### Actual Response (Success)
```json
{
    "data": {
        "id": "b43a2aa6-ac14-4355-97e1-c6721df6080c",
        "level": 1,
        "nameJp": "まどろみ見習い",
        "nameEn": "Sleeper",
        "description": "朝活の世界への第一歩。まだ眠りの世界から抜け出せないあなたに贈る称号です。",
        "requiredDays": 1,
        "imageUrl": null,
        "colorTheme": "#8B5A3C",
        "isActive": true,
        "createdAt": "2025-09-28T19:13:01.948215+09:00",
        "updatedAt": "2025-09-28T19:13:01.948215+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-05T04:10:36Z"
}
```

---

## 2. User Title Achievement Endpoints

### 2.1 GET /users/{userId}/achievements - ユーザー称号獲得履歴取得

#### Request
```http
GET /api/v1/users/{userId}/achievements
Authorization: Bearer {clerk_token}
```

#### Query Parameters Test Cases
```http
# Case 1: 基本取得
GET /api/v1/users/{userId}/achievements

# Case 2: レベル順ソート（昇順）
GET /api/v1/users/{userId}/achievements?sort_by=level&order=asc

# Case 3: レベル順ソート（降順）
GET /api/v1/users/{userId}/achievements?sort_by=level&order=desc

# Case 4: 獲得日時順ソート（最新順）
GET /api/v1/users/{userId}/achievements?sort_by=achieved_at&order=desc

# Case 5: 獲得日時順ソート（古い順）
GET /api/v1/users/{userId}/achievements?sort_by=achieved_at&order=asc
```


#### Actual Response
```json
{
    "data": {
        "user": {
            "id": "5298ae80-faeb-4f65-8894-695b652f8115",
            "displayName": "Updated Display Name",
            "username": "UpdatedUser_1759559796528",
            "avatarUrl": "https://example.com/updated-avatar.jpg"
        },
        "achievements": [
            {
                "id": "a2b8b0fc-6a5d-4095-abf2-5736db036a69",
                "userId": "5298ae80-faeb-4f65-8894-695b652f8115",
                "title": {
                    "id": "edcc23e1-f5a4-493d-a4b7-e352f96089f7",
                    "level": 3,
                    "nameJp": "朝日の使者",
                    "nameEn": "Sunrise Messenger",
                    "description": "朝日と共に動き出すあなたは、新しい一日の使者として認められました。",
                    "requiredDays": 14,
                    "imageUrl": null,
                    "colorTheme": "#F6AD55",
                    "isActive": true,
                    "createdAt": "2025-09-28T19:13:01.948215+09:00",
                    "updatedAt": "2025-09-28T19:13:01.948215+09:00"
                },
                "achievedAt": "2025-09-28T19:13:03.041582+09:00",
                "isCurrent": true,
                "createdAt": "2025-09-28T19:13:03.041582+09:00",
                "updatedAt": "2025-09-28T19:13:03.041582+09:00"
            }
        ]
    },
    "message": "success",
    "timestamp": "2025-10-05T04:14:40Z"
}
```

---

### 2.2 PUT /users/{userId}/achievements/{titleId} - 現在称号変更

#### Request
```http
PUT /api/v1/users/{userId}/achievements/{titleId}
Authorization: Bearer {clerk_token}
Content-Type: application/json
```

#### Request Body
```json
{
  "isCurrent": true
}
```

#### Test Cases
```http
# Case 1: 本人が獲得済み称号を現在称号に設定
PUT /api/v1/users/{own-userId}/achievements/{achieved-titleId}
Body: {"isCurrent": true}

# Case 2: 本人が未獲得称号を現在称号に設定（エラー）
PUT /api/v1/users/{own-userId}/achievements/{not-achieved-titleId}
Body: {"isCurrent": true}

# Case 3: 他人の称号を変更しようとする（権限エラー）
PUT /api/v1/users/{other-userId}/achievements/{titleId}
Body: {"isCurrent": true}

# Case 4: 既に現在称号に設定済み（エラー）
PUT /api/v1/users/{own-userId}/achievements/{current-titleId}
Body: {"isCurrent": true}

# Case 5: 非アクティブな称号を設定しようとする（エラー）
PUT /api/v1/users/{own-userId}/achievements/{inactive-titleId}
Body: {"isCurrent": true}
```

#### Expected Response (Success)
```json
{
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440102",
    "userId": "550e8400-e29b-41d4-a716-446655440000",
    "title": {
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "level": 1,
      "nameJp": "まどろみ見習い",
      "nameEn": "Sleeper",
      "description": "朝活の世界への第一歩。1日の参加で獲得できる称号。",
      "requiredDays": 1,
      "imageUrl": "https://cdn.ghoona-camp.com/titles/level-1.png",
      "colorTheme": "#B0BEC5",
      "isActive": true,
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z"
    },
    "achievedAt": "2024-10-15T06:30:00Z",
    "isCurrent": true,
    "createdAt": "2024-10-15T06:30:00Z",
    "updatedAt": "2025-10-05T03:45:17Z"
  },
  "message": "success",
  "timestamp": "2025-10-05T03:45:17Z"
}
```

#### Actual Response (Success)
```json
// ここに実際のレスポンスを記録
```

---

## 3. Error Response Test Cases

### 3.1 Authentication Errors

#### No Authorization Header
```http
GET /api/v1/titles
# No Authorization header
```

#### Expected Response
```json
{
  "error": {
    "code": "UNAUTHORIZED",
    "message": "Authentication required"
  },
  "timestamp": "2025-10-05T03:45:17Z"
}
```

#### Actual Response
```json
// ここに実際のレスポンスを記録
```

### 3.2 Permission Errors

#### Forbidden Access (Other User's Resource)
```http
PUT /api/v1/users/{other-userId}/achievements/{titleId}
Authorization: Bearer {clerk_token}
```

#### Expected Response
```json
{
  "error": {
    "code": "FORBIDDEN",
    "message": "Access denied: can only access own resources"
  },
  "timestamp": "2025-10-05T03:45:17Z"
}
```

#### Actual Response
```json
// ここに実際のレスポンスを記録
```

### 3.3 Validation Errors

#### Invalid UUID Format
```http
GET /api/v1/titles/invalid-uuid
Authorization: Bearer {clerk_token}
```

#### Expected Response
```json
{
  "error": {
    "code": "INVALID_TITLE_ID",
    "message": "invalid UUID format: titleId"
  },
  "timestamp": "2025-10-05T03:45:17Z"
}
```

#### Actual Response
```json
// ここに実際のレスポンスを記録
```

### 3.4 Business Logic Errors

#### Title Not Achieved
```http
PUT /api/v1/users/{userId}/achievements/{not-achieved-titleId}
Authorization: Bearer {clerk_token}
Content-Type: application/json

{
  "isCurrent": true
}
```

#### Expected Response
```json
{
  "error": {
    "code": "TITLE_NOT_ACHIEVED",
    "message": "未獲得の称号です"
  },
  "timestamp": "2025-10-05T03:45:17Z"
}
```

#### Actual Response
```json
// ここに実際のレスポンスを記録
```

#### Title Already Current
```http
PUT /api/v1/users/{userId}/achievements/{current-titleId}
Authorization: Bearer {clerk_token}
Content-Type: application/json

{
  "isCurrent": true
}
```

#### Expected Response
```json
{
  "error": {
    "code": "TITLE_ALREADY_CURRENT",
    "message": "既に現在の称号に設定済み"
  },
  "timestamp": "2025-10-05T03:45:17Z"
}
```

#### Actual Response
```json
// ここに実際のレスポンスを記録
```

---

## 4. Performance Test Notes

### Response Time Benchmarks
```
GET /titles: 目標 < 100ms
GET /titles/{titleId}: 目標 < 50ms
GET /users/{userId}/achievements: 目標 < 200ms
PUT /users/{userId}/achievements/{titleId}: 目標 < 300ms
```

### Load Test Results
```
// 同時接続数、レスポンス時間等の負荷テスト結果を記録
```

---

## 5. Integration Test Scenarios

### Scenario 1: 新規ユーザーの称号表示
1. 新規ユーザーでログイン
2. `GET /users/{userId}/achievements` → 空の配列
3. 称号一覧確認 `GET /titles`

### Scenario 2: 称号変更フロー（テストデータが必要）
1. 複数称号を獲得済みのユーザーでログイン
2. 現在の称号確認
3. 別の称号に変更
4. 変更結果確認

### Scenario 3: エラーハンドリング確認
1. 無効なリクエスト送信
2. 適切なエラーレスポンス確認
3. ログにエラー記録確認

---

## 6. Test Summary

### Test Coverage
- [ ] GET /titles (基本機能)
- [ ] GET /titles (クエリパラメータ)
- [ ] GET /titles/{titleId} (成功ケース)
- [ ] GET /titles/{titleId} (エラーケース)
- [ ] GET /users/{userId}/achievements (基本機能)
- [ ] GET /users/{userId}/achievements (ソート機能)
- [ ] PUT /users/{userId}/achievements/{titleId} (成功ケース)
- [ ] PUT /users/{userId}/achievements/{titleId} (エラーケース)
- [ ] Authentication/Authorization Tests
- [ ] Validation Error Tests
- [ ] Business Logic Error Tests

### Test Results Summary
```
Total Tests: 
Passed: 
Failed: 
Coverage: %

Failed Tests:
- 

Notes:
- 
```

---

## 7. Notes & Observations

### API Behavior Notes
```
// APIの動作で気づいた点、改善点等を記録
```

### Performance Observations
```
// パフォーマンスで気づいた点を記録
```

### Issues Found
```
// テスト中に発見したバグや問題点を記録
```

### Recommendations
```
// API改善提案等を記録
```