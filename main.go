package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	conn     net.Conn
	nickname string
}

var clients []Client
var temp_historic, temp_log string

func main() {
	//1 step launching the server
	var port string
	args := os.Args[1:]
	if len(args) > 1 {
		log.Fatalln("USAGE]: ./TCPChat $port")
	} else if len(args) == 1 {
		port = ":" + args[0]
	} else {
		port = ":8989"
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error while listening to port %s : %s", port, err)
	}
	defer listener.Close()
	fmt.Printf("Server listen on port: %s\n", port)

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error while establishing connection : %s", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	client := Client{conn: conn}

	conn.Write([]byte("welcome to TCP-tchat!\n"))
	logo, err := os.ReadFile("linux_logo.txt")
	if err != nil {
		println("error while reading linux logo")
		return
	}

	conn.Write([]byte((string(logo)) + "\n"))
	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		username, err := client.readInput()
		if err != nil {
			name_error := fmt.Sprintf("Error while receiving the username: %s\n", err)
			temp_log += name_error
			history("log", temp_log)
			fmt.Print(name_error)
			conn.Close()
			return
		}

		client.nickname = strings.TrimSpace(username)
		if client.nickname == "" {
			conn.Write([]byte("The name is not valid please try again\nOr you can quit by using the --quit flag\n"))
			time_now := time.Now().Format("2006-01-02 15:04:05")
			name_invalid := fmt.Sprintf("User %s entered an invalid name at [%s]\n", conn.LocalAddr().String(), time_now)
			temp_log += name_invalid
			history("log", temp_log)
			fmt.Print(name_invalid)
			continue
		} else if client.nickname == "--quit" {
			time_now := time.Now().Format("2006-01-02 15:04:05")
			quitmess := fmt.Sprintf("User %s left the chat at [%s]\n", conn.LocalAddr().String(), time_now)
			temp_log += quitmess
			history("log", temp_log)
			fmt.Print(quitmess)
			client.conn.Close()
			break
		} else {
			break
		}
	}

	clients = append(clients, client)
	if len(clients) > 10 {
		conn.Write([]byte("chatroom is full\n"))
		time_now := time.Now().Format("2006-01-02 15:04:05")
		chat_full_mess := fmt.Sprintf("User %s tried to connect at [%s] but the chatroom is full\n", conn.LocalAddr().String(), time_now)
		fmt.Print(chat_full_mess)
		temp_log += chat_full_mess
		history("log", temp_log)
		os.Exit(1)
	}
	joinmess := fmt.Sprintf("%s has joined our chat...\n", client.nickname)
	time_now := time.Now().Format("2006-01-02 15:04:05")
	client.sendAll(joinmess)
	joinmesslog := fmt.Sprintf("%s has joined our chat at [%s]\n", client.nickname, time_now)
	temp_log += joinmesslog
	history("log", temp_log)
	fmt.Print(joinmesslog)

	history_show, errshow := os.ReadFile("history.txt")
	if errshow != nil {
		fmt.Println("history file is created")
	}
	conn.Write(history_show)

	defer func() {
		leavemess := fmt.Sprintf("%s has left our chat...\n", client.nickname)
		client.sendAll(leavemess)
		time_now := time.Now().Format("2006-01-02 15:04:05")
		leavemesslog := fmt.Sprintf("%s has left our chat at [%s]\n", client.nickname, time_now)
		temp_log += leavemesslog
		history("log", temp_log)
		fmt.Print(leavemesslog)
		client.conn.Close()

		// Supprimer le client de la liste des clients connectés
		for i, c := range clients {
			if c.conn == client.conn {
				clients = append(clients[:i], clients[i+1:]...)
				break
			}
		}
	}()

	for {
		cmd_now := time.Now().Format("2006-01-02 15:04:05") //changing the current time formats
		client.nickname = strings.TrimSpace(client.nickname)
		prompt := fmt.Sprintf("[%s][%s]: ", cmd_now, client.nickname)
		conn.Write([]byte(prompt))

		message, err := client.readInput()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error while reading message%s : %s", client.nickname, err)
			break
		}

		if strings.TrimSpace(message) == "--quit" {
			break
		} else if strings.TrimSpace(message) == "" {
			continue
		}

		//sending datas to the users connections
		//formating the data and sending it to connections
		now := time.Now().Format("2006-01-02 15:04:05") //changing the current time format to our will
		messtext := fmt.Sprintf("[%s][%s]: - %s", now, client.nickname, message)
		temp_historic += messtext
		history("history", temp_historic)
		client.sendAll(messtext)
		mess_notif := fmt.Sprintf("%s has sent a message at [%s]\n", client.nickname, now)
		temp_log += mess_notif
		history("log", temp_log)
		fmt.Print(mess_notif)

	}
}

func history(File, mess string) {
	file, err := os.Create(File + ".txt")
	if err != nil {
		fmt.Println("Error while creation the history file")
		return
	}

	_, err = file.WriteString(mess)
	if err != nil {
		fmt.Println("Error while writing nto history")
		return
	}
}

func (c *Client) readInput() (string, error) {
	// text := make([]byte, 1024)
	// n, err := c.conn.Read(text)
	// if err != nil {
	// 	return "", err
	// }
	// return string(buffer[:n]), nil
	//receiving datas from connection
	reader := bufio.NewReader(c.conn)
	mess, errmess := reader.ReadString('\n')
	//when server is off
	if errmess == io.EOF {
		return "", errmess
	}
	return mess, nil
}

func (c *Client) sendAll(message string) {
	for _, client := range clients {
		if client.conn != c.conn {
			_, err := client.conn.Write([]byte("\n" + message))
			if err != nil {
				log.Printf("Erreur lors de l'envoi du message à %s : %s", client.nickname, err)
			}
			cmd_now := time.Now().Format("2006-01-02 15:04:05") //changing the current time formats
			name := strings.TrimSpace(client.nickname)
			prompt := fmt.Sprintf("[%s][%s]: ", cmd_now, name)
			client.conn.Write([]byte(prompt))
			continue
		}
	}
}
