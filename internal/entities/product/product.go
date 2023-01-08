package product

import (
	"encoding/json"

	"github.com/Chakravarthy7102/serverless/internal/entities"
)

type Product struct {
	//struct composition
	entities.Base
	Name string `json:"name"`
}

func (p *Product) InterfaceToModal(data interface{}) (instance Product, err error) {

}

//TODO
func (p *Product) GetFilterId() map[string]interface{} {

}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}
