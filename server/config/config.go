package config

type Config struct {
	Port int `env:"PORT" envDefault:"4664"`

	DatabaseHost     string `env:"DB_HOST"`
	DatabasePort     int    `env:"DB_PORT" envDefault:"3306"`
	DatabaseName     string `env:"DB_NAME"`
	DatabaseUsername string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`

	FrontendUrl string `env:"FRONTEND_URL" envDefault:"http://localhost"`

	OidcClientID      string `env:"OIDC_CLIENT_ID"`
	OidcClientSecret  string `env:"OIDC_CLIENT_SECRET"`
	OidcConfigURI     string `env:"OIDC_CONFIG_URI"`
	OidcUsernameClaim string `env:"OIDC_USERNAME_CLAIM" envDefault:"preferred_username"`
	OidcScopes        string `env:"OIDC_SCOPES" envDefault:"openid profile email"`
}
