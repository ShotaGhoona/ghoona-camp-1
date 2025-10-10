# DTO実装サマリー

## リビジョン履歴

| バージョン | 日付 | 変更内容 |
|-----------|------|----------|
| v1.0 | 2025-01-10 | 初版を作成 |

## 概要
本ドキュメントは、Ghoona Camp バックエンドアプリケーションにおけるDTO（Data Transfer Object）の実装状況をまとめたものです。
各APIドメインのリクエスト・レスポンス構造とフィールド定義を網羅的に記載しています。

## DTO構造の全体像

### 実装済みAPIドメイン
- Users API
- Goals API  
- Titles API
- Attendance API
- Events API

---

## 1. Users API

### Request DTOs

#### PostRivalRequest
**用途**: POST /users/{userId}/rivals
```go
type PostRivalRequest struct {
    RivalUserID string `json:"rival_user_id" validate:"required,uuid"`
}
```

#### PostSocialLinkRequest  
**用途**: POST /users/{userId}/social-links
```go
type PostSocialLinkRequest struct {
    Platform string `json:"platform" validate:"required,min=1,max=50"`
    URL      string `json:"url" validate:"required,url,max=2048"`
    Title    string `json:"title" validate:"required,min=1,max=100"`
    IsPublic bool   `json:"is_public"`
}
```

#### PutSocialLinkRequest
**用途**: PUT /users/{userId}/social-links/{linkId}
```go
type PutSocialLinkRequest struct {
    URL      *string `json:"url,omitempty" validate:"omitempty,url,max=2048"`
    Title    *string `json:"title,omitempty" validate:"omitempty,min=1,max=100"`
    IsPublic *bool   `json:"is_public,omitempty"`
}
```

#### PutUserRequest
**用途**: PUT /users/{userId}
```go
type PutUserRequest struct {
    Username  *string           `json:"username,omitempty" validate:"omitempty,min=1,max=100"`
    AvatarURL *string           `json:"avatar_url,omitempty" validate:"omitempty,url"`
    Metadata  *UserMetadataForm `json:"metadata,omitempty"`
}

type UserMetadataForm struct {
    DisplayName *string   `json:"display_name,omitempty" validate:"omitempty,min=1,max=100"`
    Tagline     *string   `json:"tagline,omitempty" validate:"omitempty,max=150"`
    Bio         *string   `json:"bio,omitempty" validate:"omitempty,max=1000"`
    Skills      *[]string `json:"skills,omitempty" validate:"omitempty,dive,min=1,max=50"`
    Interests   *[]string `json:"interests,omitempty" validate:"omitempty,dive,min=1,max=50"`
}
```

### Response DTOs

