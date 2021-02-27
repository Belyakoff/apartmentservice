package data


import (
	"fmt"
)


type Apartment struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	HREF 		  string  `json:"href" validate:"required"`
	Rayon         string  `json:"rayon"`
	Adress        string  `json:"adress"`
	Price         float32 `json:"price"`
	Subprice       string  `json:"subprice"`
	Phone_number  string  `json:"phone_number" validate:"phone_number"`
	Description_text  string `json:"description_text"`
}

type Apartments []*Apartment

var ErrApartmentNotFound = fmt.Errorf("Apartment not found")

func GetApartments() Apartments {
	return apartmentList
}

func AddApartment(p Apartment){
	//fmt.Printf("v%",p)
	maxID := apartmentList[len(apartmentList)-1].ID
	p.ID = maxID + 1
	apartmentList = append(apartmentList, &p)
}

func DeleteApartment(id int) error {
	i := findIndexByApartmentID(id)
	if i == -1 {
		return ErrApartmentNotFound
	}

	apartmentList = append(apartmentList[:i], apartmentList[i+1])

	return nil
}

func UpdateApartment(p Apartment) error {
	i := findIndexByApartmentID(p.ID)
	if i == -1 {
		return ErrApartmentNotFound
	}

	// update the apartment in the DB
	apartmentList[i] = &p

	return nil
}

func findIndexByApartmentID(id int) int {
	for i, p := range apartmentList {
		if p.ID == id {
			return i
		}
	}

	return -1
}




var apartmentList = []*Apartment {
	&Apartment{
		ID: 		  1,
		Title: 		  "kvartira 39 kv, 1 komnata",
		HREF: 		  "http://somelink",
		Rayon: 		  "Tushino",
		Adress: 	  "Trikotazhnaya",
		Price:		  40.0,
		Subprice: 	  "zalog 50",
		Phone_number: "89001001010",
		Description_text : "sdam 1 komnatnuu kvartiru",
	},
	&Apartment{
		ID: 		  2,
		Title: 		  "kvartira 20 kv, studio",
		HREF: 		  "http://somelink",
		Rayon: 		  "Hovrino",
		Adress: 	  "unknown",
		Price:		  32.0,
		Subprice: 	  "zalog 100",
		Phone_number: "89002002020",
		Description_text : "sdam studiu",
	},
}