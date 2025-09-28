/**
 * User Metadata Entity Mappers
 * 
 * APIレスポンス型からフロントエンド用型への変換関数を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * metadata-entity/index.ts からエクスポートされた関数を使用してください。
 */

import type { UserMetadata, UserMetadataDetailResponse, UserMetadataUpdateResponse } from '../model/metadata-types';

/**
 * UserMetadataDetailResponse のデータをUserMetadata型に変換
 */
export const mapUserMetadataDetailResponseToUserMetadata = (
  metadataData: UserMetadataDetailResponse['data']
): UserMetadata => {
  return {
    id: metadataData.id,
    userId: metadataData.user_id,
    displayName: metadataData.display_name,
    profileImageUrl: metadataData.profile_image_url,
    tagline: metadataData.tagline,
    bio: metadataData.bio,
    vision: metadataData.vision,
    visionPublic: metadataData.vision_public,
    timezone: metadataData.timezone,
    skills: metadataData.skills,
    interests: metadataData.interests,
    createdAt: new Date(metadataData.created_at),
    updatedAt: new Date(metadataData.updated_at),
  };
};

/**
 * UserMetadataUpdateResponse のデータをUserMetadata型に変換
 */
export const mapUserMetadataUpdateResponseToUserMetadata = (
  updateData: UserMetadataUpdateResponse['data']
): Partial<UserMetadata> => {
  return {
    id: updateData.id,
    displayName: updateData.display_name,
    tagline: updateData.tagline,
    visionPublic: updateData.vision_public,
    skills: updateData.skills,
    interests: updateData.interests,
    updatedAt: new Date(updateData.updated_at),
  };
};