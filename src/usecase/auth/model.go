package auth

type LoginRequest struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type LoginResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token Token  `json:"token"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
