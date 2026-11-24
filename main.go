package main

import (
	"authorization/model"
	"authorization/storage"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var registered []model.Registr

func main() {
	registered, err := storage.LoadUsers()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()

	tmplReg := template.Must(template.ParseFiles("template/registr2.html"))

	tmplRegMain := template.Must(template.ParseFiles("template/registr.html"))

	tmplSignInMain := template.Must(template.ParseFiles("template/vxod.html"))

	e.Static("/", "static")

	e.GET("/", func(c echo.Context) error {
		return tmplRegMain.Execute(c.Response(), model.RegLogError{})
	})

	e.GET("/registr2", func(c echo.Context) error {
		return tmplReg.Execute(c.Response(), model.RegLogError{})
	})

	e.POST("registr2", func(c echo.Context) error {

		Login := c.FormValue("login")
		pass := c.FormValue("password")
		confPass := c.FormValue("password2")

		for _, v := range registered {
			if Login == v.Login {
				return tmplRegMain.Execute(c.Response(), model.RegLogError{
					Login: Login,
					Error: "Аккаунт с таким логином уже существует!",
				})
			}
		}

		if pass != confPass {
			return tmplRegMain.Execute(c.Response(), model.RegLogError{
				Login:     Login,
				ErrorPass: "Пароли не совпадают!",
			})
		}

		registered = append(registered, model.Registr{
			ID:              len(registered) + 1,
			Login:           Login,
			Password:        pass,
			ConfirmPassword: confPass,
		})

		fmt.Println(registered)
		storage.SavedUsers(registered)

		return c.Redirect(http.StatusSeeOther, "/registr2")
	})

	e.GET("/vxod", func(c echo.Context) error {
		return tmplSignInMain.Execute(c.Response(), model.SignInError{})
	})

	e.GET("/vxod2", func(c echo.Context) error {
		return c.File("template/vxod2.html")
	})

	e.POST("vxod2", func(c echo.Context) error {
		Login := c.FormValue("login")
		Pass := c.FormValue("password")

		for _, v := range registered {
			if Login != v.Login {
				return tmplSignInMain.Execute(c.Response(), model.SignInError{
					Login:    Login,
					ErrorLog: "Пользователь с таким логином не найден!",
				})
			}

		}

		for _, v := range registered {
			if Pass != v.Password {
				return tmplSignInMain.Execute(c.Response(), model.SignInError{
					Login:     Login,
					ErrorPass: "Пароль не подходит",
				})
			}
		}

		fmt.Println(true)

		return c.Redirect(http.StatusSeeOther, "/vxod2")
	})

	e.GET("/exit", func(c echo.Context) error {
		go e.Close()
		return c.String(200, "сервер закрыт")
	})

	err = e.Start(":8080")

	if err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}
