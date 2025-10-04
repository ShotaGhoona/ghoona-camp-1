# User Domain API Endpoint Test Results

## Test Environment
- **Base URL**: `http://localhost:8080`
- **Authentication**: Mock Token (`Bearer mock-clerk-token`)
- **Test Tool**: Postman
- **Test Date**: 2025-10-04

## Authentication Settings
```
Authorization: Bearer mock-clerk-token
Content-Type: application/json
```

---

## 1. GET /api/v1/auth/me
**Description**: Get current user information

### Request
```
GET http://localhost:8080/api/v1/auth/me
Headers:
  Authorization: Bearer mock-clerk-token
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "5298ae80-faeb-4f65-8894-695b652f8115",
        "clerk_id": "clerk_dev_user_001",
        "email": "developer@ghoona.camp",
        "username": "Ghoona Developer",
        "avatar_url": "https://avatars.githubusercontent.com/u/1?v=4",
        "discord_id": null,
        "status": "active",
        "metadata": {
            "id": "fca64cbc-59ab-419b-a8c3-25b608ec631d",
            "user_id": "5298ae80-faeb-4f65-8894-695b652f8115",
            "display_name": "Ghoona Developer",
            "profile_image_url": null,
            "tagline": "朝活で人生を変える開発者",
            "bio": "Ghoona Campの開発を通じて、朝活コミュニティの価値を最大化することを目指しています。毎朝6時から開発作業を行い、生産性の高い一日をスタートしています。",
            "vision": "朝活を通じて、多くの人が充実した人生を送れる世界を作りたい。テクノロジーの力で朝活習慣を支援し、継続できる仕組みを構築する。",
            "vision_public": true,
            "timezone": "Asia/Tokyo",
            "skills": [
                "Go",
                "React",
                "TypeScript",
                "Docker",
                "AWS",
                "Clean Architecture"
            ],
            "interests": [
                "朝活",
                "プログラミング",
                "コミュニティ",
                "健康",
                "ライフハック"
            ],
            "created_at": "2025-09-28T19:13:02.192692+09:00",
            "updated_at": "2025-09-28T19:13:02.192692+09:00"
        },
        "created_at": "2025-09-28T19:13:02.104918+09:00",
        "updated_at": "2025-09-28T19:13:02.104918+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:00:22Z"
}
```

---

## 2. GET /api/v1/users
**Description**: Get list of users

