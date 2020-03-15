package model

type Token struct {
	AccessToken string
	TokenType   string
	ExpiresIN   int
}

type Policy struct {
	Role   string `json:"role"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
