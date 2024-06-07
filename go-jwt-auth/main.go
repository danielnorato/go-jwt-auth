package main

import (
	"log"

	"net/http"

	"github.com/danielnorato/go-jwt-auth/internal/models"
	"github.com/danielnorato/go-jwt-auth/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection
	log.Println("Iniciando api")
	dsn := "host=172.17.0.5 user=daniel password=daniel123 dbname=Test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	log.Println("Iniciando migrate")

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	log.Println("finalizando migrate")

	// Initialize routes
	r := routes.SetupRouter(db)
	log.Fatal(http.ListenAndServe(":3000", r))
}
