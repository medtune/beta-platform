package jsonutil

type LoginData struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type SignupData struct {
	Secret          string `json:"secret" valid:"required"`
	Email           string `json:"email" valid:"email"`
	Username        string `json:"username" valid:"required"`
	Password        string `json:"password" valid:"required"`
	PasswordConfirm string `json:"passwordc" valid:"required"`
}

type RunImageInference struct {
	Image    string `json:"image"`
	File     string `json:"file"`
	NumPreds string `json:"numpreds"`
}
