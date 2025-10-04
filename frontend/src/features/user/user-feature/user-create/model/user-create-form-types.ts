import { z } from 'zod';
import { createUserFormSchema } from '@/entities/user/user-entity';

// Entity層のvalidationを基盤として使用
export const userCreateFormSchema = createUserFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
  // 例：プロフィール画像のファイルサイズ制限、禁止ドメインチェックなど
});

export type UserCreateFormData = z.infer<typeof userCreateFormSchema>;

// UI特有の型定義
export interface UserCreateFormProps {
  onSubmit: (data: UserCreateFormData) => void;
  isLoading?: boolean;
  initialData?: Partial<UserCreateFormData>;
}

// ユーザー登録ステップ管理用の型
export interface UserRegistrationStep {
  step: 'basic' | 'verification' | 'complete';
  canProceed: boolean;
}

// アバター選択用の型
export interface AvatarUploadProps {
  value?: string;
  onChange: (url: string) => void;
  isLoading?: boolean;
  maxSize?: number; // MB
}