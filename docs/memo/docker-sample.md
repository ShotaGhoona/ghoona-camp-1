services:
  # Backend (Go + Gin + GORM)
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    container_name: eagleai_backend
    env_file:
      - ./backend/.env
    ports:
      - "8080:8080"
    volumes:
      - ./backend:/app
      - go_modules:/go/pkg/mod
    working_dir: /app
    command: go run cmd/api/main.go
    depends_on:
      db:
        condition: service_healthy
      migrate:
        condition: service_completed_successfully
    networks:
      - eagleai_network

  # MySQL Database
  db:
    image: mysql:8.0
    container_name: eagleai_mysql
    environment:
      - MYSQL_ROOT_PASSWORD=rootpassword
      - MYSQL_DATABASE=eagle_ai_db
    ports:
      - "3306:3306"
    volumes:
      - ./db_data:/var/lib/mysql
      - ./backend/my.cnf:/etc/mysql/conf.d/my.cnf
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 10s
      timeout: 20s
      retries: 10
      start_period: 30s
    networks:
      - eagleai_network

  # phpMyAdmin
  phpmyadmin:
    image: phpmyadmin/phpmyadmin
    container_name: eagleai_phpmyadmin
    platform: linux/amd64
    ports:
      - "8081:80"
    environment:
      PMA_HOST: db
      PMA_PORT: 3306
      PMA_USER: root
      PMA_PASSWORD: rootpassword
    depends_on:
      - db
    networks:
      - eagleai_network

  # Migration container (Atlas + GORM AutoMigrate)
  migrate:
    build:
      context: ./backend
      dockerfile: Dockerfile
    container_name: eagleai_migrate
    env_file:
      - ./backend/.env
    volumes:
      - ./backend:/app
    working_dir: /app
    command: go run cmd/migrate/main.go apply
    depends_on:
      db:
        condition: service_healthy
    networks:
      - eagleai_network

  # Frontend (Next.js)
  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    container_name: eagleai-frontend
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
      - eagleai_network

volumes:
  go_modules:
    driver: local

networks:
  eagleai_network:
    driver: bridge


---



# 開発用 Dockerfile
FROM node:20-alpine

# 作業ディレクトリを設定
WORKDIR /app

# 開発環境を指定
ENV NODE_ENV=development
ENV NEXT_TELEMETRY_DISABLED=1

# pnpm をインストール（より高速）
RUN npm install -g pnpm

# package.json と package-lock.json をコピー
COPY package*.json ./
COPY pnpm-lock.yaml* ./

# 依存関係をインストール（devDependencies も含む）
RUN npm ci || pnpm install --frozen-lockfile || npm install

# ポートを解放
EXPOSE 3000

# 開発サーバーをホットリロード有効で起動
CMD ["npm", "run", "dev"]


---


FROM golang:1.23-alpine

# 必要なツールのインストール
RUN apk update && apk add --no-cache \
    git \
    curl \
    mysql-client \
    ca-certificates \
    tzdata \
    make \
    && rm -rf /var/cache/apk/*

# Atlas CLIのインストール
RUN curl -sSf https://atlasgo.sh | sh

# golangci-lintのインストール (v1.64.8)
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.8

# タイムゾーン設定
ENV TZ=Asia/Tokyo
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

# 作業ディレクトリ設定
WORKDIR /app

# Go modulesの依存関係を先にコピー（キャッシュ効率化）
COPY go.mod go.sum ./
RUN go mod download

# ソースコード全体をコピー
COPY . .

# 開発用エントリーポイント
CMD ["go", "run", "cmd/api/main.go"]