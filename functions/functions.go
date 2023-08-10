package tcp

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var temp_historic, temp_log string

func Welcome(conn net.Conn) {
	conn.Write([]byte("welcome to TCP-tchat!\n"))
	logo, err := os.ReadFile("files/linux_logo.txt")
	if err != nil {
		println("❌ error while reading linux logo")
		return
	}
	conn.Write([]byte((string(logo)) + "\n"))
}

func Specify_port() string {
	port := ""
	args := os.Args[1:]
	if len(args) > 1 {
		log.Fatalln("USAGE]: ./TCPChat $port")
	} else if len(args) == 1 {
		port = ":" + args[0]
	} else {
		port = ":8989"
	}
	return port
}

func history(mess string) {
	file, err := os.Create("files/history.txt")
	if err != nil {
		fmt.Println("❌ Error while creation the history file")
		return
	}
	temp_historic += mess
	_, err = file.WriteString(temp_historic)
	if err != nil {
		fmt.Println("❌ Error while writing nto history")
		return
	}
}

func logs(notification string) {
	time := time.Now().Format("2006-01-02 15:04:05") //changing the current time format to our will

	file, err := os.Create("files/logs.txt")
	if err != nil {
		fmt.Println("❌ Error while creation the history file")
		return
	}
	temp_log += fmt.Sprintf("🔔 [%s] --> ", time) + notification
	_, err = file.WriteString(temp_log)
	if err != nil {
		fmt.Println("❌ Error while writing nto history")
		return
	}
}

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}
