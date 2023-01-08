package product

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/Chakravarthy7102/serverless/internal/entities"
	"github.com/Chakravarthy7102/serverless/internal/entities/product"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

type Rules struct{}

func NewRules() *Rules {
	return &Rules{}
}

func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("body is invalid")
	}
	return model, json.NewDecoder(data).Decode(model)
}

func (r *Rules) Migrate(connection *dynamodb.DynamoDB) error {
	return r.CreateTable(connection)
}

func (r *Rules) GetMock() interface{} {
	return product.Product{
		Base: entities.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: uuid.New().String(),
	}
}

func (r *Rules) Validate(model interface{}) error {
	var product product.Product
	productModal, err := product.InterfaceToModal(model)

	if err != nil {
		return err
	}

	return Validation.ValidateStruct(productModal,
		Validation.Filed(&productModel.ID, Validatation.Required, is.UUIDv4),
		Validation.Filed(&productModal.Name, Validatation.Required, Validation.Length(3, 50)),
	)
}

func (r *Rules) createTable() {}
