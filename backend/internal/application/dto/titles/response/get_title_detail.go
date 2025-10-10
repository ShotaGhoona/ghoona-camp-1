package response

// GetTitleDetailResponse GET /titles/{titleId} のレスポンス形式
type GetTitleDetailResponse struct {
	Data struct {
		Title GetTitleDetailTitle `json:"title"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetTitleDetailTitle GET /titles/{titleId} の称号詳細情報
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