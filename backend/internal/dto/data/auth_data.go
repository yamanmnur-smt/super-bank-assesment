package data

type JwtToken struct {
	User  UserProfileData `json:"user"`
	Token string          `json:"token"`
}
