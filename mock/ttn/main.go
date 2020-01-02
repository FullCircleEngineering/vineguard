package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"net/http"
)

const (
	deviceId  = "mock-dev-id"
	localPort = ":9060"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// note: we ignore all auth considerations in header
	e.GET("/devices", getDevices)
	e.GET(fmt.Sprintf("/query/%s", deviceId), getData)

	// Start server
	e.Logger.Fatal(e.Start(localPort))
}

// Handlers
func getDevices(c echo.Context) error {
	jsonOut, err := json.Marshal([]string{deviceId})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(jsonOut))
}

func getData(c echo.Context) error {
	if c.QueryParam("last") != "7d" {
		return c.String(http.StatusBadRequest, "query param `last` must be set to `7d`")
	}

	jsonOut, err := ioutil.ReadFile("data.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(jsonOut))
}
