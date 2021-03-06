package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args)<3{
		fmt.Println("uso; conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	}else if unidadeOrigem =="quilometros"{
		unidadeDestino = "milhas"
	}else {
		fmt.Printf("%s nao é uma unidade conhecida!\n",unidadeDestino)
		os.Exit(1)
	}

	for i, v :=range valoresOrigem{
		valorOrigem, err := strconv.ParseFloat(v,64)
		if err != nil{
			fmt.Printf(" o valor %s na posicao %d nao é um numero valido \n", v,i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem =="celsius"{
			valorDestino = valorOrigem*1.8 +32
		}else{
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s \n ", valorOrigem,unidadeOrigem,valorDestino,unidadeDestino)
	}
}