package title

import "errors"

// Title domain specific errors

// User title validation errors
var (
	ErrUserTitleUserIDRequired  = errors.New("ユーザーIDは必須です")
	ErrUserTitleTitleIDRequired = errors.New("称号IDは必須です")
)

// Title validation errors
var (
	ErrTitleNameRequired        = errors.New("称号名は必須です")
	ErrTitleDescriptionRequired = errors.New("称号説明は必須です")
)