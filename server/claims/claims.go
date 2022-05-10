package claims

type IDTokenClaims struct {
	ID       string   `json:"sub"`
	Username string   `json:"preferred_username"`
	Roles    []string `json:"roles"`
}
