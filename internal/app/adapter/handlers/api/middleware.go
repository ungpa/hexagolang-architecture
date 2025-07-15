package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(e *echo.Echo) {
	// https://echo.labstack.com/docs/middleware/trailing-slash
	e.Pre(middleware.RemoveTrailingSlash()) // make the both api/v1/about and api/v1/about/ work cause we dont need to create double instance

	// if u use the 301 redirect status code
	// e.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
	// 	RedirectCode: http.StatusMovedPermanently,
	// }))
	// then you must set the reouter like this
	// api := e.Group("/api/v1")
	// api.GET("/", echo.HandlerFunc(func(c echo.Context) error {
	// 	return c.JSON(200, map[string]string{"message": "Hello, World!"})
	// }))
	// api.GET("/about/", echo.HandlerFunc(func(c echo.Context) error {
	// 	return c.JSON(200, map[string]string{"message": "/about"})
	// }))

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
}
