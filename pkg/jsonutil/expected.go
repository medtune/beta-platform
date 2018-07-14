package jsonutil

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupData struct {
	Secret          string `json:"secret"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordc"`
}
