package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)
// definição da estrutura do jogo
type calcular struct{
	valor1 int
	valor2 int
	operacao int
	resultado int
	pontos int
}
var clear map[string]func() //create a map for storing clear funcs
// função reponsável por executar o comando de limpar a tela de acordo com sistema operacional
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

// Método que soma dois números e verifica se o resultado é correto
func (calc calcular) soma(resposta int) int{
	var resultado int
	resultado = calc.valor1 + calc.valor2
	calc.resultado = resultado

	var certo int 
	if resposta == calc.resultado{
		fmt.Println("Resposta correta")
		certo = 1
	} else {
		fmt.Println("Resposta errada")
		certo = 0
	}

	fmt.Println(calc.valor1, " + ", calc.valor2, " = ", calc.resultado)

	return certo
}
// Método que subtrai dois números e verifica se o resultado é correto
func (calc calcular) subtrai(resposta int) int{
	var resultado int
	resultado = calc.valor1 - calc.valor2
	calc.resultado = resultado

	var certo int 
	if resposta == calc.resultado{
		fmt.Println("Resposta correta")
		certo = 1
	} else {
		fmt.Println("Resposta errada")
		certo = 0
	}

	fmt.Println(calc.valor1, " - ", calc.valor2, " = ", calc.resultado)

	return certo
}
// Método que multiplica dois números e verifica se o resultado é correto
func (calc calcular) multiplica(resposta int) int{
	var resultado int
	resultado = calc.valor1 * calc.valor2
	calc.resultado = resultado

	var certo int 
	if resposta == calc.resultado{
		fmt.Println("Resposta correta")
		certo = 1
	} else {
		fmt.Println("Resposta errada")
		certo = 0
	}

	fmt.Println(calc.valor1, " * ", calc.valor2, " = ", calc.resultado)

	return certo
}
// Método principal do jogo
func (calc calcular) jogar(){

	for true{
		CallClear() //limpar a tela para o jogo começar
		
		// Nesse ponto, é gerado 2 números aleatorios
		calc.valor1 = rand.Intn(1001) //Gera um número de 0 a 1000
		calc.valor2 = rand.Intn(1001)

		calc.operacao = rand.Intn(3) //Gera um número de 0 a 2
		var resposta int
		fmt.Println("Informe o resultado para a seguinte operação:")

		if calc.operacao == 0{ //Se o número gerado em calc.operacao for 0, uma soma arbtrária é mostrada na tela
			fmt.Println("Você tem ",calc.pontos," ponto(s).")
			fmt.Print(calc.valor1, " + ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.soma(resposta) == 1{ // Aqui o método soma() verifica se a resposta dada é correta
				calc.pontos = calc.pontos + 1 // Atribui 1 ponto se a resposta for correta
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
		} else if calc.operacao == 1{ //Se o número gerado em calc.operacao for 1, uma subtração arbtrária é mostrada na tela
			fmt.Println("Você tem ",calc.pontos," ponto(s).")
			fmt.Print(calc.valor1, " - ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.subtrai(resposta) == 1{ // Aqui o método subtrai() verifica se a resposta dada é correta
				calc.pontos = calc.pontos + 1
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
		} else { //Se o número gerado em calc.operacao for 2, uma multiplicação arbtrária é mostrada na tela
			fmt.Println("Você tem ",calc.pontos," ponto(s).")
			fmt.Print(calc.valor1, " * ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.multiplica(resposta) == 1{ // Aqui o método multiplica() verifica se a resposta dada é correta
				calc.pontos = calc.pontos + 1
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
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