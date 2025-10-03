import { z } from 'zod';
import { updateUserFormSchema } from '@/entities/user/user-entity';

// Entity層のvalidationを基盤として使用
export const userBasicUpdateFormSchema = updateUserFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
});

export type UserBasicUpdateFormData = z.infer<typeof userBasicUpdateFormSchema>;

// UI特有の型定義
export interface UserBasicUpdateFormProps {
  initialData?: Partial<UserBasicUpdateFormData>;
  onSubmit: (data: UserBasicUpdateFormData) => void;
  isLoading?: boolean;
}