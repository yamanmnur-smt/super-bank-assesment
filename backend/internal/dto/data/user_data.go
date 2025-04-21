package data

type UserData struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserProfileData struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
