# オニオンアーキテクチャ設計書

このドキュメントでは、朝活コミュニティアプリGhoona Campのバックエンドで採用しているオニオンアーキテクチャの詳細な構成と、各層の責任について説明します。

## オニオンアーキテクチャとは

オニオンアーキテクチャは、ドメイン駆動設計（DDD）の原則に基づいたアーキテクチャパターンです。依存関係が内側（ドメイン層）に向かうように設計され、ビジネスロジックの独立性と保守性を高めます。

### 依存関係の方向

```
外側 → 内側
Interface → Application → Domain
Infrastructure → Application → Domain
```

## プロジェクト構成

```
backend/internal/
├── domain/           # ドメイン層（最も内側）
│   ├── entity/       # エンティティ
│   ├── repository/   # リポジトリインターフェイス
│   └── errors.go     # ドメイン固有のエラー
├── application/      # アプリケーション層
│   ├── dto/          # データ転送オブジェクト
│   ├── usecase/      # ユースケース
│   └── transaction/  # トランザクション管理
├── infrastructure/   # インフラストラクチャ層（最も外側）
│   ├── database/     # データベース接続
│   ├── gorm/         # GORM関連（モデル・リポジトリ実装）
│   ├── jwt/          # JWT認証
│   └── slack/        # 外部サービス連携
└── interface/        # インターフェイス層（最も外側）
    ├── handler/      # HTTPハンドラー
    ├── middleware/   # ミドルウェア
    └── router/       # ルーティング
```

## 各層の詳細説明

### 1. ドメイン層（Domain Layer）

**責任**: ビジネスロジックとビジネスルールの定義

#### 1.1 エンティティ（Entity）

**場所**: `internal/domain/entity/`

**責任**: 
- ビジネスオブジェクトの定義
- ドメインロジックの実装
- データ検証ルールの定義

**実装例**: `user.go`

```go
package entity

import (
    "errors"
    "time"
)

// ユーザーロールの値オブジェクト
type UserRole string

const (
    UserRoleAdmin UserRole = "ADMIN"
    UserRoleUser  UserRole = "USER"
)

// バリデーションロジック（ドメインルール）
func (r UserRole) Validate() error {
    switch r {
    case UserRoleAdmin, UserRoleUser:
        return nil
    default:
        return errors.New("invalid user role")
    }
}

// ユーザーエンティティ
type User struct {
    Id        int
    Name      string
    Email     string
    Password  string
    Role      UserRole
    PayType   PayType
    PayRate   int
    Goal      int
    CreatedAt time.Time
    UpdatedAt time.Time
}

// ファクトリーメソッド
func NewUser(name, email, password string, role UserRole, payType PayType, payRate int) (*User, error) {
    // バリデーションロジック
    if err := role.Validate(); err != nil {
        return nil, err
    }
    
    return &User{
        Name:      name,
        Email:     email,
        Password:  password,
        Role:      role,
        PayType:   payType,
        PayRate:   payRate,
        Goal:      0,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }, nil
}

// ビジネスロジック
func (u *User) IsAdmin() bool {
    return u.Role == UserRoleAdmin
}

// ドメインルール
func (u *User) Validate() error {
    if u.Name == "" {
        return errors.New("name cannot be empty")
    }
    if u.PayRate <= 0 {
        return errors.New("pay rate must be greater than zero")
    }
    return nil
}
```

**重要なポイント**:
- ビジネスルールはエンティティ内に記述
- 外部依存を持たない
- 不変条件の維持
- ファクトリーメソッドによる正しいオブジェクト生成

#### 1.2 リポジトリインターフェイス

**場所**: `internal/domain/repository/`

**責任**: データ永続化の抽象化

**実装例**: `user_repository.go`

```go
package repository

import (
    "context"
    "github.com/attendance_report_app/backend/internal/domain/entity"
)

// リポジトリインターフェイス（ドメイン層で定義）
type UserRepository interface {
    FindAll(ctx context.Context) ([]*entity.User, error)
    FindById(ctx context.Context, id int) (*entity.User, error)
    Create(ctx context.Context, user *entity.User) (*entity.User, error)
    Update(ctx context.Context, user *entity.User) (*entity.User, error)
    Delete(ctx context.Context, id int) error
    FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
```

