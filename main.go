package main

import (
	"fmt"
	"log"

	"main.go/simpleinterest"
	// "learnpackage/simpleinterest"
)

var p, r, t = -5000.0, 10.0, 1.0

func init() {
	// fmt.Println("Simple interest package initialized")
	println("Main package intitialized")
	if p < 0 {
		log.Fatal("Principal is less the zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less the zero")
	}
	if t < 0 {
		log.Fatal("Duration is less the zero")
	}
}

func main() {
	fmt.Println("Simple interest calculation")
	// p := 5000.0
	// r := 10.0
	// t := 1.0
	si := simpleinterest.Calculate(p, r, t)

	fmt.Println("Simple interest is", si)
}
