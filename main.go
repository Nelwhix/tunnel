package main

import (
	"fmt"
	"github.com/Nelwhix/tunnel/handlers"
	"github.com/Nelwhix/tunnel/pkg"
	"github.com/Nelwhix/tunnel/pkg/models"
	"github.com/go-playground/validator/v10"
	gHandlers "github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var validate *validator.Validate

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fileName := filepath.Join("logs", "tunnel.log")
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := pkg.CreateNewLogger(f)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	validate = validator.New(validator.WithRequiredStructEnabled())

	conn, err := pkg.CreateDbConn()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer conn.Close()

	model := models.Model{
		Conn: conn,
	}

	handler := handlers.Handler{
		Model:     model,
		Logger:    logger,
		Validator: validate,
	}

	m := pkg.AuthMiddleware{
		Model: model,
	}

	r := http.NewServeMux()

	// Auth routes
	r.Handle("POST /api/tunnels", m.Register(handler.CreateNewTunnel))

	fmt.Printf("iCallOn started at http://localhost:%s\n", os.Getenv("SERVER_PORT"))

	err = http.ListenAndServe(os.Getenv("SERVER_PORT"), gHandlers.CombinedLoggingHandler(os.Stdout, r))

	if err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
