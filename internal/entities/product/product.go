package product

import (
	"encoding/json"

	"github.com/Chakravarthy7102/serverless/internal/entities"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Product struct {
	//struct composition
	entities.Base
	Name string `json:"name"`
}

func (p *Product) InterfaceToModal(data interface{}) (instance Product, err error) {

	bytes, err := json.Marshal(data)

	if err != nil {
		return instance, err
	}

	return instance, json.Unmarshal(bytes, &instance)

}

//

//TODO
func (p *Product) GetFilterId() map[string]interface{} {

}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{} {

}

func ParseDynamoAttributeToStruct(Item map[string]*dynamodb.AttributeValue) (entity Product, err error) {

}
