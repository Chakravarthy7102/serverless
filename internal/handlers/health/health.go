package health

import (
	"errors"
	"net/http"

	"github.com/Chakravarthy7102/serverless/internal/handlers"
	"github.com/Chakravarthy7102/serverless/internal/repository/adapter"
	HttpStatus "github.com/Chakravarthy7102/serverless/utils/http"
)

type Handler struct {
	handlers.Interface
	Respository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Respository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Respository.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}

func Put() {}

func Delete() {}

func Options() {}
