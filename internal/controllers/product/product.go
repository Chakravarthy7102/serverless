package product

import (
	"time"

	product "github.com/Chakravarthy7102/serverless/internal/entities/product"
	product "github.com/Chakravarthy7102/serverless/internal/handlers/product"
	"github.com/Chakravarthy7102/serverless/internal/repository/adapter"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

func (c *Controller) ListAll() (entities []product.Product, err error) {

	// entities := []product.Product{}

	var entity product.Product

	filter := expression.Name("name").NotEqual(expression.Value(""))

	condition, err := expression.NewBuilder().WithFilter(filter).Build()

	if err != nil {
		return entities, err
	}

	response, err := c.repository.FindAll(condition, entity.TableName())

	if err != nil {
		return entities, err
	}

	if response != nil {
		for _, value := range response.Items {
			entity, err := product.ParseDynamoAttributeToStruct(value)

			if err != nil {
				return entities, err
			}

			entities = append(entities, entity)
		}
	}
	return entities, nil

}

func (c *Controller) Create(entity *product.Product) (ID uuid.UUID, err error) {
	entity.CreatedAt = time.Now()

	_, err = c.repository.CreateOrUpdate(entity, entity.TableName())

	if err != nil {
		return entity, nil
	}

	return entity.ID, err
}

func (c *Controller) Update(id uuid.UUID, entity *product.Product) error {
	found, err := c.ListOne(id)

	if err != nil {
		return err
	}

	found.ID = id
	found.UpdatedAt = time.Now()
	found.Name = entity.Name
	_, err = c.repository.CreateOrUpdate(found.GetMap(), entity.TableName())

	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) Remove(id uuid.UUID) error {
	entity, err := c.ListOne(id)

	if err != nil {
		return err
	}

	_, err = c.repository.Delete(entity.GetFilterId(), entity.TableName())

	if err != nil {
		return err
	}

	return nil
}
