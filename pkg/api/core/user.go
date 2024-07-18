package core

type PasswordConfirmRequest struct {
	Password string `url:"password,quoted"`
}

type PasswordConfirmResponse struct {
	SynoConfirmPWToken string `json:"SynoConfirmPWToken,omitempty"`
}
