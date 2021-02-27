package handlers

import (
	"net/http"

	"github.com/Belyakoff/apartmentservice/data"
)


func (p *Apartments) Update(rw http.ResponseWriter, r *http.Request) {

	// fetch the apartment from the context
	apartment := r.Context().Value(KeyApartment{}).(data.Apartment)
	p.l.Println("[DEBUG] updating record id", apartment.ID)

	err := data.UpdateApartment(apartment)
	if err == data.ErrApartmentNotFound {
		p.l.Println("[ERROR] apartment not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Apartment not found in database"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}