package main

import "fmt"

type Car interface {
	Brand() string
	Type() string
	Price() string
}

type Fortuner struct{}
type Pajero struct{}

func (f Fortuner) Brand() string {
	return "Toyota"
}

func (p Pajero) Brand() string {
	return "Mitsubishi"
}

func (f Fortuner) Type() string {
	return "Suv"
}

func (p Pajero) Type() string {
	return "Suv"
}

func (f Fortuner) Price() string {
	return "500.000.000"
}

func (p Pajero) Price() string {
	return "530.000.000"
}

func main() {
	var vehicle1 Car = Fortuner{}
	var vehicle2 Car = Pajero{}
	desc1 := fmt.Sprintf("Mobil pertamaku bermerek %s dan bertipe %s, yang aku beli seharga Rp %s", vehicle1.Brand(), vehicle1.Type(), vehicle1.Price())
	desc2 := fmt.Sprintf("Mobil keduaku bermerek %s dan bertipe %s, yang aku beli seharga Rp %s", vehicle2.Brand(), vehicle2.Type(), vehicle2.Price())
	fmt.Println(desc1)
	fmt.Println(desc2)
}
