# BE-01-setup-03 Docker化設定戦略

## 🎯 30秒キャッチアップ

**目標**: フロントエンド・バックエンド統合Docker環境構築  
**作業時間**: 3時間  
**対象**: アプリケーションのDocker化、コンテナ最適化  
**構成**: Frontend(Next.js) + Backend(Go) + Database(PostgreSQL)

## 📋 実装内容

### 1. Dockerfileの作成（90分）
- **バックエンド**: マルチステージビルド対応Go Dockerfile
- **フロントエンド**: Next.js用Dockerfile
- **最適化**: イメージサイズ削減、キャッシュ活用

### 2. Docker Compose設定（90分）
- **ネットワーク**: サービス間通信設定
- **ボリューム**: DB永続化
- **環境変数**: 統合管理
- **起動スクリプト**: 簡単起動

## 🏗️ Docker構成

```
project-root/
├── frontend/
│   ├── Dockerfile
│   ├── .dockerignore
│   └── .env               # フロントエンド環境変数
├── backend/
│   ├── Dockerfile
│   ├── .dockerignore
│   └── .env               # バックエンド環境変数
└── docker-compose.yml      # 統合環境
```

## 🔧 Docker設定詳細

### バックエンド Dockerfile
```dockerfile
# マルチステージビルド
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

### フロントエンド Dockerfile
```dockerfile
FROM node:20-alpine AS base
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV production
COPY --from=base /app/node_modules ./node_modules
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json
EXPOSE 3000
CMD ["npm", "start"]
```

### Docker Compose
```yaml
services:
  backend:
    build: ./backend
    container_name: ghoona_backend
    env_file:
      - ./backend/.env
    ports:
      - "8080:8080"
    volumes:
      - ./backend:/app
    working_dir: /app
    command: go run cmd/api/main.go
    depends_on:
      - db
    networks:
      - ghoona_network

  frontend:
    build: ./frontend
    container_name: ghoona_frontend
    env_file:
      - ./frontend/.env
    ports:
      - "3000:3000"
    volumes:
      - ./frontend:/app
      - /app/.next
    depends_on:
      - backend
    networks:
      - ghoona_network

  db:
    image: postgres:16-alpine
    container_name: ghoona_postgres
    environment:
      - POSTGRES_DB=ghoona_camp
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - ghoona_network

volumes:
  postgres_data:

networks:
  ghoona_network:
    driver: bridge
```

## 📦 サービス構成

### Docker環境
1. **frontend**: Next.js（:3000）
2. **backend**: Go API server（:8080）
3. **db**: PostgreSQL（:5432）

## 🚀 使用方法

### 起動
```bash
docker-compose up --build
```

### バックグラウンド起動
```bash
docker-compose up --build -d
```

### 停止
```bash
docker-compose down
```

## ✅ 完了基準

- [ ] バックエンドDockerfile作成・ビルド確認
- [ ] フロントエンドDockerfile作成・ビルド確認
- [ ] Docker Compose統合環境動作確認
- [ ] サービス間通信確認

## 🎯 次のステップ
**BE-02-arch-01**: オニオンアーキテクチャ基盤の構築

## 📝 注意事項
- **環境変数**: 各プロジェクトの.env使用
- **開発**: ローカルでの開発推奨
- **Docker**: 統合確認用途