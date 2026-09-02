package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor: ", err)
		return
	}
	defer conn.Close()

	// message := "Ola, servidor echo com buffer em go\n"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite suas mensagens (digite 'sair' para encerrar):")
	for {
		fmt.Print("calc>")
		// nao da pra passar o scanner pro buffer de leitura e escrita dos dados na conexao
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		dados := scanner.Bytes()
		dados = append(dados, '\n')

		_, err = conn.Write([]byte(dados))
		if err != nil {
			fmt.Println("Erro ao enviar dados: ", err)
			return
		}

		fmt.Printf("Enviado: %s", dados)

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("erro ao ler resposta do servidor: ", err)
			return
		}

		fmt.Printf("Echo recebido do servidor (%d bytes): %s", n, string(buf[:n]))
		if strings.ToLower(strings.TrimSpace(input)) == "sair" {
			fmt.Println("Encerrando o programa...")
			return
		}
	}

}
