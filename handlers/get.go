package handlers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/Belyakoff/apartmentservice/data"
	"github.com/Belyakoff/apartmentservice/repository"

)


func (p *Apartments) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	apartments := repository.GetItems(p.col)

	err := data.ToJSON(apartments, rw)
	
    if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing apartment", err)
	}
}


func (p *Apartments) ListByPrice(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	price, err := strconv.Atoi(vars["price"])
	if err != nil {
		panic(err)
	}
	op := vars["op"]
	op = "$"+op

	

	p.l.Println("[DEBUG] search by price ", op)

	apartments := repository.GetItemsByPrice(p.col, price, op)

	err = data.ToJSON(apartments, rw)
	
    if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing apartment", err)
	}
}