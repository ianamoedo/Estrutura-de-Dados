package main

import (
    "fmt"
)

const M = 5

func main() {
    var lista [M]int
    var x = 0

    insereOrd(&lista, &x, 4)
    insereOrd(&lista, &x, 12)
    insereOrd(&lista, &x, 2)
    insereOrd(&lista, &x, 6)
    insereOrd(&lista, &x, 17)
    insereOrd(&lista, &x, 1)
}

func insereOrd(v *[M]int, n *int, z int) {
    if *n >= M {
        fmt.Println("Overflow")
        return
    }

    
    posic := *n - 1 // posição do valor novo
    for posic >= 0 && v[posic] > z {
        posic--
    }

    
    for i := *n - 1; i > posic; i-- { // para numeros maiores, jogar para direita
        v[i+1] = v[i]
    }

    
    v[posic+1] = z // inserindo o valor novo na posiçao certa
    *n++


	fmt.Printf("inserindo %d\n", z) // inserindo o valor novo 
	fmt.Println(v) // inserindo lista atualizada
}