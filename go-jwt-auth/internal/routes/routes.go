package routes

import (
	"net/http"

	"github.com/danielnorato/go-jwt-auth/internal/controllers"
	"github.com/danielnorato/go-jwt-auth/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", controllers.Register(db))
	r.Post("/login", controllers.Login(db))

	r.Group(func(protected chi.Router) {
		protected.Use(middlewares.JWTMiddleware)
		protected.Get("/protected", controllers.ProtectedEndpoint())
	})

	return r
}
