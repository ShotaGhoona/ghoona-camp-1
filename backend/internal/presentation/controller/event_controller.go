package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// EventController イベント管理コントローラー
type EventController struct {
	// TODO: UseCase依存注入（UseCase実装後）
	// eventUseCase application.EventUseCase
}

// NewEventController EventControllerのコンストラクタ
func NewEventController( /* TODO: UseCase追加 */ ) *EventController {
	return &EventController{
		// TODO: UseCase注入
	}
}

// GetEvents 開催予定/開催中のイベント一覧を取得
// GET /events?status=&creator=&date_from=&date_to=&limit=&offset=
func (ec *EventController) GetEvents(c *gin.Context) {
	// クエリパラメータ取得
	var query struct {
		PaginationQuery
		Status   string `form:"status" binding:"omitempty,oneof=upcoming ongoing past"`
		Creator  string `form:"creator"`
		DateFrom string `form:"date_from" binding:"omitempty"`
		DateTo   string `form:"date_to" binding:"omitempty"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		SendValidationError(c, "Invalid query parameters")
		return
	}

	// TODO: UseCase呼び出し
	// events, total, err := ec.eventUseCase.GetEvents(c.Request.Context(), query)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{
	//     "events": events,
	//     "pagination": PaginationMeta{
	//         Total:   total,
	//         Limit:   query.Limit,
	//         Offset:  query.Offset,
	//         HasMore: query.Offset+query.Limit < total,
	//     },
	// })

	SendInternalError(c)
}

// PostEvent 新しい朝活イベントを作成
// POST /events
func (ec *EventController) PostEvent(c *gin.Context) {
	// ユーザーコンテキストから作成者IDを取得
	creatorID := c.GetString("user_id")
	if creatorID == "" {
		SendUnauthorizedError(c)
		return
	}

	// リクエストボディ取得
	var req struct {
		Title              string `json:"title" binding:"required,min=1,max=200"`
		Description        string `json:"description" binding:"omitempty,max=1000"`
		EventType          string `json:"event_type" binding:"omitempty,max=50"`
		ScheduledDate      string `json:"scheduled_date" binding:"required,datetime=2006-01-02"`
		StartTime          string `json:"start_time" binding:"required,datetime=15:04"`
		EndTime            string `json:"end_time" binding:"required,datetime=15:04"`
		MaxParticipants    *int   `json:"max_participants" binding:"omitempty,min=1,max=100"`
		IsRecurring        bool   `json:"is_recurring"`
		RecurrencePattern  string `json:"recurrence_pattern" binding:"omitempty,oneof=daily weekly monthly"`
		DiscordChannelID   string `json:"discord_channel_id" binding:"omitempty,max=255"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// eventID, err := ec.eventUseCase.CreateEvent(c.Request.Context(), creatorID, req)
	// if err != nil {
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusCreated, gin.H{
	//     "event_id": eventID,
	//     "message":  "Event created successfully",
	// })

	SendInternalError(c)
}

// GetEventDetail 指定イベントの詳細情報を取得
// GET /events/{eventId}
func (ec *EventController) GetEventDetail(c *gin.Context) {
	eventID := c.Param("eventId")
	if eventID == "" {
		SendValidationError(c, "Event ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// event, err := ec.eventUseCase.GetEventDetail(c.Request.Context(), eventID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrEventNotFound) {
	//         SendNotFoundError(c, "Event")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"event": event})

	SendInternalError(c)
}

// PutEvent イベント情報を更新（作成者のみ）
// PUT /events/{eventId}
func (ec *EventController) PutEvent(c *gin.Context) {
	eventID := c.Param("eventId")
	if eventID == "" {
		SendValidationError(c, "Event ID is required")
		return
	}

	// リクエストボディ取得
	var req struct {
		Title           string `json:"title" binding:"omitempty,min=1,max=200"`
		Description     string `json:"description" binding:"omitempty,max=1000"`
		MaxParticipants *int   `json:"max_participants" binding:"omitempty,min=1,max=100"`
		IsActive        *bool  `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// err := ec.eventUseCase.UpdateEvent(c.Request.Context(), eventID, req)
	// if err != nil {
	//     if errors.Is(err, domain.ErrEventNotFound) {
	//         SendNotFoundError(c, "Event")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"message": "Event updated successfully"})

	SendInternalError(c)
}

// DeleteEvent イベントを削除（作成者のみ）
// DELETE /events/{eventId}
func (ec *EventController) DeleteEvent(c *gin.Context) {
	eventID := c.Param("eventId")
	if eventID == "" {
		SendValidationError(c, "Event ID is required")
		return
	}

	// TODO: UseCase呼び出し
	// err := ec.eventUseCase.DeleteEvent(c.Request.Context(), eventID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrEventNotFound) {
	//         SendNotFoundError(c, "Event")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }

	c.Status(http.StatusNoContent)
}

// PostEventParticipant イベントに参加登録
// POST /events/{eventId}/participants
func (ec *EventController) PostEventParticipant(c *gin.Context) {
	eventID := c.Param("eventId")
	if eventID == "" {
		SendValidationError(c, "Event ID is required")
		return
	}

	// ユーザーコンテキストから参加者IDを取得
	userID := c.GetString("user_id")
	if userID == "" {
		SendUnauthorizedError(c)
		return
	}

	// TODO: UseCase呼び出し
	// err := ec.eventUseCase.JoinEvent(c.Request.Context(), eventID, userID)
	// if err != nil {
	//     if errors.Is(err, domain.ErrEventNotFound) {
	//         SendNotFoundError(c, "Event")
	//         return
	//     }
	//     if errors.Is(err, domain.ErrEventFull) {
	//         SendErrorResponse(c, 409, "EVENT_FULL", "Event is full")
	//         return
	//     }
	//     if errors.Is(err, domain.ErrAlreadyJoined) {
	//         SendErrorResponse(c, 409, "ALREADY_JOINED", "Already joined this event")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusCreated, gin.H{"message": "Successfully joined the event"})

	SendInternalError(c)
}

// PutEventParticipant 参加ステータスを更新（参加者本人またはイベント作成者）
// PUT /events/{eventId}/participants/{userId}
func (ec *EventController) PutEventParticipant(c *gin.Context) {
	eventID := c.Param("eventId")
	userID := c.Param("userId")

	if eventID == "" || userID == "" {
		SendValidationError(c, "Event ID and User ID are required")
		return
	}

	// リクエストボディ取得
	var req struct {
		Status string `json:"status" binding:"required,oneof=registered cancelled"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		SendValidationError(c, "Invalid request body")
		return
	}

	// TODO: UseCase呼び出し
	// err := ec.eventUseCase.UpdateParticipationStatus(c.Request.Context(), eventID, userID, req.Status)
	// if err != nil {
	//     if errors.Is(err, domain.ErrEventNotFound) {
	//         SendNotFoundError(c, "Event")
	//         return
	//     }
	//     if errors.Is(err, domain.ErrParticipantNotFound) {
	//         SendNotFoundError(c, "Participant")
	//         return
	//     }
	//     SendInternalError(c)
	//     return
	// }
	// SendSuccessResponse(c, http.StatusOK, gin.H{"message": "Participation status updated successfully"})

	SendInternalError(c)
}