package product

import (
	"github.com/Chakravarthy7102/serverless/internal/handlers/product"
	"github.com/Chakravarthy7102/serverless/internal/repository/adapter"
	"github.com/goolge/uuid"
)

type Controller struct {
	repository adapter.Interface
}

type Interface interface {
	ListOne(ID uuid.UUID) (entity product.Product, err error)
	ListAll() (entity []product.Product, err error)
	Create(entity *product.Product) (uuid uuid.UUID, err error)
	Update(ID uuid.UUID, entity *product.Product) error
	Remove(ID uuid.UUID) error
}

func NewController(repository adapter.Interface) Interface {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) ListOne(id uuid.UUID) (entity product.Product, err error) {
	entity.ID = id
	response, err := c.repository.FindOne(entity.GetFilterId(), entity.TableName())

	if err != nil {
		return entity, err
	}

	return product.ParseDynamoAttributeToStruct(response.Item)

}

func (c *Controller) ListAll() (entity []product.Product, err error) {

}

func (c *Controller) Create(entity *product.Product) (ID uuid.UUID, err error) {}

func (c *Controller) Update(id uuid.UUID, entity *product.Product) {}

func (c *Controller) Remove(id uuid.UUID) error {}
