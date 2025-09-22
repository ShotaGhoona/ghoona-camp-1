FROM node:20-alpine

# 作業ディレクトリを設定
WORKDIR /app

# 開発環境を指定
ENV NODE_ENV=development
ENV NEXT_TELEMETRY_DISABLED=1

# package.json と package-lock.json をコピー
COPY package*.json ./

# 依存関係をインストール（devDependencies も含む）
RUN npm ci || npm install

# ポートを解放
EXPOSE 3000

# 開発サーバーをホットリロード有効で起動
CMD ["npm", "run", "dev"]