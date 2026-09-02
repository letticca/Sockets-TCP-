package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor: ", err)
	}
	defer listener.Close()
	fmt.Println("Servidor aguardando conexões na porta 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão: ", err)
			continue
		}
		go tratarConexao(conn)
	}
}

func tratarConexao(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Cliente conectado %s\n", conn.RemoteAddr())
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Cliente desconectado ou erro de leitura: %v\n", err)
			break
		}

		//fmt.Printf("Recebido %d byte: %s", n, string(buf[:n]))

		comando := string(buf[:n])
		partes := strings.Fields(comando)

		if len(partes) == 0 {
			conn.Write([]byte("ERRO: insira um comando \n"))
			continue
		} else if len(partes) < 3 && strings.ToUpper(partes[0]) != "SAIR" {
			_, err = conn.Write([]byte("ERRO: comando invalido\n"))
			if err != nil {
				fmt.Println("Erro ao enviar dados:", err)
				return
			}
			continue

		} else if len(partes) == 1 && strings.ToLower(partes[0]) == "sair" {
			conn.Write([]byte("Encerrando programa\n"))
			return
		} else if len(partes) != 3 {
			_, err = conn.Write([]byte("ERRO: comando invalido \n"))
			if err != nil {
				fmt.Println("Erro ao enviar dados:", err)
				return
			}
			continue
		}

		// associando as partes do comando a variaveis para poder transformar em inteiro
		num1, err := strconv.Atoi(partes[1])
		if err != nil {
			conn.Write([]byte("ERRO: formato invalido\n"))
			continue
		}
		num2, err := strconv.Atoi(partes[2])
		if err != nil {
			conn.Write([]byte("ERRO: formato invalido\n"))
			continue
		}

		resposta := " "
		switch strings.ToUpper(partes[0]) {
		case "SOMA":
			resultado := num1 + num2
			resposta = fmt.Sprintf("RESULTADO: %d\n", resultado)
		case "SUB":
			resultado := num1 - num2
			resposta = fmt.Sprintf("RESULTADO: %d\n", resultado)
		case "MUL":
			resultado := num1 * num2
			resposta = fmt.Sprintf("RESULTADO: %d\n", resultado)
		case "DIV":
			if num2 == 0 {
				// erro da divisão
				_, err = conn.Write([]byte("ERRO: divisao por zero\n"))
				if err != nil {
					fmt.Println("Erro ao enviar dados:", err)
					return
				}
				continue
			}
			resultado := num1 / num2
			resposta = fmt.Sprintf("RESULTADO: %d\n", resultado)
		default:
			// erro do comando
			_, err = conn.Write([]byte("ERRO: comando invalido\n"))
			if err != nil {
				fmt.Println("Erro ao enviar dados:", err)
				return
			}
			continue

		}
		_, err = conn.Write([]byte(resposta))
		if err != nil {
			fmt.Println("Erro ao enviar dados:", err)
			return
		}

	}

}
