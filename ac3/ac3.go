package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"projeto/util"
)

func main() {
	var contatos [5]util.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite 1 -> adicionar, 2 -> remover, 3 -> editar email ou sair, 4 -> ver e-mails registrados ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Digite o nome: ")
			nome, _ = leitor.ReadString('\n')
			nome = strings.TrimSpace(nome)

			fmt.Print("Digite o email: ")
			email, _ = leitor.ReadString('\n')
			email = strings.TrimSpace(email)

			util.AdicionaContato(util.Contato{Nome: nome, Email: email}, &contatos)
		case "2":
			util.ExcluiContato(&contatos)
		case "3":
			var ind int
			fmt.Print("Digite o índice do contato para editar o email: ")
			fmt.Scanln(&ind)

			fmt.Print("Digite o email: ")
			email, _ = leitor.ReadString('\n')
			email = strings.TrimSpace(email)

			util.EditaEmail(&contatos, ind-1, email)

		case "4":
			util.ExibeContatos(&contatos)
		default:
			return
		}
	}
}