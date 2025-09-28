/**
 * User Rivals Entity Mappers
 * 
 * APIレスポンス型からフロントエンド用型への変換関数を提供します。
 * このファイルは内部実装のため、外部からは直接インポートせず、
 * rivals-entity/index.ts からエクスポートされた関数を使用してください。
 */

import type { 
  Rival, 
  RivalUser, 
  RivalsList, 
  UserRivalsListResponse, 
  RivalCreateResponse 
} from '../model/rivals-types';

/**
 * RivalResponse のデータをRivalUser型に変換
 */
const mapRivalResponseToRivalUser = (rivalData: any): RivalUser => {
  return {
    id: rivalData.rival_user.id,
    displayName: rivalData.rival_user.display_name,
    username: rivalData.rival_user.username,
    avatarUrl: rivalData.rival_user.avatar_url,
    currentTitle: {
      level: rivalData.rival_user.current_title.level,
      nameJp: rivalData.rival_user.current_title.name_jp,
      colorTheme: rivalData.rival_user.current_title.color_theme,
    },
    attendanceStats: {
      totalAttendanceDays: rivalData.rival_user.attendance_stats.total_attendance_days,
      currentStreakDays: rivalData.rival_user.attendance_stats.current_streak_days,
    },
  };
};

/**
 * UserRivalsListResponse のデータをRivalsList型に変換
 */
export const mapUserRivalsListResponseToRivalsList = (
  rivalsData: UserRivalsListResponse['data']
): RivalsList => {
  return {
    rivals: rivalsData.rivals.map((rivalResponse) => ({
      id: rivalResponse.id,
      rivalUser: mapRivalResponseToRivalUser(rivalResponse),
      createdAt: new Date(rivalResponse.created_at),
    })),
    count: rivalsData.count,
    maxRivals: rivalsData.max_rivals,
  };
};

/**
 * RivalCreateResponse のデータをRival型に変換
 */
export const mapRivalCreateResponseToRival = (
  createData: RivalCreateResponse['data']
): Rival => {
  return {
    id: createData.id,
    rivalUser: {
      id: createData.rival_user.id,
      displayName: createData.rival_user.display_name,
      username: createData.rival_user.username,
      avatarUrl: createData.rival_user.avatar_url,
      currentTitle: {
        level: createData.rival_user.current_title.level,
        nameJp: createData.rival_user.current_title.name_jp,
        colorTheme: createData.rival_user.current_title.color_theme,
      },
      attendanceStats: {
        totalAttendanceDays: 0, // CreateResponseには含まれていない
        currentStreakDays: 0,   // CreateResponseには含まれていない
      },
    },
    createdAt: new Date(createData.created_at),
  };
};