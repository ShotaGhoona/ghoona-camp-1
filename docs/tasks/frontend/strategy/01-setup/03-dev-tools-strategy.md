# FE-01-setup-03 開発ツール・Linter設定戦略

## 🎯 30秒キャッチアップ

**目標**: 開発品質とFSD構造を保つための自動化ツールを設定  
**作業時間**: 2時間  
**成果物**: ESLint + Prettier + FSD boundaries + pre-commit hooks

## 📋 実装内容

### 1. ESLint設定（30分）
- Next.js + TypeScript用ルール
- FSD import/export ルール
- コードスタイル統一

### 2. Prettier設定（15分）
- フォーマット自動化
- ESLintとの競合回避

### 3. FSD Boundaries設定（45分）
- レイヤー間依存関係チェック
- 不正なimportを防止
- `shared ← entities ← features ← widgets ← pages`

### 4. Pre-commit hooks（30分）
- husky + lint-staged
- コミット前の自動チェック
- 品質ゲートキーピング

## 🚀 設定ファイル
- `.eslintrc.json`
- `.prettierrc`
- `package.json` (scripts追加)
- `.husky/pre-commit`

## ✅ 完了基準
- `npm run lint` 実行可能
- `npm run format` 実行可能
- FSD違反でlintエラー
- コミット時自動チェック