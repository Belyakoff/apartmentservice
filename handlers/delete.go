package handlers

import (
	"net/http"

		"github.com/Belyakoff/apartmentservice/data"
)


func (p *Apartments) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getApartmentID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteApartment(id)
	if err == data.ErrApartmentNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}