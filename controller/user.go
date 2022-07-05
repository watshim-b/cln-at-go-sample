package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	form "github.com/watshim-b/cln-at-go-sample/form/web"
	uc "github.com/watshim-b/cln-at-go-sample/usecase/web/user"
)

// ユーザー情報を取得する
func GetUser(c *gin.Context) {

	// リクエストされた情報をform煮詰める
	uf := form.UserForm{}
	// err := c.BindJSON(&uf)
	// if err != nil {
	// 	print(err.Error())
	// 	c.JSON(http.StatusBadRequest, err)
	// 	return
	// }

	// usecaseの処理を投げる
	u := uc.NewUserUsecaseForLowlyEmploye()
	users, err := u.GetUser(uf)
	if err != nil {
		print("errorです")
		print(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// レスポンスを返却する
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
