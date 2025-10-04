import { z } from 'zod';
import { createUserMetadataFormSchema } from '@/entities/user/metadata-entity';

// Entity層のvalidationを基盤として使用
export const metadataCreateFormSchema = createUserMetadataFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
  // 例：スキル数の上限チェック、禁止ワードチェック、プロフィール画像解析など
});

export type MetadataCreateFormData = z.infer<typeof metadataCreateFormSchema>;

// UI特有の型定義
export interface MetadataCreateFormProps {
  onSubmit: (data: MetadataCreateFormData) => void;
  isLoading?: boolean;
  initialData?: Partial<MetadataCreateFormData>;
}

// プロフィール設定ステップ管理用の型
export interface ProfileSetupStep {
  step: 'basic' | 'skills' | 'bio' | 'vision' | 'complete';
  completedSteps: Set<string>;
  canProceed: boolean;
}

// スキル・興味入力用の型（create用に特化）
export interface SkillsInputProps {
  value: string[];
  onChange: (skills: string[]) => void;
  maxItems?: number;
  placeholder?: string;
  suggestions?: string[];
}

export interface InterestsInputProps {
  value: string[];
  onChange: (interests: string[]) => void;
  maxItems?: number;
  placeholder?: string;
  categories?: InterestCategory[];
}

// 興味カテゴリー用の型
export interface InterestCategory {
  id: string;
  name: string;
  interests: string[];
}

// プロフィール画像アップロード用の型
export interface ProfileImageUploadProps {
  value?: string;
  onChange: (url: string) => void;
  isUploading?: boolean;
  maxSize?: number; // MB
  aspectRatio?: number; // 1 for square, 16/9 for wide, etc.
}

// タイムゾーン選択用の型
export interface TimezoneSelectProps {
  value: string;
  onChange: (timezone: string) => void;
  regions?: string[];
}

// ビジョン公開設定用の型
export interface VisionPrivacyProps {
  isPublic: boolean;
  onChange: (isPublic: boolean) => void;
  description?: string;
}