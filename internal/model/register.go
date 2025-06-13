package model

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

func ValidateRegisterRequest(userInput RegisterRequest) error {
	return validate.Struct(userInput)
}
