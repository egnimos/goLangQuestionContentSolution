/*
QUESTION=>
Take the STRUCT “person” in the previous exercise and add a field called “favFood” that stores a slice of
string; for the variable “p1” use a composite literal to add values to the field favFood; print out the values
in favFood; also print out the values for “favFood” by using a for range loop
*/

package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	FavFood   []string
}

func main() {
	p1 := person{
		FirstName: "Niteesh",
		LastName:  "Dubey",
		FavFood:   []string{"Chola bhatura", "Phulaki", "Puri sabji", "Palak Paneer"},
	}

	fmt.Println("person favorite food: ", p1.FavFood)

	for _, v := range p1.FavFood {
		fmt.Println("So the Foods are=> :)", v)
	}
}
