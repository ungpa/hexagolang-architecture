package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ungpa/goon-ah-lang/internal/core/ports"
)

type Handler struct {
	urlService     ports.UrlService
	scraperService ports.ScraperService
}

func NewHandler(urlService ports.UrlService, scraperService ports.ScraperService) *Handler {
	return &Handler{
		urlService:     urlService,
		scraperService: scraperService,
	}
}

func (h *Handler) AnalyzeURLs(c echo.Context) error {
	var req AnalyzeURLRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request format",
			Success: false,
		})
	}

	if err := h.urlService.Analyze(req.URLs); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, AnalyzeURLResponse{
		Message: "URLs queued for analysis",
		Success: true,
	})
}

func (h *Handler) GetMetaTags(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "URL parameter is required",
			Success: false,
		})
	}

	metaTags, err := h.scraperService.GetMetaTags(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, MetaTagsResponse{
		URL:      url,
		MetaTags: metaTags,
		Success:  true,
	})
}
