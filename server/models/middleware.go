package models

import (
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"strings"
)

type MiddlewareType int

const (
	ErrorProvider MiddlewareType = iota
	AuthProvider
	RedirectProvider
	RedirectSchemeProvider
)

type Middleware struct {
	gorm.Model

	Type     MiddlewareType      `json:"type"`
	Settings []MiddlewareSetting `json:"settings" gorm:"foreignKey:MiddlewareReference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type MiddlewareSettingKey string

const (
	ErrorProviderStatus  MiddlewareSettingKey = "ErrorProviderStatus"
	ErrorProviderService MiddlewareSettingKey = "ErrorProviderService"
	ErrorProviderQuery   MiddlewareSettingKey = "ErrorProviderQuery"

	AuthProviderUsers MiddlewareSettingKey = "AuthProviderUsers"

	RedirectProviderRegex       MiddlewareSettingKey = "RedirectProviderRegex"
	RedirectProviderReplacement MiddlewareSettingKey = "RedirectProviderReplacement"

	RedirectSchemeProviderScheme    MiddlewareSettingKey = "RedirectSchemeProviderScheme"
	RedirectSchemeProviderPermanent MiddlewareSettingKey = "RedirectSchemeProviderPermanent"
)

type MiddlewareSetting struct {
	gorm.Model

	Key                 MiddlewareSettingKey `json:"key"`
	Value               string               `json:"value"`
	MiddlewareReference uint
}

func (middleware *Middleware) GenerateConfig(tabStart uint8) (string, string) {
	initTabs := strings.Repeat("\t", int(tabStart))

	middlewareConfig := fmt.Sprintf("%s%s:\n", initTabs, middleware.YAMLName())
	serviceConfig := ""
	switch middleware.Type {
	case ErrorProvider:
		break

	case AuthProvider:
		middlewareConfig = middlewareConfig + fmt.Sprintf("%s\tbasicAuth:\n%s\t\tusers:\n", initTabs, initTabs)
		middlewareConfig = middlewareConfig + strings.Join(lo.Map(lo.Filter(middleware.Settings, func(setting MiddlewareSetting, _ int) bool {
			return setting.Key == AuthProviderUsers
		}), func(setting MiddlewareSetting, _ int) string {
			return fmt.Sprintf("%s\t\t\t- \"%s\"", initTabs, setting.Value)
		}), "\n")
		break

	case RedirectProvider:
		regex, regexOk := lo.Find(middleware.Settings, func(setting MiddlewareSetting) bool {
			return setting.Key == RedirectProviderRegex
		})

		replacement, replacementOk := lo.Find(middleware.Settings, func(setting MiddlewareSetting) bool {
			return setting.Key == RedirectProviderReplacement
		})

		if regexOk && replacementOk {
			middlewareConfig = middlewareConfig + fmt.Sprintf("%s\tredirectRegex:\n%s\t\tregex: %s\n%s\t\treplacement: %s",
				initTabs, initTabs, regex.Value, initTabs, replacement.Value)
		}
		break

	case RedirectSchemeProvider:
		scheme, schemeOk := lo.Find(middleware.Settings, func(setting MiddlewareSetting) bool {
			return setting.Key == RedirectSchemeProviderScheme
		})

		perm, permOk := lo.Find(middleware.Settings, func(setting MiddlewareSetting) bool {
			return setting.Key == RedirectSchemeProviderPermanent
		})

		if schemeOk && permOk {
			middlewareConfig = middlewareConfig + fmt.Sprintf("%s\tredirectScheme:\n%s\t\tscheme: %s\n%s\t\tpermanent: %s",
				initTabs, initTabs, scheme.Value, initTabs, perm.Value)
		}
		break

	default:
		break
	}
	return middlewareConfig, serviceConfig
}

func (middleware *Middleware) YAMLName() string {
	return fmt.Sprintf("middleware-%d", middleware.ID)
}

func (middlewareSetting *MiddlewareSetting) Validate() bool {
	if middlewareSetting.Key == "" {
		return false
	}

	if middlewareSetting.Value == "" {
		return false
	}
	return true
}

func (middleware *Middleware) Validate() bool {
	if middleware.Type != ErrorProvider && middleware.Type != AuthProvider && middleware.Type != RedirectProvider && middleware.Type != RedirectSchemeProvider {
		return false
	}

	for _, setting := range middleware.Settings {
		if !setting.Validate() {
			return false
		}
	}

	return true
}
