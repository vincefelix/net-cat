package tcp

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	conn     net.Conn
	nickname string
}

var clients []Client

func HandleClient(conn net.Conn) {
	client := Client{conn: conn}
	if len(clients) < 2 {
		Welcome(conn)
		for {
			conn.Write([]byte("[ENTER YOUR NAME] ➡ :"))
			username, err := client.readInput()
			if err != nil {
				name_error := fmt.Sprintf("❌ Error while receiving the username: %s\n", err)
				logs(name_error)
				fmt.Print(name_error)
				conn.Close()
				return
			}

			client.nickname = strings.TrimSpace(username)
			if client.nickname == "" {
				conn.Write([]byte("❌ The name is not valid please try again\n"))
				name_invalid := fmt.Sprintf("❌ User %s entered an invalid name.\n", conn.LocalAddr().String())
				logs(name_invalid)
				fmt.Print(name_invalid)
				continue
			} else {
				break
			}
		}
		initial := client.nickname
		clients = append(clients, client)
		join_mess := fmt.Sprintf("🟢 %s has joined our chat...\n", client.nickname)
		client.Sendmess(join_mess, clients)
		logs(join_mess)
		fmt.Print(join_mess)

		history_show, _ := os.ReadFile("files/history.txt")
		conn.Write(history_show)

		defer func() {
			leave_mess := fmt.Sprintf("🔻 %s has left our chat...\n", client.nickname)
			client.Sendmess(leave_mess, clients)
			client.promptline()
			logs(leave_mess)
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
		client.Room(conn, initial)
	} else {
		conn.Write([]byte("🚫 Sorry , chatroom is full\n"))
		chat_full_mess := fmt.Sprintf("User %s tried to connect but the chatroom is full\n", conn.LocalAddr().String())
		conn.Write([]byte("tap ENTER x2 to exit\n"))
		fmt.Print(chat_full_mess)
		logs(chat_full_mess)
		conn.Close()
		return
	}

}
