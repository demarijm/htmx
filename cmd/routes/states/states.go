package states

import (
	"encoding/json"
	"htmx-go/cmd/model"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

func fetchStates() ([]model.State, error) {
	resp, err := http.Get("https://iwaste.epa.gov/api/us-states")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var states []model.State
	err = json.Unmarshal(body, &states)

	if err != nil {
		return nil, err
	}
	return states, nil
}

func StateRoute(e *echo.Echo) {
	e.GET("/states", func(c echo.Context) error {
		states, err := fetchStates()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		data := model.Data{
			Contacts: data.Contacts,
			States:   states,
		}
		return c.Render(http.StatusOK, "states", data)
	})
}
