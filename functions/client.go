package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

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

func (c *Client) sendAll(message string, b []Client) {
	for _, client := range b {
		if client.conn != c.conn {
			_, err := client.conn.Write([]byte("\n"+message))
			if err != nil {
				log.Printf("Error while sending message %s : %s", client.nickname, err)
			}
			client.promptline()
			continue
		}
	}
}
