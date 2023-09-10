package main

import (
	"fmt"
)

func achaPar(array []int, alvo int) (int, int) { // percorre para encontrar o par
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(arr); j++ {
			if array[i]+array[j] == alvo {
				return array[i], array[j]
			}
		}
	}
	return -1, -1 // neenhum par foi encontrado
}

func main() {
	array := []int{3, 5, 8, 10, 12}
	alvo := 15

	num1, num2 := achaPar(array, alvo) // verifica se o par foi encontrado
	if num1 != -1 && num2 != -1 {
		fmt.Printf("Alvo: %d, Par: (%d, %d)\n", alvo, num1, num2)
	} else {
		fmt.Println("Nenhum par encontrado.")
	}
}