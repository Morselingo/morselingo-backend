package model

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type LoginResponse struct {
	Token string
}

func ValidateLoginUserInput(userInput LoginRequest) error {
	return validate.Struct(userInput)
}
