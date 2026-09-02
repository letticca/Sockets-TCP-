# 🧮 Calculadora TCP Concorrente em Go

Este projeto implementa uma aplicação cliente-servidor em Go utilizando sockets TCP, permitindo realizar operações matemáticas básicas de forma concorrente, interativa e segura.

## 📂 Estrutura do Projeto

* 🖥️ **Servidor (`server.go`)**: Gerencia conexões simultâneas usando *goroutines*, processa comandos baseados em um protocolo de texto, realiza validações de entrada e executa os cálculos.
* 📱 **Cliente (`client.go`)**: Interface de linha de comando (`CLI`) que lê entradas do usuário via `bufio.NewScanner`, envia requisições estruturadas e exibe os resultados retornados pelo servidor.

## ⚙️ Protocolo de Comunicação

O cliente envia comandos em texto estruturado separados por espaços, seguindo o formato: `COMANDO <num1> <num2>`

| 🔤 Comando | 📝 Descrição | 💡 Exemplo de Uso |
| :--- | :--- | :--- |
| ➕ `SOMA` | Soma dois números inteiros | `SOMA 10 5` |
| ➖ `SUB` | Subtrai o segundo número do primeiro | `SUB 10 5` |
| ✖️ `MUL` | Multiplica dois números inteiros | `MUL 10 5` |
| ➗ `DIV` | Divide o primeiro pelo segundo (com proteção contra divisão por zero) | `DIV 10 5` |
| 🚪 `SAIR` | Encerra a conexão de forma limpa | `SAIR` |

## 🛡️ Robustez e Tratamento de Erros

* 🚫 **Validação de Entrada**: O servidor verifica comandos vazios, incompletos ou inválidos antes de tentar processá-los.
* ⚠️ **Tratamento de Tipos**: Conversões numéricas protegidas (`strconv.Atoi`) evitam que falhas de formato derrubem o servidor.
* 🔄 **Resiliência de Rede**: Uso estratégico de `continue` e envio de mensagens claras de erro (`ERRO: divisao por zero`, `ERRO: formato invalido`) mantêm o loop de atendimento ativo.
* ⚙️ **Multicliente**: Arquitetura concorrente baseada em *goroutines* com `net.Listen` e `net.Accept`.

## 🚀 Como Executar

1. **Compilar e Iniciar o Servidor**:
   Em um terminal, navegue até a pasta do projeto e execute:
   ```bash
   go run server.go
2. **Iniciar o Cliente**:
Em outro terminal separado (você pode abrir vários para testar a concorrência), execute:
```bash
   go run client.go