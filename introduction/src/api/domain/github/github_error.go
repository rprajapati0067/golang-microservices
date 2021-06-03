package github

type GithubErrorResponse struct {
	StatusCode       int           `json:"status_code"`
	Message          string        `json:"message"`
	DocumentationUrl string        `json:"documentation_url"`
	Errors           []GithubError `json:"errors"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	field    string `json:"field"`
	message  string `json:"message"`
}