### Request
```
GET http://localhost:8080/api/v1/users
Headers:
  Authorization: Bearer mock-clerk-token
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "users": [
            {
                "id": "879fb62c-ee02-4ff0-a1cb-005190b2c47e",
                "clerk_id": "test_clerk_001",
                "email": "test1@example.com",
                "username": "Test User 1",
                "avatar_url": "https://example.com/avatar1.jpg",
                "discord_id": null,
                "status": "active",
                "created_at": "2025-09-28T19:13:20.23238+09:00",
                "updated_at": "2025-09-28T19:13:20.23238+09:00"
            },
            {
                "id": "390b77a6-54f9-45a1-9d43-0e3142797c9b",
                "clerk_id": "test_clerk_002",
                "email": "test2@example.com",
                "username": "Test User 2",
                "avatar_url": "https://example.com/avatar2.jpg",
                "discord_id": null,
                "status": "active",
                "created_at": "2025-09-28T19:13:20.23238+09:00",
                "updated_at": "2025-09-28T19:13:20.23238+09:00"
            },
            {
                "id": "39db328a-5db9-47a2-a9f9-6721ee9d9b31",
                "clerk_id": "clerk_dev_user_003",
                "email": "designer@ghoona.camp",
                "username": "UI Designer",
                "avatar_url": "https://avatars.githubusercontent.com/u/3?v=4",
                "discord_id": null,
                "status": "active",
                "created_at": "2025-09-28T19:13:02.104918+09:00",
                "updated_at": "2025-09-28T19:13:02.104918+09:00"
            },
            {
                "id": "d48bfee7-3c73-4fa4-b90a-ed5be72a8750",
                "clerk_id": "clerk_dev_user_004",
                "email": "pm@ghoona.camp",
                "username": "Product Manager",
                "avatar_url": "https://avatars.githubusercontent.com/u/4?v=4",
                "discord_id": null,
                "status": "active",
                "created_at": "2025-09-28T19:13:02.104918+09:00",
                "updated_at": "2025-09-28T19:13:02.104918+09:00"
            },
            {
                "id": "5298ae80-faeb-4f65-8894-695b652f8115",
                "clerk_id": "clerk_dev_user_001",
                "email": "developer@ghoona.camp",
                "username": "Ghoona Developer",
                "avatar_url": "https://avatars.githubusercontent.com/u/1?v=4",
                "discord_id": null,
                "status": "active",
                "created_at": "2025-09-28T19:13:02.104918+09:00",
                "updated_at": "2025-09-28T19:13:02.104918+09:00"
            },
            {
                "id": "93cc832a-897c-407f-874a-b4473e787fd4",
                "clerk_id": "clerk_test_user_2",
                "email": "test2@ghoona.camp",
                "username": "テストユーザー2",
                "avatar_url": "https://example.com/avatar2.jpg",
                "discord_id": null,
                "status": "active",
                "created_at": "2025-09-28T19:13:02.104918+09:00",
                "updated_at": "2025-09-28T19:13:02.104918+09:00"
            },
            {
                "id": "fc96cf8a-40d5-45a1-ae40-51d042125e4b",
                "clerk_id": "clerk_test_user_1",
                "email": "test1@ghoona.camp",
                "username": "テストユーザー1",
                "avatar_url": "https://example.com/avatar1.jpg",
                "discord_id": null,
                "status": "active",
                "created_at": "2025-09-28T19:13:02.104918+09:00",
                "updated_at": "2025-09-28T19:13:02.104918+09:00"
            },
            {
                "id": "b0e2d712-002a-4073-85f2-4daefcddf249",
                "clerk_id": "clerk_dev_user_002",
                "email": "tester@ghoona.camp",
                "username": "Test Expert",
                "avatar_url": "https://avatars.githubusercontent.com/u/2?v=4",
                "discord_id": null,
                "status": "active",
                "created_at": "2025-09-28T19:13:02.104918+09:00",
                "updated_at": "2025-09-28T19:13:02.104918+09:00"
            }
        ],
        "total": 8
    },
    "message": "success",
    "timestamp": "2025-10-04T03:59:59Z"
}
```

---

## 3. GET /api/v1/users/{userId}
**Description**: Get specific user information

### Request
```
GET http://localhost:8080/api/v1/users/{userId}
Headers:
  Authorization: Bearer mock-clerk-token
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "5298ae80-faeb-4f65-8894-695b652f8115",
        "clerk_id": "clerk_dev_user_001",
        "email": "developer@ghoona.camp",
        "username": "Ghoona Developer",
        "avatar_url": "https://avatars.githubusercontent.com/u/1?v=4",
        "discord_id": null,
        "status": "active",
        "metadata": {
            "id": "fca64cbc-59ab-419b-a8c3-25b608ec631d",
            "user_id": "5298ae80-faeb-4f65-8894-695b652f8115",
            "display_name": "Ghoona Developer",
            "profile_image_url": null,
            "tagline": "朝活で人生を変える開発者",
            "bio": "Ghoona Campの開発を通じて、朝活コミュニティの価値を最大化することを目指しています。毎朝6時から開発作業を行い、生産性の高い一日をスタートしています。",
            "vision": "朝活を通じて、多くの人が充実した人生を送れる世界を作りたい。テクノロジーの力で朝活習慣を支援し、継続できる仕組みを構築する。",
            "vision_public": true,
            "timezone": "Asia/Tokyo",
            "skills": [
                "Go",
                "React",
                "TypeScript",
                "Docker",
                "AWS",
                "Clean Architecture"
            ],
            "interests": [
                "朝活",
                "プログラミング",
                "コミュニティ",
                "健康",
                "ライフハック"
            ],
            "created_at": "2025-09-28T19:13:02.192692+09:00",
            "updated_at": "2025-09-28T19:13:02.192692+09:00"
        },
        "created_at": "2025-09-28T19:13:02.104918+09:00",
        "updated_at": "2025-09-28T19:13:02.104918+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:12:20Z"
}
```