**重要なポイント**:
- インターフェイスをドメイン層で定義（依存関係逆転の原則）
- 実装はインフラストラクチャ層で行う
- ドメインエンティティのみを扱う

### 2. アプリケーション層（Application Layer）

**責任**: アプリケーションのワークフローの調整

#### 2.1 ユースケース（UseCase）

**場所**: `internal/application/usecase/`

**責任**:
- ビジネスユースケースの実装
- ドメインオブジェクトの協調
- トランザクション境界の定義

**実装例**: `user_usecase.go`

```go
package usecase

import (
    "context"
    "errors"
    "fmt"
    
    "github.com/attendance_report_app/backend/internal/application/dto"
    "github.com/attendance_report_app/backend/internal/application/dto/request"
    "github.com/attendance_report_app/backend/internal/domain/entity"
    "github.com/attendance_report_app/backend/internal/domain/repository"
)

// ユースケースインターフェイス
type UserUseCase interface {
    Login(ctx context.Context, email, password string) (*dto.LoginResponse, error)
    CreateUser(ctx context.Context, req *request.CreateUserRequest) (*dto.UserResponse, error)
    UpdateUser(ctx context.Context, userID int, req *request.UpdateUserRequest) (*dto.UserResponse, error)
}

// ユースケース実装
type userUseCase struct {
    userRepo     repository.UserRepository // ドメイン層のインターフェイス
    tokenService TokenService              // アプリケーション層のインターフェイス
}

// コンストラクタ（依存性注入）
func NewUserUseCase(userRepo repository.UserRepository, tokenService TokenService) UserUseCase {
    return &userUseCase{
        userRepo:     userRepo,
        tokenService: tokenService,
    }
}

// ログインユースケース
func (u *userUseCase) Login(ctx context.Context, email, password string) (*dto.LoginResponse, error) {
    // 1. リポジトリからユーザー取得
    user, err := u.userRepo.FindByEmail(ctx, email)
    if err != nil {
        return nil, errors.New("invalid email or password")
    }

    // 2. パスワード検証（ドメインサービス）
    if err := VerifyPassword(user.Password, password); err != nil {
        return nil, errors.New("invalid email or password")
    }

    // 3. トークン生成
    token, err := u.tokenService.GenerateToken(user.Id, string(user.Role))
    if err != nil {
        return nil, fmt.Errorf("failed to generate token: %w", err)
    }

    // 4. DTOに変換して返却
    return dto.ToLoginResponse(user, token), nil
}

// ユーザー作成ユースケース
func (u *userUseCase) CreateUser(ctx context.Context, req *request.CreateUserRequest) (*dto.UserResponse, error) {
    // 1. リクエスト検証
    if err := req.Validate(); err != nil {
        return nil, fmt.Errorf("invalid request: %w", err)
    }

    // 2. 既存ユーザーチェック
    existingUser, _ := u.userRepo.FindByEmail(ctx, req.Email)
    if existingUser != nil {
        return nil, errors.New("email already exists")
    }

    // 3. パスワードハッシュ化
    hashedPassword, err := HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    // 4. エンティティ作成（ドメインロジック使用）
    user := &entity.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashedPassword,
        Role:     entity.UserRole(req.Role),
        PayType:  entity.PayType(req.PayType),
        PayRate:  req.PayRate,
    }

    // 5. ドメインバリデーション
    if err := user.Validate(); err != nil {
        return nil, err
    }

    // 6. 永続化
    createdUser, err := u.userRepo.Create(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    // 7. DTO変換
    return dto.ToUserResponse(createdUser), nil
}
```

**重要なポイント**:
- ドメインロジックは呼び出すが実装しない
- 外部システムとの協調
- エラーハンドリング
- DTOを使用したデータ変換

#### 2.2 DTO（Data Transfer Object）

**場所**: `internal/application/dto/`

**責任**: 層間データ転送

**実装例**: `user_dto.go`

