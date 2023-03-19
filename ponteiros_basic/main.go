package main

import "fmt"

//declaração básica de ponteiros
var str string = "golang"

var ptr *string = &str

//ponteiros iniciam com valor nil
var ptrNil *string

func main() {

	//checar output declaração básica
	fmt.Println("Este é o valor de str: ", str)
	fmt.Println("Este é o endereço de str: ", &str)
	fmt.Println("Este é o valor de ptr: ", ptr)
	fmt.Println("Este é o valor ao qual ptr aponta: ", *ptr)

	//checar ponteiro sem utilização
	fmt.Println("Este é o valor de ptrNil: ", ptrNil)

	//pode-se construir ponteiros a partir da syntax curta
	str2 := "go"
	ptr2 := &str2

	fmt.Println("Este é o valor de str2: ", str2)
	fmt.Println("Este é o endereço de str2: ", &str2)
	fmt.Println("Este é o valor de ptr2: ", ptr2)
	fmt.Println("Este é o valor ao qual ptr2 aponta: ", *ptr2)

	//podemos alterar o valor da variavel chamando o ponteiro
	*ptr2 = "lang"
	fmt.Println("Este é o valor de str2: ", str2)
	fmt.Println("Este é o endereço de str2: ", &str2)
	fmt.Println("Este é o valor de ptr2: ", ptr2)
	fmt.Println("Este é o valor ao qual ptr2 aponta: ", *ptr2)

}
