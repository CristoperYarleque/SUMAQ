package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cyarleque/sumaq/config"
	"github.com/cyarleque/sumaq/internal/database"
	md "github.com/cyarleque/sumaq/internal/middlewares"
	. "github.com/cyarleque/sumaq/internal/routes/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
)

var env string

func main() {

	err := os.Setenv("env", env)
	if err != nil {
		log.Fatalf("Error setting env: %v", err.Error())
	}

	// Cargar la configuración
	configPath := config.GetConfigPath(env)
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err.Error())
	}

	factory := &database.ConnectionFactory{}
	metaDatabaseDB, err := factory.CreateConnection(cfg.MetaDatabaseDB.Adapter, &cfg.MetaDatabaseDB, nil)
	if err != nil {
		log.Fatalf("Error connecting metaDatabaseDB: %v", err.Error())
	}

	customMiddleware := md.NewMiddleware(*metaDatabaseDB)
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)
	r.Use(customMiddleware.MyMiddleware)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	config := ConnectionConfig{
		SumaqDB: *metaDatabaseDB,
	}

	v1Router := chi.NewRouter()
	v1Router.Mount("/login", config.LoginRouter())
	v1Router.Mount("/users", config.UsersRouter())
	v1Router.Mount("/entrepreneurs", config.EntrepreneursRouter())
	v1Router.Mount("/info-entrepreneurs", config.InfoEntrepreneursRouter())
	v1Router.Mount("/categories", config.CategoriesRouter())
	v1Router.Mount("/products", config.ProductsRouter())
	v1Router.Mount("/news", config.NewsRouter())
	v1Router.Mount("/promotions", config.PromotionsRouter())
	v1Router.Mount("/chatbot", config.ChatbotRouter())

	r.Mount("/v1", v1Router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}

	fmt.Println("Listening on port", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
