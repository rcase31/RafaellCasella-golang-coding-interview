package app

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"

	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/models"
)

var htmlTemplates map[string]*template.Template

type params map[string]interface{}

func home(c echo.Context) error {
	return c.HTML(http.StatusOK, execTemplateFromBase("Home", "home", params{}))
}

func getStates(c echo.Context) error {
	var states []models.State

	result := db.Find(&states)
	if result.RowsAffected == 0 {
		statesData, err := FetchStates()
		if err != nil {
			return err
		}
		states = statesData.States
	}

	statesHtml := execTemplateFromBase("States", "states", params{
		"states":       states,
		"row_template": htmlTemplates["states-row"],
	})
	return c.HTML(http.StatusOK, statesHtml)
}

func getReports(c echo.Context) error {
	reportData, err := FetchReports()
	if err != nil {
		return err
	}

	if err := c.Bind(reportData); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, reportData.Reports)
}

func getStatesJson(c echo.Context) error {
	statesData, err := FetchStates()
	if err != nil {
		return err
	}

	if err := c.Bind(statesData.States); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, statesData.States)
}

func getStatesByIDJson(c echo.Context) error {
	id := c.Param("id") // Fetch the id from the URL

	var state models.State
	result := db.First(&state, "id = ?", id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "State not found")
	}

	return c.JSON(http.StatusOK, state)
}

func fetchStates(c echo.Context) error {
	statesData, err := FetchStates()
	if err != nil {
		return err
	}

	// Populate the database with the retrieved states
	for _, state := range statesData.States {
		// Use Create with OnConflict to avoid duplicates if needed
		db.Where(models.State{ID: state.ID}).FirstOrCreate(&state)
	}

	return c.JSON(http.StatusOK, statesData.States)
}

func getCategoriesJson(c echo.Context) error {
	var categories []models.Category

	if err := c.Bind(categories); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, categories)
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
