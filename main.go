package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

// StatusRequest is the structure of the request
type StatusRequest struct {
	URL     string `json:"url"`
	Timeout int    `json:"timeout"`
}

// StatusResponse is the structure of the response
type StatusResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Datetime     string `json:"datetime"`
	TimeoutParam int    `json:"timeout"`
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

	t := 60

	if r.Timeout > 0 {
		t = r.Timeout
	}

	return c.JSON(http.StatusOK, checkWeb(r.URL, t))
}

func checkWeb(url string, timeout int) StatusResponse {
	var message = "OK"
	var code int

	var netClient = http.Client{
		Timeout: (time.Second * time.Duration(timeout)),
	}

	resp, err := netClient.Get(url)

	if err != nil {
		message = "Error"
	}

	if resp != nil {
		resp.Body.Close()
		code = resp.StatusCode
	}

	return StatusResponse{
		Code:         code,
		Message:      message,
		Datetime:     time.Now().Format("2006-01-02 15:04:05.000000"),
		TimeoutParam: timeout}
}
