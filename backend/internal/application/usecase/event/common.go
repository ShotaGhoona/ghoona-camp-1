// Package usecase 共通ヘルパー関数
package event

import (
	"context"

	"ghoona-camp-backend/internal/domain/common"
	"ghoona-camp-backend/internal/domain/event/repository"
	"ghoona-camp-backend/internal/domain/event/value"
)

// checkUserRegistration ユーザーの登録状況をチェック
func checkUserRegistration(ctx context.Context, repo repository.EventParticipantRepository, eventID common.UUID, userID string) bool {
	if userID == "" {
		return false
	}
	userUUID, err := common.ParseUUID(userID)
	if err != nil {
		return false
	}
	participant, err := repo.FindByEventAndUser(ctx, eventID, userUUID)
	return err == nil && participant != nil && participant.Status == value.ParticipantStatusRegistered
}

// getRecurrencePattern RecurrencePatternをstring pointerに変換
func getRecurrencePattern(pattern value.RecurrencePattern) *string {
	if pattern == value.RecurrencePatternNone {
		return nil
	}
	str := string(pattern)
	return &str
}

// getDiscordChannel DiscordChannelIDをstring pointerに変換
func getDiscordChannel(channelID string) *string {
	if channelID == "" {
		return nil
	}
	return &channelID
}