// Package shared イベントDTO間で共通利用される構造体を定義する
// 複数のAPIエンドポイントで使用される共通データ構造を集約
package dto

// CreatorDTO イベント作成者情報
type CreatorDTO struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	AvatarURL   string `json:"avatar_url"`
}

// UserDTO ユーザー情報
type UserDTO struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	AvatarURL   string `json:"avatar_url"`
}
