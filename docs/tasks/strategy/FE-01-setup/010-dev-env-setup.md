# フロントエンド開発環境構築戦略

## 概要
朝活コミュニティアプリ「Ghoona Camp」のフロントエンド開発環境を構築する。Next.js 15、TypeScript、FSD（Feature-Sliced Design）アーキテクチャを採用し、効率的で保守しやすい開発基盤を整備する。

## 実装戦略

### 1. 技術スタック選定（30分）
- **Next.js 15**: App Router採用、最新安定版
- **TypeScript**: 厳密な型安全性で品質向上
- **Node.js**: LTS版（20.x系）推奨
- **パッケージマネージャー**: npm または pnpm

### 2. 開発環境セットアップ（90分）
```bash
# Node.js環境確認
node --version  # v20.x.x
npm --version   # 10.x.x

# Next.jsプロジェクト作成
npx create-next-app@latest frontend \
  --typescript \
  --tailwind \
  --eslint \
  --app \
  --src-dir \
  --import-alias "@/*"

# 依存関係インストール
cd frontend
npm install
```

### 3. 開発ツール設定（60分）
- **VS Code Extensions**:
  - TypeScript Importer
  - Tailwind CSS IntelliSense
  - ES7+ React/Redux/React-Native snippets
  - Prettier - Code formatter

- **開発サーバー確認**:
```bash
npm run dev  # http://localhost:3000
```

## 成功基準
- [ ] Next.js開発サーバーが正常起動
- [ ] TypeScriptコンパイルエラーなし
- [ ] ホットリロード機能が動作
- [ ] Tailwind CSSスタイルが適用される

## リスク & 対策
- **Node.jsバージョン不適合**: nvm使用で環境管理
- **ポート競合**: `--port 3001`で代替ポート使用
- **権限エラー**: `sudo`使用せず、npmグローバル設定確認

## 次のステップ
FE-01-setup-02「プロジェクト初期化とFSD構造作成」で、FSDアーキテクチャのディレクトリ構造を構築する。
