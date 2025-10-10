package response

// GetTitlesResponse GET /titles のレスポンス形式
type GetTitlesResponse struct {
	Data struct {
		Titles []GetTitlesTitle `json:"titles"`
	} `json:"data"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// GetTitlesTitle GET /titles の個別称号情報
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