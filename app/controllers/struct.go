package controllers

type User struct {
	IsVerified bool   `json:"is_verified,omitempty"`
	Username   string `json:"username"`
	FacebookId string `json:"facebook_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}
