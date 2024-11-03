package handlers

import (
	"github.com/Nelwhix/tunnel/pkg/models"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type Handler struct {
	Model     models.Model
	Logger    *slog.Logger
	Validator *validator.Validate
}

func (h *Handler) CreateNewTunnel(w http.ResponseWriter, r *http.Request) {

}
