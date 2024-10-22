package responses

type UserTokenResponse struct {
	Id           string `json:"id" gorm:"column:id"`
	UserId       string `json:"user_id" gorm:"column:user_id"`
	AccessToken  string `json:"access_token" gorm:"column:access_token"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
}
