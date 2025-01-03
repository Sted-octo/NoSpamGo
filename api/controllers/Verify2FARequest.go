package controllers

type Verify2FARequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
