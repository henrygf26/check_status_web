package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

// StatusRequest is the structure of the request
type StatusRequest struct {
	URL string `json:"url"`
}

// StatusResponse is the structure of the response
type StatusResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Datetime string `json:"datetime"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/status", status)
	e.Start(":8085")
}

func status(c echo.Context) error {
	r := new(StatusRequest)
	c.Bind(r)
	return c.JSON(http.StatusOK, checkWeb(r.URL))
}

func checkWeb(url string) StatusResponse {
	var message = "OK"
	var code int

	resp, err := http.Get(url)

	if err != nil {
		message = "Error"
	}

	if resp != nil {
		resp.Body.Close()
		code = resp.StatusCode
	}

	return StatusResponse{Code: code, Message: message, Datetime: time.Now().Format("2006-01-02 15:04:05.000000")}
}
