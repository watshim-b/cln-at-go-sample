package response

import "github.com/watshim-b/cln-at-go-sample/ent"

type UserResponse struct {
	Users     []*ent.User
	ListCount int
}
