package services

import (
	"log"

	"github.com/ungpa/goon-ah-lang/internal/core/models"
	"github.com/ungpa/goon-ah-lang/internal/core/ports"
)

type UrlService struct {
	Scraper    ports.ScraperService
	Repository ports.UrlRepository
}

func NewUrlService(scraper ports.ScraperService, repository ports.UrlRepository) *UrlService {
	return &UrlService{
		Scraper:    scraper,
		Repository: repository,
	}
}

func (dcs *UrlService) Analyze(urls []string) error {
	for _, link := range urls {
		metaTags, err := dcs.Scraper.GetMetaTags(link)
		if err != nil {
			log.Printf("Error scraping %s: %v", link, err)
			continue
		}

		analysisResult := models.UrlAnalysisResult{
			Link:     link,
			MetaTags: metaTags,
		}

		if err = dcs.Repository.Create(analysisResult); err != nil {
			log.Printf("Error saving result for %s: %v", link, err)
			continue
		}
	}

	return nil
}
