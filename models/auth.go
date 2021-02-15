package models

// Roll type rolls available
type Roll uint8

// Roll defines the structure for available rolls
const (
	AdminRoll Roll = iota
	StudentRoll
	CompanyRoll
)

// EmailAndPassword defines the structure for login body
type EmailAndPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ErrorResponse defines the structure for error response
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Key     string `json:"key"`
}

// LoginResponse defines the structure for login response
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// UnAuthorizeError defines the structure for unauthorize error response
type UnAuthorizeError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