```go
package dto

import (
    "time"
    "github.com/attendance_report_app/backend/internal/domain/entity"
)

// レスポンス用DTO
type UserResponse struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Role      string    `json:"role"`
    PayType   string    `json:"pay_type"`
    PayRate   int       `json:"pay_rate"`
    Goal      int       `json:"goal"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// エンティティからDTOへの変換
func ToUserResponse(user *entity.User) *UserResponse {
    return &UserResponse{
        Id:        user.Id,
        Name:      user.Name,
        Email:     user.Email,
        Role:      string(user.Role),
        PayType:   string(user.PayType),
        PayRate:   user.PayRate,
        Goal:      user.Goal,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }
}
```

#### 2.3 トランザクション管理

**場所**: `internal/application/transaction/`

**責任**: データベーストランザクションの管理

**実装例**: `manager.go`

```go
package transaction

import (
    "context"
    "gorm.io/gorm"
)

type Manager interface {
    ExecuteInTx(ctx context.Context, fn func(context.Context) error) error
}

type manager struct {
    db *gorm.DB
}

func NewManager(db *gorm.DB) Manager {
    return &manager{db: db}
}

func (m *manager) ExecuteInTx(ctx context.Context, fn func(context.Context) error) error {
    return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        // コンテキストにトランザクションを設定
        txCtx := context.WithValue(ctx, "tx", tx)
        return fn(txCtx)
    })
}
```

### 3. インフラストラクチャ層（Infrastructure Layer）

**責任**: 外部システムとの連携

#### 3.1 データベース関連

**場所**: `internal/infrastructure/gorm/`

##### GORM モデル

**場所**: `internal/infrastructure/gorm/model/`

**責任**: データベーステーブルとの O/R マッピング

**実装例**: `user.go`

```go
package model

import (
    "time"
    "github.com/attendance_report_app/backend/internal/domain/entity"
)

