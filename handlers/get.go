package handlers

import (
	"net/http"

	"github.com/Belyakoff/apartmentservice/data"
)


func (p *Apartments) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	apartments := data.GetApartments()

	err := data.ToJSON(apartments, rw)
	
    if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
