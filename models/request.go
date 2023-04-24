package models

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
}

type SaveWordRequest struct {
	WordID uint `json:"word_id" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
}

type SaveNoteRequest struct {
	WordID uint   `json:"word_id" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
	Note   string `json:"note" binding:"required"`
}

type OCRRequest struct {
	Image string `json:"image" binding:"required"`
}