// GORMモデル（データベーススキーマ）
type User struct {
    Id        int       `gorm:"primaryKey;autoIncrement" json:"id"`
    Name      string    `gorm:"type:varchar(100);not null" json:"name"`
    Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
    Password  string    `gorm:"type:varchar(255);not null" json:"-"`
    Role      string    `gorm:"type:enum('ADMIN','USER');not null;default:'USER'" json:"role"`
    PayType   string    `gorm:"type:enum('HOURLY','MONTHLY');not null;default:'HOURLY'" json:"pay_type"`
    PayRate   int       `gorm:"not null;default:1000" json:"pay_rate"`
    Goal      int       `gorm:"default:0" json:"goal"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// テーブル名指定
func (User) TableName() string {
    return "users"
}

// エンティティからモデルへの変換
func FromUserEntity(user *entity.User) *User {
    return &User{
        Id:        user.Id,
        Name:      user.Name,
        Email:     user.Email,
        Password:  user.Password,
        Role:      string(user.Role),
        PayType:   string(user.PayType),
        PayRate:   user.PayRate,
        Goal:      user.Goal,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }
}

// モデルからエンティティへの変換
func ToUserEntity(user *User) *entity.User {
    return &entity.User{
        Id:        user.Id,
        Name:      user.Name,
        Email:     user.Email,
        Password:  user.Password,
        Role:      entity.UserRole(user.Role),
        PayType:   entity.PayType(user.PayType),
        PayRate:   user.PayRate,
        Goal:      user.Goal,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }
}
```

##### リポジトリ実装

**場所**: `internal/infrastructure/gorm/repository/`

**責任**: ドメイン層リポジトリインターフェイスの実装

**実装例**: `user_repository_impl.go`

```go
package repository

import (
    "context"
    "gorm.io/gorm"
    
    "github.com/attendance_report_app/backend/internal/domain/entity"
    "github.com/attendance_report_app/backend/internal/domain/repository"
    "github.com/attendance_report_app/backend/internal/infrastructure/gorm/model"
)

// リポジトリ実装
type userRepository struct {
    db *gorm.DB
}

// コンストラクタ
func NewUserRepository(db *gorm.DB) repository.UserRepository {
    return &userRepository{db: db}
}

// トランザクション対応のDB取得
func (r *userRepository) getDB(ctx context.Context) *gorm.DB {
    if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
        return tx // トランザクション中
    }
    return r.db // 通常のDB接続
}

// 全ユーザー取得
func (r *userRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
    var users []model.User
    if err := r.getDB(ctx).Find(&users).Error; err != nil {
        return nil, err
    }
    return model.ToUserEntities(users), nil
}

// ID検索
func (r *userRepository) FindById(ctx context.Context, id int) (*entity.User, error) {
    var user model.User
    if err := r.getDB(ctx).First(&user, id).Error; err != nil {
        return nil, err
    }
    return model.ToUserEntity(&user), nil
}

// ユーザー作成
func (r *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
    userModel := model.FromUserEntity(user)
    if err := r.getDB(ctx).Create(&userModel).Error; err != nil {
        return nil, err
    }
    return model.ToUserEntity(userModel), nil
}
```

#### 3.2 外部サービス連携

**場所**: `internal/infrastructure/jwt/`, `internal/infrastructure/slack/`

**実装例**: JWT サービス (`jwt/token_service.go`)

```go
package jwt

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
    secretKey string
}

func NewTokenService(secretKey string) *TokenService {
    return &TokenService{secretKey: secretKey}
}

func (s *TokenService) GenerateToken(userID int, role string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.secretKey))
}
```

### 4. インターフェイス層（Interface Layer）

**責任**: 外部からのリクエスト受付とレスポンス

#### 4.1 ハンドラー

**場所**: `internal/interface/handler/`

**責任**: HTTPリクエスト処理

**実装例**: `user_handler.go`

```go
package handler

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "github.com/attendance_report_app/backend/internal/application/usecase"
    "github.com/attendance_report_app/backend/internal/application/transaction"
)

type UserHandler struct {
    userUseCase usecase.UserUseCase
    txManager   transaction.Manager
}

func NewUserHandler(userUseCase usecase.UserUseCase, txManager transaction.Manager) *UserHandler {
    return &UserHandler{
        userUseCase: userUseCase,
        txManager:   txManager,
    }
}

// ユーザー作成ハンドラー
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req request.CreateUserRequest
    
    // 1. リクエスト解析
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // 2. リクエスト検証
    if err := req.Validate(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 3. トランザクション内でユースケース実行
    var user *dto.UserResponse
    err := h.txManager.ExecuteInTx(c.Request.Context(), func(ctx context.Context) error {
        var err error
        user, err = h.userUseCase.CreateUser(ctx, &req)
        return err
    })

    // 4. エラーハンドリング
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 5. 成功レスポンス
    c.JSON(http.StatusCreated, user)
}
```

#### 4.2 ミドルウェア

**場所**: `internal/interface/middleware/`

**実装例**: 認証ミドルウェア (`auth.go`)

```go
package middleware

import (
    "net/http"
    "strings"
    
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(secretKey), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims := token.Claims.(jwt.MapClaims)
        c.Set("userID", int(claims["user_id"].(float64)))
        c.Set("role", claims["role"].(string))
        
        c.Next()
    }
}
```

## 共通基盤コンポーネント

### 1. エラーハンドリング

**場所**: `internal/domain/errors.go`

```go
package domain

import "errors"

// ドメインエラー定義
var (
    ErrUserNotFound     = errors.New("user not found")
    ErrInvalidEmail     = errors.New("invalid email format")
    ErrEmailExists      = errors.New("email already exists")
    ErrUnauthorized     = errors.New("unauthorized access")
    ErrInvalidRole      = errors.New("invalid user role")
)

// エラータイプ定義
type DomainError struct {
    Code    string
    Message string
}

func (e DomainError) Error() string {
    return e.Message
}

func NewDomainError(code, message string) *DomainError {
    return &DomainError{
        Code:    code,
        Message: message,
    }
}
```

### 2. ロギング

**場所**: `internal/interface/middleware/logger.go`

```go
package middleware

import (
    "time"
    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        return fmt.Sprintf("[%s] %s %s %d %s %s\n",
            param.TimeStamp.Format(time.RFC3339),
            param.Method,
            param.Path,
            param.StatusCode,
            param.Latency,
            param.ErrorMessage,
        )
    })
}
```

### 3. バリデーション

**場所**: `internal/application/dto/request/`

```go
package request

import (
    "errors"
    "regexp"
)

type CreateUserRequest struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
    Role     string `json:"role" binding:"required"`
    PayType  string `json:"pay_type" binding:"required"`
    PayRate  int    `json:"pay_rate" binding:"required"`
}

func (r *CreateUserRequest) Validate() error {
    if len(r.Password) < 8 {
        return errors.New("password must be at least 8 characters")
    }
    
    emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    if match, _ := regexp.MatchString(emailRegex, r.Email); !match {
        return errors.New("invalid email format")
    }
    
    return nil
}
```

## Docker設定

### Dockerfile

```dockerfile
FROM golang:1.23-alpine AS builder

# 依存関係インストール
RUN apk add --no-cache git

# 作業ディレクトリ設定
WORKDIR /app

# Go modules
COPY go.mod go.sum ./
RUN go mod download

# ソースコードコピー
COPY . .

# アプリケーションビルド
RUN go build -o main cmd/api/main.go

# 実行用軽量イメージ
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app

# バイナリコピー
COPY --from=builder /app/main .
COPY --from=builder /app/.env* ./

# ポート公開
EXPOSE 8080

# 実行
CMD ["./main"]
```

### docker-compose.yml

```yaml
version: '3.8'

services:
  # バックエンドAPI
  backend:
    env_file:
      - ./backend/.env
    image: golang:1.23-alpine
    ports:
      - "8088:8080"
    environment:
      - DB_HOST=db
      - DB_PORT=3306
    volumes:
      - ./backend:/app
    working_dir: /app
    command: >
      sh -c "cd /app && 
             go run cmd/migrate/main.go && 
             go run cmd/seed/main.go && 
             go run cmd/api/main.go"
    depends_on:
      db:
        condition: service_healthy

  # MySQL データベース
  db:
    image: mysql:8.0
    container_name: mysql_container_attendance
    env_file:
      - ./backend/.env
    environment:
      - MYSQL_ROOT_PASSWORD=yourpassword
      - MYSQL_DATABASE=attendance_db
    ports:
      - "3306:3306"
    volumes:
      - ./backend/db_data:/var/lib/mysql
      - ./backend/my.cnf:/etc/mysql/conf.d/my.cnf
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      timeout: 20s
      retries: 10

  # phpMyAdmin (開発用)
  phpmyadmin:
    image: phpmyadmin/phpmyadmin
    container_name: phpmyadmin_attendance
    ports:
      - "8081:80"
    environment:
      PMA_HOST: db
      PMA_PORT: 3306
    depends_on:
      - db
```

## アーキテクチャの利点

### 1. 保守性
- 各層の責任が明確
- 依存関係が整理されている
- ビジネスロジックが分離されている

### 2. テスタビリティ
- 依存性注入によりモックが容易
- 各層を独立してテスト可能
- ユニットテストとインテグレーションテストの分離

### 3. 拡張性
- 新しい要件に対する変更が局所化
- 外部システム変更の影響を最小化
- 新しいインターフェイス（CLI、gRPC等）の追加が容易

### 4. 再利用性
- ドメインロジックの再利用
- 複数のインターフェイスでの同一ユースケース利用

## 開発時のベストプラクティス

### 1. 依存関係の管理
- 内側の層は外側の層を知らない
- インターフェイスを活用した疎結合
- 依存性注入の活用

### 2. エラーハンドリング
- 各層でのエラー変換
- ドメインエラーの定義
- 適切なHTTPステータスコードの返却

### 3. テスト戦略
- ドメインロジックの単体テスト
- ユースケースの統合テスト
- ハンドラーのAPIテスト

### 4. パフォーマンス考慮
- 必要最小限のデータ取得
- トランザクション範囲の最適化
- キャッシュ戦略の検討

このオニオンアーキテクチャにより、保守性が高く、テストしやすく、拡張可能なアプリケーションを構築できます。