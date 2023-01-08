package product

import (
	"errors"
	"net/http"
	"time"

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
		h.GetAll(w, r)
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

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	response, err := h.Controller.ListAll()

	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)

}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {

	productBody, err := h.getBodyAndValidate(r, uuid.Nil)

	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	ID, err := h.Controller.Create(productBody)

	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, map[string]interface{}{
		"id": ID.String(),
	})

}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))

	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not a valide uuid"))
		return
	}

	productBody, err := h.getBodyAndValidate(r, ID)

	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	if err := h.Controller.Update(ID, productBody); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	ID, err := uuid.Parse(chi.URLParam(r, "ID"))

	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, errors.New("Not a valid uuid"))
		return
	}

	if err := h.Controller.Remove(ID); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)

}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) getBodyAndValidate(r *http.Request, ID uuid.UUID) (*EntityProduct.Product, error) {
	productBody := &EntityProduct.Product{}

	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody)

	if err != nil {
		return &EntityProduct.Product{}, errors.New("Body is required.")
	}

	productParsed, err := EntityProduct.InterfaceToModel(body)

	if err != nil {
		return &EntityProduct.Product{}, errror.New("Error on converting the body to modal")
	}

	setDefaultValues(productParsed, ID)

	return productParsed, h.Rules.Validate(productParsed)

}

func setDefaultValues(product *EntityProduct.Product, ID uuid.UUID) {
	product.UpdatedAt = time.Now()
	if ID == uuid.Nil {
		product.ID = uuid.New()
		product.CreatedAt = time.Now()
	} else {
		product.ID = ID
	}

}
