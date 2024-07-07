package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

var ErrProductNotFound = fmt.Errorf("Product not found")

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// FromJson Decode the request body the defined struct of the product
func (p *Product) FromJSON(r io.Reader) error{
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

/**
 * AddProduct adds a new product to the productList.
 * The product's ID is automatically generated using getNextID().
*/
func AddProduct(p Product){
	p.ID = getNextID()
	productList = append(productList, &p)

}

func getNextID() int{
	return productList[len(productList)-1].ID +1
}

/**
 * UpdateProduct updates product to the productList.
 * The product's ID is fetch using fecthID() which is used to find if the id,
 * is present in the productList
*/
func UpdateProduct(id int, p Product) error{
	// find the product
	_, pos, err := fetchID(id)
	if err != nil {
		return err
	}
	productList[pos] = &p
	return nil
}

func fetchID(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound 
}



// productList is a hard coded list of products for this
func newProduct(id int, name, description string, price float32, sku string) *Product {
	now := time.Now().UTC().String()
	return &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		SKU:         sku,
		CreatedOn:   now,
		UpdatedOn:   now,
	}
}

var productList = []*Product{
	newProduct(1, "Latte", "Frothy milky coffee", 2.45, "abc323"),
	newProduct(2, "Espresso", "Short and strong coffee without milk", 1.99, "fjd34"),
}