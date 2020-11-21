    package main
     
    import (
    	"bufio"
    	"io"
    	"log"
    	"net"
    	"strings"
    )
     
    func main() {
    	escuta, err := net.Listen("tcp", "0.0.0.0:9999")
    	if err != nil {
    		log.Fatalln(err)
    	}
    	defer escuta.Close()
     
    	for {
    		con, err := escuta.Accept()
    		if err != nil {
    			log.Println(err)
    			continue
    		}
     
    		go handleclienteSolicita(con)
    	}
    }
     
    func handleclienteSolicita(con net.Conn) {
    	defer con.Close()
     
    	clienteLeitura := bufio.NewReader(con)
     
    	for {
    		// Espera respota do cliente ao conectar
    		clienteSolicita, err := clienteLeitura.ReadString('\n')
     
    		switch err {
    		case nil:
    			clienteSolicita := strings.TrimSpace(clienteSolicita)
    			if clienteSolicita == ":QUIT" {
    				log.Println("Cliente solicita ao servidor o fechamento da conexao")
    				return
    			} else {
    				log.Println(clienteSolicita)
    			}
    		case io.EOF:
    			log.Println("Cliente fecha a conexao apos o terminio do processo")
    			return
    		default:
    			log.Printf("Erro: %v\n", err)
    			return
    		}
     
    		// Cliente responde a requisasao da conexao
    		if _, err = con.Write([]byte("GOT IT!\n")); err != nil {
    			log.Printf("Falha na reposda do Cliente: %v\n", err)
    		}
    	}
    }