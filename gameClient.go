package main

import (
	// "os"
	"bufio"
	"fmt"
	"net"
	"strings"
	"strconv"
)

const (
	AWAITING	= 0
	RUNNING		= 1

	UNAUTHED	= 10
	AUTHED		= 11
)

type User struct {
	username string
	password string
	status int
}

type GameData struct {
	round int
	max int
	user *User
	op int
	val1 int
	val2 int
}

type GameManager struct {
	conn net.Conn
	status int
	data *GameData
}

func (manager *GameManager) clearData() {
	manager.data.round = 0
	manager.data.max = 0
	manager.data.op = 0
	manager.data.val1 = 0
	manager.data.val2 = 0
}

func (manager *GameManager) parseData(message string) {
	manager.clearData()
	var data = strings.Split(message, "|")

	manager.data.op, _ = strconv.Atoi(data[0])
	manager.data.val1, _ = strconv.Atoi(data[1])
	manager.data.val2, _ = strconv.Atoi(data[2])
	manager.data.round, _ = strconv.Atoi(data[3])
	manager.data.max, _ = strconv.Atoi(data[4])
}

func (manager *GameManager) connect() {
	c, _ := net.Dial("tcp", "localhost:3333")
	manager.conn = c
}

func (manager *GameManager) request(f string, args []string) string{
	fmt.Fprintf(manager.conn, f + "::" + strings.Join(args, "|") + "\n")
	message, _ := bufio.NewReader(manager.conn).ReadString('\n')
	return message
}

func (manager *GameManager) login() {
	for manager.data.user.status == UNAUTHED {
		fmt.Print("Digite seu usuário:\t")
		fmt.Scanln(&manager.data.user.username)
		fmt.Print("Digite sua senha:\t")
		fmt.Scanln(&manager.data.user.password)
		
		var message = manager.request("LOGIN", []string{manager.data.user.username, manager.data.user.password})
		if message == "SUCCESS\n" {
			manager.data.user.status = AUTHED
		}
		if manager.data.user.status == UNAUTHED {
			fmt.Print("Acesso negado! Tente novamente.\n")
		}
	}
}

func (manager *GameManager) start() {
	manager.login()
	var message = manager.request("START", []string{manager.data.user.username})
	manager.parseData(message)
	fmt.Println("\n----- INÍCIO DE PARTIDA -----")
	fmt.Println("Seu recorde é de", manager.data.max, "ponto(s).")
	fmt.Println("-----------------------------\n")
}

func (manager *GameManager) verifyAnswer(answer int) bool {
	var message = manager.request("VERIFY", []string{manager.data.user.username, strconv.Itoa(answer)})
	if message == "SUCCESS\n"{
		return true
	}
	return false
}

func (manager *GameManager) nextRound() bool{
	var message = manager.request("NEXT", []string{manager.data.user.username})
	manager.parseData(message)

	var op_str = "+"
	if manager.data.op == 1 {
		op_str = "-"
	} else if manager.data.op == 2 {
		op_str = "*"
	}
	fmt.Println("Você tem", manager.data.round, "ponto(s).")
	fmt.Println("Informe o resultado para a seguinte operação:")
	fmt.Print(manager.data.val1, op_str, manager.data.val2, " = ")
	
	var answer int
	fmt.Scanln(&answer)
	return manager.verifyAnswer(answer)		
}

func (manager *GameManager) finish() {
	var message = manager.request("FINISH", []string{manager.data.user.username})
	if message == "SUCCESS\n" {
		manager.clearData()
	}
	
	fmt.Println("\n----- FIM DE JOGO -----")
	fmt.Println("Você fez", manager.data.round, "ponto(s).")
	fmt.Println("Seu recorde é de", manager.data.max, "ponto(s).")
	fmt.Println("-----------------------\nPressione [ENTER] para finalizar.")
	fmt.Scanln()
}

func (manager *GameManager) run() {
	manager.data = &GameData{
		round: 0,
		user: &User{
			username: "",
			password: "",
			status: UNAUTHED,
		},
		op: 0,
		val1: 0,
		val2: 0,
	}

	manager.start()
	for manager.nextRound() {
		fmt.Println()
	}
	manager.finish()
}


func main() {
	manager := GameManager{
		conn: nil,
		status: AWAITING,
	}

	manager.connect()
	manager.run()
}
    