---

## 4. POST /api/v1/users
**Description**: Create a new user

### Request
```
POST http://localhost:8080/api/v1/users
Headers:
  Authorization: Bearer mock-clerk-token
  Content-Type: application/json
```

**Request Body**:
```json
{
    "clerk_id": "clerk_new_user_001",
    "email": "newuser@example.com",
    "username": "New User",
    "avatar_url": "https://example.com/avatar.jpg",
    "discord_id": "discord_123456"
}
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "30f173da-65f1-4c61-a0a8-552da28406b3",
        "clerk_id": "clerk_new_user_001",
        "email": "newuser@example.com",
        "username": "New User",
        "avatar_url": "https://example.com/avatar.jpg",
        "discord_id": "discord_123456",
        "status": "active",
        "created_at": "2025-10-04T13:17:53.735391+09:00",
        "updated_at": "2025-10-04T13:17:53.735391+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:17:54Z"
}
```

---

## 5. PUT /api/v1/users/{userId}
**Description**: Update user basic information

### Request
```
PUT http://localhost:8080/api/v1/users/{userId}
Headers:
  Authorization: Bearer mock-clerk-token
  Content-Type: application/json
```

**Request Body**:
```json
{
    "username": "Updated Username",
    "avatar_url": "https://example.com/new-avatar.jpg",
    "discord_id": "new_discord_123"
}
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "5298ae80-faeb-4f65-8894-695b652f8115",
        "clerk_id": "clerk_dev_user_001",
        "email": "developer@ghoona.camp",
        "username": "Updated Username",
        "avatar_url": "https://example.com/new-avatar.jpg",
        "discord_id": "new_discord_123",
        "status": "active",
        "created_at": "2025-09-28T19:13:02.104918+09:00",
        "updated_at": "2025-09-28T19:13:02.104918+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:19:28Z"
}
```

---

## 6. GET /api/v1/users/{userId}/metadata
**Description**: Get user metadata (extended profile info)

### Request
```
GET http://localhost:8080/api/v1/users/{userId}/metadata
Headers:
  Authorization: Bearer mock-clerk-token
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "fca64cbc-59ab-419b-a8c3-25b608ec631d",
        "user_id": "5298ae80-faeb-4f65-8894-695b652f8115",
        "display_name": "Ghoona Developer",
        "profile_image_url": null,
        "tagline": "朝活で人生を変える開発者",
        "bio": "Ghoona Campの開発を通じて、朝活コミュニティの価値を最大化することを目指しています。毎朝6時から開発作業を行い、生産性の高い一日をスタートしています。",
        "vision": "朝活を通じて、多くの人が充実した人生を送れる世界を作りたい。テクノロジーの力で朝活習慣を支援し、継続できる仕組みを構築する。",
        "vision_public": true,
        "timezone": "Asia/Tokyo",
        "skills": [
            "Go",
            "React",
            "TypeScript",
            "Docker",
            "AWS",
            "Clean Architecture"
        ],
        "interests": [
            "朝活",
            "プログラミング",
            "コミュニティ",
            "健康",
            "ライフハック"
        ],
        "created_at": "2025-09-28T19:13:02.192692+09:00",
        "updated_at": "2025-09-28T19:13:02.192692+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:12:51Z"
}
```

---

## 7. POST /api/v1/users/{userId}/metadata
**Description**: Create user metadata

### Request
```
POST http://localhost:8080/api/v1/users/{userId}/metadata
Headers:
  Authorization: Bearer mock-clerk-token
  Content-Type: application/json
```

