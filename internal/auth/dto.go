package auth

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,min=1,max=255"`
	Name     string `json:"name" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,min=1,max=255"`
	Password string `json:"password" binding:"required,min=1,max=255"`
}

// AuthResponse represents the response after successful login/register
type AuthResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
