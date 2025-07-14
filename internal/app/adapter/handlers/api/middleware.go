package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// is like the clean-architecture because the Middleware as dependency
// type Middleware struct {
// 	// define middleware fields here
// }

// // i think this is a constructor (maybeeehh)
// func NewMiddleware() *Middleware {
// 	return &Middleware{}
// }

// // just dumbast think ever
// func (m *Middleware) SetupMiddleware(e *echo.Echo) {
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())
// 	e.Use(middleware.CORS())
// }

func SetupMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
}
