package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "strings"
    "math/rand"
    "time"
    "strconv"
)

type ServerGameData struct {
	round int
	max int
	op int
	val1 int
	val2 int
	res int
}

func (data ServerGameData) toString() string {
    return strconv.Itoa(data.op) + "|" + strconv.Itoa(data.val1) + "|" + strconv.Itoa(data.val2) + "|" + strconv.Itoa(data.round) + "|"+ strconv.Itoa(data.max) + "|"
}

func login(username string, password string) string{
    switch username {
    case "glins":
        if password == "1234" {
            return "SUCCESS"
        }
    case "glins2":
        if password == "12345" {
            return "SUCCESS"
        }
    }
    return "FAIL"
}

var rounds map[string] *ServerGameData 

func start(username string) string{
    if obj, ok := rounds[username]; ok {
        obj.round = 0
    } else {
        rounds[username] = &ServerGameData {
            round: 0,
            op: 0,
            val1: 0,
            val2: 0,
            res: 0,
        }
    }

    return rounds[username].toString()
}

func finish(username string) string {
    if obj, ok := rounds[username]; ok {
        obj.round = 0
    } else {
        rounds[username] = &ServerGameData {
            round: 0,
            op: 0,
            val1: 0,
            val2: 0,
            res: 0,
        }
    }

    return rounds[username].toString()
}

func nextRound(username string) string{
    var data = rounds[username]
    data.op = rand.Intn(3) 
	if data.op == 0 { 
		// 0pts => from 1 to 10		
		// 1pts => from 11 to 20		
		// 2pts => from 21 to 30 
		// [...]		
		var base int = data.round * 10 + 1
		data.val1 = rand.Intn(base + 9) + base
        data.val2 = rand.Intn(base + 9) + base
        data.res = data.val1 + data.val2
	} else if data.op == 1 {
		// 0pts => from 1 to 10		
		// 1pts => from 11 to 20		
		// 2pts => from 21 to 30 
		// [...]		
		var base int = data.round * 10
		data.val1 = rand.Intn(base + 9) + base
		data.val2 = rand.Intn(base + 9) + base
        
		if data.val2 > data.val1 {
			var temp int = data.val1;
			data.val1 = data.val2;
			data.val2 = temp;
		}
        data.res = data.val1 - data.val2
        } else { 
		// 0pts => from 0 to 2		
		// 1pts => from 3 to 5		
		// 2pts => from 6 to 8 
		// [...]		
		var base int = data.round * 3
		data.val1 = rand.Intn(base + 2) + base
		data.val2 = rand.Intn(base + 2) + base
        data.res = data.val1 * data.val2
    }
    var res = data.toString()
    return res
}

func verify(username string, answer string) string{
    var value, _ = strconv.Atoi(answer)
    var data = rounds[username]
    if data.res == value {
        data.round++
        if data.round > data.max {
            data.max = data.round
        }
        return "SUCCESS"
    }
    return "FAIL"
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        message := scanner.Text()
        fmt.Println(">> RECEIVED", message)

        data := strings.Split(message, "::")
        if data[0] == "LOGIN" {
            var args = strings.Split(data[1], "|")
            conn.Write([]byte(login(args[0], args[1]) + "\n"))
            continue
        }
        if data[0] == "START" {
            conn.Write([]byte(start(data[1]) + "\n"))
            continue
        }
        if data[0] == "NEXT" {
            conn.Write([]byte(nextRound(data[1]) + "\n"))
            continue
        }
        if data[0] == "VERIFY" {
            var args = strings.Split(data[1], "|")
            fmt.Print(data[1])
            conn.Write([]byte(verify(args[0], args[1]) + "\n"))
            continue
        }
        if data[0] == "FINISH" {
            conn.Write([]byte(finish(data[1]) + "\n"))
            continue
        }

        conn.Write([]byte(strings.ToUpper(message) + "\n"))
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("error:", err)
    }
}

func main() {
	rand.Seed(time.Now().UnixNano())
    rounds = make(map[string] *ServerGameData)

    ln, err := net.Listen("tcp", "127.0.0.1:3333")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("----- SERVER STARTED SUCCESSFULLY -----")
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(" >> ----- NEW CLIENT CONNECTED -----")
        go handleConnection(conn)
    }
}