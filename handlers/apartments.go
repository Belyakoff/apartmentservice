package handlers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	
)


// Apartments handler for getting and updating apartments
type Apartments struct {
	l    *log.Logger
	col  *mongo.Collection
}

// NewApartments returns a new apartments handler with the given logger
func NewApartments(l *log.Logger, col *mongo.Collection) *Apartments {
	return &Apartments{l,col}
}

type KeyApartment struct{}


// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func getApartmentID(r *http.Request) int {
	// parse the apartment id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}


