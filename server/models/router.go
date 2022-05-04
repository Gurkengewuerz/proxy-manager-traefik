package models

import (
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Router struct {
	gorm.Model

	Name            string     `json:"name"`
	RedirectToHTTPs bool       `json:"redirectToHTTPs"`
	Locations       []Location `json:"locations" gorm:"foreignKey:RouterReference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (router *Router) GenerateConfig(tabStart uint8) (string, string) {
	routerConfig := ""
	serviceConfig := ""

	defaultLoc, foundDefault := lo.Find(router.Locations, func(loc Location) bool {
		return loc.IsDefault
	})

	if !foundDefault {
		return routerConfig, serviceConfig
	}

	for _, location := range router.Locations {
		rC, sC := location.GenerateConfig(tabStart, defaultLoc, router)

		routerConfig = routerConfig + rC + "\n"
		serviceConfig = serviceConfig + sC + "\n"
	}

	return routerConfig, serviceConfig
}

func (router *Router) YAMLName() string {
	return fmt.Sprintf("router-%d", router.ID)
}

func (router *Router) Validate() bool {
	if len(router.Name) == 0 {
		return false
	}

	if router.RedirectToHTTPs != true && router.RedirectToHTTPs != false {
		return false
	}

	if len(router.Locations) == 0 {
		return false
	}

	foundDefault := false
	for _, location := range router.Locations {
		if !location.Validate() {
			return false
		}
		if location.IsDefault && foundDefault {
			return false
		}

		if location.IsDefault && !foundDefault {
			foundDefault = true
		}
	}

	if !foundDefault {
		return false
	}

	return true
}
