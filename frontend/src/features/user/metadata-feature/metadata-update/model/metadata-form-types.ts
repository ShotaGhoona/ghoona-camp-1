import { z } from 'zod';
import { updateUserMetadataFormSchema } from '@/entities/user/metadata-entity';

// Entity層のvalidationを基盤として使用
export const metadataUpdateFormSchema = updateUserMetadataFormSchema.extend({
  // Feature特有のバリデーション（必要に応じて拡張）
  // 例：スキル数の上限チェック、禁止ワードチェック、変更差分チェックなど
});

export type MetadataUpdateFormData = z.infer<typeof metadataUpdateFormSchema>;

// UI特有の型定義
export interface MetadataUpdateFormProps {
  initialData?: Partial<MetadataUpdateFormData>;
  onSubmit: (data: MetadataUpdateFormData) => void;
  isLoading?: boolean;
}

// スキル・興味入力用の型（update用に特化）
export interface SkillsSelectorProps {
  value: string[];
  onChange: (skills: string[]) => void;
  maxItems?: number;
  placeholder?: string;
  existingSkills?: string[];
  onSkillRemove?: (skill: string) => void;
}

export interface InterestsSelectorProps {
  value: string[];
  onChange: (interests: string[]) => void;
  maxItems?: number;
  placeholder?: string;
  existingInterests?: string[];
  onInterestRemove?: (interest: string) => void;
}

// 編集フォーム用の型
export interface EditableSectionProps {
  isEditing: boolean;
  onEdit: () => void;
  onSave: () => void;
  onCancel: () => void;
  isLoading?: boolean;
}

// プロフィール画像更新用の型
export interface ProfileImageUpdateProps {
  currentImageUrl?: string;
  onImageChange: (url: string) => void;
  onImageRemove: () => void;
  isUploading?: boolean;
  maxSize?: number;
}

// ビジョン編集用の型
export interface VisionEditorProps {
  value: string;
  onChange: (value: string) => void;
  isPublic: boolean;
  onPublicChange: (isPublic: boolean) => void;
  maxLength?: number;
  placeholder?: string;
}

// タイムゾーン更新用の型
export interface TimezoneUpdateProps {
  currentTimezone: string;
  onTimezoneChange: (timezone: string) => void;
  autoDetect?: boolean;
}

// 変更追跡用の型
export interface MetadataChanges {
  hasChanges: boolean;
  changedFields: Set<keyof MetadataUpdateFormData>;
  originalData: Partial<MetadataUpdateFormData>;
}