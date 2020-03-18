package model

type Token struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIN   int    `json:"expiresIn"`
}

type Policy struct {
	Role   string `json:"role"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
