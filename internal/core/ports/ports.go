package ports

import "github.com/ungpa/goon-ah-lang/internal/core/models"

type ScraperService interface {
	GetMetaTags(url string) ([]string, error)
}

type UrlService interface {
	Analyze(urls []string) error
}

type UrlRepository interface {
	Create(data models.UrlAnalysisResult) error
}
