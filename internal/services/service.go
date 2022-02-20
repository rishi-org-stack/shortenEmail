package services

type (
	GetTokenRequest struct {
		Code         string `json:"code"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RedirectUri  string `json:"redirect_uri"`
		GrantType    string `json:"grant_type"`
	}
	GetTokenResponse struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
	}
	RefreshTokenRequest struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RefreshToken string `json:"refresh_token"`
		GrantType    string `json:"grant_type"`
	}

	RefreshTokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
	}
)
