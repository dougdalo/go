package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibir logs...")
		case 0:
			fmt.Println("Sair do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não foi possível exibir o monitoramento")
			os.Exit(-1)
		}

	}
}
func exibeIntroducao() {
	name := "Douglas"
	versao := 1.1
	fmt.Println("Olá! Sr.", name)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"http://hering.com.br", "http://hering.com", "http://hering.net"}

	for i, site := range sites {
		fmt.Println("testando site", i, ":", site)
		testeSite(site)
	}

}

func testeSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Monitorando resposta de", site, "Site foi carregado com sucesso")
	} else {
		fmt.Println("Monitorando resposta de", site, "Site Está com problema... Verificar Logs", resp.StatusCode)
	}
}
