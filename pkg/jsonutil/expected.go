package jsonutil

type LoginData struct {
	Username string
	Password string
}

type SignupData struct {
	Secret          string
	Username        string
	Password        string
	PasswordConfirm string `json:"passwordc"`
}
