package main

import "errors"

type Product struct {
	ID          string
	Description string
	Stock       uint
}

func (p Product) IsInStock() bool {
	return p.Stock > 0
}

func (p *Product) SetDescription(description string) error {
	if len(description) <= 10 || len(description) >= 250 {
		return errors.New("length of the description is not correct")
	}
	p.Description = description
	return nil
}

func main() {

	p := Product{
		ID:    "42",
		Stock: 10,
	}
	//fmt.Println(p.IsInStock())
	err := p.SetDescription("sjdqjdisqdjiodjidjqdjiqjsjdqidji")
	if err != nil {
		//log.Fatalf("%s",err)
	}
	//fmt.Println("Success")

}
