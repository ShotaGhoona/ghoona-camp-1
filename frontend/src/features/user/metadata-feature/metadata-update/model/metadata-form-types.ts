import { z } from 'zod';
import { updateUserMetadataFormSchema } from '@/entities/user/metadata-entity';

// Entity層のvalidationを基盤として使用
export const metadataUpdateFormSchema = updateUserMetadataFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
  // 例：スキル数の上限チェック、禁止ワードチェックなど
});

export type MetadataUpdateFormData = z.infer<typeof metadataUpdateFormSchema>;

// UI特有の型定義
export interface MetadataUpdateFormProps {
  initialData?: Partial<MetadataUpdateFormData>;
  onSubmit: (data: MetadataUpdateFormData) => void;
  isLoading?: boolean;
}

// スキル・興味入力用の型
export interface SkillsSelectorProps {
  value: string[];
  onChange: (skills: string[]) => void;
  maxItems?: number;
  placeholder?: string;
}