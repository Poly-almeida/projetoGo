    package main
     
    import (
    	"bufio"
    	"io"
    	"log"
    	"net"
    	"os"
    	"strings"
    )
     
    func main() {
    	con, err := net.Dial("tcp", "0.0.0.0:9999")
    	if err != nil {
    		log.Fatalln(err)
    	}
    	defer con.Close()
     
    	clienteLeitura := bufio.NewReader(os.Stdin)
    	serverLeitura := bufio.NewReader(con)
     
    	for {
    		// Espera a requisao da conexao do cliente
    		clientePede, err := clienteLeitura.ReadString('\n')
     
    		switch err {
    		case nil:
    			clientePede := strings.TrimSpace(clientePede)
    			if _, err = con.Write([]byte(clientePede + "\n")); err != nil {
    				log.Printf("Falha no envio da requisao ao Cliente: %v\n", err)
    			}
    		case io.EOF:
    			log.Println("Conexao fechada pelo CLiente")
    			return
    		default:
    			log.Printf("Erro do Cliente: %v\n", err)
    			return
    		}
     
    		// Espera a resposta do sercidor 
    		serverResponde, err := serverLeitura.ReadString('\n')
     
    		switch err {
    		case nil:
    			log.Println(strings.TrimSpace(serverResponde))
    		case io.EOF:
    			log.Println("Conexao fechada pelo servidor")
    			return
    		default:
    			log.Printf("Erro no Servidor: %v\n", err)
    			return
    		}
    	}
    }