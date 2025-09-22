FROM golang:1.25-alpine

# 必要なツールのインストール
RUN apk update && apk add --no-cache \
    git \
    curl \
    ca-certificates \
    tzdata \
    && rm -rf /var/cache/apk/*

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