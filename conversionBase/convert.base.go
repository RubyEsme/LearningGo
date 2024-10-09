package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Este es el algoritmo para convertir un numero de base 10 a cualquier otra base n.")

	var n, c, b int64
	var m int64
	var restos []int64 // Slice para almacenar los restos
	var hexChar string
	var hex int64
	var restosHex []string

	// n = número entero, c = cociente, b = base a la que quiere llegar, m = modulus

	fmt.Println("Ingrese el numero que quiere convertir: ")
	fmt.Scan(&n)
	fmt.Println("A que base la quiere convertir?:")
	fmt.Scan(&b)

	for n != 0 {
		if b == 16 {

			c = n / b                              // Dividir el entero entre la base que se quiere llegar y obtener el cociente
			hex = n % b                            // Sacar el resto de la división
			hexChar = strconv.FormatInt(hex, 16)   //parece que esta funcion ya los convierte a valores necesitados en hexadecimal
			restosHex = append(restosHex, hexChar) // Guardar el resto en el slice
			n = c                                  // Igualar el número al cociente para continuar con la siguiente división

		} else {
			c = n / b                  // Dividir el entero entre la base que se quiere llegar y obtener el cociente
			m = n % b                  // Sacar el resto de la división
			restos = append(restos, m) // Guardar el resto en el slice
			n = c                      // Igualar el número al cociente para continuar con la siguiente división
		}

	}
	if b == 16 {
		// Imprimir los restos en orden inverso
		fmt.Println("El número convertido es:")
		//el como funciona la funcion es que toma el lenght del array, y si esta en la posicion n-1 (por el orden de index en el que se cuenta) entonces recorre desde el punto mas lejano, osea el 3, y luego va al 2 y asi
		for i := len(restosHex) - 1; i >= 0; i-- {
			fmt.Printf("%s", restosHex[i])
		}

		fmt.Println() // Salto de línea
	} else {
		// Imprimir los restos en orden inverso
		fmt.Println("El número convertido es:")
		//el como funciona la funcion es que toma el lenght del array, y si esta en la posicion n-1 (por el orden de index en el que se cuenta) entonces recorre desde el punto mas lejano, osea el 3, y luego va al 2 y asi
		for i := len(restos) - 1; i >= 0; i-- {
			fmt.Printf("%d", restos[i])
		}

		fmt.Println() // Salto de línea

	}

}
