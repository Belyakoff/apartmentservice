package handlers

import (
	"net/http"
	"github.com/Belyakoff/apartmentservice/data"
)


func (p *Apartments) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Apartment")

	apartment := r.Context().Value(KeyApartment{}).(data.Apartment)
	p.l.Printf("[DEBUG] Inserting apartment: %#v\n", apartment.Title)
	data.AddApartment(apartment)
}