package schemas

type Register struct {
	Username string `form:"username" json:"username" required:"true" binding:"required"`
	Password string `form:"password" json:"password" required:"true" binding:"required"`
	Email    string `form:"email" json:"email" required:"true" binding:"required"`
}
