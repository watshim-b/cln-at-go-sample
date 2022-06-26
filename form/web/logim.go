package form

type LoginForm struct {
	ID       string `form:"id"`
	Password string `form:"password"`
}
