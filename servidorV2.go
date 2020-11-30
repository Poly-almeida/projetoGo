package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
)

func main() {

  fmt.Println("Aguardando conexões.....")

  // A parti da 3000 via protocolo tcp/ip
  ln, ERRO_1 := net.Listen("tcp", ":3000")
  if ERRO_1 != nil {
    fmt.Println(ERRO_1)
    /* Neste nosso exemplo vamos convencionar que a saída 3 está reservada para erros de conexão.
    IMPORTANTE: defers não serão executados quando utilizamos os.Exit() e a saída será imediata */
    os.Exit(3)
  }

  // aceitando conexões
  conexao, ERRO_2 := ln.Accept()
  if ERRO_2 != nil {
    fmt.Println(ERRO_2)
    os.Exit(3)
  }
//Fecha a conexao 
  defer ln.Close()

  fmt.Println("Conexão aceita...")
  // Servidor fica ativo de modo contínuo. (Para parar usar o comando no termina ctrl-c)
  for {
    // Receba a mensagem e pula uma linha com a expressão = (\n)
    mensagem, ERRO_3 := bufio.NewReader(conexao).ReadString('\n')
    //Havendo erro, a execução é parada
    if ERRO_3 != nil {
      //Mostra na tela o erro
      fmt.Println(ERRO_3)
      os.Exit(3)
    }

    // escreve no terminal a mensagem recebida
    fmt.Print("Mensagem recebida:", string(mensagem))

    // Para cada mensagem excrita no cliente, o servidor converte a msg recebida para caixa alta com o comando (ToUpper)
    novamensagem := strings.ToUpper(mensagem)

    // Devolva a mgs excrita pelo servidor, apos ser recebida pelo servidor.
    conexao.Write([]byte(novamensagem + "\n"))
  }
}
