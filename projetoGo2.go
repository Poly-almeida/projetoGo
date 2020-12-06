package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)
// definição da estrutura do jogo
type calcular struct{
	valor1 int
	valor2 int
	operacao int
	operacao_str string
	resultado int
	pontos int
}
var clear map[string]func() //create a map for storing clear funcs
func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") // comando para sistema Unix
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") // comando para o windows
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func (calc *calcular) rodada() {
	calc.operacao = rand.Intn(3) //Gera um número de 0 a 2
	if calc.operacao == 0 { // Se o número gerado em calc.operacao for 0, uma soma arbtrária é mostrada na tela
		fmt.Println("OP 0")
		// valores para soma são da forma 10 * pontuação + 9 		
		// 0pts => de 0 a 9		
		// 1pts => de 10 a 19		
		// 2pts => de 20 a 29 
		// [...]		
		var base int = 10 * calc.pontos
		calc.valor1 = rand.Intn(base + 9) + base
		calc.valor2 = rand.Intn(base + 9) + base
		calc.resultado = calc.valor1 + calc.valor2
		calc.operacao_str = " + "

	} else if calc.operacao == 1{ //Se o número gerado em calc.operacao for 1, uma subtração arbtrária é mostrada na tela
		fmt.Println("OP 1")
		// valores para subtração são da forma 10 * pontuação + 9 		
		// 0pts => de 0 a 9		
		// 1pts => de 10 a 19		
		// 2pts => de 20 a 29 
		// [...]		
		var base int = 10 * calc.pontos
		calc.valor1 = rand.Intn(base + 9) + base
		calc.valor2 = rand.Intn(base + 9) + base

		// evitar subtrações com resultado negativo
		if calc.valor2 > calc.valor1 {
			var temp int = calc.valor1;
			calc.valor1 = calc.valor2;
			calc.valor2 = temp;
		}
		calc.resultado = calc.valor1 - calc.valor2
		calc.operacao_str = " - "

	} else { //Se o número gerado em calc.operacao for 2, uma multiplicação arbtrária é mostrada na tela
		// valores para multiplicação são da forma 3 * pontuação + 2 		
		// 0pts => de 0 a 2		
		// 1pts => de 3 a 5		
		// 2pts => de 6 a 8 
		// [...]		
		fmt.Println("OP 2")
		var base int = 3 * calc.pontos
		calc.valor1 = rand.Intn(base + 2) + base
		calc.valor2 = rand.Intn(base + 2) + base
		calc.resultado = calc.valor1 * calc.valor2
		calc.operacao_str = " * "
	}
}

func (calc *calcular) finalizar() {

}

// Método principal do jogo
func (calc calcular) jogar(){
	rand.Seed(time.Now().UnixNano())
	calc.pontos = 0
	for true{
		CallClear() //limpar a tela para o jogo começar
		
		calc.rodada()
		var resposta int
		fmt.Println("Operação atual:  ", calc.operacao)
		fmt.Println("Você tem ", calc.pontos," ponto(s).")
		fmt.Println("Informe o resultado para a seguinte operação:")
		fmt.Print(calc.valor1, calc.operacao_str, calc.valor2, " = ")
		fmt.Scanln(&resposta)
		
		if resposta == calc.resultado {
			calc.pontos += 1
		} else {
			calc.finalizar()
		}

		if calc.pontos == 10{
			fmt.Println("Parabéns! Vc venceu!")
			break
		}
	}
}

func main() {  // Função principal que chama o método jogar() onde toda a lógica do jogo acontece
	var calc calcular
	calc.jogar()
}