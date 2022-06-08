package main

import (
	"fmt"
)

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {

	timmy := &person{
		name: "timothy",
		age:  10,
	}
	timmy.superpower = "flying"
	(*timmy).superpower = "flying"
	fmt.Printf("%+v\n", timmy)

	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println(*superpowers)
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])
	fmt.Println((*superpowers)[2:])
	fmt.Println(superpowers[2:])

	rebecca := person{
		name:       "rebeccca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)

}
