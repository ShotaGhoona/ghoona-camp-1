/**
 * User Social Links Entity Mappers
 * 
 * APIレスポンス型からフロントエンド用型への変換関数を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * social-links-entity/index.ts からエクスポートされた関数を使用してください。
 */

import type { 
  SocialLink, 
  UserSocialLinksListResponse, 
  SocialLinkCreateResponse, 
  SocialLinkUpdateResponse 
} from '../model/social-links-types';

/**
 * SocialLinkResponse のデータをSocialLink型に変換
 */
const mapSocialLinkResponseToSocialLink = (linkData: any): SocialLink => {
  return {
    id: linkData.id,
    platform: linkData.platform,
    url: linkData.url,
    title: linkData.title,
    isPublic: linkData.is_public,
    createdAt: new Date(linkData.created_at),
    updatedAt: new Date(linkData.updated_at),
  };
};

/**
 * UserSocialLinksListResponse のデータをSocialLink配列に変換
 */
export const mapUserSocialLinksListResponseToSocialLinks = (
  linksData: UserSocialLinksListResponse['data']
): SocialLink[] => {
  return linksData.social_links.map(mapSocialLinkResponseToSocialLink);
};

/**
 * SocialLinkCreateResponse のデータをSocialLink型に変換
 */
export const mapSocialLinkCreateResponseToSocialLink = (
  createData: SocialLinkCreateResponse['data']
): SocialLink => {
  return {
    id: createData.id,
    platform: createData.platform,
    url: createData.url,
    title: createData.title,
    isPublic: createData.is_public,
    createdAt: new Date(createData.created_at),
    updatedAt: new Date(createData.created_at), // created_at をupdated_atとして使用
  };
};

/**
 * SocialLinkUpdateResponse のデータをSocialLink型に変換
 */
export const mapSocialLinkUpdateResponseToSocialLink = (
  updateData: SocialLinkUpdateResponse['data']
): SocialLink => {
  return {
    id: updateData.id,
    platform: updateData.platform,
    url: updateData.url,
    title: updateData.title,
    isPublic: updateData.is_public,
    createdAt: new Date(), // UpdateResponseには含まれていない
    updatedAt: new Date(updateData.updated_at),
  };
};