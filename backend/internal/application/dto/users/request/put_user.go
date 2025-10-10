package request

// PutUserRequest PUT /users/{userId} のリクエストボディ
type PutUserRequest struct {
	Username  *string           `json:"username,omitempty" validate:"omitempty,min=1,max=100"`
	AvatarURL *string           `json:"avatar_url,omitempty" validate:"omitempty,url"`
	Metadata  *UserMetadataForm `json:"metadata,omitempty"`
}

// UserMetadataForm ユーザーメタデータフォームデータ
type UserMetadataForm struct {
	DisplayName *string   `json:"display_name,omitempty" validate:"omitempty,min=1,max=100"`
	Tagline     *string   `json:"tagline,omitempty" validate:"omitempty,max=150"`
	Bio         *string   `json:"bio,omitempty" validate:"omitempty,max=1000"`
	Skills      *[]string `json:"skills,omitempty" validate:"omitempty,dive,min=1,max=50"`
	Interests   *[]string `json:"interests,omitempty" validate:"omitempty,dive,min=1,max=50"`
}