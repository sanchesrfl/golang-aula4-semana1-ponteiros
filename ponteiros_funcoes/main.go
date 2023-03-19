package main

import "fmt"

//funcao que nao trabalha com ponteiros
func parseSemPonteiro(s string) {

	//recebe input e altera valor
	s = "Novo Valor Sem Ponteiro"

	//checa o endereço de "s"
	fmt.Println("O valor de s no escopo da função 'parseSemPonteiro' é: ", s)

}

func parseComPonteiro(s *string) {

	//recebe input e altera valor
	*s = "Novo Valor Sem Ponteiro"

	//checa o endereço de "s"
	fmt.Println("O valor de s no escopo da função 'parseSemPonteiro' é: ", &s)

}

func main() {

	//variavel inicial para manipualação sem e com ponteiro
	s := "golang"
	fmt.Println("O endereço da string no escopo de main é: ", &s)
	fmt.Println("O valor de s no escopo de main é: ", s)

	//chamamos função sem ponteiro
	parseSemPonteiro(s)
	fmt.Println("O valor de s no escopo de main é: ", s)

	//chamamos função com ponteiro
	parseComPonteiro(&s)
	fmt.Println("O valor de s no escopo de main é: ", s)
	fmt.Println("O endereço da string no escopo de main é: ", &s)

}
