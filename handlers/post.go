package handlers

import (
	"net/http"
	"github.com/Belyakoff/apartmentservice/data"
	"github.com/Belyakoff/apartmentservice/repository"
)


func (p *Apartments) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Apartment")

	apartment := r.Context().Value(KeyApartment{}).(data.Apartment)
	p.l.Printf("[DEBUG] Inserting apartment: %#v\n", apartment.Title)
	repository.InsertItem(p.col, apartment)
}