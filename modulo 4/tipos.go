package main

import "fmt"

func main() {
	// Booleans: True or False
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	// Strings: Sêquencia de bytes
	fmt.Printf("Type: %T - Value: %v\n", "Gabriel", "Gabriel")
	// Int: Numeros Inteiros
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	// Strings: Sêquencia de bytes - ao colocar o "" o numero 1 que até então era inteiro se torna uma String
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")
	// Float: Numeros Decimais (Float 64 / Float 32) Float 64 e o mais ultilizado dentro da linguagem de programação Go
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)

}

// Tipos:
// Booleans: True or False
// Strings: Sêquencia de bytes
// Int: Numeros Inteiros
// Float: Numeros Decimais (Float 64 / Float 32)
