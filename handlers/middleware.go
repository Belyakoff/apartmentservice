package handlers 

import (
	"context"
	"net/http"
	"fmt"
	"github.com/Belyakoff/apartmentservice/data"
)

func (p *Apartments) MiddlewareValidateApartment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){


		apartment := &data.Apartment{}              

         
		err := data.FromJSON(apartment, r.Body) 
		if err != nil {
			// we should never be here but log the error just incase
			p.l.Println("[ERROR] deserializing apartment", err)


			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		v := data.NewValidation()

	  

		errs := v.Validate(apartment)

		fmt.Printf("Middle v%", apartment)
		
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating apartment", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		} 


		ctx := context.WithValue(r.Context(), KeyApartment{}, *apartment)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)

	})
}