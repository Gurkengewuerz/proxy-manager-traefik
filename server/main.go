package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"log"
	"strings"
	"traefikmanager/server/claims"
	"traefikmanager/server/config"
	"traefikmanager/server/database"
	"traefikmanager/server/routes"
	"traefikmanager/server/utils"
)

var oauth2Config oauth2.Config
var verifier *oidc.IDTokenVerifier
var cfg config.Config
var sessStore = session.New()

func CheckOidc(c *fiber.Ctx) error {
	rawAccessToken, authHeaderFound := c.GetReqHeaders()["Authorization"]
	if !authHeaderFound || rawAccessToken == "" {
		return c.SendStatus(401)
	}
	parts := strings.Split(rawAccessToken, " ")
	if len(parts) != 2 {
		return c.SendStatus(400)
	}

	ctx := context.Background()
	idToken, err := verifier.Verify(ctx, parts[1])
	if err != nil {
		return c.SendStatus(401)
	}

	resp := struct {
		IDTokenClaims *claims.IDTokenClaims
	}{}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		return c.Status(500).SendString("No claims found " + err.Error())
	}
	c.Locals("claims", resp.IDTokenClaims)

	return c.Next()
}

func RedirectLogin(c *fiber.Ctx) error {
	sess, err := sessStore.Get(c)
	if err != nil {
		return c.SendStatus(500)
	}
	state := utils.RandSeq(12)
	sess.Set("state", state)
	_ = sess.Save()
	return c.Redirect(oauth2Config.AuthCodeURL(state))
}

func LoginCallback(c *fiber.Ctx) error {
	sess, err := sessStore.Get(c)
	if err != nil {
		return c.SendStatus(500)
	}
	state := sess.Get("state")
	if c.Query("state", "") != state {
		return c.Status(500).SendString("Invalid state")
	}

	ctx := context.Background()
	oauth2Token, err := oauth2Config.Exchange(ctx, c.Query("code", ""))
	if err != nil {
		return c.Status(500).SendString("Failed to exchange token: " + err.Error())
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return c.Status(500).SendString("No id_token field in oauth2 token")
	}

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return c.Status(500).SendString("Failed to verify ID Token: " + err.Error())
	}

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *claims.IDTokenClaims
	}{oauth2Token, new(claims.IDTokenClaims)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		return c.Status(500).SendString("No claims found " + err.Error())
	}

	return c.Redirect(fmt.Sprintf("%s/login/callback?token=%s", cfg.FrontendUrl, rawIDToken))
}

func setUpRoutes(app *fiber.App) {
	app.Get("/config", CheckOidc, routes.GenerateConfig)
	app.Get("/commit", CheckOidc, routes.Commit)
	app.Get("/stats", CheckOidc, routes.Stats)

	app.Get("/router", CheckOidc, routes.GetRouter)
	app.Post("/router", CheckOidc, routes.PostRouter)
	app.Put("/router", CheckOidc, routes.PutRouter)
	app.Delete("/router", CheckOidc, routes.DeleteRouter)

	app.Get("/middleware", CheckOidc, routes.GetMiddleware)
	app.Post("/middleware", CheckOidc, routes.PostMiddleware)
	app.Put("/middleware", CheckOidc, routes.PutMiddleware)
	app.Delete("/middleware", CheckOidc, routes.DeleteMiddleware)

	app.Get("/audit", CheckOidc, routes.GetAudit)

	app.Get("/auth/info", CheckOidc, routes.AuthInfoMiddleware)

	app.Get("/login", RedirectLogin)
	app.Get("/oidc/callback", LoginCallback)

	app.Static("/", "./web/build")
	app.Static("*", "./web/build/index.html")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg = config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, cfg.OidcConfigURI)
	if err != nil {
		panic(err)
	}

	oauth2Config = oauth2.Config{
		ClientID:     cfg.OidcClientID,
		ClientSecret: cfg.OidcClientSecret,
		RedirectURL:  fmt.Sprintf("%s/oidc/callback", cfg.BackendURL),
		Endpoint:     provider.Endpoint(),
		Scopes:       strings.Split(cfg.OidcScopes, " "),
	}
	oidcConfig := &oidc.Config{
		ClientID: cfg.OidcClientID,
	}
	verifier = provider.Verifier(oidcConfig)

	database.ConnectDb(&cfg)
	app := fiber.New()

	app.Use(cors.New())

	setUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 Not Found")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Port)))

}
