package form

type UserForm struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active int    `json:"active"`
}
