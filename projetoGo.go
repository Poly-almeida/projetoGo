package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)
// definição da estrutura do jogo
type calcular struct{
	valor1 int
	valor2 int
	operacao int
	resultado int
	pontos int
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
	var continua bool = true

	for continua == true{
		cmd := exec.Command("cmd", "/c", "cls") // comando para limpar a tela no windows
        cmd.Stdout = os.Stdout
		cmd.Run()
		
		// Nesse ponto, é gerado 2 números aleatorios
		calc.valor1 = rand.Intn(1001) //Gera um número de 0 a 1000
		calc.valor2 = rand.Intn(1001)

		calc.operacao = rand.Intn(3) //Gera um número de 0 a 2
		var resposta int
		fmt.Println("Informe o resultado para a seguinte operação:")

		if calc.operacao == 0{ //Se o número gerado em calc.operacao for 0, uma soma arbtrária é mostrada na tela
			fmt.Print(calc.valor1, " + ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.soma(resposta) == 1{ // Aqui o método soma() verifica se a resposta dada é correta
				calc.pontos = calc.pontos + 1 // Atribui 1 ponto se a resposta for correta
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
		} else if calc.operacao == 1{ //Se o número gerado em calc.operacao for 1, uma subtração arbtrária é mostrada na tela
			fmt.Print(calc.valor1, " - ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.subtrai(resposta) == 1{ // Aqui o método subtrai() verifica se a resposta dada é correta
				calc.pontos = calc.pontos + 1
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
		} else { //Se o número gerado em calc.operacao for 2, uma multiplicação arbtrária é mostrada na tela
			fmt.Print(calc.valor1, " * ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.multiplica(resposta) == 1{ // Aqui o método multiplica() verifica se a resposta dada é correta
				calc.pontos = calc.pontos + 1
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
		}

		fmt.Println("Deseja continuar jogando? (true) (false)")
		fmt.Scanln(&continua)
	}
}

func main() {  // Função principal que chama o método jogar() onde toda a lógica do jogo acontece
	var calc calcular
	calc.jogar()
}