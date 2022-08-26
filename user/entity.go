package user

type User struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password_hash   string `json:"password_has"`
	Avatar_filename string `json:"avatar_filename"`
	Occupation      string `json:"occupation"`
	Role            string `json:"role"`
}