#### GetAuthMeResponse
**用途**: GET /auth/me
```go
type GetAuthMeResponse struct {
    Data struct {
        User GetAuthMeUser `json:"user"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetAuthMeUser struct {
    ID          string `json:"id"`
    ClerkID     string `json:"clerk_id"`
    Email       string `json:"email"`
    Username    string `json:"username"`
    DisplayName string `json:"display_name,omitempty"`
    AvatarURL   string `json:"avatar_url,omitempty"`
    DiscordID   string `json:"discord_id,omitempty"`
    IsActive    bool   `json:"is_active"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}
```

#### GetUsersResponse
**用途**: GET /users
```go
type GetUsersResponse struct {
    Data struct {
        Users      []GetUsersUser     `json:"users"`
        Pagination *common.Pagination `json:"pagination,omitempty"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetUsersUser struct {
    ID          string `json:"id"`
    AvatarURL   string `json:"avatar_url,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Tagline     string `json:"tagline,omitempty"`
    IsActive    bool   `json:"is_active"`
    CreatedAt   string `json:"created_at"`
}
```

#### GetUserDetailResponse
**用途**: GET /users/{userId}
```go
type GetUserDetailResponse struct {
    Data struct {
        User GetUserDetailUser `json:"user"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetUserDetailUser struct {
    ID          string                       `json:"id"`
    Username    string                       `json:"username"`
    DisplayName string                       `json:"display_name"`
    AvatarURL   string                       `json:"avatar_url,omitempty"`
    IsActive    bool                         `json:"is_active"`
    CreatedAt   string                       `json:"created_at"`
    UpdatedAt   string                       `json:"updated_at"`
    Metadata    *GetUserDetailMetadata       `json:"metadata,omitempty"`
    SocialLinks []GetUserDetailSocialLink    `json:"social_links,omitempty"`
}

type GetUserDetailMetadata struct {
    ID          string   `json:"id"`
    DisplayName string   `json:"display_name"`
    Tagline     string   `json:"tagline,omitempty"`
    Bio         string   `json:"bio,omitempty"`
    Skills      []string `json:"skills,omitempty"`
    Interests   []string `json:"interests,omitempty"`
    CreatedAt   string   `json:"created_at"`
    UpdatedAt   string   `json:"updated_at"`
}

type GetUserDetailSocialLink struct {
    ID        string `json:"id"`
    Platform  string `json:"platform"`
    URL       string `json:"url"`
    Title     string `json:"title"`
    IsPublic  bool   `json:"is_public"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
```

#### GetUserRivalsResponse
**用途**: GET /users/{userId}/rivals
```go
type GetUserRivalsResponse struct {
    Data struct {
        Rivals []GetUserRivalsRival `json:"rivals"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetUserRivalsRival struct {
    ID          string                    `json:"id"`
    UserID      string                    `json:"user_id"`
    RivalUserID string                    `json:"rival_user_id"`
    CreatedAt   string                    `json:"created_at"`
    UpdatedAt   string                    `json:"updated_at"`
    RivalUser   GetUserRivalsRivalUser    `json:"rival_user"`
}

type GetUserRivalsRivalUser struct {
    ID        string `json:"id"`
    Username  string `json:"username"`
    AvatarURL string `json:"avatar_url,omitempty"`
}
```

#### PostRivalResponse
**用途**: POST /users/{userId}/rivals
```go
type PostRivalResponse struct {
    Data struct {
        Rival PostRivalRival `json:"rival"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PostRivalRival struct {
    ID          string `json:"id"`
    UserID      string `json:"user_id"`
    RivalUserID string `json:"rival_user_id"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}
```

#### PostSocialLinkResponse
**用途**: POST /users/{userId}/social-links
```go
type PostSocialLinkResponse struct {
    Data struct {
        SocialLink PostSocialLinkSocialLink `json:"social_link"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PostSocialLinkSocialLink struct {
    ID        string `json:"id"`
    UserID    string `json:"user_id"`
    Platform  string `json:"platform"`
    URL       string `json:"url"`
    Title     string `json:"title"`
    IsPublic  bool   `json:"is_public"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
```

#### PutSocialLinkResponse
**用途**: PUT /users/{userId}/social-links/{linkId}
```go
type PutSocialLinkResponse struct {
    Data struct {
        SocialLink PutSocialLinkSocialLink `json:"social_link"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PutSocialLinkSocialLink struct {
    ID        string `json:"id"`
    UserID    string `json:"user_id"`
    Platform  string `json:"platform"`
    URL       string `json:"url"`
    Title     string `json:"title"`
    IsPublic  bool   `json:"is_public"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
```

#### PutUserResponse
**用途**: PUT /users/{userId}
```go
type PutUserResponse struct {
    Data struct {
        User PutUserUser `json:"user"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PutUserUser struct {
    ID          string              `json:"id"`
    ClerkID     string              `json:"clerk_id"`
    Email       string              `json:"email"`
    Username    string              `json:"username"`
    DisplayName string              `json:"display_name,omitempty"`
    AvatarURL   string              `json:"avatar_url,omitempty"`
    DiscordID   string              `json:"discord_id,omitempty"`
    IsActive    bool                `json:"is_active"`
    CreatedAt   string              `json:"created_at"`
    UpdatedAt   string              `json:"updated_at"`
    Metadata    *PutUserMetadata    `json:"metadata,omitempty"`
}

type PutUserMetadata struct {
    ID          string   `json:"id"`
    DisplayName string   `json:"display_name,omitempty"`
    Tagline     string   `json:"tagline,omitempty"`
    Bio         string   `json:"bio,omitempty"`
    Skills      []string `json:"skills,omitempty"`
    Interests   []string `json:"interests,omitempty"`
    CreatedAt   string   `json:"created_at"`
    UpdatedAt   string   `json:"updated_at"`
}
```

---

## 2. Goals API

### Request DTOs

#### PostGoalRequest
**用途**: POST /goals
```go
type PostGoalRequest struct {
    UserID      string `json:"user_id" validate:"required,uuid"`
    Title       string `json:"title" validate:"required,min=1,max=200"`
    Description string `json:"description" validate:"required,min=1,max=1000"`
    StartedAt   string `json:"started_at" validate:"required,date"`
    EndedAt     string `json:"ended_at" validate:"omitempty,date"`
    IsPublic    bool   `json:"is_public"`
}
```

#### PutGoalRequest
**用途**: PUT /goals/{goalId}
```go
type PutGoalRequest struct {
    Title       *string `json:"title,omitempty" validate:"omitempty,min=1,max=200"`
    Description *string `json:"description,omitempty" validate:"omitempty,min=1,max=1000"`
    EndedAt     *string `json:"ended_at,omitempty" validate:"omitempty,date"`
    IsPublic    *bool   `json:"is_public,omitempty"`
    IsActive    *bool   `json:"is_active,omitempty"`
}
```

### Response DTOs

#### GetGoalsMeResponse
**用途**: GET /goals/me
```go
type GetGoalsMeResponse struct {
    Data struct {
        Goals      []GetGoalsMeGoal   `json:"goals"`
        Pagination *common.Pagination `json:"pagination,omitempty"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetGoalsMeGoal struct {
    ID          string `json:"id"`
    UserID      string `json:"user_id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    StartedAt   string `json:"started_at"`
    EndedAt     string `json:"ended_at,omitempty"`
    IsActive    bool   `json:"is_active"`
    IsPublic    bool   `json:"is_public"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}
```

#### GetGoalsPublicResponse
**用途**: GET /goals/public
```go
type GetGoalsPublicResponse struct {
    Data struct {
        Goals      []GetGoalsPublicGoal `json:"goals"`
        Pagination *common.Pagination   `json:"pagination,omitempty"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetGoalsPublicGoal struct {
    ID          string                  `json:"id"`
    UserID      string                  `json:"user_id"`
    Title       string                  `json:"title"`
    Description string                  `json:"description"`
    StartedAt   string                  `json:"started_at"`
    EndedAt     string                  `json:"ended_at,omitempty"`
    IsActive    bool                    `json:"is_active"`
    IsPublic    bool                    `json:"is_public"`
    CreatedAt   string                  `json:"created_at"`
    UpdatedAt   string                  `json:"updated_at"`
    User        GetGoalsPublicUser      `json:"user"`
}

type GetGoalsPublicUser struct {
    ID        string `json:"id"`
    Username  string `json:"username"`
    AvatarURL string `json:"avatar_url,omitempty"`
}
```

#### PostGoalResponse
**用途**: POST /goals
```go
type PostGoalResponse struct {
    Data struct {
        Goal PostGoalGoal `json:"goal"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PostGoalGoal struct {
    ID          string `json:"id"`
    UserID      string `json:"user_id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    StartedAt   string `json:"started_at"`
    EndedAt     string `json:"ended_at,omitempty"`
    IsActive    bool   `json:"is_active"`
    IsPublic    bool   `json:"is_public"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}
```

#### PutGoalResponse
**用途**: PUT /goals/{goalId}
```go
type PutGoalResponse struct {
    Data struct {
        Goal PutGoalGoal `json:"goal"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PutGoalGoal struct {
    ID          string `json:"id"`
    UserID      string `json:"user_id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    StartedAt   string `json:"started_at"`
    EndedAt     string `json:"ended_at,omitempty"`
    IsActive    bool   `json:"is_active"`
    IsPublic    bool   `json:"is_public"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}
```

---

## 3. Titles API

### Response DTOs

#### GetTitlesResponse
**用途**: GET /titles
```go
type GetTitlesResponse struct {
    Data struct {
        Titles []GetTitlesTitle `json:"titles"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetTitlesTitle struct {
    ID          string `json:"id"`
    Level       int    `json:"level"`
    NameJP      string `json:"name_jp"`
    NameEN      string `json:"name_en"`
    Description string `json:"description"`
    RequiredDays int   `json:"required_days"`
    ImageURL    string `json:"image_url,omitempty"`
    ColorTheme  string `json:"color_theme,omitempty"`
    IsActive    bool   `json:"is_active"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}
```

#### GetTitleDetailResponse
**用途**: GET /titles/{titleId}
```go
type GetTitleDetailResponse struct {
    Data struct {
        Title GetTitleDetailTitle `json:"title"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetTitleDetailTitle struct {
    ID               string  `json:"id"`
    Level            int     `json:"level"`
    NameJP           string  `json:"name_jp"`
    NameEN           string  `json:"name_en"`
    Description      string  `json:"description"`
    RequiredDays     int     `json:"required_days"`
    ImageURL         string  `json:"image_url,omitempty"`
    ColorTheme       string  `json:"color_theme,omitempty"`
    IsActive         bool    `json:"is_active"`
    CreatedAt        string  `json:"created_at"`
    UpdatedAt        string  `json:"updated_at"`
    HolderCount      int     `json:"holder_count"`
    RarityPercentage float64 `json:"rarity_percentage"`
}
```

#### GetUserAchievementsResponse
**用途**: GET /users/{userId}/achievements
```go
type GetUserAchievementsResponse struct {
    Data struct {
        Achievements []GetUserAchievementsAchievement `json:"achievements"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetUserAchievementsAchievement struct {
    ID          string                         `json:"id"`
    UserID      string                         `json:"user_id"`
    TitleID     string                         `json:"title_id"`
    AchievedAt  string                         `json:"achieved_at"`
    IsCurrent   bool                           `json:"is_current"`
    CreatedAt   string                         `json:"created_at"`
    UpdatedAt   string                         `json:"updated_at"`
    Title       GetUserAchievementsTitle       `json:"title"`
}

type GetUserAchievementsTitle struct {
    ID          string `json:"id"`
    Level       int    `json:"level"`
    NameJP      string `json:"name_jp"`
    NameEN      string `json:"name_en"`
    Description string `json:"description"`
    RequiredDays int   `json:"required_days"`
    ImageURL    string `json:"image_url,omitempty"`
    ColorTheme  string `json:"color_theme,omitempty"`
}
```

#### PutUserAchievementResponse
**用途**: PUT /users/{userId}/achievements/{titleId}
```go
type PutUserAchievementResponse struct {
    Data struct {
        Achievement PutUserAchievementAchievement `json:"achievement"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PutUserAchievementAchievement struct {
    ID         string `json:"id"`
    UserID     string `json:"user_id"`
    TitleID    string `json:"title_id"`
    AchievedAt string `json:"achieved_at"`
    IsCurrent  bool   `json:"is_current"`
    CreatedAt  string `json:"created_at"`
    UpdatedAt  string `json:"updated_at"`
}
```

---

## 4. Attendance API

### Response DTOs

#### GetAttendanceSummariesResponse
**用途**: GET /users/{userId}/attendance/summaries
```go
type GetAttendanceSummariesResponse struct {
    Data struct {
        Summaries []GetAttendanceSummariesSummary `json:"summaries"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetAttendanceSummariesSummary struct {
    ID                   string `json:"id"`
    UserID               string `json:"user_id"`
    Date                 string `json:"date"`
    TotalDurationMinutes int    `json:"total_duration_minutes"`
    SessionCount         int    `json:"session_count"`
    FirstJoinTime        string `json:"first_join_time,omitempty"`
    LastLeaveTime        string `json:"last_leave_time,omitempty"`
    IsMorningActive      bool   `json:"is_morning_active"`
    CreatedAt            string `json:"created_at"`
    UpdatedAt            string `json:"updated_at"`
}
```

#### GetAttendanceStatisticsResponse
**用途**: GET /users/{userId}/attendance/statistics
```go
type GetAttendanceStatisticsResponse struct {
    Data struct {
        Statistics GetAttendanceStatisticsStatistics `json:"statistics"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetAttendanceStatisticsStatistics struct {
    ID                   string `json:"id"`
    UserID               string `json:"user_id"`
    TotalAttendanceDays  int    `json:"total_attendance_days"`
    CurrentStreakDays    int    `json:"current_streak_days"`
    MaxStreakDays        int    `json:"max_streak_days"`
    LastAttendanceDate   string `json:"last_attendance_date,omitempty"`
    FirstAttendanceDate  string `json:"first_attendance_date,omitempty"`
    TotalDurationMinutes int    `json:"total_duration_minutes"`
    CreatedAt            string `json:"created_at"`
    UpdatedAt            string `json:"updated_at"`
}
```

#### GetRankingMonthlyResponse
**用途**: GET /ranking/monthly
```go
type GetRankingMonthlyResponse struct {
    Data struct {
        Ranking    []GetRankingMonthlyRanking `json:"ranking"`
        Pagination *common.Pagination         `json:"pagination,omitempty"`
        Month      int                        `json:"month"`
        Year       int                        `json:"year"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetRankingMonthlyRanking struct {
    Rank                    int    `json:"rank"`
    UserID                  string `json:"user_id"`
    Username                string `json:"username"`
    AvatarURL               string `json:"avatar_url,omitempty"`
    MonthlyAttendanceDays   int    `json:"monthly_attendance_days"`
    TotalDurationMinutes    int    `json:"total_duration_minutes"`
}
```

#### GetRankingTotalResponse
**用途**: GET /ranking/total
```go
type GetRankingTotalResponse struct {
    Data struct {
        Ranking    []GetRankingTotalRanking `json:"ranking"`
        Pagination *common.Pagination       `json:"pagination,omitempty"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetRankingTotalRanking struct {
    Rank                 int    `json:"rank"`
    UserID               string `json:"user_id"`
    Username             string `json:"username"`
    AvatarURL            string `json:"avatar_url,omitempty"`
    TotalAttendanceDays  int    `json:"total_attendance_days"`
    TotalDurationMinutes int    `json:"total_duration_minutes"`
    FirstAttendanceDate  string `json:"first_attendance_date,omitempty"`
}
```

#### GetRankingStreakResponse
**用途**: GET /ranking/streak
```go
type GetRankingStreakResponse struct {
    Data struct {
        Ranking    []GetRankingStreakRanking `json:"ranking"`
        Pagination *common.Pagination        `json:"pagination,omitempty"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetRankingStreakRanking struct {
    Rank                int    `json:"rank"`
    UserID              string `json:"user_id"`
    Username            string `json:"username"`
    AvatarURL           string `json:"avatar_url,omitempty"`
    CurrentStreakDays   int    `json:"current_streak_days"`
    MaxStreakDays       int    `json:"max_streak_days"`
    LastAttendanceDate  string `json:"last_attendance_date,omitempty"`
}
```

---

## 5. Events API

### Request DTOs

#### PostEventRequest
**用途**: POST /events
```go
type PostEventRequest struct {
    Title               string  `json:"title" validate:"required,max=200"`
    Description         string  `json:"description" validate:"required"`
    EventType           string  `json:"event_type" validate:"required,max=50"`
    ScheduledDate       string  `json:"scheduled_date" validate:"required"`
    StartTime           string  `json:"start_time" validate:"required"`
    EndTime             string  `json:"end_time" validate:"required"`
    MaxParticipants     int     `json:"max_participants" validate:"required,min=1"`
    IsRecurring         bool    `json:"is_recurring"`
    RecurrencePattern   *string `json:"recurrence_pattern,omitempty"`
    DiscordChannelID    string  `json:"discord_channel_id" validate:"required,max=255"`
}
```

#### PutEventRequest
**用途**: PUT /events/{eventId}
```go
type PutEventRequest struct {
    Title           *string `json:"title,omitempty" validate:"omitempty,max=200"`
    Description     *string `json:"description,omitempty"`
    MaxParticipants *int    `json:"max_participants,omitempty" validate:"omitempty,min=1"`
    IsActive        *bool   `json:"is_active,omitempty"`
}
```

#### PutEventParticipantRequest
**用途**: PUT /events/{eventId}/participants/{userId}
```go
type PutEventParticipantRequest struct {
    Status string `json:"status" validate:"required,oneof=registered cancelled"`
}
```

### Response DTOs

#### GetEventsResponse
**用途**: GET /events
```go
type GetEventsResponse struct {
    Data struct {
        Events     []GetEventsEvent      `json:"events"`
        Pagination *common.Pagination    `json:"pagination,omitempty"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetEventsEvent struct {
    ID                string               `json:"id"`
    CreatorID         string               `json:"creator_id"`
    Title             string               `json:"title"`
    Description       string               `json:"description"`
    EventType         string               `json:"event_type"`
    ScheduledDate     string               `json:"scheduled_date"`
    StartTime         string               `json:"start_time"`
    EndTime           string               `json:"end_time"`
    MaxParticipants   int                  `json:"max_participants"`
    IsRecurring       bool                 `json:"is_recurring"`
    RecurrencePattern *string              `json:"recurrence_pattern"`
    DiscordChannelID  string               `json:"discord_channel_id"`
    IsActive          bool                 `json:"is_active"`
    CreatedAt         string               `json:"created_at"`
    UpdatedAt         string               `json:"updated_at"`
    Creator           GetEventsCreator     `json:"creator"`
}

type GetEventsCreator struct {
    ID        string `json:"id"`
    Username  string `json:"username"`
    AvatarURL string `json:"avatar_url,omitempty"`
}
```

#### GetEventDetailResponse
**用途**: GET /events/{eventId}
```go
type GetEventDetailResponse struct {
    Data struct {
        Event GetEventDetailEvent `json:"event"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type GetEventDetailEvent struct {
    ID                string                        `json:"id"`
    CreatorID         string                        `json:"creator_id"`
    Title             string                        `json:"title"`
    Description       string                        `json:"description"`
    EventType         string                        `json:"event_type"`
    ScheduledDate     string                        `json:"scheduled_date"`
    StartTime         string                        `json:"start_time"`
    EndTime           string                        `json:"end_time"`
    MaxParticipants   int                           `json:"max_participants"`
    IsRecurring       bool                          `json:"is_recurring"`
    RecurrencePattern *string                       `json:"recurrence_pattern"`
    DiscordChannelID  string                        `json:"discord_channel_id"`
    IsActive          bool                          `json:"is_active"`
    CreatedAt         string                        `json:"created_at"`
    UpdatedAt         string                        `json:"updated_at"`
    Creator           GetEventDetailCreator         `json:"creator"`
    Participants      []GetEventDetailParticipant   `json:"participants"`
}

type GetEventDetailCreator struct {
    ID        string `json:"id"`
    Username  string `json:"username"`
    AvatarURL string `json:"avatar_url,omitempty"`
}

type GetEventDetailParticipant struct {
    ID        string `json:"id"`
    Username  string `json:"username"`
    AvatarURL string `json:"avatar_url,omitempty"`
    Status    string `json:"status"`
}
```

#### PostEventResponse
**用途**: POST /events
```go
type PostEventResponse struct {
    Data struct {
        Event PostEventEvent `json:"event"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PostEventEvent struct {
    ID                string  `json:"id"`
    CreatorID         string  `json:"creator_id"`
    Title             string  `json:"title"`
    Description       string  `json:"description"`
    EventType         string  `json:"event_type"`
    ScheduledDate     string  `json:"scheduled_date"`
    StartTime         string  `json:"start_time"`
    EndTime           string  `json:"end_time"`
    MaxParticipants   int     `json:"max_participants"`
    IsRecurring       bool    `json:"is_recurring"`
    RecurrencePattern *string `json:"recurrence_pattern"`
    DiscordChannelID  string  `json:"discord_channel_id"`
    IsActive          bool    `json:"is_active"`
    CreatedAt         string  `json:"created_at"`
    UpdatedAt         string  `json:"updated_at"`
}
```

#### PutEventResponse
**用途**: PUT /events/{eventId}
```go
type PutEventResponse struct {
    Data struct {
        Event PutEventEvent `json:"event"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PutEventEvent struct {
    ID                string  `json:"id"`
    CreatorID         string  `json:"creator_id"`
    Title             string  `json:"title"`
    Description       string  `json:"description"`
    EventType         string  `json:"event_type"`
    ScheduledDate     string  `json:"scheduled_date"`
    StartTime         string  `json:"start_time"`
    EndTime           string  `json:"end_time"`
    MaxParticipants   int     `json:"max_participants"`
    IsRecurring       bool    `json:"is_recurring"`
    RecurrencePattern *string `json:"recurrence_pattern"`
    DiscordChannelID  string  `json:"discord_channel_id"`
    IsActive          bool    `json:"is_active"`
    CreatedAt         string  `json:"created_at"`
    UpdatedAt         string  `json:"updated_at"`
}
```

#### PostEventParticipantResponse
**用途**: POST /events/{eventId}/participants
```go
type PostEventParticipantResponse struct {
    Data struct {
        Participation PostEventParticipation `json:"participation"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PostEventParticipation struct {
    ID        string `json:"id"`
    EventID   string `json:"event_id"`
    UserID    string `json:"user_id"`
    Status    string `json:"status"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
```

#### PutEventParticipantResponse
**用途**: PUT /events/{eventId}/participants/{userId}
```go
type PutEventParticipantResponse struct {
    Data struct {
        Participation PutEventParticipation `json:"participation"`
    } `json:"data"`
    Message   string `json:"message"`
    Timestamp string `json:"timestamp"`
}

type PutEventParticipation struct {
    ID        string `json:"id"`
    EventID   string `json:"event_id"`
    UserID    string `json:"user_id"`
    Status    string `json:"status"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
```

---

## 共通パターン

### バリデーションタグ
- `required`: 必須フィールド
- `omitempty`: 空値時は無視
- `uuid`: UUID形式の検証
- `email`: メールアドレス形式の検証
- `url`: URL形式の検証
- `min=N,max=N`: 長さの制限
- `oneof=value1 value2`: 列挙値の制限
- `dive`: 配列要素への検証適用

### JSONタグパターン
- `,omitempty`: 空値時はJSONから除外
- スネークケース命名規則に従ったフィールド名

### 共通レスポンス構造
全てのレスポンスで以下の構造を共有：
```go
type Response struct {
    Data      interface{} `json:"data"`
    Message   string      `json:"message"`
    Timestamp string      `json:"timestamp"`
}
```

### ページネーション
一覧系APIでは`common.Pagination`を使用：
```go
type Pagination struct {
    Total   int  `json:"total"`
    Limit   int  `json:"limit"`
    Offset  int  `json:"offset"`
    HasMore bool `json:"has_more"`
}
```

---

## ディレクトリ構造

```
backend/internal/application/dto/
├── attendance/
│   └── response/
├── events/
│   ├── request/
│   └── response/
├── goals/
│   ├── request/
│   └── response/
├── titles/
│   └── response/
├── users/
│   ├── request/
│   └── response/
└── common/
    └── pagination.go
```

---

## 備考

- 全てのDTOは仕様書との整合性が確認済み
- バリデーションルールは業務要件に基づいて設定
- omitemptyタグはオプショナルフィールドに適用
- 日時フィールドはISO 8601形式の文字列として扱い
- UUIDはstring型で定義し、バリデーションで形式チェック