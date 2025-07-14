package api

import (
	"github.com/labstack/echo/v4"
)

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{handler: handler}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
	api.GET("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello, World!"})
	}))
	api.POST("/analyze", r.handler.AnalyzeURLs)
	api.GET("/meta-tags", r.handler.GetMetaTags)

}

func RegisterRoutes(e *echo.Echo, router *Router) {
	router.RegisterRoutes(e)
}
