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
		p.l.Println("[ERROR] serializing apartment", err)
	}
}


// ListSingle handles GET requests
func (p *Apartments) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getApartmentID(r)

	p.l.Println("[DEBUG] get record id", id)

	apartment, err := data.GetApartmentByID(id)

	switch err {
	case nil:

	case data.ErrApartmentNotFound:
		p.l.Println("[ERROR] fetching apartment", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching apartment", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(apartment, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing apartment", err)
	}
}
