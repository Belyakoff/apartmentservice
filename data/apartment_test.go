package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApartmentMissingTitleReturnsErr(t *testing.T) {
	p := Apartment{
		Price: 10.22,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestApartmentMissingPriceReturnsErr(t *testing.T) {
	p := Apartment{
		Title:  "kvartira 20 kv, studio",
		Price: -1,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestApartmentInvalidPhoneNumReturnsErr(t *testing.T) {
	p := Apartment{
		Title:  "abc",
		Price: 1.22,
		Phone_number:   "abc",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestValidApartmentDoesNOTReturnsErr(t *testing.T) {
	p := Apartment{
		Title:  "abc",
		Price: 1.22,
		Phone_number:   "89003003030",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestApartmentsToJSON(t *testing.T) {
	ps := []*Apartment{
		&Apartment{
			Title: "abc",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(ps, b)
	assert.NoError(t, err)
}