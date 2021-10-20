package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/models"

	"github.com/labstack/echo/v4"
)

var htmlTemplates map[string]*template.Template

type params map[string]interface{}

func home(c echo.Context) error {
	return c.HTML(http.StatusOK, execTemplate("home", params{"content": ""}))
}

func displayStates(c echo.Context) error {
	var states []models.State

	jsonData, err := FetchStates()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(jsonData), &states)
	if err != nil {
		return err
	}

	statesHtml := execTemplate("statesSection", params{})
	return c.HTML(http.StatusOK, statesHtml)
}

func execTemplate(templateName string, p params) string {
	t := htmlTemplates[templateName]
	if t == nil {
		return ""
	}
	var res bytes.Buffer
	err := t.Execute(&res, p)
	if err != nil {
		fmt.Println("ERROR in execTemplate:", err.Error())
		return ""
	}
	return res.String()
}
