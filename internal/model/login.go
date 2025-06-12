package model

type LoginRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type LoginResponse struct {
	Token string
}

func ValidateLoginUserInput(userInput LoginRequest) error {
	return validate.Struct(userInput)
}
