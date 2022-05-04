package models

import (
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"strings"
)

type LocationSchema string

const (
	SchemaHTTP  LocationSchema = "http"
	SchemaHTTPS LocationSchema = "https"
	SchemaTCP   LocationSchema = "tcp"
	SchemaUDP   LocationSchema = "udp"
)

type Location struct {
	gorm.Model

	Host            string         `json:"host"`
	PathPrefix      string         `json:"pathPrefix"`
	Service         string         `json:"service"`
	Port            int            `json:"port"`
	Schema          LocationSchema `json:"schema"`
	IsDefault       bool           `json:"isDefault"`
	Middlewares     []Middleware   `json:"middlewares" gorm:"many2many:location_middlewares;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RouterReference uint
}

func (loc *Location) Rule(defaultLocation Location) string {
	hosts := strings.TrimSpace(loc.Host)
	if len(hosts) == 0 {
		hosts = defaultLocation.Host
	} else {
		if !loc.IsDefault {
			hosts = hosts + "," + defaultLocation.Host
		}
	}

	allHosts := strings.Split(hosts, ",")
	hostsRules := strings.Join(lo.Map(allHosts, func(host string, _ int) string {
		return fmt.Sprintf("Host(`%s`)", host)
	}), " || ")

	if len(allHosts) > 1 {
		hostsRules = "(" + hostsRules + ")"
	}

	pathPrefix := strings.TrimSpace(loc.PathPrefix)
	if len(pathPrefix) == 0 {
		return hostsRules
	}

	return hostsRules + " && " + fmt.Sprintf("PathPrefix(`%s`)", pathPrefix)
}

func (loc *Location) URL() string {
	return fmt.Sprintf("%s://%s:%d", loc.Schema, loc.Service, loc.Port)
}

func (loc *Location) YAMLRouterName(router *Router) string {
	return fmt.Sprintf("%s-%d", router.YAMLName(), loc.ID)
}

func (loc *Location) YAMLServiceName() string {
	return fmt.Sprintf("service-%d", loc.ID)
}

func (loc *Location) GenerateConfig(tabStart uint8, defaultLocation Location, parentRouter *Router) (string, string) {
	initTabs := strings.Repeat("\t", int(tabStart))

	serviceConfig := fmt.Sprintf("%s%s:\n%s\tloadBalancer:\n%s\t\tservers:\n%s\t\t\t- url: \"%s\"",
		initTabs, loc.YAMLServiceName(), initTabs, initTabs, initTabs, loc.URL())

	middlewareYAMLNames := lo.Map(loc.Middlewares, func(middleware Middleware, _ int) string {
		return middleware.YAMLName()
	})
	if parentRouter.RedirectToHTTPs {
		middlewareYAMLNames = append(middlewareYAMLNames, "manager-https-redir")
	}

	middlewareYAMLList := strings.Join(lo.Map(middlewareYAMLNames, func(middlewareName string, _ int) string {
		return fmt.Sprintf("\n%s\t\t- \"%s\"", initTabs, middlewareName)
	}), "")
	routerConfig := fmt.Sprintf("%s%s:\n%s\tservice: %s\n%s\trule: \"%s\"\n%s\tmiddlewares: %s",
		initTabs, loc.YAMLRouterName(parentRouter), initTabs, loc.YAMLServiceName(), initTabs, loc.Rule(defaultLocation), initTabs, middlewareYAMLList)

	return routerConfig, serviceConfig
}

func (loc *Location) Validate() bool {
	if loc.IsDefault && len(loc.Host) == 0 {
		return false
	}

	if !loc.IsDefault && len(loc.PathPrefix) == 0 {
		return false
	}

	if len(loc.Service) == 0 {
		return false
	}

	if loc.Port <= 0 || loc.Port > 65535 {
		return false
	}

	if loc.Schema != SchemaHTTP && loc.Schema != SchemaHTTPS && loc.Schema != SchemaTCP && loc.Schema != SchemaUDP {
		return false
	}

	return true
}
