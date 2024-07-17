package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"os/exec"
	"runtime"
	"time"
	"math/rand"
)

func clear() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao limpar o terminal:", err)
	}
}

func printar_tabela() {
	regioes := []string{
		"Rio Grande do Sul",
		"Distrito Federal, Goiás, Mato Grosso, Mato Grosso do Sul e Tocantins",
		"Pará, Amazonas, Acre, Amapá, Rondônia e Roraima",
		"Ceará, Maranhão e Piauí",
		"Pernambuco, Rio Grande do Norte, Paraíba e Alagoas",
		"Bahia e Sergipe",
		"Minas Gerais",
		"Rio de Janeiro e Espírito Santo",
		"São Paulo",
		"Paraná e Santa Catarina"}

	tabela := table.NewWriter()
	tabela.SetOutputMirror(os.Stdout)
	tabela.AppendHeader(table.Row{"CÓDIGO", "REGIÃO"})
	for i := range regioes {
		tabela.AppendRow([]interface{}{i, regioes[i]})
		tabela.AppendSeparator()
	}
	tabela.SetStyle(table.StyleColoredDark)
	tabela.Render()
}

func gerar_cpf(codigo int) string {
	senha := []int{}
	var senha_concatenada string
	var soma int
	var multi int

	for i := 0; i < 8; i++ {
		random_number := rand.Intn(10)
		senha = append(senha, random_number)
	}
	senha = append(senha, codigo)

	multi = 10
	for j := 0; j < len(senha); j++ {
		soma += senha[j] * multi
		multi--
	}

	resto1 := soma % 11

	if resto1 < 2 {
		resto1 = 0
	} else {
		resto1 = 11 - resto1
	}

	senha = append(senha, resto1)

	soma = 0
	multi = 11
	for j := 0;  j < len(senha); j++ {
		soma += senha[j] * multi
		multi--
	}

	resto2 := soma % 11

	if resto2 < 2 {
		resto2 = 0
	} else {
		resto2 = 11 - resto2
	}

	senha = append(senha, resto2)

	for i := range senha {
		senha_concatenada += fmt.Sprintf("%d", senha[i])
	}

	return senha_concatenada
}

func main() {
	var cidade int

	for {
		fmt.Println("╭────────────────╮")
		fmt.Println("│ GERADOR DE CPF │")
		fmt.Println("╰────────────────╯")
		fmt.Println("")

		printar_tabela()

		fmt.Println("\nDigite o código da sua cidade:")
		_, err := fmt.Scan(&cidade)
		if err != nil {
			panic(err.Error())
		}

		if cidade < 0 || cidade > 9 {
			fmt.Println("\nCodigo invalido!")
			time.Sleep(2 * time.Second)
			clear()
		}

		fmt.Println("\nSeu CPF:", gerar_cpf(cidade))
		os.Exit(0)
	}
}