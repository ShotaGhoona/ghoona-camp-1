package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotificationController 通知管理コントローラー
type NotificationController struct {
	// TODO: UseCase依存注入（UseCase実装後）
	// notificationUseCase application.NotificationUseCase
}

// NewNotificationController NotificationControllerのコンストラクタ
func NewNotificationController( /* TODO: UseCase追加 */ ) *NotificationController {
	return &NotificationController{
		// TODO: UseCase注入
	}
}

// PutNotification 通知の既読ステータスを更新
// PUT /notifications/{notificationId}
func (nc *NotificationController) PutNotification(c *gin.Context) {
	notificationID := c.Param("notificationId")
	if notificationID == "" {
		SendValidationError(c, "Notification ID is required")
		return
	}

	// リクエストボディ取得
	var req struct {
		IsRead *bool `json:"is_read"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// err := nc.notificationUseCase.UpdateNotification(c.Request.Context(), notificationID, req)
	// if err != nil {
	//     if errors.Is(err, domain.ErrNotificationNotFound) {
	//         SendNotFoundError(c, "Notification")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"message": "Notification updated successfully"})

	SendInternalError(c)
}

// DeleteNotification 不要な通知を削除
// DELETE /notifications/{notificationId}
func (nc *NotificationController) DeleteNotification(c *gin.Context) {
	notificationID := c.Param("notificationId")
	if notificationID == "" {
		SendValidationError(c, "Notification ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// err := nc.notificationUseCase.DeleteNotification(c.Request.Context(), notificationID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrNotificationNotFound) {
	//         SendNotFoundError(c, "Notification")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }

	c.Status(http.StatusNoContent)
}