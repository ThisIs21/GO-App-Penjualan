package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
type UpdateUserRequest struct {
    Name     string `json:"name,omitempty"`
    Email    string `json:"email,omitempty" validate:"email"`
    Password string `json:"password,omitempty" validate:"min=6"`
    Role     string `json:"role,omitempty"`
	 Status   string `json:"status,omitempty"`
}