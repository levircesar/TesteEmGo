/*** Levir César Ribeiro Lemos ***/ 

package main 

import "fmt"

import "math"

//Criar funcao
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
		total +=numero
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
	for i:= 1; i<=3 ; i= i +1{
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
	a:= []float64{10.5,20,30.5}

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

	//Input 
	var entrada float64
	fmt.Printf("Digite um numero: ")
	fmt.Scan(&entrada)
	fmt.Println("Você escreveu : ",entrada)
	
	//Adicionar valor no vetor usando append
	var b []float64
	fmt.Println("valores iniciais do vetor b = ", b)
	fmt.Println("adicione 3 valores ao array : ")
	for i:= 0; i<=2 ; i= i +1{
		var j float64;
		fmt.Print("incluir no vetor : ")
		fmt.Scan(&j)
		b = append(b,j)
	}
	fmt.Println("valores finais do vetor b = ", b)
	fmt.Println("Soma dos valores do vetor b =  " , somaVariavel(b...))

	//Adicionar no vetor usando switch
	var c []float64
	var k float64 
	var caso string
	var flag bool = true
	for flag ==true{
	fmt.Println("digita add/end")
	fmt.Scan(&caso)
		switch {
			case caso =="add":
				fmt.Print("incluir no vetor : ")
				fmt.Scan(&k)
				c = append(c,k)
			case caso =="end":
				flag = false
			default:
				fmt.Print("use add ou end ")
		}
	}
	fmt.Println("valores finais do vetor c = ", c)
	fmt.Println("Soma dos valores do vetor c =  " , somaVariavel(c...))

	//fim
	fmt.Println("fim")
}