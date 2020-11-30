package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // Realizando a conexão na porta 3000 via protocolo TCP/IP (teste local)
  conexao, erro1 := net.Dial("tcp", "127.0.0.1:3000")
  if ERRO_01 != nil {
    fmt.Println(ERRO_01)
    os.Exit(3)
  }
//Executando o lanço ou uma condição para envio de msg via terminal
  for {
    // passando a mensagem via terminal ao servidor
    leitura := bufio.NewReader(os.Stdin)
    fmt.Print("O texto foi enviando ao servidor: ")
    texto, ERRO_2 := leitura.ReadString('\n')
    if ERRO_2 != nil {
      fmt.Println(ERRO_2)
      os.Exit(3)
    }

    // enviando a msg do cliente conectado via socket
    fmt.Fprintf(conexao, texto+"\n")

    // Resposta do servidor ao envio da msg 
    mensagem, ERRO_3 := bufio.NewReader(conexao).ReadString('\n')
    if ERRO_3 != nil {
      fmt.Println(ERRO_3)
      os.Exit(3)
    }
    // Devolução da resposta do escrita do servidor ao cliente
    fmt.Print("Mensagem do Servidor: " + mensagem)
  }
}