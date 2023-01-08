package product

import (
	"errors"
	"net/http"

	"github.com/Chakravarthy7102/serverless/internal/handlers"
	"github.com/Chakravarthy7102/serverless/internal/repository/adapter"
	HttpStatus "github.com/Chakravarthy7102/serverless/utils/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Handler struct {
	handlers.Interface
	Controller product.Interface
	Rules      rules.Interface
}

func NewHandler(respository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(respository),
		Rules:      product.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "ID") != "" {
		h.GetOne(w, r)
	} else {
		h.GetAll()
	}
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))

	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not valid uuid"))
	}

	response, err := h.Controller.ListOne(ID)

	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) GetAll() {}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {

	productBody, err := h.getBodyAndValidate(r, uuid.Nil)

	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	ID, err := h.Controller.Create(productBody)

	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
	}

	HttpStatus.StatusOK(w, r, map[string]interface{}{
		"id": ID.String(),
	})

}

func (h *Handler) Put() {}

func (h *Handler) Delete() {}

func (h *Handler) Options() {}

func (h *Handler) getBodyAndValidate() {}
