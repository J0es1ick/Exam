package product

import "errors"

type Product struct {
	name   string
	cost   float64
	amount float64
}

func (p Product) GetName() string {
	return p.name
}

func (p Product) GetCost() float64 {
	return p.cost
}

func (p Product) GetAmount() float64 {
	return p.amount
}

func (p *Product) SetAmonut(amount float64) float64 {
	if amount >= 0 && amount <= 1000 {
		p.amount = amount
		return nil
	}
	ErrWrongAmount := errors.New("Impossible amount")
	return ErrWrongAmount
}

func NewProduct(name string, amount float64) (Product error) {
	var p Product = Product{
		name:   name,
		amount: amount,
	}
	if name == "phone" {
		p.cost = 25000
		return p.cost
	} else if name == "laptop" {
		p.cost = 65000
		return p.cost
	} else if name == "television" {
		p.cost = 100000
		return p.cost
	}
	if p.cost < 0 || p.cost > 100000 {
		var err = NewError("Cost is too low or too high")
		return p, err
	}
	return p, err
}

func Sum(all []Product) int {
	result := 0
	for _, x := range all {
		result += x.GetCost()
	}

	return result
}

func Average(all []Product) float64 {
	av := float64(Sum(all)) / float64(len(all))
	return av
}

func TryAdd(all *[]Product, p Product) bool {
	for _, y := range *all {
		if y.GetName() == p.GetName() && y.GetAge() == p.GetAge() {
			return false
		}
	}
	*all = append(*all, p)
	return true
}
