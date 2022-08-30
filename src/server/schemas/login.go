package schemas

type Login struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
}

type OAuthLogin struct {
	Username  string `form:"username" json:"username" required:"true"`
	Password  string `form:"password" json:"password" required:"true"`
	GrantType string `form:"grant_type" json:"grant_type"`
}

type LoginTest struct {
	Key string `json:"key" form:"key"`
}
