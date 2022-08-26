package user

type formatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

func Formatter(user User, token string) formatter {
	jsonFormat := formatter{
		ID:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}

	return jsonFormat
}
