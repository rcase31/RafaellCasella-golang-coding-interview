package app

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

var htmlTemplates map[string]*template.Template

type params map[string]interface{}

func home(c echo.Context) error {
	return c.HTML(http.StatusOK, execTemplateFromBase("Home", "home", params{}))
}

func getStates(c echo.Context) error {
	statesData, err := FetchStates()
	if err != nil {
		return err
	}

	statesHtml := execTemplateFromBase("States", "states", params{
		"states":       statesData.States,
		"row_template": htmlTemplates["states-row"],
	})
	return c.HTML(http.StatusOK, statesHtml)
}

// Helpers

func execTemplateFromBase(title, templateName string, p params) string {
	return execTemplate(
		"base-template",
		params{
			"title": title,
			"body":  execTemplate(templateName, p),
		},
	)
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
