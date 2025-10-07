package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dto "ghoona-camp-backend/internal/application/dto/event"
	eventUseCase "ghoona-camp-backend/internal/application/usecase/event"
	"ghoona-camp-backend/internal/interface/controller/common"
	"ghoona-camp-backend/internal/interface/middleware"
)

// EventController handles event management operations
type EventController struct {
	getEventsUseCase             eventUseCase.GetEventsUseCase
	postEventsUseCase            eventUseCase.PostEventsUseCase
	getEventByIDUseCase          eventUseCase.GetEventByIDUseCase
	putEventByIDUseCase          eventUseCase.PutEventByIDUseCase
	deleteEventByIDUseCase       eventUseCase.DeleteEventByIDUseCase
	getEventParticipantsUseCase  eventUseCase.GetEventParticipantsUseCase
	postEventParticipantsUseCase eventUseCase.PostEventParticipantsUseCase
	putEventParticipantUseCase   eventUseCase.PutEventParticipantUseCase
}

// NewEventController creates a new EventController instance
func NewEventController(
	getEventsUseCase eventUseCase.GetEventsUseCase,
	postEventsUseCase eventUseCase.PostEventsUseCase,
	getEventByIDUseCase eventUseCase.GetEventByIDUseCase,
	putEventByIDUseCase eventUseCase.PutEventByIDUseCase,
	deleteEventByIDUseCase eventUseCase.DeleteEventByIDUseCase,
	getEventParticipantsUseCase eventUseCase.GetEventParticipantsUseCase,
	postEventParticipantsUseCase eventUseCase.PostEventParticipantsUseCase,
	putEventParticipantUseCase eventUseCase.PutEventParticipantUseCase,
) *EventController {
	return &EventController{
		getEventsUseCase:             getEventsUseCase,
		postEventsUseCase:            postEventsUseCase,
		getEventByIDUseCase:          getEventByIDUseCase,
		putEventByIDUseCase:          putEventByIDUseCase,
		deleteEventByIDUseCase:       deleteEventByIDUseCase,
		getEventParticipantsUseCase:  getEventParticipantsUseCase,
		postEventParticipantsUseCase: postEventParticipantsUseCase,
		putEventParticipantUseCase:   putEventParticipantUseCase,
	}
}

// GetEvents retrieves a list of events
// GET /events?date_from=2024-01-01&date_to=2024-12-31&event_type=study&search=keyword&page=1&limit=20
func (c *EventController) GetEvents(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	// Query parameters
	dateFrom := getOptionalQuery(ctx, "date_from")
	dateTo := getOptionalQuery(ctx, "date_to")
	eventType := getOptionalQuery(ctx, "event_type")
	search := getOptionalQuery(ctx, "search")

	// Pagination parameters with defaults
	page := 1
	limit := 20

	if p, err := parseIntQuery(ctx, "page"); err == nil && p >= 1 {
		page = p
	}

	if l, err := parseIntQuery(ctx, "limit"); err == nil && l >= 1 && l <= 50 {
		limit = l
	}

	// Get optional user ID from middleware (認証任意)
	userID, _ := middleware.GetUserID(ctx)

	events, err := c.getEventsUseCase.Execute(reqCtx, dateFrom, dateTo, eventType, search, page, limit, userID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

// PostEvents creates a new event
// POST /events
func (c *EventController) PostEvents(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	var req dto.PostEventsRequest
	if !bindJSON(ctx, &req) {
		return
	}

	// Get required user ID from middleware (認証必須)
	userID, ok := middleware.RequireUserID(ctx)
	if !ok {
		return
	}

	event, err := c.postEventsUseCase.Execute(reqCtx, userID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

// GetEventByID retrieves a specific event by ID
// GET /events/{eventId}
func (c *EventController) GetEventByID(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	eventID, err := common.ParseUUIDParam(ctx, "eventId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "eventId", err)
		return
	}

	// Get optional user ID from middleware (認証任意)
	userID, _ := middleware.GetUserID(ctx)

	event, err := c.getEventByIDUseCase.Execute(reqCtx, eventID, userID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

// PutEventByID updates a specific event
// PUT /events/{eventId}
func (c *EventController) PutEventByID(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	eventID, err := common.ParseUUIDParam(ctx, "eventId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "eventId", err)
		return
	}

	// Get required user ID from middleware (認証必須)
	userID, ok := middleware.RequireUserID(ctx)
	if !ok {
		return
	}

	var req dto.PutEventByIDRequest
	if !bindJSON(ctx, &req) {
		return
	}

	event, err := c.putEventByIDUseCase.Execute(reqCtx, eventID, userID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

// DeleteEventByID deletes a specific event
// DELETE /events/{eventId}
func (c *EventController) DeleteEventByID(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	eventID, err := common.ParseUUIDParam(ctx, "eventId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "eventId", err)
		return
	}

	// Get required user ID from middleware (認証必須)
	userID, ok := middleware.RequireUserID(ctx)
	if !ok {
		return
	}

	err = c.deleteEventByIDUseCase.Execute(reqCtx, eventID, userID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetEventParticipants retrieves participants of a specific event
// GET /events/{eventId}/participants?status=registered
func (c *EventController) GetEventParticipants(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	eventID, err := common.ParseUUIDParam(ctx, "eventId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "eventId", err)
		return
	}

	status := getOptionalQuery(ctx, "status")

	participants, err := c.getEventParticipantsUseCase.Execute(reqCtx, eventID, status)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, participants)
}

// PostEventParticipants registers user participation in an event
// POST /events/{eventId}/participants
func (c *EventController) PostEventParticipants(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	eventID, err := common.ParseUUIDParam(ctx, "eventId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "eventId", err)
		return
	}

	// Get required user ID from middleware (認証必須)
	userID, ok := middleware.RequireUserID(ctx)
	if !ok {
		return
	}

	participant, err := c.postEventParticipantsUseCase.Execute(reqCtx, eventID, userID)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, participant)
}

// PutEventParticipant updates participant status
// PUT /events/{eventId}/participants/{userId}
func (c *EventController) PutEventParticipant(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	eventID, err := common.ParseUUIDParam(ctx, "eventId")
	if err != nil {
		common.HandleUUIDParamError(ctx, "eventId", err)
		return
	}

	participantUserID := ctx.Param("userId")
	if participantUserID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID parameter"})
		return
	}

	// Get required user ID from middleware (認証必須)
	authUserID, ok := middleware.RequireUserID(ctx)
	if !ok {
		return
	}

	var req dto.PutEventParticipantRequest
	if !bindJSON(ctx, &req) {
		return
	}

	participant, err := c.putEventParticipantUseCase.Execute(reqCtx, eventID, participantUserID, authUserID, &req)
	if err != nil {
		common.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, participant)
}

// Helper functions

func getOptionalQuery(ctx *gin.Context, key string) *string {
	value := ctx.Query(key)
	if value == "" {
		return nil
	}
	return &value
}

func parseIntQuery(ctx *gin.Context, key string) (int, error) {
	value := ctx.Query(key)
	if value == "" {
		return 0, nil
	}
	return strconv.Atoi(value)
}

func bindJSON(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return false
	}
	return true
}
