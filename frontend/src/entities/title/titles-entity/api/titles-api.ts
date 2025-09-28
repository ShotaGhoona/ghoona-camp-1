/**
 * Titles API Functions
 * 
 * 称号関連のAPI呼び出しを提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * titles-entity/index.ts からエクスポートされた関数を使用してください。
 */

import { apiClient } from '@/shared/api';
import type { 
  TitlesListResponse,
  TitleDetailResponse,
  TitlesQueryParams
} from '../model/titles-types';

/** 全称号（8段階）の一覧を取得 */
/** GET /api/v1/titles */
export const getTitles = async (params?: TitlesQueryParams): Promise<TitlesListResponse['data']> => {
  const { data } = await apiClient.get<TitlesListResponse>('/titles', { params });
  return data.data;
};

/** 指定称号の詳細情報を取得 */
/** GET /api/v1/titles/{titleId} */
export const getTitleDetail = async (titleId: string): Promise<TitleDetailResponse['data']> => {
  const { data } = await apiClient.get<TitleDetailResponse>(`/titles/${titleId}`);
  return data.data;
};