**Request Body**:
```json
{
    "display_name": "My Display Name",
    "profile_image_url": "https://example.com/profile.jpg",
    "tagline": "A passionate developer",
    "bio": "I love building amazing applications and contributing to open source projects.",
    "vision": "To create technology that makes people's lives better.",
    "vision_public": true,
    "timezone": "Asia/Tokyo",
    "skills": ["JavaScript", "Go", "React", "Node.js"],
    "interests": ["Programming", "Technology", "Open Source", "Learning"]
}
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "error": {
        "code": "USER_METADATA_ALREADY_EXISTS",
        "message": "ユーザーメタデータが既に存在します"
    },
    "timestamp": "2025-10-04T04:20:36Z"
}
```

---

## 8. PUT /api/v1/users/{userId}/metadata
**Description**: Update user metadata

### Request
```
PUT http://localhost:8080/api/v1/users/{userId}/metadata
Headers:
  Authorization: Bearer mock-clerk-token
  Content-Type: application/json
```

**Request Body**:
```json
{
    "display_name": "Updated Display Name",
    "tagline": "Updated tagline",
    "bio": "Updated bio information.",
    "vision_public": false,
    "skills": ["Go", "React", "TypeScript", "Docker"],
    "interests": ["朝活", "Programming", "Health"]
}
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "fca64cbc-59ab-419b-a8c3-25b608ec631d",
        "user_id": "5298ae80-faeb-4f65-8894-695b652f8115",
        "display_name": "Updated Display Name",
        "profile_image_url": "https://example.com/profile.jpg",
        "tagline": "Updated tagline",
        "bio": "Updated bio information.",
        "vision": "To create technology that makes people's lives better.",
        "vision_public": false,
        "timezone": "Asia/Tokyo",
        "skills": [
            "Go",
            "React",
            "TypeScript",
            "Docker"
        ],
        "interests": [
            "朝活",
            "Programming",
            "Health"
        ],
        "created_at": "2025-09-28T19:13:02.192692+09:00",
        "updated_at": "2025-10-04T13:21:10.743519+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:21:20Z"
}
```

---

## 9. GET /api/v1/users/{userId}/social-links
**Description**: Get user's social links

### Request
```
GET http://localhost:8080/api/v1/users/{userId}/social-links
Headers:
  Authorization: Bearer mock-clerk-token
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "social_links": [
            {
                "id": "24afb786-065e-4b25-a0ba-4f812d2b4765",
                "user_id": "5298ae80-faeb-4f65-8894-695b652f8115",
                "platform": "github",
                "url": "https://github.com/ghoona-developer",
                "title": "Ghoona Camp Repository",
                "is_public": true,
                "created_at": "2025-09-28T19:13:02.450473+09:00",
                "updated_at": "2025-09-28T19:13:02.450473+09:00"
            }
        ],
        "total": 1
    },
    "message": "success",
    "timestamp": "2025-10-04T04:13:35Z"
}
```

---

## 10. POST /api/v1/users/{userId}/social-links
**Description**: Add a new social link

### Request
```
POST http://localhost:8080/api/v1/users/{userId}/social-links
Headers:
  Authorization: Bearer mock-clerk-token
  Content-Type: application/json
```

**Request Body**:
```json
{
    "platform": "twitter",
    "url": "https://twitter.com/myusername",
    "title": "My Twitter Account",
    "is_public": true
}
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "71cfa2c6-091d-4edd-a09c-388257b8a231",
        "user_id": "5298ae80-faeb-4f65-8894-695b652f8115",
        "platform": "twitter",
        "url": "https://twitter.com/myusername",
        "title": "My Twitter Account",
        "is_public": true,
        "created_at": "2025-10-04T13:21:45.832944592+09:00",
        "updated_at": "2025-10-04T13:21:45.832944717+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:21:46Z"
}
```

---

## 11. PUT /api/v1/users/{userId}/social-links/{linkId}
**Description**: Update social link

### Request
```
PUT http://localhost:8080/api/v1/users/{userId}/social-links/{linkId}
Headers:
  Authorization: Bearer mock-clerk-token
  Content-Type: application/json
```

