/**
 * User Social Links API Functions
 * 
 * ユーザーソーシャルリンク関連のAPI呼び出しを提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * social-links-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { apiClient } from '@/shared/api';
import type { 
  UserSocialLinksListResponse,
  SocialLinkCreateResponse,
  SocialLinkUpdateResponse,
  CreateSocialLinkDto,
  UpdateSocialLinkDto
} from '../model/social-links-types';

/** ユーザーのSNSリンク一覧取得 */
/** GET /api/v1/users/{userId}/social-links */
export const getUserSocialLinks = async (userId: string): Promise<UserSocialLinksListResponse['data']['social_links']> => {
  const { data } = await apiClient.get<UserSocialLinksListResponse>(`/users/${userId}/social-links`);
  return data.data.social_links;
};

/** 新しいSNSリンクを追加 */
/** POST /api/v1/users/{userId}/social-links */
export const createSocialLink = async ({
  userId,
  data: linkData,
}: {
  userId: string;
  data: CreateSocialLinkDto;
}): Promise<SocialLinkCreateResponse['data']> => {
  const { data } = await apiClient.post<SocialLinkCreateResponse>(`/users/${userId}/social-links`, linkData);
  return data.data;
};

/** 既存のSNSリンクを更新 */
/** PUT /api/v1/users/{userId}/social-links/{linkId} */
export const updateSocialLink = async ({
  userId,
  linkId,
  data: linkData,
}: {
  userId: string;
  linkId: string;
  data: UpdateSocialLinkDto;
}): Promise<SocialLinkUpdateResponse['data']> => {
  const { data } = await apiClient.put<SocialLinkUpdateResponse>(`/users/${userId}/social-links/${linkId}`, linkData);
  return data.data;
};

/** SNSリンクを削除 */
/** DELETE /api/v1/users/{userId}/social-links/{linkId} */
export const deleteSocialLink = async ({
  userId,
  linkId,
}: {
  userId: string;
  linkId: string;
}): Promise<void> => {
  await apiClient.delete(`/users/${userId}/social-links/${linkId}`);
};