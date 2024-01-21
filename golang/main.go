package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/bilet1/bilet1.go"
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

	person2, err := bilet1.NewPerson("John", 64)
	if err != nil {
		log.Fatal(err)
	}

	collection := make([]bilet1.Person, 0)
	fmt.Println(bilet1.TryAdd(&collection, person1))
	fmt.Println(bilet1.TryAdd(&collection, person2))

	fmt.Println(bilet1.CalculateAvgAge(collection))
}
