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

func (c *Client) Room(conn net.Conn, initial string) {
	for {
		c.promptline()
		message, err := c.readInput()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("❌ Error while reading message%s : %s", c.nickname, err)
			break
		}

		if strings.TrimSpace(message) == "--change" {
			c.change(conn, initial)
			continue
		} else if strings.TrimSpace(message) == "--avatar" {
			c.change_av(conn, initial)
			continue
		} else if strings.TrimSpace(message) == "--remove=avatar" {
			c.remove_avatar(initial)
			continue
		} else if strings.TrimSpace(message) == "" {
			continue
		}

		//sending datas to the users connections
		//formating the data and sending it to connections
		now := time.Now().Format("2006-01-02 15:04:05") //changing the current time format to our will
		messtext := fmt.Sprintf("[%s][%s]: - %s", now, c.nickname, message)
		c.Sendmess(messtext, clients)

		history(messtext)
		mess_notif := fmt.Sprintf("🟢 %s has sent a message ... 💬\n", c.nickname)
		logs(mess_notif)
		fmt.Print(mess_notif)
	}
}

func (c *Client) promptline() {
	cmd_now := time.Now().Format("2006-01-02 15:04:05") //changing the current time formats
	c.nickname = strings.TrimSpace(c.nickname)
	prompt := fmt.Sprintf("[%s][%s]: ", cmd_now, c.nickname)
	c.conn.Write([]byte(prompt))
}

func (c *Client) readInput() (string, error) {
	reader := bufio.NewReader(c.conn)
	mess, errmess := reader.ReadString('\n')
	//when server is off
	if errmess == io.EOF {
		return "", errmess
	}
	return mess, nil
}

func (c *Client) Sendmess(message string, b []Client) {
	for _, client := range b {
		if client.conn != c.conn {
			_, err := client.conn.Write([]byte("\n" + message))
			if err != nil {
				log.Printf("❌ Error while sending message %s : %s", client.nickname, err)
				continue
			}
			client.promptline()
			continue
		}
	}
}


