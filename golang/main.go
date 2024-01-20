package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/bilet1"
)

func main() {
	person1, err := bilet1.NewPerson("Bob", 33)
	if err != nil {
		log.Fatal(err)
	}

	err = person1.SetAge(32)
	if err != nil {
		log.Fatal(err)
	}

	person2, err := bilet1.NewPerson("John",52)
	if err != nil {
		log.Fatal(err)
	}
	
	collection := make([]bilet1.Person,0)
	fmt.Println(bilet1.TryAdd(&collection, person1))
	fmt.Println(bilet1.TryAdd(&collection, person2))

	fmt.Println(bilet1.CalculateAvgAge(collection))

	person3, err := bilet1.NewPerson("asd", 25)
	fmt.Println(person3)

	person3.SetAge(61)

	fmt.Println(person3)

	fmt.Println(bilet1.CalculateSum(collection))
}