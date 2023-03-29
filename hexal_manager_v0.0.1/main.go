package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"hexal_manager_v0.0.1/app/config"
	"hexal_manager_v0.0.1/app/models"
	"hexal_manager_v0.0.1/app/routes"
	"net/http"
	"os"
	"crypto/tls"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	// Create new echo instance
	e := echo.New()
	// Configurar certificado y clave SSL
	configSSL := &tls.Config{}
	configSSL.Certificates = make([]tls.Certificate, 1)
	configSSL.Certificates[0], err = tls.LoadX509KeyPair("/etc/letsencrypt/live/www.spinworld.top/fullchain.pem", "/etc/letsencrypt/live/www.spinworld.top/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}

	// Agregar puerto HTTPS
	e.Server.Addr = ":8023"
	e.Server.TLSConfig = configSSL

	// Set log level
	e.Logger.SetLevel(log.DEBUG)
	// Log Echo output to standard out
	e.Logger.SetOutput(os.Stdout)

	// Database connection
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error to load configuration %v", err)
	}
	db, err := config.NewDatabase(conf)
	if err != nil {
		log.Fatalf("Error to connect database %v", err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error obteniendo la conexión de base de datos: %v", err)
		}

		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Error cerrando la conexión de base de datos: %v", err)
		}
	}()

	if err := config.Migrate(db, &models.Area{}); err != nil {
		log.Fatalf("Error realizando migraciones: %v", err)
	}
	// Resgiter routes
	routes.RegisterAreaRoutes(e, db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Please, enter the endpoint that you have")
	})
	// Iniciar el servidor HTTPS
	e.Logger.Fatal(e.StartTLS(":8023", "/etc/letsencrypt/live/www.spinworld.top/fullchain.pem", "/etc/letsencrypt/live/www.spinworld.top/privkey.pem"))
	//e.Logger.Fatal(e.Start(":8023"))
}
