package services

type (
	GetTokenRequest struct {
		Code          string `json:"code"`
		ClientID      string `json:"client_id"`
		CliientSecret string `json:"client_secret"`
		RedirectUri   string `json:"redirect_uri"`
		GrantType     string `json:"grant_type"`
	}

	GetTokenResponse struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
	}
)
