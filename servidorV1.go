package main
// Codigo simples para simular um servidor, na linguagem GO
import (
	"fmt"
	"net/http"
)
//Texto que de teste na pagina do navegador, no Titulo
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Teste de servidor</h1>")
	fmt.Fprintf(w, "<br>")
	fmt.Fprintf(w, "Teste OK")
	fmt.Fprintf(w, "<br>")
	fmt.Fprintf(w, "<h2>Criacao do servidor OK</h2>")
}
//Checagem do teste 
func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Checagem do teste</h1>")
}
//Funcao para conectar ao servidor, usando o navegador na url: http://localhost:3000
//Note que uma vez que o codigo estiver rodando (depois que ele é copilado), ele fica 
//rodando até que termine ou faça o stop na copilação
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/checagem_teste", check)
	fmt.Println("Start servidor...")
	http.ListenAndServe(":3000", nil)
}