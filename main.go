package main 

import "fmt"

//criar funcao
func soma(x int,y int) int{
	return x + y
}


func main() {
	//declarar variaveis
	var(
		nome string= "Levir"
		idade int = 21
		sobrenome string = "cesar"
	)

	//laço de repeticao absico
	fmt.Println("Hello,World")
	for i:= 1; i<=10 ; i= i +1{
		fmt.Println("numero : ",i)
	}

	//chamando funcao
	fmt.Println("A soma de 5 + 10 = " , soma(5,10))

	fmt.Println("Meu nome é : " ,nome, "e tenho ",idade," anos")

	//laco de repeticao para array
	a:= []int{10,20,30,40,50}

	for i, v:=range a{
		fmt.Println("posicao no vetor: ",i," valor: ",v)
	}

	//concatenar string
	fmt.Println("meu nick no githube: " ,nome+sobrenome)

	//fim
	fmt.Println("fim")

}
