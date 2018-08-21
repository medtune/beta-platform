package jsonutil

// LoginData .
type LoginData struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// SignupData .
type SignupData struct {
	Secret          string `json:"secret" valid:"required"`
	Email           string `json:"email" valid:"email"`
	Username        string `json:"username" valid:"required"`
	Password        string `json:"password" valid:"required"`
	PasswordConfirm string `json:"passwordc" valid:"required"`
}

// RunImageInference .
type RunImageInference struct {
	Id       int    `json:"id"`
	Image    string `json:"image"`
	File     string `json:"file"`
	NumPreds string `json:"numpreds"`
}

// RunImageCam .
type RunImageCam struct {
	Id            int    `json:"id"`
	Image         string `json:"image"`
	File          string `json:"file"`
	CamOutputFile string `json:"cam_output_dir"`
}

// TestRequest .
type TestRequest struct {
	Test  bool `json:"test"`
	Sleep int  `json:"sleep"`
}
