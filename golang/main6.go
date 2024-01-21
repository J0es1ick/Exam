package main

import (
	"fmt"
	"log"
	"isuct.ru/informatics2022/bilet1/str6.go"

func main() {
	product1, err := product.NewProduct("laptop", 2)
	if err != nil {
		log.Fatal(err)
	}

	err = product1.SetAmount(32)
	if err != nil {
		log.Fatal(err)
	}

	product2, err := product.NewPerson("John", 64)
	if err != nil {
		log.Fatal(err)
	}

	all := make([]product.Person, 0)
	fmt.Println(product.TryAdd(&all, product1))
	fmt.Println(product.TryAdd(&all, product2))

	fmt.Println(product.Average(all))
}
