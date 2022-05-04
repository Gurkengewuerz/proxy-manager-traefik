package traefik

import (
	"fmt"
	"strings"
	"traefikmanager/server/database"
	"traefikmanager/server/models"
)

func GenerateConfig() string {
	var routers []models.Router
	database.DBConn.Preload("Locations").Preload("Locations.Middlewares").Preload("Locations.Middlewares.Settings").Find(&routers)

	middlewaresUnique := make(map[uint]*models.Middleware)

	routerConfig := ""
	serviceConfig := ""
	for _, router := range routers {
		rC, sC := router.GenerateConfig(2)
		routerConfig = routerConfig + "\n" + rC
		serviceConfig = serviceConfig + "\n" + sC

		for _, loc := range router.Locations {
			for _, middleware := range loc.Middlewares {
				middlewaresUnique[middleware.ID] = &middleware
			}
		}
	}
	if len(routers) == 0 {
		routerConfig = " {}"
		serviceConfig = " {}"
	}

	middlewareConfig := ""
	for _, middleware := range middlewaresUnique {
		mC, sC := middleware.GenerateConfig(2)
		middlewareConfig = middlewareConfig + "\n" + mC

		if sC == "" {
			serviceConfig = serviceConfig + "\n" + sC
		}
	}
	middlewareConfig = middlewareConfig + "\n" + GetStaticProviders(2)

	if len(middlewaresUnique) == 0 {
		middlewareConfig = " {}"
	}

	tabbedConfig := fmt.Sprintf("http:\n\trouters:%s\n\tmiddlewares:%s\n\tservices:%s\n\n", routerConfig, serviceConfig, middlewareConfig)
	return strings.ReplaceAll(tabbedConfig, "\t", "  ")
}
