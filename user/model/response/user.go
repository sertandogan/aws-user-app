package response

type UserResponse struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}
