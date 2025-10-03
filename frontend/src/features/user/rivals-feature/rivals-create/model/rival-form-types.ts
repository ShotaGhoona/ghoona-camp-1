import { z } from 'zod';
import { 
  createRivalFormSchema, 
  validateRivalCount,
  validateNotSelfRival,
  validateDuplicateRival,
  type Rival
} from '@/entities/user/rivals-entity';

// Entity層のvalidationを基盤として使用
export const rivalCreateFormSchema = createRivalFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
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
}

// ビジネスロジック用のヘルパー型
export interface RivalValidationContext {
  currentUserId: string;
  existingRivals: Rival[];
  targetUserId: string;
}

// ビジネスルール検証関数（Feature層特有のロジック）
export const validateRivalBusiness = (context: RivalValidationContext) => {
  const { currentUserId, existingRivals, targetUserId } = context;
  
  // 上限チェック
  validateRivalCount(existingRivals.length);
  
  // 自分自身チェック
  validateNotSelfRival(currentUserId, targetUserId);
  
  // 重複チェック
  validateDuplicateRival(existingRivals, targetUserId);
};

// UI状態管理用の型
export interface RivalCreateModalState {
  isOpen: boolean;
  step: 'search' | 'confirm';
  selectedUserId: string | null;
}

// 定数の再エクスポート（isolatedModules対応）
export { MAX_RIVALS_COUNT } from '@/entities/user/rivals-entity';