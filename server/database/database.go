package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"traefikmanager/server/config"
	"traefikmanager/server/database/gormigrate"
	"traefikmanager/server/models"
)

var (
	DBConn *gorm.DB
)

func ConnectDb(cfg *config.Config) {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		{
			ID: "202205041435",
			Migrate: func(tx *gorm.DB) error {
				type LogEntry struct {
					gorm.Model

					User     string `json:"user"`
					Action   string `json:"action"`
					Metadata string `json:"metadata"`
				}

				type MiddlewareSetting struct {
					gorm.Model

					Key                 models.MiddlewareSettingKey `json:"key"`
					Value               string                      `json:"value"`
					MiddlewareReference uint
				}

				type Middleware struct {
					gorm.Model

					Type     models.MiddlewareType `json:"type"`
					Settings []MiddlewareSetting   `json:"settings" gorm:"foreignKey:MiddlewareReference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				type Location struct {
					gorm.Model

					Host            string                `json:"Host"`
					PathPrefix      string                `json:"PathPrefix"`
					Service         string                `json:"service"`
					Port            int                   `json:"port"`
					Schema          models.LocationSchema `json:"schema"`
					IsDefault       bool                  `json:"isDefault"`
					Middlewares     []Middleware          `json:"middlewares" gorm:"many2many:location_middlewares;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
					RouterReference uint
				}

				type Router struct {
					gorm.Model

					Name            string     `json:"name"`
					RedirectToHTTPs bool       `json:"redirectToHTTPs"`
					Locations       []Location `json:"locations" gorm:"foreignKey:RouterReference;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
				}

				return tx.AutoMigrate(&LogEntry{}, &Middleware{}, &MiddlewareSetting{}, &Router{}, &Location{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("router", "middleware", "middleware_setting", "log_entry")
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	DBConn = db
}
