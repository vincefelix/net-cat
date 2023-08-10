package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// Room manages the messages and notifications broadcast
func (c *Client) Room(conn net.Conn, initial string) {
	for {
		//--receiving user's entry
		c.promptline()
		message, err := c.readInput()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("❌ Error while reading message%s : %s", c.nickname, err)
			break
		}
		//flag cases
		if strings.TrimSpace(message) == "--nick" { // user wants to change hisusername
			c.change(conn, initial)
			continue
		} else if strings.TrimSpace(message) == "--av" { // user wants to change his avatar
			c.change_av(conn, c.nickname)
			continue
		} else if strings.TrimSpace(message) == "--rmav" { // user wants to remove his avatar
			c.remove_avatar(c.nickname)
			continue
		} else if strings.TrimSpace(message) == "" { //cannot send an empty message
			continue //pursue the loop
		}

		//sending datas to the users connections
		//formating the data and sending it to connections
		now := time.Now().Format("2006-01-02 15:04:05") //changing the current time format to our will
		messtext := fmt.Sprintf("[%s][%s]: - %s", now, c.nickname, message)
		c.Sendmess(messtext, clients) //broadcastinfg the messages to all chatroom members

		history(messtext) //store the message in the history
		mess_notif := fmt.Sprintf("🟢 %s has sent a message ... 💬\n", c.nickname)
		logs(mess_notif)      // storing activity in a log file
		fmt.Print(mess_notif) //terminal logs
	}
}

// promptline display the command line with the time and the user's name
func (c *Client) promptline() {
	cmd_now := time.Now().Format("2006-01-02 15:04:05") //changing the current time formats
	c.nickname = strings.TrimSpace(c.nickname)
	prompt := fmt.Sprintf("[%s][%s]: ", cmd_now, c.nickname)
	c.conn.Write([]byte(prompt))
}

// readInput returns the users input and return an error when it encounters a problem
func (c *Client) readInput() (string, error) {
	reader := bufio.NewReader(c.conn)
	mess, errmess := reader.ReadString('\n')
	if errmess == io.EOF { // there is no more input available
		return "", errmess
	}
	return mess, nil
}

// Sendmess broadcasts messages and notifications to all users present in the chatroom
func (c *Client) Sendmess(message string, b []Client) {
	for _, client := range b { // range the tab storing the connected users
		if client.conn != c.conn { // sending to all users except the sender
			_, err := client.conn.Write([]byte("\n" + message))
			if err != nil {
				log.Printf("❌ Error while sending message %s : %s", client.nickname, err)
				continue
			}
			client.promptline() //displaying back th promptline
			continue            //pursue the loop
		}
	}
}
