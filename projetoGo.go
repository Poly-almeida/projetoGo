package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type calcular struct{
	dificuldade int
	valor1 int
	valor2 int
	operacao int
	resultado int
	pontos int
}

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

func (calc calcular) jogar(){
	var dificuldade int
	var continua bool = true

	for continua == true{
		cmd := exec.Command("cmd", "/c", "cls") 
        cmd.Stdout = os.Stdout
		cmd.Run()
		
		fmt.Println("Informe o nível de dificuldade desejado [1, 2, 3 ou 4]:")
		fmt.Scanln(&dificuldade)
		calc.dificuldade = dificuldade

		if calc.dificuldade == 1{
			calc.valor1 = rand.Intn(11)
			calc.valor2 = rand.Intn(11)
		} else if calc.dificuldade == 2{
			calc.valor1 = rand.Intn(101)
			calc.valor2 = rand.Intn(101)
		} else if calc.dificuldade == 3{
			calc.valor1 = rand.Intn(1001)
			calc.valor2 = rand.Intn(1001)
		} else if calc.dificuldade == 4{
			calc.valor1 = rand.Intn(10001)
			calc.valor2 = rand.Intn(10001)
		} else{
			calc.valor1 = rand.Intn(100001)
			calc.valor2 = rand.Intn(100001)
		}

		calc.operacao = rand.Intn(3)
		var resposta int
		fmt.Println("Informe o resultado para a seguinte operação:")

		if calc.operacao == 0{
			fmt.Println(calc.valor1, " + ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.soma(resposta) == 1{
				calc.pontos = calc.pontos + 1
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
		} else if calc.operacao == 1{
			fmt.Println(calc.valor1, " - ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.subtrai(resposta) == 1{
				calc.pontos = calc.pontos + 1
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
		} else if calc.operacao == 2{
			fmt.Println(calc.valor1, " * ", calc.valor2, " = ")
			fmt.Scanln(&resposta)

			if calc.multiplica(resposta) == 1{
				calc.pontos = calc.pontos + 1
				fmt.Println("Você tem ",calc.pontos," ponto(s).")
			}
		} else {
			fmt.Println("Essa operação não é válida")
		}

		fmt.Println("Deseja continuar jogando? (true) (false)")
		fmt.Scanln(&continua)
	}
}

func main() {
	var calc calcular
	calc.jogar()
}