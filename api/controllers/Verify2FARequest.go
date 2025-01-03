package controllers

type Verify2FARequest struct {
	Mail  string `json:"mail"`
	Token string `json:"token"`
}
