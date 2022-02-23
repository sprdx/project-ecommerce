package requests

type RegisterRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
