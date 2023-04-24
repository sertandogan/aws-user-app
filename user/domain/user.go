package domain

type UserEntity struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}
