package main

import (
	"fmt"
	"html/template"

	"io"
	"net/http"

	"htmx-go/cmd/model"
	"htmx-go/cmd/routes/states"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newContact(name string, email string) model.Contact {
	return model.Contact{
		Name:  name,
		Email: email,
	}
}

func newData() model.Data {
	return model.Data{
		Contacts: model.Contacts{
			newContact("John Doe", "johndoe@hogor.com"),
			newContact("Clara Dow", "jane@domw.dwoim"),
		},
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	data := newData()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", data)
	})
	states.StateRoute(e)

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		data.Contacts = append(data.Contacts, newContact(name, email))
		return c.Render(http.StatusOK, "display", data)
	})

	e.Logger.Fatal(e.Start(":42069"), fmt.Sprintf("Server started at port %d", 42069))
}
