/***Algoritmo quicksort! modificacao: entrada aceita float64***/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	entrada := os.Args[1:]
	numeros := make([]float64 , len(entrada))

	for i,n := range entrada {
		numero, err := strconv.ParseFloat(n,64)
		if err != nil{
			fmt.Printf("%s nao é um numero valido! \n",n)
			os.Exit(1)
		}
		numeros[i] = numero
	}

	fmt.Println(quicksort(numeros))
} 

func quicksort(numeros []float64) []float64 {
	if len(numeros )<= 1 {
		return numeros
	}

	n := make ([]float64 , len(numeros))
	copy(n,numeros)

	indicePivo := len(n)/2
	pivo := n[indicePivo]
	n = append(n[:indicePivo], n[indicePivo+1:]...)

	menores,maiores := particionar(n, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(numeros []float64 , pivo float64) (menores []float64 , maiores []float64){
	for _,n := range numeros{
		if n<= pivo{
			menores = append(menores,n)
		}else {
			maiores = append(maiores,n)
		}
	}
	return menores,maiores
}
