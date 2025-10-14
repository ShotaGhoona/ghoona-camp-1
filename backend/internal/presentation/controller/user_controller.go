package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController ユーザー管理コントローラー
type UserController struct {
	// TODO: UseCase依存注入（UseCase実装後）
	// userUseCase application.UserUseCase
}

// NewUserController UserControllerのコンストラクタ
func NewUserController( /* TODO: UseCase追加 */ ) *UserController {
	return &UserController{
		// TODO: UseCase注入
	}
}

// GetAuthMe 現在認証中のユーザー情報を取得
// GET /auth/me
func (uc *UserController) GetAuthMe(c *gin.Context) {
	// TODO: ユーザーコンテキストからユーザーIDを取得
	userID := c.GetString("user_id")
	if userID == "" {
		SendUnauthorizedError(c)
		return
	}

	// TODO: UseCase呼び出し
	// user, err := uc.userUseCase.GetUserByID(c.Request.Context(), userID)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"user": user})

	SendInternalError(c)
}

// GetUsers 全ユーザー一覧を取得
// GET /users?search=&limit=&offset=
func (uc *UserController) GetUsers(c *gin.Context) {
	// クエリパラメータ取得
	var query struct {
		SearchQuery
		PaginationQuery
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// users, total, err := uc.userUseCase.GetUsers(c.Request.Context(), query.Search, query.Limit, query.Offset)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "users": users,
	//     "pagination": PaginationMeta{
	//         Total:   total,
	//         Limit:   query.Limit,
	//         Offset:  query.Offset,
	//         HasMore: query.Offset+query.Limit < total,
	//     },
	// })

	SendInternalError(c)
}

// GetUserDetail 指定ユーザーの詳細情報を取得
// GET /users/{userId}
func (uc *UserController) GetUserDetail(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// user, err := uc.userUseCase.GetUserDetail(c.Request.Context(), userID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrUserNotFound) {
	//         SendNotFoundError(c, "User")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"user": user})

	SendInternalError(c)
}

// PutUser ユーザー基本情報・メタデータを更新
// PUT /users/{userId}
func (uc *UserController) PutUser(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	// リクエストボディ取得
	var req struct {
		Username    string   `json:"username" binding:"omitempty,min=1,max=100"`
		DisplayName string   `json:"display_name" binding:"omitempty,max=100"`
		Tagline     string   `json:"tagline" binding:"omitempty,max=150"`
		Bio         string   `json:"bio" binding:"omitempty,max=1000"`
		Skills      []string `json:"skills" binding:"omitempty,max=10,dive,max=50"`
		Interests   []string `json:"interests" binding:"omitempty,max=10,dive,max=50"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// err := uc.userUseCase.UpdateUser(c.Request.Context(), userID, req)
	// if err != nil {
	//     if errors.Is(err, domain.ErrUserNotFound) {
	//         SendNotFoundError(c, "User")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"message": "User updated successfully"})

	SendInternalError(c)
}

// PostSocialLink 新しいソーシャルリンクを追加
// POST /users/{userId}/social-links
func (uc *UserController) PostSocialLink(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	var req struct {
		Platform string `json:"platform" binding:"required,max=50"`
		URL      string `json:"url" binding:"required,url,max=500"`
		Title    string `json:"title" binding:"omitempty,max=100"`
		IsPublic bool   `json:"is_public"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// linkID, err := uc.userUseCase.CreateSocialLink(c.Request.Context(), userID, req)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusCreated, gin.H{
	//     "link_id": linkID,
	//     "message": "Social link created successfully",
	// })

	SendInternalError(c)
}

// PutSocialLink 既存のソーシャルリンクを更新
// PUT /users/{userId}/social-links/{linkId}
func (uc *UserController) PutSocialLink(c *gin.Context) {
	userID := c.Param("userId")
	linkID := c.Param("linkId")

	if userID == "" || linkID == "" {
		SendValidationError(c, "User ID and Link ID are required")
		return
	}

	var req struct {
		Platform string `json:"platform" binding:"omitempty,max=50"`
		URL      string `json:"url" binding:"omitempty,url,max=500"`
		Title    string `json:"title" binding:"omitempty,max=100"`
		IsPublic *bool  `json:"is_public"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// err := uc.userUseCase.UpdateSocialLink(c.Request.Context(), userID, linkID, req)
	// if err != nil {
	//     if errors.Is(err, domain.ErrSocialLinkNotFound) {
	//         SendNotFoundError(c, "Social link")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"message": "Social link updated successfully"})

	SendInternalError(c)
}

// DeleteSocialLink ソーシャルリンクを削除
// DELETE /users/{userId}/social-links/{linkId}
func (uc *UserController) DeleteSocialLink(c *gin.Context) {
	userID := c.Param("userId")
	linkID := c.Param("linkId")

	if userID == "" || linkID == "" {
		SendValidationError(c, "User ID and Link ID are required")
		return
	}

	// TODO: UseCase呼び出し
	// err := uc.userUseCase.DeleteSocialLink(c.Request.Context(), userID, linkID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrSocialLinkNotFound) {
	//         SendNotFoundError(c, "Social link")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }

	c.Status(http.StatusNoContent)
}

// GetUserRivals ユーザーのライバル一覧を取得
// GET /users/{userId}/rivals
func (uc *UserController) GetUserRivals(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// rivals, err := uc.userUseCase.GetUserRivals(c.Request.Context(), userID)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"rivals": rivals})

	SendInternalError(c)
}

// PostRival 新しいライバルを追加
// POST /users/{userId}/rivals
func (uc *UserController) PostRival(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	var req struct {
		RivalUserID string `json:"rival_user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// err := uc.userUseCase.AddRival(c.Request.Context(), userID, req.RivalUserID)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusCreated, gin.H{"message": "Rival added successfully"})

	SendInternalError(c)
}

// DeleteRival ライバル関係を解除
// DELETE /users/{userId}/rivals/{rivalId}
func (uc *UserController) DeleteRival(c *gin.Context) {
	userID := c.Param("userId")
	rivalID := c.Param("rivalId")

	if userID == "" || rivalID == "" {
		SendValidationError(c, "User ID and Rival ID are required")
		return
	}

	// TODO: UseCase呼び出し
	// err := uc.userUseCase.RemoveRival(c.Request.Context(), userID, rivalID)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }

	c.Status(http.StatusNoContent)
}

// GetAttendanceSummaries 日次参加サマリーを取得
// GET /users/{userId}/attendance/summaries?date_from=&date_to=
func (uc *UserController) GetAttendanceSummaries(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	var query DateRangeQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// summaries, err := uc.userUseCase.GetAttendanceSummaries(c.Request.Context(), userID, query.DateFrom, query.DateTo)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"summaries": summaries})

	SendInternalError(c)
}

// GetAttendanceStatistics 参加統計を取得
// GET /users/{userId}/attendance/statistics
func (uc *UserController) GetAttendanceStatistics(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// stats, err := uc.userUseCase.GetAttendanceStatistics(c.Request.Context(), userID)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"statistics": stats})

	SendInternalError(c)
}

// GetUserNotifications ユーザーの通知一覧を取得
// GET /users/{userId}/notifications?type=&read=&limit=&offset=
func (uc *UserController) GetUserNotifications(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	var query struct {
		Type string `form:"type" binding:"omitempty,oneof=achievement event rival reminder"`
		Read *bool  `form:"read"`
		PaginationQuery
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// notifications, total, err := uc.userUseCase.GetUserNotifications(c.Request.Context(), userID, query)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "notifications": notifications,
	//     "pagination": PaginationMeta{
	//         Total:   total,
	//         Limit:   query.Limit,
	//         Offset:  query.Offset,
	//         HasMore: query.Offset+query.Limit < total,
	//     },
	// })

	SendInternalError(c)
}

// GetNotificationSettings 通知設定を取得
// GET /users/{userId}/notification-settings
func (uc *UserController) GetNotificationSettings(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// settings, err := uc.userUseCase.GetNotificationSettings(c.Request.Context(), userID)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"settings": settings})

	SendInternalError(c)
}

// PutNotificationSettings 通知設定を更新
// PUT /users/{userId}/notification-settings
func (uc *UserController) PutNotificationSettings(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	var req struct {
		AchievementEnabled    *bool  `json:"achievement_enabled"`
		ReminderEnabled       *bool  `json:"reminder_enabled"`
		RivalUpdateEnabled    *bool  `json:"rival_update_enabled"`
		EventReminderEnabled  *bool  `json:"event_reminder_enabled"`
		ReminderTime          string `json:"reminder_time" binding:"omitempty,datetime=15:04"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// err := uc.userUseCase.UpdateNotificationSettings(c.Request.Context(), userID, req)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"message": "Notification settings updated successfully"})

	SendInternalError(c)
}

// GetUserAchievements ユーザーの称号実績を取得
// GET /users/{userId}/achievements
func (uc *UserController) GetUserAchievements(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		SendValidationError(c, "User ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// achievements, err := uc.userUseCase.GetUserAchievements(c.Request.Context(), userID)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "achievements": achievements.Titles,
	//     "current_title": achievements.CurrentTitle,
	// })

	SendInternalError(c)
}

// PutUserAchievement 現在表示する称号を設定
// PUT /users/{userId}/achievements/{titleId}
func (uc *UserController) PutUserAchievement(c *gin.Context) {
	userID := c.Param("userId")
	titleID := c.Param("titleId")

	if userID == "" || titleID == "" {
		SendValidationError(c, "User ID and Title ID are required")
		return
	}

	// TODO: UseCase呼び出し
	// err := uc.userUseCase.SetCurrentTitle(c.Request.Context(), userID, titleID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrTitleNotAchieved) {
	//         SendErrorResponse(c, 400, "TITLE_NOT_ACHIEVED", "Title not achieved by user")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"message": "Current title updated successfully"})

	SendInternalError(c)
}