import { z } from 'zod';
import { 
  createRivalFormSchema, 
  validateRivalCount,
  validateNotSelfRival,
  validateDuplicateRival,
  type Rival,
  MAX_RIVALS_COUNT
} from '@/entities/user/rivals-entity';

// Entity層のvalidationを基盤として使用
export const rivalCreateFormSchema = createRivalFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
  // 例：ライバル候補のアクティビティ状態チェック、相互フォロー推奨など
});

export type RivalCreateFormData = z.infer<typeof rivalCreateFormSchema>;

// UI特有の型定義
export interface RivalCreateFormProps {
  onSubmit: (data: RivalCreateFormData) => void;
  isLoading?: boolean;
  existingRivals?: Rival[];
  currentUserId?: string;
}

// ライバル検索・選択用の型
export interface RivalSearchFormProps {
  onUserSelect: (userId: string) => void;
  excludeUserIds?: string[];
  isLoading?: boolean;
  searchPlaceholder?: string;
}

// ユーザー検索結果用の型
export interface UserSearchResult {
  id: string;
  username: string | null;
  email: string;
  avatarUrl: string | null;
  displayName?: string;
  isAlreadyRival: boolean;
  isSelf: boolean;
}

// ライバル候補表示用の型
export interface RivalCandidateProps {
  user: UserSearchResult;
  onSelect: (userId: string) => void;
  onViewProfile: (userId: string) => void;
  isSelected?: boolean;
  disabled?: boolean;
}

// ビジネスロジック用のヘルパー型
export interface RivalValidationContext {
  currentUserId: string;
  existingRivals: Rival[];
  targetUserId: string;
}

// ビジネスルール検証関数（Feature層特有のロジック）
export const validateRivalBusiness = (context: RivalValidationContext): void => {
  const { currentUserId, existingRivals, targetUserId } = context;
  
  // 上限チェック
  if (!validateRivalCount(existingRivals.length)) {
    throw new Error(`ライバルは最大${MAX_RIVALS_COUNT}人まで設定できます`);
  }
  
  // 自分自身チェック
  if (!validateNotSelfRival(currentUserId, targetUserId)) {
    throw new Error('自分自身をライバルに設定することはできません');
  }
  
  // 重複チェック
  if (!validateDuplicateRival(existingRivals, targetUserId)) {
    throw new Error('このユーザーは既にライバルに設定されています');
  }
};

// UI状態管理用の型
export interface RivalCreateModalState {
  isOpen: boolean;
  step: 'search' | 'confirm' | 'success';
  selectedUserId: string | null;
  selectedUser: UserSearchResult | null;
}

// ライバル追加確認用の型
export interface RivalConfirmationProps {
  selectedUser: UserSearchResult;
  onConfirm: () => void;
  onCancel: () => void;
  isLoading?: boolean;
}

// 検索フィルター用の型
export interface UserSearchFilters {
  query: string;
  excludeUserIds: string[];
  includeInactive: boolean;
  sortBy: 'username' | 'email' | 'createdAt';
  sortOrder: 'asc' | 'desc';
}

// ライバル推奨システム用の型
export interface RivalSuggestion {
  user: UserSearchResult;
  reason: 'mutual_skills' | 'similar_interests' | 'activity_level' | 'community';
  score: number;
  description: string;
}

export interface RivalSuggestionsProps {
  suggestions: RivalSuggestion[];
  onSelectSuggestion: (userId: string) => void;
  isLoading?: boolean;
}

// 定数の再エクスポート
export { MAX_RIVALS_COUNT };