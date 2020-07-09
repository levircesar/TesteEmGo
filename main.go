package main

import "fmt"

func main() {
	var(
		nome = "levir"
		idade int = 21
	)
	fmt.Println("Hello,World")
	for i:= 1; i<100 ; i= i +1{
		fmt.Println("numero : ",i)
	}

	fmt.Println("Meu nome é : " ,nome, "e tenho ",idade," anos")

	fmt.Println("fim")
}
