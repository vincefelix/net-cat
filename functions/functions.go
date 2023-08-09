package tcp

import (
	"fmt"
	"log"
	"os"
	"time"
)

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

func history(mess, receiver string) {
	file, err := os.Create("history.txt")
	if err != nil {
		fmt.Println("Error while creation the history file")
		return
	}
	receiver += mess
	_, err = file.WriteString(receiver)
	if err != nil {
		fmt.Println("Error while writing nto history")
		return
	}
}

func logs(notification, receiver string) {
	time := time.Now().Format("2006-01-02 15:04:05") //changing the current time format to our will

	file, err := os.Create("logs.txt")
	if err != nil {
		fmt.Println("Error while creation the history file")
		return
	}
	receiver += fmt.Sprintf("[%s] --> ", time) + notification
	_, err = file.WriteString(receiver)
	if err != nil {
		fmt.Println("Error while writing nto history")
		return
	}
}
