package handler

import (
	"log"
	"microservice/data"
	"net/http"
	"regexp"
	"strconv"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle the post request
	if(r.Method == http.MethodPost){
		p.addProduct(rw, r)
		return
	}

	if(r.Method == http.MethodPut){
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProduct(rw, r, id)
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products...")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json ", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST Products...")

	prod:= data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
	http.Error(rw, "Unable to Unmarshall the JSON, Check the parameters provided ", http.StatusBadRequest)
	}

	data.AddProduct(prod)
	p.l.Println("Product Added Successfully...")
}

func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request, id int){
	p.l.Println("Handle PUT Products...")
	prod:= data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
	http.Error(rw, "Unable to Unmarshall the JSON, Check the parameters provided ", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id,prod)
	if err == data.ErrProductNotFound {
		http.Error(rw,"Product Not Found in the list, provide correct id ",http.StatusBadRequest)
	}else{
		p.l.Println("Product ID %v updated successfully...")
	}
}