package controllers

type User struct {
	AppUserId  int64  `json:"app_user_id"`
	IsVerified bool   `json:"is_verified"`
	Username   string `json:"username"`
	FacebookId string `json:"facebook_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Link       string `json:"link"`
}
