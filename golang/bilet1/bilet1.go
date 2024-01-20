package bilet1

import (
	"errors"
)

type Person struct {
	age int
	name string
	status string
}

func (p Person) GetAge() int {
	return p.age
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) GetStatus() string {
	return p.status
}

func (p* Person) SetAge(age int) error {
	if age >= 0 && age <= 150 {
		p.age = age
		switch {
		case p.age < 18:
			p.status = "child"
	 	case p.age >= 18 && p.age < 60: 
			p.status = "adult"
		case p.age > 60:
			p.status = "elderly"
		}
		return nil
	}
	var ErrWrongAge = errors.New("age is higher or lower than bound")
	return ErrWrongAge
}

func NewPerson(name string, age int) (Person, error) {
	var p Person = Person {
		age: age,
		name: name,
	}
	var err = p.SetAge(age)
	return p, err 
}

func CalculateSum(collection []Person) int {
	result := 0
	for _, v := range collection {
		result += v.GetAge()
	}

	return result
}

func CalculateAvgAge(collection []Person) float64 {
	avg := float64(CalculateSum(collection)) / float64(len(collection))
	return avg
}

func TryAdd(collection *[]Person, p Person) bool {
	for _, x := range *collection {
		if x.GetName() == p.GetName() && x.GetAge() == p.GetAge() {
			return false
		}
	}
	*collection = append(*collection, p)
	return true
}