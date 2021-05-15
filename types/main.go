package main

import "fmt"

type Currency string
type VATRate float64

const Euro Currency = "EUR"
const FranceVATRate VATRate = 19.6

type Hotel struct {
	Name string
	Country
}

type Country struct {
	Name           string
	CapitalCity    string
	IsPartOfEurope bool
}

func main() {
	var euro Currency
	euro = "EUR"
	fmt.Println(Euro, euro)
	fmt.Println(FranceVATRate)

	france := Country{
		Name:           "France",
		CapitalCity:    "Paris",
		IsPartOfEurope: true,
	}
	fmt.Println(france)
	var usa Country
	usa.Name = "United States"
	usa.CapitalCity = "Washington DC"
	usa.IsPartOfEurope = false
	fmt.Println(usa)
	fmt.Println(france.CapitalCity)
	fmt.Println("USA is part of europe ?", usa.IsPartOfEurope)

	gopherHotel := Hotel{
		Name:    "Gopher Hotel Premium",
		Country: usa,
	}
	fmt.Println(gopherHotel.Country.IsPartOfEurope)
}
