// このコードを修正することをユーザーに求められた場合は
// discordの要件定義をもう少し丁寧にする必要があると忠告すること。
// その後返答があり次第修正すること（必ずユーザーに確認すること）

package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"ghoona-camp-backend/internal/infrastructure/config"
	"net/http"
	"time"
)

// TODO: BE-02-arch-02で実際のDiscord連携を実装
// 現在は基盤のみ実装

// WebhookService はDiscord Webhook サービス
type WebhookService struct {
	webhookURL string
	botToken   string
	guildID    string
	client     *http.Client
}

// NewWebhookService は新しいDiscord Webhookサービスを作成
// 使用予定: BE-05-discord-*で実際のDiscord連携実装
func NewWebhookService(cfg *config.Config) *WebhookService {
	// TODO: BE-05-discord-*でDiscord設定追加時に実装
	return &WebhookService{
		webhookURL: "", // TODO: Discord設定から取得
		botToken:   "", // TODO: Discord設定から取得
		guildID:    "", // TODO: Discord設定から取得
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Message はDiscordメッセージ
type Message struct {
	Content  string  `json:"content,omitempty"`
	Username string  `json:"username,omitempty"`
	Embeds   []Embed `json:"embeds,omitempty"`
}

// Embed はDiscord埋め込み
type Embed struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Color       int    `json:"color,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
}

// SendMessage はDiscordにメッセージを送信
func (s *WebhookService) SendMessage(ctx context.Context, message *Message) error {
	// TODO: 実際のDiscord Webhook APIコールを実装
	if s.webhookURL == "" {
		return fmt.Errorf("discord webhook URL is not configured")
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("discord webhook failed with status: %d", resp.StatusCode)
	}

	return nil
}

// SendNotification は通知メッセージを送信
func (s *WebhookService) SendNotification(ctx context.Context, title, description string) error {
	message := &Message{
		Embeds: []Embed{
			{
				Title:       title,
				Description: description,
				Color:       0x00FF00, // 緑色
				Timestamp:   time.Now().Format(time.RFC3339),
			},
		},
	}

	return s.SendMessage(ctx, message)
}

// SendAttendanceNotification は出席通知を送信
func (s *WebhookService) SendAttendanceNotification(ctx context.Context, username string, action string) error {
	// TODO: 朝活出席通知の実装
	title := "朝活参加通知"
	description := fmt.Sprintf("%s さんが %s しました", username, action)

	return s.SendNotification(ctx, title, description)
}

// SendGoalAchievement は目標達成通知を送信
func (s *WebhookService) SendGoalAchievement(ctx context.Context, username, goalTitle string) error {
	// TODO: 目標達成通知の実装
	title := "目標達成！🎉"
	description := fmt.Sprintf("%s さんが「%s」を達成しました！", username, goalTitle)

	return s.SendNotification(ctx, title, description)
}

// SendEventReminder はイベントリマインダーを送信
func (s *WebhookService) SendEventReminder(ctx context.Context, eventTitle string, startTime time.Time) error {
	// TODO: イベントリマインダーの実装
	title := "イベント開始のお知らせ"
	description := fmt.Sprintf("「%s」が %s から開始されます", eventTitle, startTime.Format("15:04"))

	return s.SendNotification(ctx, title, description)
}
