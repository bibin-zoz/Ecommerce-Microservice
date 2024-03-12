package models

type UserLogin struct {
	Email    string
	Password string
}

type UserDetails struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
}

type UserSignUp struct {
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Password  string
}
type UserDetail struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Password  string
}
type TokenUser struct {
	User         UserDetails
	AccessToken  string
	RefreshToken string
}
type UserDetailsResponse struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
