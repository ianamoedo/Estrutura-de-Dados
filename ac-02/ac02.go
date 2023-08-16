package main

import "fmt"

func main() {
	for {
		switch escolha{
		case nome, email:
			var nome, email string
			fmt.Println("Nome: ")
			fmt.Scanln(&nome)
			fmt.Println("Email: ")
			fmt.Scanln(&email)

			c := Contato{
				Nome: nome,
				Email: email,
			}

		}
	}

	c.alteraEmail("ianamoedo123@gmail.com")
	fmt.Println(c)


}

type Contato struct {
	Nome	string
	Email	string
}

func (c *Contato) alterarEmail(email string) {
	c.Email = email
}


func adicionaContato(c Contato, a [5]Contato) [5]Contato {
	for indice := range a {
		if (contato == Contato{}) {
			a[i] = c
			break
		}
	return a
	}
}

func excluiContato(a [5]Contato) [5]Contato {
	for indice, contato := range a {
		if (contato == Contato{}) {
			a[i-1] = c
			break
		}
	return a
	}
}



