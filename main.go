package main 

import "fmt"
import "math"

//criar funcao
func soma(x int,y int) int{
	return x + y
}
//retornar variavel na funcao
func diminuir(x float64,y float64) (total float64){
	total = x - y
	return
}

//funcao com entrada variavel
func somaVariavel(numeros ...float64)(total float64){
	for _, numero:= range numeros{
		total = numero+numero
	}
	return
}

func main() {
	//declarar variaveis
	var(
		nome string= "Levir"
		idade int = 21
		sobrenome string = "cesar"
	)

	//laço de repeticao básico
	fmt.Println("Hello,World")
	for i:= 1; i<=10 ; i= i +1{
		fmt.Println("numero : ",i)
	}

	//chamando funcao simples
	fmt.Println("A soma de 5 + 10 = " , soma(5,10))
	fmt.Println("Meu nome é : " ,nome, "e tenho ",idade," anos")

	//chamada de funcao que retorna variável
	fmt.Println("A operacao de 5.25 - 10 = " , diminuir(5.25,10))

	//chamada de funcao com parametro variável
	fmt.Println("A somaVariavel de 1,2,3,4,5 = " , somaVariavel(1,2,3,4,5))

	//laco de repeticao para array
	a:= []float64{10,20,30,40,50}

	for i, v:=range a{
		fmt.Println("posicao no vetor: ",i," valor: ",v)
	}

	//Passar array para funcao
	fmt.Println("Soma dos valores do array =  " , somaVariavel(a...))

	//concatenar string
	fmt.Println("meu nick no githube: " ,nome+sobrenome)

	//criar Slice vazio
	meuSlice := make([]int,0)
	meuSlice = append(meuSlice,3)
	meuSlice = append(meuSlice,74)
	meuSlice = append(meuSlice,53)
	fmt.Println("tamanho do slice : ", len(meuSlice))
	fmt.Println("capacidade do slice : ", cap(meuSlice))
	fmt.Println("tudo do slice : ",meuSlice)

	//importado da math
	fmt.Println("cos de 30 : ", math.Cos(30))

	//fim
	fmt.Println("fim")
}