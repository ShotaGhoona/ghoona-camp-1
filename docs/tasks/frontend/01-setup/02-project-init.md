# フロントエンドプロジェクト初期化とFSD構造作成戦略

**タスクID**: FE-01-setup-02  
**見積**: 4時間  
**優先度**: 🔴 高

## 概要
Next.jsプロジェクトの初期化とFSD（Feature-Sliced Design）アーキテクチャのディレクトリ構造を構築する。将来的な保守性と拡張性を考慮した設計基盤を整備する。

## 実装戦略

### 1. FSDアーキテクチャ設計（60分）
FSDの7層構造を採用：
- **app**: Next.js App Router統合
- **pages**: ページコンポーネント（App Routerとの橋渡し）
- **widgets**: 複合UIブロック
- **features**: ビジネス機能
- **entities**: ドメインエンティティ
- **shared**: 共有コンポーネント・ユーティリティ

### 2. ディレクトリ構造構築（120分）
```bash
cd frontend/src

# FSD基本構造作成
mkdir -p {app,pages,widgets,features,entities,shared}/{ui,api,model,lib,config}

# Shared層詳細構造
mkdir -p shared/{ui/components,lib/utils,api/client,config/constants}

# 各エンティティ作成
mkdir -p entities/{user,attendance,goal,event,title,notification}

# 主要機能作成
mkdir -p features/{auth,attendance-tracking,goal-management}

# ウィジェット作成
mkdir -p widgets/{header,sidebar,dashboard}
```

### 3. 基本設定ファイル作成（60分）
```typescript
// shared/config/constants.ts
export const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'
export const APP_NAME = 'Ghoona Camp'

// shared/lib/utils/index.ts
export * from './cn'
export * from './date'
export * from './validation'

// app/globals.css (Tailwind設定)
@tailwind base;
@tailwind components;
@tailwind utilities;
```

### 4. TypeScript設定最適化（40分）
```json
// tsconfig.json パス追加
{
  "compilerOptions": {
    "paths": {
      "@/*": ["./src/*"],
      "@/shared/*": ["./src/shared/*"],
      "@/entities/*": ["./src/entities/*"],
      "@/features/*": ["./src/features/*"],
      "@/widgets/*": ["./src/widgets/*"],
      "@/pages/*": ["./src/pages/*"]
    }
  }
}
```

## 成功基準
- [ ] FSD 7層ディレクトリ構造が正しく作成される
- [ ] TypeScript パスエイリアス設定完了
- [ ] 基本設定ファイルが作成される
- [ ] Next.js開発サーバーがエラーなく起動する
- [ ] インポートパスが正しく解決される

## FSD依存関係ルール
```
app ← pages ← widgets ← features ← entities ← shared
```
- 上位層は下位層のみに依存
- 同一層間の依存は禁止
- shared層はすべての層から利用可能

## リスク & 対策
- **ディレクトリ深すぎエラー**: Windows環境での長いパス名対応
- **インポートパス解決失敗**: TypeScript設定確認とVS Code再起動
- **FSD境界違反**: ESLintルール設定で自動検出

## 次のステップ
FE-01-setup-03「開発ツール・Linter設定」で、FSD境界ルールを強制するESLint設定を行う。

---
*作成日: 2025-01-21*