**Request Body**:
```json
{
    "url": "https://twitter.com/newusername",
    "title": "Updated Twitter Account",
    "is_public": false
}
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "71cfa2c6-091d-4edd-a09c-388257b8a231",
        "user_id": "5298ae80-faeb-4f65-8894-695b652f8115",
        "platform": "twitter",
        "url": "https://twitter.com/newusername",
        "title": "Updated Twitter Account",
        "is_public": false,
        "created_at": "2025-10-04T13:21:45.832944+09:00",
        "updated_at": "2025-10-04T13:21:45.832944+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:22:40Z"
}
```

---

## 12. DELETE /api/v1/users/{userId}/social-links/{linkId}
**Description**: Delete social link

### Request
```
DELETE http://localhost:8080/api/v1/users/{userId}/social-links/{linkId}
Headers:
  Authorization: Bearer mock-clerk-token
```

### Response
**Status Code**: 

**Response Body**:
```json

```

---

## 13. GET /api/v1/users/{userId}/rivals
**Description**: Get user's rivals list

### Request
```
GET http://localhost:8080/api/v1/users/{userId}/rivals
Headers:
  Authorization: Bearer mock-clerk-token
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "rivals": [],
        "total": 0,
        "max_rivals": 3
    },
    "message": "success",
    "timestamp": "2025-10-04T04:10:43Z"
}
```

---

## 14. POST /api/v1/users/{userId}/rivals
**Description**: Add a rival

### Request
```
POST http://localhost:8080/api/v1/users/{userId}/rivals
Headers:
  Authorization: Bearer mock-clerk-token
  Content-Type: application/json
```

**Request Body**:
```json
{
    "rival_user_id": "b0e2d712-002a-4073-85f2-4daefcddf249"
}
```

### Response
**Status Code**: 

**Response Body**:
```json
{
    "data": {
        "id": "25fbb5d1-f06d-4b6b-9ae8-4671c4a9c52b",
        "user_id": "5298ae80-faeb-4f65-8894-695b652f8115",
        "rival_user": {
            "id": "b0e2d712-002a-4073-85f2-4daefcddf249",
            "clerk_id": "clerk_dev_user_002",
            "email": "tester@ghoona.camp",
            "username": "Test Expert",
            "avatar_url": "https://avatars.githubusercontent.com/u/2?v=4",
            "discord_id": null,
            "status": "active",
            "created_at": "2025-09-28T19:13:02.104918+09:00",
            "updated_at": "2025-09-28T19:13:02.104918+09:00"
        },
        "created_at": "2025-10-04T13:25:40.918967048+09:00"
    },
    "message": "success",
    "timestamp": "2025-10-04T04:25:41Z"
}
```

---

## 15. DELETE /api/v1/users/{userId}/rivals/{rivalId}
**Description**: Remove rival

### Request
```
DELETE http://localhost:8080/api/v1/users/{userId}/rivals/{rivalId}
Headers:
  Authorization: Bearer mock-clerk-token
```

### Response
**Status Code**: 

**Response Body**:
```json

```

---

## Test Results Summary

| Endpoint | Status | Notes |
|---|---|---|
| GET /api/v1/auth/me | | |
| GET /api/v1/users | | |
| GET /api/v1/users/{userId} | | |
| POST /api/v1/users | | |
| PUT /api/v1/users/{userId} | | |
| GET /api/v1/users/{userId}/metadata | | |
| POST /api/v1/users/{userId}/metadata | | |
| PUT /api/v1/users/{userId}/metadata | | |
| GET /api/v1/users/{userId}/social-links | | |
| POST /api/v1/users/{userId}/social-links | | |
| PUT /api/v1/users/{userId}/social-links/{linkId} | | |
| DELETE /api/v1/users/{userId}/social-links/{linkId} | | |
| GET /api/v1/users/{userId}/rivals | | |
| POST /api/v1/users/{userId}/rivals | | |
| DELETE /api/v1/users/{userId}/rivals/{rivalId} | | |

## Verification Points
- [ ] Authentication works properly
- [ ] Response format matches specification
- [ ] Error handling is appropriate
- [ ] Validation functions correctly
- [ ] Database integration works properly

## Issues & Improvements


## Next Test Schedule


---
*Last Updated: 2025-10-04*