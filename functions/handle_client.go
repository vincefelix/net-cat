package tcp

import (
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

var (
	clients                 []Client
	temp_historic, temp_log string
)

func HandleClient(conn net.Conn) {
	client := Client{conn: conn}
	clients = append(clients, client)
	if len(clients) < 11 {
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
				logs(name_error, temp_log)
				fmt.Print(name_error)
				conn.Close()
				return
			}

			client.nickname = strings.TrimSpace(username)
			if client.nickname == "" {
				conn.Write([]byte("The name is not valid please try again\n"))
				name_invalid := fmt.Sprintf("User %s entered an invalid name.\n", conn.LocalAddr().String())
				logs(name_invalid, temp_log)
				fmt.Print(name_invalid)
				continue
			} else {
				break
			}
		}

		join_mess := fmt.Sprintf("%s has joined our chat...\n", client.nickname)
		client.sendAll(join_mess, clients)
		logs(join_mess, temp_log)
		fmt.Print(join_mess)

		history_show, _ := os.ReadFile("history.txt")
		conn.Write(history_show)

		defer func() {
			leave_mess := fmt.Sprintf("%s has left our chat...\n", client.nickname)
			client.sendAll(leave_mess, clients)
			client.promptline()
			logs(leave_mess, temp_log)
			fmt.Print(leave_mess)
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
			client.promptline()
			message, err := client.readInput()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Printf("Error while reading message%s : %s", client.nickname, err)
				break
			}

			if strings.TrimSpace(message) == "--change" {
				initial := client.nickname
				for {
					conn.Write([]byte("[ENTER YOUR NAME]: "))
					username, err := client.readInput()
					if err != nil {
						name_error := fmt.Sprintf("Error while receiving the username: %s\n", err)
						logs(name_error, temp_log)
						fmt.Print(name_error)
						conn.Close()
						return
					}

					client.nickname = strings.TrimSpace(username)
					if client.nickname == "" {
						conn.Write([]byte("The name is not valid please try again\n"))
						name_invalid := fmt.Sprintf("User %s entered an invalid name.\n", conn.LocalAddr().String())
						logs(name_invalid, temp_log)
						fmt.Print(name_invalid)
						continue
					} else {
						break
					}
				}
				change_mess := fmt.Sprintf("%s changed his username to %s ...\n", initial, client.nickname)
				client.sendAll("\n"+change_mess, clients)
				logs(change_mess, temp_log)
				fmt.Print(change_mess)
				continue
			} else if strings.TrimSpace(message) == "" {
				continue
			}

			//sending datas to the users connections
			//formating the data and sending it to connections
			now := time.Now().Format("2006-01-02 15:04:05") //changing the current time format to our will
			messtext := fmt.Sprintf("[%s][%s]: - %s", now, client.nickname, message)
			client.sendAll(messtext, clients)

			history(messtext, temp_historic)
			mess_notif := fmt.Sprintf("%s has sent a message ...\n", client.nickname)
			logs(mess_notif, temp_log)
			fmt.Print(mess_notif)
		}
	} else {
		conn.Write([]byte("chatroom is full\n"))
		chat_full_mess := fmt.Sprintf("User %s tried to connect but the chatroom is full\n", conn.LocalAddr().String())
		fmt.Print(chat_full_mess)
		logs(chat_full_mess, temp_log)
		conn.Close()
		return
	}
}
