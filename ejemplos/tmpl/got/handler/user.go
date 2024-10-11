package handler

import (
	"tmplexample/model"
	"tmplexample/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	use := model.User{
		Email: "andres.kenji@outlook.com",
	}

	return render(c, user.Show(use))
}
