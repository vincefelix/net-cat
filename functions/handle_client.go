package tcp

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// this structure stores datas related to a user
type Client struct {
	conn     net.Conn
	nickname string
}

var clients []Client // stores the connected users

func HandleClient(conn net.Conn) {
	client := Client{conn: conn}
	if len(clients) < 10 { //delimiting the number of connections
		Welcome(conn) // welcome message

		//--receiving the username
		for {
			conn.Write([]byte("[ENTER YOUR NAME] ➡ :"))
			username, err := client.readInput()
			if err != nil {
				name_error := fmt.Sprintf("❌ Error while receiving the username: %s\n", err)
				logs(name_error)
				fmt.Print(name_error)
				conn.Close() //shutting the connection
				return
			}

			client.nickname = strings.TrimSpace(username)
			if client.nickname == "" { // username can't be empty
				conn.Write([]byte("❌ The name is not valid please try again\n"))
				name_invalid := fmt.Sprintf("❌ User %s entered an invalid name.\n", conn.LocalAddr().String())
				logs(name_invalid)      // storing activity in a log file
				fmt.Print(name_invalid) //terminal logs
				continue
			} else {
				break
			}
		}
		initial := client.nickname        // stores the first name entered by the user
		clients = append(clients, client) // adding the client to our array
		join_mess := fmt.Sprintf("🟢 %s has joined our chat...\n", client.nickname)
		client.Sendmess(join_mess, clients) // notifying al the chat members
		logs(join_mess)                     // storing activity in a log file
		fmt.Print(join_mess)                //terminal logs

		history_show, _ := os.ReadFile("files/history.txt") //displaying the chat history for a newcomer
		conn.Write(history_show)                            //writing to user's connection address

		defer func() {
			leave_mess := fmt.Sprintf("🔻 %s has left our chat...\n", client.nickname)
			client.Sendmess(leave_mess, clients) // notifying all te chat members
			client.promptline()                  //displaying the promptline under after the notification
			logs(leave_mess)                     // storing activity in a log file
			fmt.Print(leave_mess)                //terminal logs
			client.conn.Close()                  // shutting the connection

			// deleting the disconneted user from the clients array
			for i, c := range clients {
				if c.conn == client.conn {
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
		}()
		client.Room(conn, initial) // managing all the user's activities in the chatroom
	} else {
		conn.Write([]byte("🚫 Sorry , chatroom is full\n")) // notifying the user that limit is reached
		chat_full_mess := fmt.Sprintf("User %s tried to connect but the chatroom is full\n", conn.LocalAddr().String())
		conn.Write([]byte("tap ENTER to exit\n"))
		fmt.Print(chat_full_mess) //terminal logs
		logs(chat_full_mess)      // storing activity in a log file
		conn.Close()              //shutting the connection
		return
	}

}
