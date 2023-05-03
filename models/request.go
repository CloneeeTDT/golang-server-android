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

type UpdateUserRequest struct {
	Name string `json:"name"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
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

type Speech2TextRequest struct {
	Audio string `json:"audio"`
	From  string `json:"from"`
	To    string `json:"to"`
}

type Speech2TextResponse struct {
	Content   string `json:"content"`
	Translate string `json:"translate"`
	From      string `json:"from"`
	To        string `json:"to"`
	Audio     string `json:"audio"`
}
