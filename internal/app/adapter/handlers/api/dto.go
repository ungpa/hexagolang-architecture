package api

type AnalyzeURLRequest struct {
	URLs []string `json:"urls" validate:"required,dive,url"`
}

type AnalyzeURLResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MetaTagsResponse struct {
	URL      string   `json:"url"`
	MetaTags []string `json:"meta_tags"`
	Success  bool     `json:"success"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}
