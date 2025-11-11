package server

import "github.com/labstack/echo/v4"

func Run() (err error) {
	mux := echo.New()

	return mux.Start(":8000")
}
