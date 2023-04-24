package request

type UserCreateRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	UserId      string `json:"userId"`
}
